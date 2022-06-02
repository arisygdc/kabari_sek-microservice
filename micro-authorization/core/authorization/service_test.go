package authorization

import (
	"chat-in-app_microservices/micro-authorization/config"
	"chat-in-app_microservices/micro-authorization/db"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	TestRoleBasic       string = "basic"
	TestRoleAdvance     string = "advance"
	TestPermissionText  string = "chat-text"
	TestPermissionEmote string = "chat-special-emote"
)

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Config{}
	err := config.LoadConfig("../", "config", &cfg)
	assert.NoError(t, err)

	repos, err := db.NewRepository(ctx, cfg.Database)
	assert.NoError(t, err)

	// Insert User Role
	type userRole struct {
		user_id uuid.UUID
		role    string
	}

	testTableUserRole := []userRole{
		{user_id: uuid.New(), role: TestRoleBasic},
		{user_id: uuid.New(), role: TestRoleBasic},
		{user_id: uuid.New(), role: TestRoleAdvance},
	}

	svcAuth := NewAuthorization(repos)
	for _, v := range testTableUserRole {
		err = svcAuth.InsertUserRole(ctx, v.user_id, v.role)
		assert.NoError(t, err)
	}

	// User get permission
	type userGetPermission struct {
		userRole
		request_permission string
		expected           error
	}

	testTableGetPermission := []userGetPermission{
		{userRole: testTableUserRole[0], request_permission: TestPermissionText, expected: nil},
		{userRole: testTableUserRole[0], request_permission: TestPermissionEmote, expected: ErrPermissionDenied},
		{userRole: testTableUserRole[1], request_permission: TestPermissionText, expected: nil},
		{userRole: testTableUserRole[1], request_permission: TestPermissionEmote, expected: ErrPermissionDenied},
		{userRole: testTableUserRole[2], request_permission: TestPermissionEmote, expected: nil},
		{userRole: testTableUserRole[2], request_permission: TestPermissionText, expected: nil},
	}

	for _, v := range testTableGetPermission {
		_, err = svcAuth.UserGetPermission(ctx, v.user_id, v.request_permission)
		assert.ErrorIs(t, v.expected, err)
	}

	// User get banned
	type UserGetBann struct {
		userGetPermission
	}

	err = svcAuth.BannPermissions(ctx, testTableUserRole[1].user_id, 1*time.Minute, TestPermissionText)
	assert.NoError(t, err)

	testTableGetBann := []UserGetBann{
		{userGetPermission: userGetPermission{
			userRole:           testTableUserRole[0],
			request_permission: TestPermissionText,
			expected:           ErrUserNotBann,
		}},
		{userGetPermission: userGetPermission{
			userRole:           testTableUserRole[1],
			request_permission: TestPermissionText,
			expected:           nil,
		}},
	}

	for _, v := range testTableGetBann {
		_, err = svcAuth.UserGetBannedPermission(ctx, v.user_id, v.request_permission)
		assert.ErrorIs(t, v.expected, err)
	}
}
