package controllers

import (
	"context"
	"errors"
	"user-service/proto/user"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/metadata"
)

// Controllers
func GetUser(ctx context.Context, db *pgx.Conn, req *user.EmptyMessage) (*user.UserResponse, error) {
	username, ok := getUsernameMetadata(ctx)
	if !ok {
		return nil, errors.New("cannot retrieve username")
	}

	// select without pw 
	rows, err := db.Query(ctx, `SELECT id, email, username, full_name, resume_link, portfolio_link, is_blocked, blocked_till, block_history, is_deactivated, preferences, currency FROM end_user WHERE username = $1`, username)
	if err != nil {
		return nil, err
	}
	
	var endUser user.UserResponse
	if err := pgxscan.ScanOne(&endUser, rows); err != nil {
		return nil, err
	}
	
	return &endUser, nil
}

// Metadata handling
func getUsernameMetadata(ctx context.Context) (string, bool) {
	jwt, ok := getJwtMetadata(ctx)
	if !ok {
		return "", false
	}

	claims, err := extractClaims(jwt)
	if err != nil {
		return "", false
	}

	return claims["username"].(string), true 
}

func getJwtMetadata(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	return md.Get("authorization")[0], ok
}

func extractClaims(tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(jwtKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			 return hmacSecret, nil
	})

	if err != nil {
			return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, nil
	} else {
			return nil, err
	}
}