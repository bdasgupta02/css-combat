package config

import (
	"user-service/db/models"

	"github.com/jackc/pgx/v4"
)

type Config struct {
	WebPort string
	Models  models.Models
	DB      *pgx.Conn
}

// FIXME: add proper model init with sql
func LoadConfig(db *pgx.Conn, mod models.Models) Config {
	return Config{
		":8020",
		mod,
		db,
	}
}
