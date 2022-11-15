package main

import (
	"api-gateway-service/proto/auth"
	"api-gateway-service/proto/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client auth.AuthServiceClient
	ctx    context.Context
}

type userClient struct {
	client user.UserServiceClient
	ctx    context.Context
}

var userConn *grpc.ClientConn
var authC authClient
var userC userClient

// TODO TLS credentials
func CreateRouter() http.Handler {
	router := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	log.Println("Connecting to User Service via gRPC")
	var err error = nil
	for err != nil {
		userConn, err = grpc.Dial("localhost:8020", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not open gRPC client from API Gateway to User Service: %v, retrying!", err)
			time.Sleep(2 * time.Second)
		}
	}

	authC = authClient{client: auth.NewAuthServiceClient(userConn)}
	userC = userClient{client: user.NewUserServiceClient(userConn)}

	log.Println("gRPC connections successful")
	router.Use(middleware.Logger)
	router.Use(c.Handler)

	router.Group(ProtectedRoutes)
	router.Group(PublicRoutes)

	return router
}

func PublicRoutes(router chi.Router) {
	router.Post("/auth/sign-in", conf.LoginViaGRPC)
	router.Post("/auth/sign-up", conf.RegisterViaGRPC)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Gateway is online"))
	})
}

func ProtectedRoutes(router chi.Router) {
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator)

	router.Get("/user/get", conf.GetUser)

	router.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You're authorized!"))
	})
}
