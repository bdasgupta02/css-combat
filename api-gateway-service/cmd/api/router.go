package main

import (
	"api-gateway-service/proto/auth"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client auth.AuthServiceClient
	ctx    context.Context
}

var userConn *grpc.ClientConn
var authC authClient

// TODO TLS credentials
func CreateRouter() http.Handler {
	router := chi.NewRouter()

	log.Println("Connecting to User Service via gRPC")
	var err error
	userConn, err = grpc.Dial("localhost:8020", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not open gRPC client from API Gateway to User Service: %v", err)
	}

	log.Println("Starting Auth client")
	authC = authClient{client: auth.NewAuthServiceClient(userConn)}

	log.Println("gRPC connections successful")
	router.Use(middleware.Logger)

	router.Group(ProtectedRoutes)
	router.Group(PublicRoutes)

	return router
}

func PublicRoutes(router chi.Router) {
	router.Post("/auth/login", conf.LoginViaGRPC)
	router.Post("/auth/register", conf.RegisterViaGRPC)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service is online"))
	})
}

func ProtectedRoutes(router chi.Router) {
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator)
}
