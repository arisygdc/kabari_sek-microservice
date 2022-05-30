package db

import (
	"context"
	"fmt"

	"chat-in-app_microservices/micro-authorization/config"
	"chat-in-app_microservices/micro-authorization/db/postgres"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
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
		querier: postgres.New(conn),
	}
	return repos, nil
}

// Use query function
func (r Repository) Q() postgres.Querier {
	return r.querier
}
