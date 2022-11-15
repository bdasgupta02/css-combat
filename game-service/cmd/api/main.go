package main

import (
	"context"
	"game-service/config"
	"game-service/proto/problem"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serverConfig config.Config

var tokenAuth *jwtauth.JWTAuth
var conf serverConfig
var connCount uint32

var mHub *matchHub
var gHub *gameHub

type probClient struct {
	client problem.ProblemClient
	ctx    context.Context
}

var probConn *grpc.ClientConn
var probC probClient

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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	var err error = nil
	for err != nil {
		probConn, err = grpc.Dial("localhost:8040", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not open gRPC client from Game Service to Problem Service: %v", err)
			time.Sleep(2 * time.Second)
		}
	}

	probC = probClient{client: problem.NewProblemClient(probConn)}
	log.Println("gRPC connections successful")

	q := newMatchQueue()
	mHub = newMatchHub(q)
	gHub = newGameHub()
	go mHub.run(ctx)
	go gHub.run(ctx)

	router.Use(middleware.Logger)
	router.Use(c.Handler)

	router.HandleFunc("/ws/match", func(w http.ResponseWriter, r *http.Request) {
		serveMatchWS(mHub, w, r)
	})

	router.HandleFunc("/ws/game/{id}", func(w http.ResponseWriter, r *http.Request) {
		serveGameWS(gHub, w, r)
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
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
	dsn := "postgres://admin:password@host.docker.internal:5432/game_db"
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
