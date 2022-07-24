// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: auth.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
)

const getAuth = `-- name: GetAuth :one
SELECT id, username, password, created_at, updated_at FROM auth WHERE username = $1
`

func (q *Queries) GetAuth(ctx context.Context, username string) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuth, username)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertAuth = `-- name: InsertAuth :exec
INSERT INTO auth (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
`

type InsertAuthParams struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

func (q *Queries) InsertAuth(ctx context.Context, arg InsertAuthParams) error {
	_, err := q.db.Exec(ctx, insertAuth,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}