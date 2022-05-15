-- name: GetAuth :one
SELECT * FROM auth WHERE username = $1;