package core

import (
	"chat-in-app_microservices/micro-user/config"
	"chat-in-app_microservices/micro-user/pkg/db"
	"chat-in-app_microservices/micro-user/pkg/db/postgres"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db db.DB
}

// If using implicit connection string
// postgresql://%s:%s@%s:%d/%s?sslmode=disable
// user:pass@host:port/dbName
func NewRepository(conf config.DbConfig) (Repository, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("connecting to postgresql %s:%d", conf.Host, conf.Port)

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Name, conf.SslMode)

	// Additional options
	connString = fmt.Sprintf("%s pool_max_conns=%d pool_min_conns=%d",
		connString, conf.Pool.MaxSize, conf.Pool.MinSize)

	conn, err := db.NewConnection(ctx, connString)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		db: conn,
	}, nil
}

func (r Repository) Query() postgres.Querier {
	return r.db.Querier
}

func (r Repository) Tx(ctx context.Context, statement func(postgres.Querier) error) error {
	pgTx, err := r.db.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	tx := r.db.Querier.WithTx(pgTx)
	err = statement(tx)
	if err != nil {
		rllbkErr := pgTx.Rollback(ctx)
		if rllbkErr != nil {
			return err
		}
		return err
	}
	err = pgTx.Commit(ctx)
	return err
}
