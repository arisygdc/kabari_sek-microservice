-- name: UserBannedPermission :one
SELECT * FROM user_banned_permission WHERE user_id = $1 AND permission_id = $2 AND banned_exp > @now::bigint;

-- name: Permission :one
SELECT * FROM permission WHERE name = $1;

-- name: Role :one
SELECT * FROM role WHERE name = $1;

-- name: UserPermit :one
SELECT rp.* FROM user_role ur
INNER JOIN role_permission rp ON ur.role_id = rp.role_id
WHERE ur.user_id = $1 AND rp.permission_id = $2;

-- name: BannUserPermission :exec
INSERT INTO user_banned_permission (user_id, permission_id, banned_at, banned_exp) VALUES ($1, $2, $3, $4);

-- name: RoleGrantPermission :exec
INSERT INTO role_permission (role_id, permission_id) VALUES ($1, $2);

-- name: InsertUserRole :exec
INSERT INTO user_role (user_id, role_id) VALUES ($1, $2);