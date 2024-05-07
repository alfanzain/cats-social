package helpers

import (
	"catssocial/configs"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		configs.Config{}.BcryptSalt,
	)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
