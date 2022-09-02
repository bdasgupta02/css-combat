package main

import (
	"game-service/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

type serverConfig config.Config

var tokenAuth *jwtauth.JWTAuth
var conf serverConfig

func init() {
	conf = serverConfig(config.LoadConfig())
	tokenAuth = jwtauth.New("HS256", conf.JwtKey, nil)
}

// chat message and game event are the 2 types of events
// remove state files if not needed later (esp matchmaking)
// matchmaking room, join room with encrypted string
func main() {
	log.Println("Starting Game Service")
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// authenticates socket connection with JWT
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(jwtauth.Authenticator)
	
	
		router.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You're authorized!"))
		})
	})

	log.Fatal(http.ListenAndServe(conf.WebPort, router))
}
