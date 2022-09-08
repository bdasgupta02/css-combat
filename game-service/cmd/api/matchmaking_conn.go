package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
)

const (
	timedOut = "timed out"
	start    = "start"
	stop     = "stop"
	found    = "found"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// FIXME: IMPORTANT: should not send "start" message
// FIXME: ONLY TESTING make this authoratiative through sql query for low and high
// FIXME: use algorithm to find the bracket for mmr
// FIXME: high and low cant be below 0 in logic
// FIXME: use channels more than direct
// FIXME: timeout for total connection: 3 mins
// FIXME: when client disconnects this doesnt
// FIXME: send to channel pump instead

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveMatchWS(hub *matchHub, w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	username := claims["username"].(string)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &matchClient{
		hub:      hub,
		username: username,
		conn:     conn,
		send:     make(chan []byte, 256),
		update:   make(chan int, 16),
	}
	client.hub.register <- client

	go client.healthPump()
	go client.readPump(hub.queue)
	go client.matchmakingFinder(hub.queue)
	go client.timeoutChecker()
}

func (c *matchClient) timeoutChecker() {
	time.Sleep(4 * time.Minute)
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err == nil {
		w.Write([]byte(timedOut))
		w.Close()
	}
}

// "start" and "stop" to start or stop matchmaking
func (c *matchClient) readPump(q *matchQueue) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		msgstr := string(message)
		userPos := q.findQueueItem(c)

		if msgstr == stop {
			q.c.L.Lock()
			q.removeClientFromQ(&q.data[userPos])
			c.hub.unregister <- c
			c.conn.Close()
			q.c.Signal()
			q.c.L.Unlock()
			log.Printf("User with username: %v wants to stop matchmaking", c.username)
			break
		}

		words := strings.Fields(msgstr)
		if len(words) != 3 || words[0] != start {
			log.Println("Error parsing start string")
			break
		}

		l, err := strconv.Atoi(words[1])
		if err != nil {
			log.Println("Error parsing low MMR string")
			break
		}

		h, err := strconv.Atoi(words[2])
		if err != nil {
			log.Println("Error parsing high MMR string")
			break
		}

		q.data[userPos].low = l
		q.data[userPos].high = h
	}
}

// needs to send update every 5 seconds for alive status: "alive"
// finally writes roomId for generated room (randomized string - check if exists in rooms in games or not): "found <id>"
// stopped message once stopped
// return when unregistered
func (c *matchClient) healthPump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			break
		}

		w.Write([]byte("searching"))
		w.Close()
		time.Sleep(2 * time.Second)
	}
}

func (c *matchClient) matchmakingFinder(q *matchQueue) {
	log.Printf("Started matchmaking for username: %v", c.username)

	for {
		userPos := q.findQueueItem(c)
		if userPos == -1 {
			break
		} else if q.data[userPos].low == -1 || q.data[userPos].high == -1 {
			time.Sleep(2 * time.Second)
		} else {
			q.c.L.Lock()
			res := q.findMatch(&q.data[userPos])
			if res != nil {
				completeMatchmaking(res, q)
				q.c.Signal()
				q.c.L.Unlock()
				break
			} else {
				q.c.Signal()
				q.c.L.Unlock()
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func completeMatchmaking(res []*matchQueueItem, q *matchQueue) {
	for i := 0; i < len(res); i++ {
		c := res[i].client
		log.Printf("Complete matchmaking for %v", c.username)
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			break
		}

		w.Write([]byte(found + " <room>"))
		w.Close()
		c.hub.unregister <- c
		c.conn.Close()
	}

	// could even use channel for this
	q.removeMultipleClientsFromQ(res)
}
