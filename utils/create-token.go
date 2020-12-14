package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken helper for creating jwt tokens
func CreateToken(userID uint) (string, error) {
	var err error

	os.Setenv("JWT_SECRET", "secret") // should be in a ENV file in a real application
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user-id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix() // 1h
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
