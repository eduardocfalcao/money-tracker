// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repositoy

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type RawTransaction struct {
	ID                pgtype.UUID      `db:"id" json:"id"`
	AccountID         pgtype.Text      `db:"account_id" json:"accountId"`
	DataPosted        pgtype.Timestamp `db:"data_posted" json:"dataPosted"`
	TransactionAmount pgtype.Numeric   `db:"transaction_amount" json:"transactionAmount"`
	FitID             int32            `db:"fit_id" json:"fitId"`
	Checknum          string           `db:"checknum" json:"checknum"`
	Memo              string           `db:"memo" json:"memo"`
}
