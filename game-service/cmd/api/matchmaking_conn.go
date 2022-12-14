package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
)

const (
	timedOut  = "104"
	stop      = "103"
	found     = "102 "
	searching = "101"
	search    = "100"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO change check origin condition
func serveMatchWS(hub *matchHub, w http.ResponseWriter, r *http.Request) {
	// TODO change condition
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var userId uint64
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			log.Printf("Error at matchmaking connection: %v", err)
			return
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		msgstr := string(message)
		msg_parts := strings.Split(msgstr, " ")
		if msg_parts[0] == search {
			token, err := jwtauth.VerifyToken(tokenAuth, msg_parts[1])
			if err != nil {
				conn.Close()
				log.Printf("Error at matchmaking connection: %v", err)
				return
			}

			userIdstr, ok := token.Get("userId")
			if !ok {
				conn.Close()
				log.Printf("Error at matchmaking connection: %v", err)
				return
			}

			userId = uint64(userIdstr.(float64))

			break
		}
	}

	client := &matchClient{
		hub:        hub,
		userId:     userId,
		conn:       conn,
		send:       make(chan []byte, 256),
		quitSearch: make(chan bool),
	}

	client.hub.register <- client

	go client.healthPump()
	go client.readPump(hub.queue)
	go client.matchmakingFinder(hub.queue)
	go client.timeoutChecker()
}

func (c *matchClient) timeoutChecker() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	time.Sleep(4 * time.Minute)

	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err == nil {
		w.Write([]byte(timedOut))
		w.Close()
	}
}

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
				log.Printf("Read error for matchmaking: %v", err)
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
			log.Printf("User with id: %v wants to stop matchmaking", c.userId)
			break
		}
	}
}

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

		w.Write([]byte(searching))
		w.Close()
		time.Sleep(2 * time.Second)
	}
}

func (c *matchClient) matchmakingFinder(q *matchQueue) {
	log.Printf("Started matchmaking for id: %v", c.userId)

	for {
		select {
		case q := <-c.quitSearch:
			if q {
				log.Printf("Ended matchmaking for id: %v", c.userId)
				return
			}
		default:
			userPos := q.findQueueItem(c)
			if userPos == -1 {
				time.Sleep(2 * time.Second)
			} else {
				q.c.L.Lock()
				res := q.findMatch(&q.data[userPos])
				if res != nil {
					l := len(res)
					completeMatchmaking(&l, res, q)
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
}

func completeMatchmaking(num *int, res []*matchQueueItem, q *matchQueue) {
	g := createGame(num)

	for i := 0; i < len(res); i++ {
		c := res[i].client
		log.Printf("Complete matchmaking for User with id %v", c.userId)
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			break
		}

		w.Write([]byte(found + g))
		w.Close()
		c.hub.unregister <- c
		c.conn.Close()
	}

	// could even use channel for this
	q.removeMultipleClientsFromQ(res)
}
