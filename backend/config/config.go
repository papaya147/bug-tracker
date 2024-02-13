package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DSN         string        `mapstructure:"POSTGRES_DSN"`
	HTTP_SERVER_PORT     int           `mapstructure:"HTTP_SERVER_PORT"`
	SENDER_EMAIL         string        `mapstructure:"SENDER_EMAIL"`
	SENDER_PASSWORD      string        `mapstructure:"SENDER_PASSWORD"`
	EMAIL_HOST           string        `mapstructure:"EMAIL_HOST"`
	EMAIL_HOST_PORT      int           `mapstructure:"EMAIL_HOST_PORT"`
	EMAIL_DURATION       time.Duration `mapstructure:"EMAIL_DURATION"`
	API_VERSION          int           `mapstructure:"API_VERSION"`
	API_PREFIX           string        `mapstructure:"API_PREFIX"`
	SESSION_DURATION     time.Duration `mapstructure:"SESSION_DURATION"`
	CLIENT_PREFIX        string        `mapstructure:"CLIENT_PREFIX"`
	FORGOT_PASSWORD_PATH string        `mapstructure:"FORGOT_PASSWORD_PATH"`
}

func NewConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("error loading env vars:", err)
		return nil
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Panic("error unmarshalling config:", err)
		return nil
	}

	return &config
}
