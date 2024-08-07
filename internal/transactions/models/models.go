package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Amount float64

type RawTransaction struct {
	DatePosted        time.Time   `db:"date_posted" json:"date_posted"`
	Type              pgtype.Text `db:"transaction_type" json:"type"`
	AccountID         pgtype.Text `db:"account_id" json:"account_id"`
	Checknum          pgtype.Text `db:"checknum" json:"checknum"`
	Memo              pgtype.Text `db:"memo" json:"memo"`
	Description       pgtype.Text `db:"description" json:"description"`
	FitID             pgtype.Text `db:"fit_id" json:"fit_id"`
	TransactionAmount float64     `db:"transaction_amount" json:"transaction_amount"`
	ID                int32       `db:"id" json:"id"`
	UserID            int32       `db:"user_id" json:"user_id"`
}
