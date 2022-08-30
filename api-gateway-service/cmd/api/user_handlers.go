package main

import (
	"api-gateway-service/proto/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc/metadata"
)

func (app *serverConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jwtToken := jwtauth.TokenFromHeader(r)
	
	//header := metadata.New(map[string]string{"authorization": jwtToken})
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", jwtToken)

	var cancel context.CancelFunc
	userC.ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := userC.client.GetUser(userC.ctx, &user.EmptyMessage{})
	if err != nil {
		log.Printf("Cannot get user via User Service: %v", err)
		app.errorJSON(w, err)
		return
	}

	writeJSONResponse(w, resp)
}
