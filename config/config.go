package config

import (
	"os"
)

type Conf struct {
	XAPIKey string
	Port    string
	BaseURL string

	DB DBConfig
}

func NewConf() Conf {
	return Conf{
		XAPIKey: os.Getenv("X_API_KEY"),
		Port:    os.Getenv("PORT"),
		BaseURL: os.Getenv("BASE_URL"),
		DB: DBConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			DBHost:   os.Getenv("DB_HOST"),
			DBPort:   os.Getenv("DB_PORT"),
		},
	}
}
