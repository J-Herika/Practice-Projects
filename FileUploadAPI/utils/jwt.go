package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string) string {
	claim := jwt.MapClaims{
		"Username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	strToken, err := token.SignedString([]byte("verySecure"))
	if err != nil {
		log.Printf("COULD NOT STRINGIFY TOKEN")
		return ""
	}

	return strToken
}
