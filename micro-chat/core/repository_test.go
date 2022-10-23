package core

import (
	"chat-in-app_microservices/micro-chat/config"
	"chat-in-app_microservices/micro-chat/db"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	cfg, err := config.LoadConfig("../", "config")
	assert.NoError(t, err)

	repo, err := NewRepository(cfg.Db)
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	defer func(repo *Repository) {
		if err := repo.Disconnect(); err != nil {
			panic(err)
		}
	}(&repo)

	ctx := context.TODO()

	testTable := []db.MessageParam{
		{
			Sender:   "user1",
			Receiver: "user2",
			Message:  "woii",
		},
		{
			Sender:   "user1",
			Receiver: "user2",
			Message:  "kamu",
		},
		{
			Sender:   "user2",
			Receiver: "user1",
			Message:  "ketemuan yuk",
		},
	}

	for _, v := range testTable {
		id, err := repo.Q().SendMessage(ctx, db.MessageParam{
			Sender:   v.Sender,
			Receiver: v.Receiver,
			Message:  v.Message,
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	}
}
