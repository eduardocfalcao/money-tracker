package container

import (
	"github.com/eduardocfalcao/money-tracker/database/queries"
	"github.com/eduardocfalcao/money-tracker/internal/auth"
	"github.com/eduardocfalcao/money-tracker/internal/users"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	JWTMiddlewareService auth.JWTMiddlewareService
	UsersHanders         *users.Handlers
}

func NewContainer(secret string, conn *pgx.Conn) (*Container, error) {
	jwtService := auth.NewJWTService([]byte(secret))
	queries := queries.New(conn)
	usersService := users.NewService(queries)

	c := &Container{
		// Db: db,
		JWTMiddlewareService: jwtService,
		UsersHanders:         users.NewHandler(jwtService, usersService),
	}

	return c, nil
}
