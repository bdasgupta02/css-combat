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

// should only return username, fullname, avatar 
// min 2 max 4
func GetMatchUsers(ctx context.Context, db *pgx.Conn, req *user.MatchUsersRequest) (*user.MultipleUserResponse, error) {

	// if 2 if 3 if 4
	// join cos item and end user
	var rows pgx.Rows
	var err error
	if req.GetNum() == 2 {
		rows, err = db.Query(ctx, `SELECT username, full_name, avatar_img FROM end_users WHERE id = $1 OR id = $2`, req.GetId_1(), req.GetId_2())
	} else if req.GetNum() == 3 {
		rows, err = db.Query(ctx, `SELECT username, full_name, avatar_img FROM end_users WHERE id = $1 OR id = $2 OR id = $3`, req.GetId_1(), req.GetId_2(), req.GetId_3())
	} else if req.GetNum() == 4 {
		rows, err = db.Query(ctx, `SELECT username, full_name, avatar_img FROM end_users WHERE id = $1 OR id = $2 OR id = $3`, req.GetId_1(), req.GetId_2(), req.GetId_3())
	} else {
		return nil, errors.New("number of users should be from 2 to 4")
	}

	if rows == nil {}
	// get matcher values, if not in body or null, return null

	if err != nil {
		return nil, err
	}
	return nil, nil
}