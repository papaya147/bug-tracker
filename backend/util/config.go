package util

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_DSN         string        `mapstructure:"POSTGRES_DSN"`
	HTTP_SERVER_PORT     string        `mapstructure:"HTTP_SERVER_PORT"`
	SENDER_EMAIL         string        `mapstructure:"SENDER_EMAIL"`
	SENDER_PASSWORD      string        `mapstructure:"SENDER_PASSWORD"`
	EMAIL_HOST           string        `mapstructure:"EMAIL_HOST"`
	EMAIL_HOST_PORT      string        `mapstructure:"EMAIL_HOST_PORT"`
	EMAIL_DURATION       time.Duration `mapstructure:"EMAIL_DURATION"`
	API_VERSION          string        `mapstructure:"API_VERSION"`
	API_PREFIX           string        `mapstructure:"API_PREFIX"`
	SESSION_DURATION     time.Duration `mapstructure:"SESSION_DURATION"`
	CLIENT_PREFIX        string        `mapstructure:"CLIENT_PREFIX"`
	FORGOT_PASSWORD_PATH string        `mapstructure:"FORGOT_PASSWORD_PATH"`
	MIGRATION_URL        string
	COOKIE_STORE_SECRET  string
}

func NewConfig(path string) Config {
	godotenv.Load(path + ".env")

	emailDuration, _ := time.ParseDuration(os.Getenv("EMAIL_DURATION"))
	sessionDuration, _ := time.ParseDuration(os.Getenv("SESSION_DURATION"))

	return Config{
		POSTGRES_DSN:         os.Getenv("POSTGRES_DSN"),
		HTTP_SERVER_PORT:     os.Getenv("HTTP_SERVER_PORT"),
		SENDER_EMAIL:         os.Getenv("SENDER_EMAIL"),
		SENDER_PASSWORD:      os.Getenv("SENDER_PASSWORD"),
		EMAIL_HOST:           os.Getenv("EMAIL_HOST"),
		EMAIL_HOST_PORT:      os.Getenv("EMAIL_HOST_PORT"),
		EMAIL_DURATION:       emailDuration,
		API_VERSION:          os.Getenv("API_VERSION"),
		API_PREFIX:           os.Getenv("API_PREFIX"),
		SESSION_DURATION:     sessionDuration,
		CLIENT_PREFIX:        os.Getenv("CLIENT_PREFIX"),
		FORGOT_PASSWORD_PATH: os.Getenv("FORGOT_PASSWORD_PATH"),
		MIGRATION_URL:        os.Getenv("MIGRATION_URL"),
		COOKIE_STORE_SECRET:  os.Getenv("COOKIE_STORE_SECRET"),
	}
}
