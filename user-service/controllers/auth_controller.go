package controllers

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"user-service/db/models"
	"user-service/proto/auth"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, db *pgx.Conn, req *auth.AuthRegister) (*auth.AuthToken, error) {
	salt := randomizedString(6)
	hashed, err := hashPassword(req.GetPassword() + salt)
	if err != nil {
		return nil, err
	}

	stmt := `INSERT INTO end_user (email, username, pass_hash, pass_salt, full_name) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(ctx, stmt, req.GetEmail(), req.GetUsername(), hashed, salt, req.GetFullName())
	if err != nil {
		return nil, err
	}

	jwt, err := createJWT(generateClaims(req.GetUsername()))
	if err != nil {
		return nil, err
	}

	return &auth.AuthToken{Token: jwt}, nil
}

func Login(ctx context.Context, db *pgx.Conn, req *auth.AuthLogin) (*auth.AuthToken, error) {
	var rows pgx.Rows
	var err error

	if req.GetType() == "username" {
		rows, err = db.Query(ctx, `SELECT username, email, pass_hash, pass_salt FROM end_user WHERE username = $1`, req.GetIdentifier())
		if err != nil {
			return nil, err
		}
	} else if req.GetType() == "email" {
		rows, err = db.Query(ctx, `SELECT username, email, pass_hash, pass_salt FROM end_user WHERE email = $1`, req.GetIdentifier())
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid type")
	}

	var user models.EndUser
	if err := pgxscan.ScanOne(&user, rows); err != nil {
		return nil, err
	}

	hashCheck := checkPasswordHash(req.GetPassword()+*user.PassSalt, *user.PassHash)
	if !hashCheck {
		return nil, errors.New("invalid password")
	}

	jwt, err := createJWT(generateClaims(*user.Username))
	if err != nil {
		return nil, err
	}

	return &auth.AuthToken{Token: jwt}, nil
}

// Password hashing (a bit slower because of bcrypt)
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Salt generation
func randomizedString(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
