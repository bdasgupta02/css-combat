package main

import (
	"api-gateway-service/proto/auth"
	"context"
	"net/http"
	"time"
)

type RequestPayload struct {
	Action       string          `json:"action"`
	AuthRegister RegisterPayload `json:"authRegister,omitempty"`
	AuthLogin    LoginPayload    `json:"authLogin,omitempty"`
}

type LoginPayload struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type RegisterPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (app *serverConfig) LoginViaGRPC(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	var cancel context.CancelFunc
	authC.ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorJSON(w, err)
		return
	}

	authResponse, err := authC.client.Login(authC.ctx, &auth.AuthLogin{
		Type:       requestPayload.AuthLogin.Type,
		Identifier: requestPayload.AuthLogin.Identifier,
		Password:   requestPayload.AuthLogin.Password,
	})

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	writeJSONResponse(w, authResponse)
}

func (app *serverConfig) RegisterViaGRPC(w http.ResponseWriter, r *http.Request) {

}
