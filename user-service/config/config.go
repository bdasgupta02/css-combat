package config

import (
	"github.com/jackc/pgx/v4"
)

type Config struct {
	WebPort string
	DB      *pgx.Conn
}

// FIXME: add proper model init with sql
func LoadConfig(db *pgx.Conn) Config {
	return Config{
		":8020",
		db,
	}
}
