package main

import (
	"context"
	"log"

	"user-service/controllers"
	"user-service/proto/auth"
	"user-service/proto/user"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
}

type UserServer struct {
	user.UnimplementedUserServiceServer
}

// Auth
func (a *AuthServer) Register(ctx context.Context, req *auth.AuthRegister) (*auth.AuthToken, error) {
	logGRPC("Called Register", "")
	return controllers.Register(ctx, conf.DB, req)
}

func (a *AuthServer) Login(ctx context.Context, req *auth.AuthLogin) (*auth.AuthToken, error) {
	logGRPC("Called Login", "")
	return controllers.Login(ctx, conf.DB, req)
}

// User
func (a *UserServer) GetUser(ctx context.Context, req *user.EmptyMessage) (*user.UserResponse, error) {
	logGRPC("Called GetUser", "")
	return controllers.GetUser(ctx, conf.DB, req)
}

func (a *UserServer) GetMatchUsers(ctx context.Context, req *user.MatchUsersRequest) (*user.MatchUsersResponse, error) {
	logGRPC("Called GetMatchUsers", "")
	return controllers.GetMatchUsers(ctx, conf.DB, req)
}

func (a *UserServer) EditUser(ctx context.Context, req *user.EditUserRequest) (*user.UserResponse, error) {
	logGRPC("Called EditUser", "")
	return nil, nil
}

func (a *UserServer) GetCosmetics(ctx context.Context, req *user.EmptyMessage) (*user.CosmeticsResponse, error) {
	logGRPC("Called GetCosmetics", "")
	return nil, nil
}

func (a *UserServer) GetCosmetic(ctx context.Context, req *user.IDRequest) (*user.CosmeticResponse, error) {
	logGRPC("Called GetCosmetic", "")
	return nil, nil
}

func (a *UserServer) GetCosmeticType(ctx context.Context, req *user.CosmeticTypeRequest) (*user.CosmeticsResponse, error) {
	logGRPC("Called GetCosmeticType", "")
	return nil, nil
}

func (a *UserServer) PurchaseCosmetic(ctx context.Context, req *user.IDRequest) (*user.IDRequest, error) {
	logGRPC("Called PurchaseCosmetic", "")
	return nil, nil
}

func (a *UserServer) GetInventory(ctx context.Context, req *user.EmptyMessage) (*user.InventoryItemsResponse, error) {
	logGRPC("Called GetInventory", "")
	return nil, nil
}

func (a *UserServer) GetInventoryItem(ctx context.Context, req *user.IDRequest) (*user.InventoryItemResponse, error) {
	logGRPC("Called GetInventoryItem", "")
	return nil, nil
}

func (a *UserServer) EquipInventoryItem(ctx context.Context, req *user.IDRequest) (*user.InventoryItemResponse, error) {
	logGRPC("Called EquipInventoryItem", "")
	return nil, nil
}

func logGRPC(msg string, addOn string) {
	log.Printf("gRPC: %v %v", msg, addOn)
}
