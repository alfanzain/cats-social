package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBName     string
	DBPort     string
	DBHost     string
	DBUsername string
	DBPassword string

	APPPort string
	ENV     string

	JWTSecret  string
	BcryptSalt int
}

func LoadConfig() (Config, error) {
	config := Config{
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		APPPort: os.Getenv("APP_PORT"),
		ENV:     os.Getenv("APP_ENV"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}
	salt, err := strconv.Atoi("8")
	if err != nil {
		return Config{}, fmt.Errorf("failed get bcrypt salt %v", err)
	}

	if os.Getenv("APP_PORT") == "" {
		config.APPPort = "8080"
	}

	config.BcryptSalt = salt

	return config, nil
}
