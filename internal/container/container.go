package container

import (
	"github.com/eduardocfalcao/money-tracker/internal/auth"
	"github.com/eduardocfalcao/money-tracker/internal/users"
	"github.com/eduardocfalcao/money-tracker/internal/users/repository"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	JWTMiddlewareService auth.JWTMiddlewareService
	UsersHanders         *users.Handlers
}

func NewContainer(secret string, conn *sqlx.DB) (*Container, error) {
	jwtService := auth.NewJWTService([]byte(secret))

	userRepository := repository.New(conn)
	usersService := users.NewService(userRepository)

	c := &Container{
		// Db: db,
		JWTMiddlewareService: jwtService,
		UsersHanders:         users.NewHandler(jwtService, usersService),
	}

	return c, nil
}
