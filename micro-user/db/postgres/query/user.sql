-- name: InsertUser :exec
INSERT INTO users (id, firstname, lastname, birth) VALUES (@id, @firstname, @lastname, @birth);

-- name: GetUser :one
SELECT * FROM users WHERE id = @id;