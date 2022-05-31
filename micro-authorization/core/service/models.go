package service

import "github.com/google/uuid"

type Permission struct {
	ID   int16
	Name string
}

type UserBannedPermission struct {
	ID           int64
	UserID       uuid.UUID
	PermissionID int16
	BannedAt     int64
	BannedExp    int64
}
