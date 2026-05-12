package models

import "time"

const DefaultRealmId = "00000000-0000-0000-0000-000000000001"

type Realm struct {
	Id          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Slug        string    `json:"slug" db:"slug"`
	Description string    `json:"description" db:"description"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type RealmDTO struct {
	Id          string `json:"id" db:"id"`
	Slug        string `json:"slug" db:"slug" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"isActive" db:"is_active"`
}

type UserRealm struct {
	Id        string    `json:"id" db:"id"`
	UserId    string    `json:"userId" db:"user_id"`
	RealmId   string    `json:"realmId" db:"realm_id"`
	RoleId    string    `json:"roleId" db:"role_id"`
	Role      *Role     `json:"role,omitempty"`
	Realm     *Realm    `json:"realm,omitempty"`
	IsActive  bool      `json:"isActive" db:"is_active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserRealmDTO struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"userId" db:"user_id" binding:"required"`
	RealmId   string `json:"realmId" db:"realm_id" binding:"required"`
	RoleId    string `json:"roleId" db:"role_id" binding:"required"`
	IsActive  bool   `json:"isActive" db:"is_active"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

type UserRealmsDTO struct {
	UserId string `json:"userId" db:"user_id" binding:"required"`
}
