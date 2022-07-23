-- name: InsertUser :exec
INSERT INTO users (id, firstname, lastname, birth, created_at, updated_at) VALUES (@id, @firstname, @lastname, @birth, @created_at, @updated_at);

-- name: GetUser :one
SELECT * FROM users WHERE id = @id;