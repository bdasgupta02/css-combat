package controllers

import (
	"context"
	"game-service/db/models"
	"game-service/mmr"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

func GetOrCreateMMR(ctx context.Context, db *pgx.Conn, id *uint64) (*mmr.Rating, error) {
	var userMMR models.UserMMR
	pgxscan.Get(ctx, db, &userMMR, `SELECT * FROM user_mmr WHERE end_user_id = $1`, id)

	exists := userMMR.Id != nil && userMMR.MMR != nil && userMMR.Deviation != nil && userMMR.Volatility != nil
	if !exists {
		r := mmr.DefaultRating()
		userMMR = models.UserMMR{
			MMR:        &r.Rating,
			Deviation:  &r.Deviation,
			Volatility: &r.Volatility,
		}
		if _, err := db.Exec(
			ctx,
			`INSERT INTO user_mmr (end_user_id, mmr, deviation, volatility) VALUES ($1, $2, $3, $4)`,
			id,
			*userMMR.MMR,
			*userMMR.Deviation,
			*userMMR.Volatility,
		); err != nil {
			return nil, err
		}
	}

	return &mmr.Rating{
		Rating:     *userMMR.MMR,
		Deviation:  *userMMR.Deviation,
		Volatility: *userMMR.Volatility,
		TwoSigma:   *userMMR.Deviation,
	}, nil
}

// Important: IDs should be in order of rank
func UpdateGameMMR(id *[]uint64) {
	// use for loop to scan one buy one
	// just placeholder
	mmr.CalculateMatchMMRs([]*mmr.Rating{})
}
