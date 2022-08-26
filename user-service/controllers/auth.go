package controllers

import (
	"user-service/db/models"

	"github.com/go-chi/jwtauth"
)

var secret = []byte("RPFwhPSTpd2fp1PeGftiAM810EYTYn8AvZrVefiBzlMjp8LZxekK88xkcEivUEYJ3rD8C0UUS7Eq07sxko7Yf75hNl8QqwpQCFoSArJpRjEogdeDXCTwBi2JgRMz9Ufr")

func createJWT() *jwtauth.JWTAuth {
	return nil
}

func verifyJWT() bool {
	return false
}

// If email null or not
func Login(user *models.EndUser) error {
	return nil
}

func Register(user *models.EndUser) error {
	return nil
}
