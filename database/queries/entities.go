// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID      int32  `db:"id" json:"id"`
	UserID  int32  `db:"user_id" json:"userId"`
	Name    string `db:"name" json:"name"`
	Enabled bool   `db:"enabled" json:"enabled"`
}

type RawTransaction struct {
	ID                int32            `db:"id" json:"id"`
	UserID            int32            `db:"user_id" json:"userId"`
	AccountID         pgtype.Text      `db:"account_id" json:"accountId"`
	DatePosted        pgtype.Timestamp `db:"date_posted" json:"date_posted"`
	TransactionAmount pgtype.Numeric   `db:"transaction_amount" json:"transaction_amount"`
	FitID             pgtype.Int4      `db:"fit_id" json:"fitId"`
	Checknum          pgtype.Text      `db:"checknum" json:"checknum"`
	Memo              pgtype.Text      `db:"memo" json:"memo"`
	Description       pgtype.Text      `db:"description" json:"description"`
}

type RawTransactionsCategory struct {
	RawTransactionID int32 `db:"raw_transaction_id" json:"rawTransactionId"`
	CategoryID       int32 `db:"category_id" json:"categoryId"`
}

type User struct {
	ID           int32            `db:"id" json:"id"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"createdAt"`
	Name         string           `db:"name" json:"name"`
	Email        string           `db:"email" json:"email"`
	Passwordhash string           `db:"passwordhash" json:"passwordhash"`
	Salt         string           `db:"salt" json:"salt"`
}
