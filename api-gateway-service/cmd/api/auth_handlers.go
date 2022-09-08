package main

import (
	"api-gateway-service/proto/auth"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (app *serverConfig) LoginViaGRPC(w http.ResponseWriter, r *http.Request) {
	var cancel context.CancelFunc
	authC.ctx, cancel = context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var loginObj auth.AuthLogin
	err := json.NewDecoder(r.Body).Decode(&loginObj)
	if err != nil {
		log.Printf("Cannot login via User Service: %v", err)
		app.errorJSON(w, err)
		return
	}

	authResponse, err := authC.client.Login(authC.ctx, &loginObj)
	if err != nil {
		log.Printf("Cannot login via User Service: %v", err)
		app.errorJSON(w, err)
		return
	}

	writeJSONResponse(w, authResponse)
}

func (app *serverConfig) RegisterViaGRPC(w http.ResponseWriter, r *http.Request) {

	var cancel context.CancelFunc
	authC.ctx, cancel = context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var registerObj auth.AuthRegister
	err := json.NewDecoder(r.Body).Decode(&registerObj)
	if err != nil {
		log.Printf("Cannot register via User Service: %v", err)
		app.errorJSON(w, err)
		return
	}

	authResponse, err := authC.client.Register(authC.ctx, &registerObj)
	if err != nil {
		log.Printf("Cannot register via User Service: %v", err)
		app.errorJSON(w, err)
		return
	}

	writeJSONResponse(w, authResponse)
}

