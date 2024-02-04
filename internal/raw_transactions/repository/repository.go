package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repositoy interface {
	Get(ctx context.Context, id pgtype.UUID) (RawTransaction, error)
	Tx(tx pgx.Tx) Repositoy
}

func (q *Queries) Tx(tx pgx.Tx) Repositoy {
	return &Queries{
		db: tx,
	}
}

var _ Repositoy = (*Queries)(nil)
