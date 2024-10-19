package helper

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTKey(username string, email string) (string, error) {
	secretKey := os.Getenv("TOKEN_AUTHORIZE_KEY")
	log.Println(secretKey)

	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"roles":    "SYSTEM_ADMIN",                        // <-- Hard code for sample
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time 1d
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("TOKEN_AUTHORIZE_KEY")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
