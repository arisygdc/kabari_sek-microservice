package token

import (
	"fmt"

	jwtProvider "github.com/golang-jwt/jwt/v4"
)

var signingMethod = jwtProvider.SigningMethodHS256

type JWToken struct {
	secret string
}

func NewJWT(secret string) JWToken {
	return JWToken{
		secret: secret,
	}
}

func (jwt JWToken) Generate(payload Payload) (string, error) {
	claim := jwtProvider.NewWithClaims(signingMethod, payload)
	return claim.SignedString([]byte(jwt.secret))
}

func (jwt JWToken) ParseWithClaimToken(signedToken string, payloadClaim Payload) (*jwtProvider.Token, error) {
	keyFunc := func(t *jwtProvider.Token) (interface{}, error) {
		if t.Method.Alg() != signingMethod.Alg() {
			return nil, fmt.Errorf("signing method not provide")
		}
		return []byte(jwt.secret), nil
	}
	return jwtProvider.ParseWithClaims(signedToken, jwtProvider.Claims(payloadClaim), keyFunc)
}

func (jwt JWToken) ParseWithClaimAuthPayload(signedToken string, payload *AuthenticationPayload) error {
	token, err := jwt.ParseWithClaimToken(signedToken, payload)
	if err != nil {
		return err
	}

	var ok bool
	payload, ok = token.Claims.(*AuthenticationPayload)
	if !ok {
		return jwtProvider.ErrTokenInvalidClaims
	}

	return nil
}
