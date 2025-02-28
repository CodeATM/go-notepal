package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresServer   string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	JwtSecret        string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PostgresServer:   os.Getenv("DB_HOST"),
		PostgresPort:     os.Getenv("DB_PORT"),
		PostgresUser:     os.Getenv("DB_USER"),
		PostgresPassword: os.Getenv("DB_PASSWORD"),
		PostgresDB:       os.Getenv("DB_NAME"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
	}
}
