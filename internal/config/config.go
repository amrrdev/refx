package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl     string
	RedisConnection string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		DatabaseUrl:     os.Getenv("DATABASE_URL"),
		RedisConnection: os.Getenv("REDIS_CONNECTION"),
	}, nil

}
