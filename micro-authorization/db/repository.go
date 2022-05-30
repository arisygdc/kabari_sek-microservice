package db

import (
	"context"
	"fmt"

	"chat-in-app_microservices/micro-authorization/config"
	"chat-in-app_microservices/micro-authorization/db/postgres"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	conn    *pgx.Conn
	querier *postgres.Queries
}

func newConnection(ctx context.Context, config config.DbConfig) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Pass, config.Host, config.Port, config.Name)
	return pgx.Connect(ctx, connString)
}

func NewRepository(ctx context.Context, config config.DbConfig) (*Repository, error) {
	conn, err := newConnection(ctx, config)
	if err != nil {
		return nil, err
	}
	repos := &Repository{
		conn:    conn,
		querier: postgres.New(conn),
	}
	return repos, nil
}

// Use query function
func (r Repository) Q() postgres.Querier {
	return r.querier
}

// Query with transaction
func (r Repository) TX(ctx context.Context, stmt func(q postgres.Querier) error) error {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	qTx := r.querier.WithTx(tx)
	err = txErrAction(ctx, tx, stmt(qTx))
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func txErrAction(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		RlbErr := tx.Rollback(ctx)
		if RlbErr != nil {
			return RlbErr
		}
		return err
	}
	return nil
}
