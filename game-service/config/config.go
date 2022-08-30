package config

type Config struct {
	WebPort string
	JwtKey []byte
}

func LoadConfig() Config {
	return Config{
		":8010",
		[]byte("RPFwhPSTpd2fp1PeGftiAM810EYTYn8AvZrVefiBzlMjp8LZxekK88xkcEivUEYJ3rD8C0UUS7Eq07sxko7Yf75hNl8QqwpQCFoSArJpRjEogdeDXCTwBi2JgRMz9Ufr"),
	}
}
