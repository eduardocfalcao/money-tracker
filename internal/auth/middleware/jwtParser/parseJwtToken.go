package jwtParser

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/eduardocfalcao/money-tracker/internal/auth"
)

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func Parse(tokenString string) *jwt.Token {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := auth.GetPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})
	return token
}
