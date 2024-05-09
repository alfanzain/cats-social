package helpers

import (
	"catssocial/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateAccessToken(userID string) (string, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return "", err
	}
	var secretKey = []byte(config.JWTSecret)

	// expiredAt := time.Now().Add(8 * time.Hour)
	expiredAt := time.Now().Add(2 * time.Minute)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
