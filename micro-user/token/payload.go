package token

import (
	"fmt"
	"time"

	jwtProvider "github.com/golang-jwt/jwt/v4"
)

type Payload interface {
	Valid() error
}

type AuthenticationPayload struct {
	Username   string
	Issued_at  int64
	Expired_at int64
}

func NewAuthPayload(username string, duration time.Duration) AuthenticationPayload {
	now := time.Now()
	nowAddDur := now.Add(duration)
	return AuthenticationPayload{
		Username:   username,
		Issued_at:  now.Unix(),
		Expired_at: nowAddDur.Unix(),
	}
}

func (payload AuthenticationPayload) Valid() error {
	issued := time.Unix(payload.Issued_at, 0)
	exp := time.Unix(payload.Expired_at, 0)
	if issued.After(exp) {
		return fmt.Errorf("token issue after expired")
	}

	if time.Now().After(exp) {
		return jwtProvider.ErrTokenExpired
	}

	return nil
}
