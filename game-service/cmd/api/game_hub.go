package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type gameMessage struct {
	data string
	room string
}

type gameClient struct {
	hub     *gameHub
	conn    *websocket.Conn
	room    string
	userId  uint64
	maxAcc  float64
	bestImg *string

	// map of ids with game update data
	send chan string

	// quit channels
	quitWrite chan bool
}

// how to create games? create nil
type gameHub struct {
	expected   map[string]int
	games      map[string]map[*gameClient]bool
	broadcast  chan *gameMessage
	register   chan *gameClient
	unregister chan *gameClient
}

func newGameHub() *gameHub {
	return &gameHub{
		expected:   make(map[string]int),
		games:      make(map[string]map[*gameClient]bool),
		broadcast:  make(chan *gameMessage),
		register:   make(chan *gameClient),
		unregister: make(chan *gameClient),
	}
}

// TODO: Generate image by calling server here
// (and store image for the game using the game id string - either locally [best] or url string or in struct via base64 [ram risk])
func createGame(num *int) string {
	for {
		s := randomizedString()
		if !checkGameExists(&s) {
			gHub.games[s] = nil
			gHub.expected[s] = *num
			return s
		}
	}
}

func checkGameExists(g *string) bool {
	_, ok := gHub.games[*g]
	return ok
}

// for h.register check if room exists, or else dont connect and close connection for client
func (h *gameHub) run(ctx context.Context) {
	for {
		select {
		case client := <-h.register:
			clients := h.games[client.room]
			if clients == nil {
				clients = make(map[*gameClient]bool)
				h.games[client.room] = clients
			}
			h.games[client.room][client] = true
		// only called from game ending, automatically removes game
		case client := <-h.unregister:
			clients := h.games[client.room]
			if clients != nil {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					if len(clients) == 0 {
						h.remove(&client.room)
					}
				}
			}
		case message := <-h.broadcast:
			clients := h.games[message.room]
			for client := range clients {
				select {
				case client.send <- message.data:
				default:
					close(client.send)
					delete(clients, client)
					if len(clients) == 0 {
						h.remove(&client.room)
					}
				}
			}
		}
	}
}

func (h *gameHub) remove(room *string) {
	delete(h.games, *room)
	delete(h.expected, *room)
}

// RNG string
func randomizedString() string {
	return stringWithCharset(20, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
