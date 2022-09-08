package config

import "github.com/jackc/pgx/v4"

type Config struct {
	WebPort string
	JwtKey  []byte
	DB      *pgx.Conn
}

// TODO: DB connection
func LoadConfig(db *pgx.Conn) Config {
	return Config{
		":8030",
		[]byte("RPFwhPSTpd2fp1PeGftiAM810EYTYn8AvZrVefiBzlMjp8LZxekK88xkcEivUEYJ3rD8C0UUS7Eq07sxko7Yf75hNl8QqwpQCFoSArJpRjEogdeDXCTwBi2JgRMz9Ufr"),
		db,
	}
}
