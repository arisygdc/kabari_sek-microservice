-- name: InsertUser :exec
INSERT INTO users (id, firstname, lastname, birth) VALUES (@id, @firstname, @lastname, @birth);