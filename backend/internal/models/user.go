package models

import "time"

type User struct {
	Id          string              `json:"id" db:"id"`
	Role        string              `json:"role"`
	Name        string              `json:"name" db:"username"`
	Permissions map[string][]string `json:"permissions"`
	Realms      []*UserRealm        `json:"realms,omitempty"`

	AccessToken  string `json:"token"`
	RefreshToken string `json:"-"`
}

type UserDTO struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username" binding:"required"`
	Email     string `json:"email" db:"email"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}

type UserShort struct {
	Id        string `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

type UserRole struct {
	UserId   string
	RealmId  string
	RoleName string
}

type KeycloakUser struct {
	Id        *string `json:"id"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type UserData struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	IsActive  bool      `json:"isActive" db:"is_active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`

	Realms []*UserRealm `json:"realms,omitempty"`
}

type UserDataDTO struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	IsActive  bool   `json:"isActive" db:"is_active"`

	Realms []*UserRealmDTO `json:"realms"`
}
