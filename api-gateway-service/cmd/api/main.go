package main

import (
	"api-gateway-service/config"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type serverConfig config.Config

var tokenAuth *jwtauth.JWTAuth
var conf serverConfig

func init() {
	conf = serverConfig(config.LoadConfig())
	tokenAuth = jwtauth.New("HS256", conf.JwtKey, nil)
}

func main() {
	log.Println("Starting API Gateway Service")
	router := CreateRouter()
	defer userConn.Close()

	log.Fatal(http.ListenAndServe(conf.WebPort, router))
}
