package user

import (
	"chat-in-app_microservices/micro-user/config"
	"chat-in-app_microservices/micro-user/core"
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUtility(t *testing.T) {
	password := "abdsiuhfweaofnwe"
	hashed, err := HashPassword(password)
	assert.NoError(t, err)

	err = CheckHashPassword(hashed, password)
	assert.NoError(t, err)

	timeParsed, err := ParseTime("2017-06-21")
	assert.NoError(t, err)
	log.Println(timeParsed)
}

func GetConfig() (config.Config, error) {
	cfg := config.Config{}
	err := config.LoadConfig("../../", "config", &cfg)
	return cfg, err
}

func TestService(t *testing.T) {
	cfg, err := GetConfig()
	assert.NoError(t, err)
	secretKey := "68f82936c216ddc153bd16a52fd54c43b70271e2ff9b4885dfd6f6114f0e1565"

	repos, err := core.NewRepository(cfg.Database)
	assert.NoError(t, err)

	svc := NewService(repos)
	ctx := context.Background()

	birth, err := ParseTime("2017-06-21")
	assert.NoError(t, err)

	auth := Auth{
		username: "simons",
		password: "the hasingin",
	}

	err = svc.Register(ctx, RegisterUser{
		Auth: auth,
		User: User{
			Firstname: "simon",
			Lastname:  "celi",
			Birth:     birth,
		},
	})
	assert.NoError(t, err)

	sessionString, err := svc.Login(ctx, auth, secretKey, 5*time.Minute)
	assert.NoError(t, err)

	err = svc.VerifySession(ctx, secretKey, sessionString)
	assert.NoError(t, err)
}
