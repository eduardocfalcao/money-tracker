-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: SearchUsers :many
SELECT * FROM users u
where CASE WHEN LENGTH(@name::text) != 0 THEN t.name LIKE '%'+@name::text +'%' ELSE TRUE END;

-- name: CreateUser :exec
INSERT INTO users (name, email, passwordHash, salt)
VALUES ($1, $2, $3, $4);
