package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		App: App{
			Host:    os.Getenv("APP_HOST"),
			Port:    os.Getenv("APP_PORT"),
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
		},
	}
}
