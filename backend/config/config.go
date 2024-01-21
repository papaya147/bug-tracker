package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_DSN     string `mapstructure:"POSTGRES_DSN"`
	HTTP_SERVER_PORT int    `mapstructure:"HTTP_SERVER_PORT"`
}

func NewConfig() *Config {
	godotenv.Load("./.env")
	godotenv.Load("../.env")
	godotenv.Load("../../.env")

	port, _ := strconv.Atoi(os.Getenv("HTTP_SERVER_PORT"))

	return &Config{
		POSTGRES_DSN:     os.Getenv("POSTGRES_DSN"),
		HTTP_SERVER_PORT: port,
	}
}
