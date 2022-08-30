package controllers

import (
	"context"
	"reflect"
	"user-service/db/models"
	"user-service/proto/user"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

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

// Converters
func userToResponse(u models.EndUser) *user.UserResponse {
	var r = user.UserResponse{}
	vu := reflect.ValueOf(u)
	vr := reflect.ValueOf(&r).Elem()

	addIfExists(&vu, &vr)

	return &r
}

func addIfExists(origin *reflect.Value, resp *reflect.Value) {
	type_origin := origin.Type()

	for i := 0; i < origin.NumField(); i++ {
		if !origin.Field(i).IsNil() {
			resp.FieldByName(type_origin.Field(i).Name).Set(reflect.ValueOf(origin.Field(i).Elem().Interface()))
		}
	}
}
