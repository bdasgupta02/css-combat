package main

import (
	"context"
	"log"
	"user-service/controllers"
	"user-service/db/models"
	"user-service/proto/auth"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	Models models.Models
}

func (a *AuthServer) Register(ctx context.Context, req *auth.AuthRegister) (*auth.AuthToken, error) {
	logGRPC("Called Register", "")
	return controllers.Register(ctx, conf.DB, req)
}

func (a *AuthServer) Login(ctx context.Context, req *auth.AuthLogin) (*auth.AuthToken, error) {
	logGRPC("Called Login", "")
	return controllers.Login(ctx, conf.DB, req)
}

func logGRPC(msg string, addOn string) {
	log.Printf("gRPC: %v %v", msg, addOn)
}
