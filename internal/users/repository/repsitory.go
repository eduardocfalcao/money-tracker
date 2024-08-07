package repository

import (
	"context"

	"github.com/eduardocfalcao/money-tracker/internal/users/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(ctx context.Context, id int32) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	SearchUsers(ctx context.Context, name string) ([]models.User, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db,
	}
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, email, password_hash, salt)
VALUES ($1, $2, $3, $4)
`

func (q *repository) CreateUser(ctx context.Context, user models.User) error {
	_, err := q.db.ExecContext(ctx, createUser,
		user.Name,
		user.Email,
		user.Passwordhash,
		user.Salt,
	)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, name, email, password_hash, salt FROM users
WHERE id = $1 LIMIT 1
`

func (q *repository) GetUser(ctx context.Context, id int32) (models.User, error) {
	var i models.User
	err := q.db.GetContext(ctx, &i, getUser, id)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, name, email, password_hash, salt FROM users
WHERE email = $1 LIMIT 1
`

func (q *repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var i models.User
	err := q.db.GetContext(ctx, &i, getUserByEmail, email)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, created_at, name, email, password_hash, salt FROM users
`

func (q *repository) ListUsers(ctx context.Context) ([]models.User, error) {
	r := []models.User{}
	err := q.db.SelectContext(ctx, &r, listUsers)
	return r, err
}

const searchUsers = `-- name: SearchUsers :many
SELECT id, created_at, name, email, password_hash, salt FROM users u
where CASE WHEN LENGTH($1::text) != 0 THEN t.name LIKE '%'+@name::text +'%' ELSE TRUE END
`

func (q *repository) SearchUsers(ctx context.Context, name string) ([]models.User, error) {
	r := []models.User{}
	err := q.db.SelectContext(ctx, &r, searchUsers, name)
	return r, err
}
