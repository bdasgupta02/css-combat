package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
)

/*
TODO
- exit functions

Requirements
- create game function with locks
- timeout for game rooms if client doesn't connect in time
- reconnection (based on jwt user id)
- needs to maintain state (max acc per person - in gameClient, time, best img)
- chat

readPump
- no need to time anything (timing of base64 transmissions will be done via frontend)
- flow: gets base64, processes metrics via grpc, sends broadcast using hub

timePump
- calculates total time once game has started (how to make game start? - loop to check if all players connect or hub.run to calculate number of players on register[preferred but need expected players])

writePump
- use broadcast to give updates to all players using hub.run
*/

const (
	start      = "201"
	img        = "202"
	chat       = "203"
	wait       = "204"
	updateImg  = "205" // format: id1: accuracy1; id2: accuracy2; ..
	updateChat = "206"
	timeUpdate = "207"
	endTime    = "208"
	endWin     = "209"
)

// reconnections -> find in hubs -> replace connection
// don't delete connections on unregister -> disconnect on game finish
func serveGameWS(hub *gameHub, w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "id")
	if !checkGameExists(&room) {
		http.Error(w, "Game not found", http.StatusNotFound)
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := uint64(claims["userId"].(float64))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &gameClient{
		hub:       hub,
		conn:      conn,
		room:      room,
		userId:    userId,
		maxAcc:    0,
		bestImg:   nil,
		send:      make(chan string),
		quitWrite: make(chan bool),
	}

	client.hub.register <- client

	go client.timePump()
	go client.readPump()
	go client.writePump()
}

// 2 stages:
//  1. Waiting for players: 30s timeout then all are unregistered and all client funcs are closed, then TODO: send image
//  2. Counting ticker for 30m (listen to game ending), TODO: regular second by second update with [207 time]
func (c *gameClient) timePump() {
	defer c.conn.Close()

	// Stage 1: waiting for all players
	for {
		if len(c.hub.games[c.room]) == c.hub.expected[c.room] {
			c.hub.broadcast <- &gameMessage{start, c.room}
			break
		}

		time.Sleep(gameFreq)
	}

	ticker := time.NewTicker(gameFreqDouble)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				if err := c.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%v %02d:%02d", timeUpdate, t.Minute(), t.Second()))); err != nil {
					log.Printf("User with id %v had an issue in writePump: %v", c.userId, err)
					return
				}
			}
		}
	}()

	time.Sleep(gamePeriod)
	ticker.Stop()
	done <- true
}

// One reader for client:
//  1. Base64 image with 202
//  2. Chat message with 203
//
// Both broadcast reply to every member
func (c *gameClient) readPump() {
	defer c.conn.Close()

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
		code := msgstr[:3]

		switch code {
		case img:
			// TODO: analytics server calls here
			c.hub.broadcast <- &gameMessage{"img update here", c.room}
		case chat:
			c.hub.broadcast <- &gameMessage{msgstr[4:], c.room}
		}
	}
}

// Broadcasting: does not discriminate or check if chat or game update
func (c *gameClient) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Printf("User with id %v had an issue in writePump: %v", c.userId, err)
				return
			}
		case q := <-c.quitWrite:
			if q {
				return
			}
		}
	}
}
