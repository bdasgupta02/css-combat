package main

import (
	"context"
	"user-service/db/models"
	"user-service/proto/auth"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	Models models.Models
}

func CreateJWT() {}

func (a *AuthServer) Register(ctx context.Context, req *auth.AuthRegister) (*auth.AuthToken, error) {
	// email := req.GetEmail()
	// username := req.GetUsername()
	// fullName := req.GetFullName()
	// password := req.GetUsername()

	return nil, nil
}

func (a *AuthServer) Login(ctx context.Context, req *auth.AuthLogin) (*auth.AuthToken, error) {

	return nil, nil
}