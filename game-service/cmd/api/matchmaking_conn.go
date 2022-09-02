package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var match_upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type match_conn struct {
	ws   *websocket.Conn
	send chan []byte
}

func serveMatchmaking(w http.ResponseWriter, r *http.Request) {
	ws, err := match_upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Cannot upgrade HTTP to WS: %v", err)
		return
	}

	c := &match_conn{send: make(chan []byte, 256), ws: ws}
	s := match_subscription{c}
	match_h.register <- s
	go s.writePump()
	go s.readPump()
}
