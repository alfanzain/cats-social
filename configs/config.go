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
		DBName:     "",
		DBHost:     "",
		DBPort:     "",
		DBUsername: "",
		DBPassword: "",

		APPPort: "3000",
		ENV:     "development",

		JWTSecret: "secret",
	}
	salt, err := strconv.Atoi("8")
	if err != nil {
		return Config{}, fmt.Errorf("failed get bcrypt salt %v", err)
	}

	if os.Getenv("APP_PORT") == "" {
		config.APPPort = "3000"
	}

	config.BcryptSalt = salt

	return config, nil
}
