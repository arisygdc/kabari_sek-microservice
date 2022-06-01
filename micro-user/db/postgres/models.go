// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birth     time.Time `json:"birth"`
	CreatedAt time.Time `json:"created_at"`
}
