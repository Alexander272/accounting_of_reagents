package models

type RolePermission struct {
	RoleId       string `json:"roleId" db:"role_id"`
	PermissionId string `json:"permissionId" db:"permission_id"`
}

type RolePermissionDTO struct {
	ActorId      string `json:"actorId" db:"actor_id"`
	RoleId       string `json:"roleId" db:"role_id"`
	PermissionId string `json:"permissionId" db:"permission_id"`
}
