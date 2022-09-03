package main

import "github.com/gorilla/websocket"

// between the websocket connection and the hub.
type matchClient struct {
	hub      *matchHub
	conn     *websocket.Conn
	username string
	send     chan []byte
	update   chan int
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

func (h *matchHub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.queue.addToQ(&matchQueueItem{client: client, low: -1, high: -1})
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
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
