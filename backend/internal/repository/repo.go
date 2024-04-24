package repository

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Role interface {
	postgres.Role
}
type MenuItem interface {
	postgres.MenuItem
}
type Menu interface {
	postgres.Menu
}

type ReagentType interface {
	postgres.ReagentType
}

type Repository struct {
	Role
	MenuItem
	Menu

	ReagentType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Role:     postgres.NewRoleRepo(db),
		MenuItem: postgres.NewMenuItemRepo(db),
		Menu:     postgres.NewMenuRepo(db),

		ReagentType: postgres.NewReagentTypeRepo(db),
	}
}
