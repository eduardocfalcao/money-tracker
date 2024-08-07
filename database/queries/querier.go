// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetRawTransaction(ctx context.Context, id int32) (RawTransaction, error)
	GetTransactionsByPeriod(ctx context.Context, arg GetTransactionsByPeriodParams) ([]RawTransaction, error)
	GetTransactionsByUser(ctx context.Context, arg GetTransactionsByUserParams) ([]RawTransaction, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	InsertRawTransaction(ctx context.Context, arg InsertRawTransactionParams) error
	ListUsers(ctx context.Context) ([]User, error)
	SearchUsers(ctx context.Context, name string) ([]User, error)
}

var _ Querier = (*Queries)(nil)