package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPort    string
	JWTSecret string
	User      string
	Name      string
	Host      string
	Password  string
	Port      string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, err
	}
	return &Config{
		DBPort:    os.Getenv("DB_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		User:      os.Getenv("DB_USER"),
		Name:      os.Getenv("DB_NAME"),
		Host:      os.Getenv("DB_HOST"),
		Password:  os.Getenv("DB_PASSWORD"),
		Port:      os.Getenv("PORT"),
	}, nil
}
