package main

import (
	"log"
	"net/http"
	"api-gateway-service/config"

	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth
var serverConfig config.Config

func init() {
	serverConfig = config.LoadConfig()
	TokenAuth = jwtauth.New("HS256", serverConfig.JwtKey, nil)
}

func main() {
	log.Println("Starting API Gateway Service")
	router := CreateRouter()

	log.Fatal(http.ListenAndServe(serverConfig.WebPort, router))
}