package pq_models

import (
	"database/sql"
	"time"
)

type User struct {
	Id               string         `db:"id"`
	Username         string         `db:"username"`
	Email            string         `db:"email"`
	FirstName        string         `db:"first_name"`
	LastName         string         `db:"last_name"`
	CreatedAt        time.Time      `db:"created_at"`
	IsActive         sql.NullBool   `db:"is_active"`
	UserRealmId      sql.NullString `db:"ur_id"`
	RoleId           sql.NullString `db:"role_id"`
	RoleSlug         sql.NullString `db:"role_slug"`
	RoleName         sql.NullString `db:"role_name"`
	RoleDescription  sql.NullString `db:"role_description"`
	RoleLevel        sql.NullInt64  `db:"role_level"`
	RoleIsActive     sql.NullBool   `db:"role_is_active"`
	RoleIsEditable   sql.NullBool   `db:"role_is_editable"`
	RealmId          sql.NullString `db:"realm_id"`
	RealmName        sql.NullString `db:"realm_name"`
	RealmDescription sql.NullString `db:"realm_description"`
	RealmCreatedAt   sql.NullTime   `db:"realm_created_at"`
}
