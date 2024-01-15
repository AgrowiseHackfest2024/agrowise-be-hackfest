package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/joho/godotenv"
)

var SecretKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	SecretKey = os.Getenv("SECRET_KEY")
	if SecretKey == "" {
		panic("Failed to get Secret Key from .env")
	}
}

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}
