package token

import (
	"fmt"
	"time"

	jwtProvider "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Payload interface {
	Valid() error
}

type AuthenticationPayload struct {
	Id         uuid.UUID
	Username   string
	Issued_at  int64
	Expired_at int64
}

func NewAuthPayload(id uuid.UUID, username string, duration time.Duration) AuthenticationPayload {
	now := time.Now()
	id = HashPayloadID(id, username)
	nowAddDur := now.Add(duration)
	return AuthenticationPayload{
		Id:         id,
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

func (Payload AuthenticationPayload) CompareID(user_id uuid.UUID, username string) error {
	if CompareHashPayloadID(Payload.Id, user_id, username) {
		return nil
	}

	return fmt.Errorf("payload id and user id not match")
}
