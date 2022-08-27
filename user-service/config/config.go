package config

import "user-service/db/models"

type Config struct {
	WebPort string
	Models  models.Models
}

//FIXME: add proper model init with sql
func LoadConfig() Config {
	return Config{
		":8020",
		models.Models{},
	}
}