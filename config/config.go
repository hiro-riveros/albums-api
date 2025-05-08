package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadDBConfig() (DatabaseConf, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return loadDbConfigFromEnv()
	}

	return loadDbConfigFromEnv()
}

func loadDbConfigFromEnv() (DatabaseConf, error) {
	return DatabaseConf{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}, nil
}
