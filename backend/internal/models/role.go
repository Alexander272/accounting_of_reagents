package models

import "time"

type Role struct {
	Id          string    `json:"id" db:"id"`
	Slug        string    `json:"slug" db:"slug"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Level       int       `json:"level" db:"level"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	IsSystem    bool      `json:"isSystem" db:"is_system"`
	IsEditable  bool      `json:"isEditable" db:"is_editable"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type RoleWithStats struct {
	Role
	Children   []string       `json:"children"`
	PermsCount PermsWithCount `json:"perms"`
	UserCount  int            `json:"userCount"`
}

type RoleShort struct {
	Id   string `json:"id" db:"id"`
	Slug string `json:"slug" db:"slug"`
	Name string `json:"name" db:"name"`
}

type GetRoleDTO struct {
	Id   string `json:"id" db:"id"`
	Slug string `json:"slug" db:"slug"`
}

type RoleDTO struct {
	Id string `json:"id" db:"id"`
	// Actor       Actor
	Slug        string    `json:"slug" db:"slug"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Level       int       `json:"level" db:"level"`
	IsSystem    bool      `json:"isSystem" db:"is_system"`
	Permissions []string  `json:"permissions" db:"permissions"`
	Inherits    []string  `json:"inherits" db:"inherits"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type DeleteRoleDTO struct {
	Id string `json:"id" db:"id"`
	// Actor Actor
}

type RoleInheritance struct {
	ParentRole string
	ChildRole  string
	Realm      string
}

type RoleWithPerms struct {
	Role
	Inherits []string                  `json:"inherits"`
	Perms    []*RolePermissionsGrouped `json:"perms"`
}
