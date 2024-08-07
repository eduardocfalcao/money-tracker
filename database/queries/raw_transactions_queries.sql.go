// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: raw_transactions_queries.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getRawTransaction = `-- name: GetRawTransaction :one
SELECT id, user_id, account_id, date_posted, transaction_amount, fit_id, checknum, memo, description FROM raw_transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRawTransaction(ctx context.Context, id int32) (RawTransaction, error) {
	row := q.db.QueryRow(ctx, getRawTransaction, id)
	var i RawTransaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccountID,
		&i.DatePosted,
		&i.TransactionAmount,
		&i.FitID,
		&i.Checknum,
		&i.Memo,
		&i.Description,
	)
	return i, err
}

const getTransactionsByPeriod = `-- name: GetTransactionsByPeriod :many
SELECT id, user_id, account_id, date_posted, transaction_amount, fit_id, checknum, memo, description FROM raw_transactions
WHERE user_id = $1
 AND date_posted BETWEEN $2 and $3
order by date_posted DESC
LIMIT $4
OFFSET $5
`

type GetTransactionsByPeriodParams struct {
	UserID       int32            `db:"user_id" json:"userId"`
	DatePosted   pgtype.Timestamp `db:"date_posted" json:"datePosted"`
	DatePosted_2 pgtype.Timestamp `db:"date_posted_2" json:"datePosted2"`
	Limit        int32            `db:"limit" json:"limit"`
	Offset       int32            `db:"offset" json:"offset"`
}

func (q *Queries) GetTransactionsByPeriod(ctx context.Context, arg GetTransactionsByPeriodParams) ([]RawTransaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByPeriod,
		arg.UserID,
		arg.DatePosted,
		arg.DatePosted_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RawTransaction{}
	for rows.Next() {
		var i RawTransaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.AccountID,
			&i.DatePosted,
			&i.TransactionAmount,
			&i.FitID,
			&i.Checknum,
			&i.Memo,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionsByUser = `-- name: GetTransactionsByUser :many
SELECT id, user_id, account_id, date_posted, transaction_amount, fit_id, checknum, memo, description FROM raw_transactions
WHERE user_id = $1
order by date_posted DESC
LIMIT $2
OFFSET $3
`

type GetTransactionsByUserParams struct {
	UserID int32 `db:"user_id" json:"userId"`
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetTransactionsByUser(ctx context.Context, arg GetTransactionsByUserParams) ([]RawTransaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RawTransaction{}
	for rows.Next() {
		var i RawTransaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.AccountID,
			&i.DatePosted,
			&i.TransactionAmount,
			&i.FitID,
			&i.Checknum,
			&i.Memo,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertRawTransaction = `-- name: InsertRawTransaction :exec
INSERT INTO raw_transactions (
  user_id,
  account_id,
  date_posted,
  transaction_amount,
  fit_id,
  checknum,
  memo,
  description
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (user_id, account_id, date_posted, transaction_amount, fit_id, checknum, memo)
DO NOTHING
`

type InsertRawTransactionParams struct {
	UserID            int32            `db:"user_id" json:"userId"`
	AccountID         pgtype.Text      `db:"account_id" json:"accountId"`
	DatePosted        pgtype.Timestamp `db:"date_posted" json:"datePosted"`
	TransactionAmount pgtype.Numeric   `db:"transaction_amount" json:"transactionAmount"`
	FitID             pgtype.Int4      `db:"fit_id" json:"fitId"`
	Checknum          pgtype.Text      `db:"checknum" json:"checknum"`
	Memo              pgtype.Text      `db:"memo" json:"memo"`
	Description       pgtype.Text      `db:"description" json:"description"`
}

func (q *Queries) InsertRawTransaction(ctx context.Context, arg InsertRawTransactionParams) error {
	_, err := q.db.Exec(ctx, insertRawTransaction,
		arg.UserID,
		arg.AccountID,
		arg.DatePosted,
		arg.TransactionAmount,
		arg.FitID,
		arg.Checknum,
		arg.Memo,
		arg.Description,
	)
	return err
}
