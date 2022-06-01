package user

import (
	"chat-in-app_microservices/micro-user/core"
	"chat-in-app_microservices/micro-user/db/postgres"
	"chat-in-app_microservices/micro-user/token"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrIncorrectUserOrPass = fmt.Errorf("incorrect username or password")
)

type Service struct {
	repo core.Repository
}

func NewService(repo core.Repository) Service {
	return Service{repo: repo}
}

// Login return nil on successed and error on failure
func (s Service) Login(ctx context.Context, auth Auth, tokenSecretKey string, tokenDuration time.Duration) (string, error) {
	authenticate, err := s.repo.Query().GetAuth(ctx, auth.username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrIncorrectUserOrPass
		}

		return "", err
	}

	err = CheckHashPassword(authenticate.Password, auth.password)
	if err != nil {
		return "", ErrIncorrectUserOrPass
	}

	token, err := s.token(tokenSecretKey, authenticate.ID, auth.username, tokenDuration)
	return token, err
}

// return jwt token string
func (s Service) token(secretKey string, id uuid.UUID, username string, duration time.Duration) (string, error) {
	payload := token.NewAuthPayload(id, username, duration)
	jwToken := token.NewJWT(secretKey)
	return jwToken.Generate(payload)
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