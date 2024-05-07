package helpers

import (
	"catssocial/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(userID string) (string, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return "", err
	}
	var secretKey = []byte(config.JWTSecret)

	// expirationTime := time.Now().Add(8 * time.Hour)
	expirationTime := time.Now().Add(2 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"expiration": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
