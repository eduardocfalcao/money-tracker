package queries

import "github.com/jackc/pgx/v5"

type QuerierTx interface {
	Querier
	Tx(tx pgx.Tx) QuerierTx
}

func (q *Queries) Tx(tx pgx.Tx) QuerierTx {
	return q.WithTx(tx)
}
