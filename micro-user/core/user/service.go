package user

import (
	"chat-in-app_microservices/micro-user/core"
	"chat-in-app_microservices/micro-user/db/postgres"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo core.Repository
}

func NewService(repo core.Repository) Service {
	return Service{repo: repo}
}

// Login return nil on successed and error on failure
func (s Service) Login(ctx context.Context, auth Auth) error {
	authenticate, err := s.repo.Query().GetAuth(ctx, auth.username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("auth not registered")
		}
		return err
	}

	err = CheckHashPassword(authenticate.Password, auth.password)
	return err
}

// Register user needs authentication and user information
func (s Service) Register(ctx context.Context, user RegisterUser) error {
	password, err := HashPassword(user.password)
	if err != nil {
		return err
	}

	authParam := postgres.InsertAuthParams{
		ID:       uuid.New(),
		Username: user.username,
		Password: password,
	}

	userParam := postgres.InsertUserParams{
		ID:        uuid.New(),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Birth:     user.Birth,
	}

	return s.repo.Tx(ctx, func(q postgres.Querier) error {
		err := q.InsertAuth(ctx, authParam)

		if err != nil {
			return err
		}

		err = q.InsertUser(ctx, userParam)

		return err
	})
}

// Get user information by id
func (s Service) GetUser(ctx context.Context, id uuid.UUID) (postgres.User, error) {
	return s.repo.Query().GetUser(ctx, id)
}
