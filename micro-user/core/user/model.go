package user

import (
	"time"
)

type Auth struct {
	username, password string
}

type User struct {
	Firstname string
	Lastname  string
	Birth     time.Time
}

type RegisterUser struct {
	Auth
	User
}
