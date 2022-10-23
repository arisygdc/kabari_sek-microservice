-- name: GetAuth :one
SELECT * FROM auth WHERE username = $1;

-- name: InsertAuth :exec
INSERT INTO auth (id, username, password, created_at, updated_at) VALUES (@id, @username, @password, @created_at, @updated_at);