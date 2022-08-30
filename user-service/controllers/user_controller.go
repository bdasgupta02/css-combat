package controllers

import (
	"context"
	"errors"
	"user-service/db/models"
	"user-service/proto/user"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
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

	var endUser models.EndUser
	if err := pgxscan.ScanOne(&endUser, rows); err != nil {
		return nil, err
	}

	return userToResponse(endUser), nil
}
