package container

import (
	"github.com/eduardocfalcao/money-tracker/internal/auth"
	"github.com/eduardocfalcao/money-tracker/internal/users"
)

type Container struct {
	JWTMiddlewareService auth.JWTMiddlewareService
	UsersHanders         *users.Handlers
}

func NewContainer(secret string) (*Container, error) {
	// db := database.New()
	// _ := database.TestConnection() // treat error

	jwtService := auth.NewJWTService([]byte(secret))
	c := &Container{
		// Db: db,
		JWTMiddlewareService: jwtService,
		UsersHanders:         users.NewHandler(jwtService),
	}

	return c, nil
}
