package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	PostgresDSN string
	JwtSecret   string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:        os.Getenv("Port"),
		PostgresDSN: os.Getenv("PostgresDSN"),
		JwtSecret:   os.Getenv("JwtSecret"),
	}
}
