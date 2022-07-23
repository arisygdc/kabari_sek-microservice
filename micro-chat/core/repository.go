package core

import (
	"chat-in-app_microservices/micro-chat/config"
	"chat-in-app_microservices/micro-chat/db"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DriverMongo      string = "mongodb"
	DriverCloudMongo string = "mongodb+srv"
)

type Repository struct {
	conn *db.Connection
	ctx  context.Context
}

func NewRepository(cfg config.DbConfig) (Repository, error) {
	var serverAPIOptions *options.ServerAPIOptions = nil
	var connString = fmt.Sprintf("%s://%s:%s@%s", cfg.Driver, cfg.User, cfg.Password, cfg.Host)

	if cfg.Driver == DriverMongo {
		connString = fmt.Sprintf("%s:%d", connString, cfg.Port)
	}

	if cfg.Driver == DriverCloudMongo {
		serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)
		connString = fmt.Sprintf("%s/?retryWrites=true&w=majority", connString)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := db.NewConnection(ctx, connString, cfg.Name, cfg.Collection, serverAPIOptions)
	if err := conn.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	repo := Repository{
		conn: conn,
		ctx:  ctx,
	}

	return repo, err
}

func (r *Repository) Disconnect() error {
	if err := r.conn.Disconnect(r.ctx); err != nil {
		return err
	}

	r.ctx.Done()
	return nil
}

func (r Repository) Q() db.Querier {
	return r.conn
}
