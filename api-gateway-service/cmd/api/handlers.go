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
	authC.ctx, cancel = context.WithTimeout(context.Background(), time.Second)
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
	authC.ctx, cancel = context.WithTimeout(context.Background(), time.Second)
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

func writeJSONResponse(w http.ResponseWriter, jsonData any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (app *serverConfig) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	writeJSONResponse(w, payload)
}


type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
