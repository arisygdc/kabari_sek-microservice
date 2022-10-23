package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	testLogin := []struct {
		*LoginRequest
		error
	}{
		{&LoginRequest{Username: " keatn ", Password: "uhofwauohfaew"}, ErrUsernameLeng},
		{&LoginRequest{Username: " keatngajbwef ", Password: "uhofwauohfaew"}, nil},
		{&LoginRequest{Username: " ;;;;;;;;", Password: "uhofwauohfaew"}, ErrInvalidChar},
		{&LoginRequest{Username: "Pisan09", Password: "uhofwauohfaew"}, nil},
		{&LoginRequest{Username: "Pisan09", Password: " egjh "}, ErrPassLeng},
		{&LoginRequest{Username: "Pisan09", Password: " e'gjh "}, ErrInvalidChar},
		{&LoginRequest{Username: "Pisan09", Password: "/\\ egjh "}, ErrInvalidChar},
		{&LoginRequest{Username: "Pisan09", Password: "\"egjh "}, ErrInvalidChar},
	}

	for _, v := range testLogin {
		err := ValidateRequest(v)
		assert.ErrorIs(t, err, v.error, fmt.Sprintf("param: usr:\"%s\", pwd:\"%s\" \nexpected error %v, got value %v", v.Username, v.Password, v.error, err))
	}

	a := &RegisterRequest{
		LoginRequest: LoginRequest{
			Username: "arisygdc",
			Password: "iawgfiewagbfw",
		},
		Name: Name{
			Firstname: "woiyoo",
			Lastname:  "iuwaefig",
		},
		Birth: "1998-7-27",
	}
	ValidateRequest(a)
}
