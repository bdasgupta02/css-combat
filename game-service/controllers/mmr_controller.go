package helpers

import (
	"context"
	"game-service/mmr"

	"github.com/jackc/pgx/v4"
)

// returns (low, high) MMR bracket based on username
func GetOrCreateMMR(ctx context.Context, db *pgx.Conn, id *uint64) *mmr.Rating {
  // if doesn't exist -> create
	// default rating

  // 1. Check if MMR exists

  // 2. Create user_mmr entry if it doesn't

  // 3. Return rating
  return &mmr.Rating{}
}

// Important: IDs should be in order of rank
func UpdateGameMMR(id *[]uint64) {
  // use for loop to scan one buy one
  // just placeholder
  calculateMatchMMRs([]*mmr.Rating{})
}

// Important: players array must be sorted in order of game rank
// Assumes there cannot be ties (if people have same accuracy, whoever came earlier must win)
func calculateMatchMMRs(ratings []*mmr.Rating) ([]*mmr.Rating, error) {
  var err error
  for i, p := range(ratings) {
    res := genRatings(len(ratings), i)
    o := append(ratings[:i], ratings[i+1:]...)
    ratings[i], err = mmr.CalculateRating(p, o, res)
    if err != nil {
      return nil, err
    }
  }
  return ratings, nil
}

func genRatings(num int, pos int) []mmr.Result {
  res := make([]mmr.Result, num - 1)
  for i := range(res) {
    if i >= pos {
      res[i] = mmr.Win
    } else {
      res[i] = mmr.Loss
    }
  }
  return res
}
