-- name: InsertUser :exec
INSERT INTO users (id, firstname, lastname, birth) VALUES (@id, @firstname, @lastname, @birth);

-- name: InsertAuthUser :exec
INSERT INTO auth_user (auth_id, user_id) VALUES (@auth_id, @user_id);