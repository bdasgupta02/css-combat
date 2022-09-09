package models

type UserMMR struct {
	Id           *uint64
	EndUserId    *uint64
	PlayerRankId *uint64
	MMR          *float64
	Deviation    *float64
	Volatility   *float64
	NumGames     *uint
	GameHistory  *[]string
}
