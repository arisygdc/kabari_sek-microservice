package db

import (
	"chat-in-app_microservices/micro-user/pkg/db/postgres"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Conn    *pgxpool.Pool
	Querier *postgres.Queries
}

// Explicit connection string
func NewConnection(ctx context.Context, connString string) (DB, error) {
	var db DB
	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return db, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return db, err
	}

	q := postgres.New(conn)
	db = DB{
		Conn:    conn,
		Querier: q,
	}

	log.Println("connected to postgresql database")
	return db, nil
}
