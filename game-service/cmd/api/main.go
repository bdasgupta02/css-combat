package main

import (
	"context"
	"game-service/config"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v4"
)

type serverConfig config.Config

var tokenAuth *jwtauth.JWTAuth
var conf serverConfig
var connCount uint32

func init() {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to pSQL DB")
	}

	conf = serverConfig(config.LoadConfig(conn))
	tokenAuth = jwtauth.New("HS256", conf.JwtKey, nil)
}

// chat message and game event are the 2 types of events
// remove state files if not needed later (esp matchmaking)
// matchmaking room, join room with encrypted string
func main() {
	log.Println("Starting Game Service")
	ctx := context.Background()
	router := chi.NewRouter()

	q := newMatchQueue()
	m := newMatchHub(q)
	go m.run(ctx)

	router.Use(middleware.Logger)

	// authenticates socket connection with JWT
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(jwtauth.Authenticator)

		router.HandleFunc("/ws/match", func(w http.ResponseWriter, r *http.Request) {
			serveMatchWS(m, w, r)
		})

		router.HandleFunc("/ws/game/:id", func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})

		router.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You're authorized!"))
		})
	})

	log.Fatal(http.ListenAndServe(conf.WebPort, router))
}

func openDB(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

// TODO shift to env
func connectToDB() *pgx.Conn {
	dsn := "postgres://admin:password@localhost:5432/game_db"
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Failed connection attempt to pSQL: %v", err)
			connCount++
		} else {
			log.Println("Successfully connected to pSQL database")
			return connection
		}

		if connCount > 20 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds..")
		time.Sleep(2 * time.Second)
		continue
	}
}
