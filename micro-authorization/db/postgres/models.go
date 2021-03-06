// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package postgres

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID   int16
	Name string
}

type Role struct {
	ID   int16
	Name string
}

type RolePermission struct {
	RoleID       int16
	PermissionID int16
}

type UserBannedPermission struct {
	ID           int64
	UserID       uuid.UUID
	PermissionID int16
	BannedAt     int64
	BannedExp    int64
}

type UserRole struct {
	UserID uuid.UUID
	RoleID int16
}
