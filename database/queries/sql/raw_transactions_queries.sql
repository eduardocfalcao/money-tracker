-- name: InsertRawTransaction :exec
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
DO NOTHING;

-- name: GetRawTransaction :one
SELECT * FROM raw_transactions
WHERE id = $1 LIMIT 1;

-- name: GetTransactionsByUser :many
SELECT * FROM raw_transactions
WHERE user_id = $1
order by date_posted DESC
LIMIT $2
OFFSET $3;

-- name: GetTransactionsByPeriod :many
SELECT * FROM raw_transactions
WHERE user_id = $1
 AND date_posted BETWEEN $2 and $3
order by date_posted DESC
LIMIT $4
OFFSET $5;
