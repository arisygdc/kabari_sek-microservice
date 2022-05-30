package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var secret = "8be78fda62b3842d8047e2396a0751e3aa7ef78bc23e1c694fe36dfca3080aa0"

func TestJWT(t *testing.T) {
	dur, err := time.ParseDuration("15m")
	assert.NoError(t, err)

	// Create auth payload
	idUser := uuid.New()
	username := "arisygdc"
	payload := NewAuthPayload(idUser, username, dur)
	err = payload.Valid()
	assert.NoError(t, err)

	// Generate jwt signed string
	jwt := NewJWT(secret)
	signedString, err := jwt.Generate(payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, signedString)

	// Claim jwt
	payload = AuthenticationPayload{}
	jwt.ParseWithClaimAuthPayload(signedString, &payload)
	assert.NotEmpty(t, payload.Username)

	// Compare hash ID
	err = payload.CompareID(idUser, username)
	assert.NoError(t, err)
	err = payload.CompareID(uuid.New(), username)
	assert.ErrorIs(t, err, ErrTokenInvalidId)
}
