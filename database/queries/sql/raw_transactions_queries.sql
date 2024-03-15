-- name: GetRawTransaction :one
SELECT * FROM raw_transactions
WHERE id = $1 LIMIT 1;
