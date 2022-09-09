package main

import (
	"context"
	"game-service/controllers"
	"log"

	"github.com/gorilla/websocket"
)

type matchClient struct {
	hub        *matchHub
	conn       *websocket.Conn
	username   string
	userId     uint64
	send       chan []byte
	update     chan int
	quitSearch chan bool
}

type matchHub struct {
	clients    map[*matchClient]bool
	queue      *matchQueue
	broadcast  chan []byte
	register   chan *matchClient
	unregister chan *matchClient
}

func newMatchHub(q *matchQueue) *matchHub {
	return &matchHub{
		broadcast:  make(chan []byte),
		register:   make(chan *matchClient),
		unregister: make(chan *matchClient),
		clients:    make(map[*matchClient]bool),
		queue:      q,
	}
}

func (h *matchHub) run(ctx context.Context) {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			r, err := controllers.GetOrCreateMMR(ctx, conf.DB, &client.userId)
			if err != nil {
				log.Printf("Error connecting user with id %v: %v", client.userId, err)
				h.unregister <- client
			} else {
				log.Printf("User with id %v started matchmaking with low: %v, high: %v", client.userId, r.Rating-r.TwoSigma, r.Rating+r.TwoSigma)
				h.queue.addToQ(&matchQueueItem{
					client: client,
					low:    r.Rating - r.TwoSigma,
					high:   r.Rating + r.TwoSigma,
				})
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				client.quitSearch <- true
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
