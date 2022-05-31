package service

import (
	"chat-in-app_microservices/micro-authorization/db"
	"chat-in-app_microservices/micro-authorization/db/postgres"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type (
	PermissionId   int16
	PermissionName string
)

type AuthorizationService struct {
	repos           db.IRepository
	permissionCache map[PermissionName]PermissionId
}

var (
	ErrPermissionNotValid   = fmt.Errorf("permission not valid")
	ErrRoleNotValid         = fmt.Errorf("role not valid")
	ErrUserNotBann          = fmt.Errorf("user are not in bann list")
	ErrPermissionDenied     = fmt.Errorf("permission denied")
	ErrPermissionNotProvide = fmt.Errorf("permission not provide")
)

func NewAuthorization(repos db.IRepository) AuthorizationService {
	return AuthorizationService{
		repos: repos,
	}
}

func (svc AuthorizationService) UserGetPermission(ctx context.Context, user_id uuid.UUID, permission_name string) (Permission, error) {
	var permission Permission
	checkPermision, err := svc.getPermission(ctx, permission_name)
	if err != nil {
		return permission, ErrPermissionNotValid
	}

	_, err = svc.repos.Q().UserPermit(ctx, postgres.UserPermitParams{
		UserID:       user_id,
		PermissionID: checkPermision.ID,
	})

	if err != nil {
		if err == pgx.ErrNoRows {
			return permission, ErrPermissionDenied
		}
		return permission, err
	}

	permission = Permission(checkPermision)
	return permission, nil
}

func (svc AuthorizationService) UserGetBannedPermission(ctx context.Context, user_id uuid.UUID, permission_name string) (UserBannedPermission, error) {
	var bann UserBannedPermission
	checkPermision, err := svc.getPermission(ctx, permission_name)
	if err != nil {
		return bann, ErrPermissionNotValid
	}

	bannRow, err := svc.repos.Q().UserBannedPermission(ctx, postgres.UserBannedPermissionParams{
		UserID:       user_id,
		PermissionID: checkPermision.ID,
		Now:          time.Now().Unix(),
	})

	if err != nil {
		if err == pgx.ErrNoRows {
			return bann, ErrUserNotBann
		}
		return bann, err
	}

	bann = UserBannedPermission(bannRow)
	return bann, nil
}

func (svc AuthorizationService) RoleGrantPermissions(ctx context.Context, role_name string, permission_name ...string) error {

	if len(permission_name) < 1 {
		return ErrPermissionNotProvide
	}

	stmt := func(q postgres.Querier) error {
		role, err := q.Role(ctx, role_name)
		if err != nil {
			if err == pgx.ErrNoRows {
				return ErrRoleNotValid
			}
		}

		for _, v := range permission_name {
			permission, err := svc.getPermission(ctx, v)
			if err != nil {
				return err
			}

			err = q.RoleGrantPermission(ctx, postgres.RoleGrantPermissionParams{
				RoleID:       role.ID,
				PermissionID: permission.ID,
			})

			if err != nil {
				return err
			}
		}

		return nil
	}

	return svc.repos.TX(ctx, stmt)
}

func (svc AuthorizationService) BannPermissions(ctx context.Context, user_id uuid.UUID, duration time.Duration, permission_name ...string) error {

	if len(permission_name) < 1 {
		return ErrPermissionNotProvide
	}

	stmt := func(q postgres.Querier) error {
		now := time.Now()
		for _, v := range permission_name {
			permission, err := svc.getPermission(ctx, v)
			if err != nil {
				return err
			}

			err = q.BannUserPermission(ctx, postgres.BannUserPermissionParams{
				UserID:       user_id,
				PermissionID: permission.ID,
				BannedAt:     now.Unix(),
				BannedExp:    now.Add(duration).Unix(),
			})

			if err != nil {
				return err
			}
		}

		return nil
	}

	return svc.repos.TX(ctx, stmt)
}

func (svc AuthorizationService) getPermission(ctx context.Context, permission_name string) (postgres.Permission, error) {
	permissionRow, ok := svc.getPermissionCache(permission_name)
	if !ok {
		permissionRow, err := svc.repos.Q().Permission(ctx, permission_name)
		if err != nil {
			return permissionRow, err
		}
		svc.setPermissionCache(permissionRow.ID, permissionRow.Name)
	}

	return permissionRow, nil
}

func (svc *AuthorizationService) setPermissionCache(permission_id int16, permission_name string) {
	svc.permissionCache[PermissionName(permission_name)] = PermissionId(permission_id)
}

func (svc AuthorizationService) getPermissionCache(permission_name string) (postgres.Permission, bool) {
	var permission postgres.Permission
	id, ok := svc.permissionCache[PermissionName(permission_name)]
	if !ok {
		return permission, ok
	}

	permission = postgres.Permission{
		ID:   int16(id),
		Name: permission_name,
	}

	return permission, ok
}
