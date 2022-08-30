package main

import (
	"game-service/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type serverConfig config.Config

var tokenAuth *jwtauth.JWTAuth
var conf serverConfig

func init() {
	conf = serverConfig(config.LoadConfig())
	tokenAuth = jwtauth.New("HS256", conf.JwtKey, nil)
}

// TODO need to change this for CORS
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

// chat message and game event are the 2 types of events
// remove state files if not needed later (esp matchmaking)
// matchmaking room, join room with encrypted string
func main() {
	log.Println("Starting Game Service")
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/",  func(c socketio.Conn) error {
		return nil
	})

	// authenticates socket connection with JWT
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(jwtauth.Authenticator)
	
		router.Handle("/socket.io/", server)
	
		router.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You're authorized!"))
		})
	})
}
