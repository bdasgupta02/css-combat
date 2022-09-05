package helpers

import (
	"game-service/mmr"
)

// Important: players array must be sorted in order of game rank
// Assumes there cannot be ties (if people have same accuracy, whoever came earlier must win)
func CalculateMatchMMRs(ratings []*mmr.Rating) ([]*mmr.Rating, error) {
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
