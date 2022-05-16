package user

import (
	"chat-in-app_microservices/micro-user/core"
	"chat-in-app_microservices/micro-user/db/postgres"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	err = bcrypt.CompareHashAndPassword([]byte(authenticate.Password), []byte(auth.password))
	return err
}

func (s Service) Register(ctx context.Context, auth Auth, user User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(auth.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authParam := postgres.InsertAuthParams{
		ID:       uuid.New(),
		Username: auth.username,
		Password: string(password),
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

		if err != nil {
			return err
		}

		return q.InsertAuthUser(ctx, postgres.InsertAuthUserParams{
			AuthID: authParam.ID,
			UserID: userParam.ID,
		})
	})
}
