package user

import (
	"chat-in-app_microservices/micro-user/core"
	"chat-in-app_microservices/micro-user/pkg/db/postgres"
	"chat-in-app_microservices/micro-user/pkg/token"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
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
		if err == pgx.ErrNoRows {
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
	generatedToken, err := jwToken.Generate(payload)
	return generatedToken, err
}

// Register user needs authentication and user information
// Return username
// TODO Should return inserted username
func (s Service) Register(ctx context.Context, user RegisterUser) (string, error) {
	password, err := HashPassword(user.password)
	if err != nil {
		return "", err
	}

	registerTime := time.Now().Unix()

	authParam := postgres.InsertAuthParams{
		ID:        uuid.New(),
		Username:  user.username,
		Password:  password,
		CreatedAt: registerTime,
		UpdatedAt: registerTime,
	}

	userParam := postgres.InsertUserParams{
		ID:        uuid.New(),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Birth:     user.Birth,
		CreatedAt: registerTime,
		UpdatedAt: registerTime,
	}

	execTx := s.repo.Tx(ctx, func(q postgres.Querier) error {
		err := q.InsertAuth(ctx, authParam)

		if err != nil {
			return err
		}

		err = q.InsertUser(ctx, userParam)

		return err
	})

	return authParam.Username, execTx
}

// Get user information by id
func (s Service) GetUser(ctx context.Context, id uuid.UUID) (postgres.User, error) {
	return s.repo.Query().GetUser(ctx, id)
}

func (s Service) VerifySession(ctx context.Context, tokenSecretKey, tokenString string) error {
	var payload token.AuthenticationPayload

	tokenHandler := token.NewJWT(tokenSecretKey)
	err := tokenHandler.ParseWithClaimAuthPayload(tokenString, &payload)
	if err != nil {
		return err
	}

	auth, err := s.repo.Query().GetAuth(ctx, payload.Username)
	if err != nil {
		return err
	}

	return payload.CompareID(auth.ID, auth.Username)
}
