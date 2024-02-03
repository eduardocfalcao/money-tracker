package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	service struct {
		// store the secret in some parameter store/receive from env
		SecretKey []byte
	}

	CreateTokenArgs struct {
		Username string
		Email    string
	}

	JWTToken struct {
		Token     string `json:"token"`
		ExpiresIn int64  `json:"expires_in"` // unit timestamp
	}
)

type JWTService interface {
	CreateToken(args CreateTokenArgs) (string, error)
	VerifyToken(tokenString string) error
}

type JWTMiddlewareService interface {
	VerifyTokenMiddleware(next http.Handler) http.Handler
}

func NewJWTService(secret []byte) *service {
	return &service{
		SecretKey: secret,
	}
}

func (j *service) CreateToken(args CreateTokenArgs) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": args.Username,
			"email":    args.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *service) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil // combine the secret key with some
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
