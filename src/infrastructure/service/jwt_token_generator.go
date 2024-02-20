package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenGenerator struct {
	SecretKey string
}

func NewJWTTokenGenerator(secretKey string) *JWTTokenGenerator {
	return &JWTTokenGenerator{
		SecretKey: secretKey,
	}
}

func (j *JWTTokenGenerator) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
