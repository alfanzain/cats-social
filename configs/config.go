package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
	Logger     *logrus.Logger
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		APPPort: os.Getenv("APP_PORT"),
		ENV:     os.Getenv("APP_ENV"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		Logger:    LoggerConfig(),
	}
	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		return Config{}, fmt.Errorf("failed get bcrypt salt %v", err)
	}

	if os.Getenv("APP_PORT") == "" {
		config.APPPort = "8080"
	}

	config.BcryptSalt = salt

	return config, nil
}
