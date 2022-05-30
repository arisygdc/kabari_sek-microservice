-- name: UserRoles :many
SELECT r.* FROM user_role ur
INNER JOIN role r ON ur.role_id = r.id
WHERE ur.user_id = $1;

-- name: UserBannedPermissions :many
SELECT * FROM user_banned_permission WHERE user_id = $1 AND banned_exp > @now::bigint;

-- name: UserBannedRoles :many
SELECT * FROM user_banned_role WHERE user_id = $1 AND banned_exp > @now::bigint;

-- name: RolePermissions :many
SELECT * FROM role_permission rp
INNER JOIN permission p ON rp.permission_id = p.id
WHERE rp.role_id = $1;
