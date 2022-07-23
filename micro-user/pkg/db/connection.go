package db

import (
	"chat-in-app_microservices/micro-user/pkg/db/postgres"
	"context"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn    *pgx.Conn
	Querier *postgres.Queries
}

func NewConnection(ctx context.Context, connString string) (DB, error) {
	var db DB
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return db, err
	}

	q := postgres.New(conn)
	db = DB{
		Conn:    conn,
		Querier: q,
	}

	return db, nil
}
