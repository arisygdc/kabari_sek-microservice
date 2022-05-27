package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var secret = "8be78fda62b3842d8047e2396a0751e3aa7ef78bc23e1c694fe36dfca3080aa0"

func TestJWT(t *testing.T) {
	dur, err := time.ParseDuration("15m")
	assert.NoError(t, err)
	payload := NewAuthPayload("arisygdc", dur)
	err = payload.Valid()
	assert.NoError(t, err)

	jwt := NewJWT(secret)
	signedString, err := jwt.Generate(payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, signedString)

	payload = AuthenticationPayload{}
	jwt.ParseWithClaimAuthPayload(signedString, &payload)
	assert.NotEmpty(t, payload.Username)
}
