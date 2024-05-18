package utils

import (
	"errors"
	// "fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lpernett/godotenv"
)

func GenerateToken(email, userId string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(tkn *jwt.Token) (any, error) {
		_, isValidSigningMethod := tkn.Method.(*jwt.SigningMethodHMAC)
		if !isValidSigningMethod {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return "", errors.New("could not parse token")
	}
	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	// fmt.Println("Claims", claims)
	// email := claims["email"].(string)
	userId := claims["userId"].(string)
	return userId, nil
}
