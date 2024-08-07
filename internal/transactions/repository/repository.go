package repository

import (
	"context"

	"github.com/eduardocfalcao/money-tracker/internal/transactions/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InsertRawTransaction(ctx context.Context, transaction models.RawTransaction) error
	GetTransactionsByUser(ctx context.Context, params GetTransactionsByUserParams) ([]*models.RawTransaction, error)
}

type repository struct {
	db *sqlx.DB
}

type PageParams struct {
	Limit  int
	Offset int
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (q *repository) InsertRawTransaction(ctx context.Context, transaction models.RawTransaction) error {
	_, err := q.db.ExecContext(ctx, insertRawTransaction,
		transaction.UserID,
		transaction.AccountID,
		transaction.DatePosted,
		transaction.TransactionAmount,
		transaction.FitID,
		transaction.Checknum,
		transaction.Memo,
		transaction.Description,
	)
	return err
}

func (r *repository) GetTransactionById(ctx context.Context, id int32) (models.RawTransaction, error) {
	var t models.RawTransaction
	err := r.db.Get(&t, getTransactionById, id)
	return t, err
}

type GetTransactionsByUserParams struct {
	PageParams
	UserID int32
}

func (r *repository) GetTransactionsByUser(ctx context.Context, params GetTransactionsByUserParams) ([]*models.RawTransaction, error) {
	var list []*models.RawTransaction
	err := r.db.SelectContext(ctx, &list, getTransactionsByUser,
		params.UserID,
		params.Limit,
		params.Offset,
	)

	return list, err
}
