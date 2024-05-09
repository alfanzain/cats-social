package helpers

import (
	"catssocial/configs"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	config, _ := configs.LoadConfig()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptSalt)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
