package repository

const insertRawTransaction = ` 
INSERT INTO raw_transactions (
	user_id,
  account_id,
  date_posted,
  transaction_amount,
  fit_id,
  checknum,
  memo,
  description
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

//ON CONFLICT (account_id, date_posted, transaction_amount, fit_id, checknum, memo)
//DO NOTHING;
//`

const getTransactionById = `
SELECT * FROM raw_transactions
WHERE id = $1 LIMIT 1;
`

const getTransactionsByUser = `-- name: GetTransactionsByUser :many
SELECT id, user_id, account_id, date_posted, transaction_amount, fit_id, checknum, memo, description FROM raw_transactions
WHERE user_id = $1
order by date_posted DESC
LIMIT $2
OFFSET $3
`
