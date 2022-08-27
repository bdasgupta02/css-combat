package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("RPFwhPSTpd2fp1PeGftiAM810EYTYn8AvZrVefiBzlMjp8LZxekK88xkcEivUEYJ3rD8C0UUS7Eq07sxko7Yf75hNl8QqwpQCFoSArJpRjEogdeDXCTwBi2JgRMz9Ufr")

type claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// TODO refresh token (to be done in API gateway layer)
func generateClaims(username string) *claims {
	expirationTime := time.Now().Add(60 * 24 * time.Hour)
	return &claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
}

func createJWT(c *claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jwtKey)
}
