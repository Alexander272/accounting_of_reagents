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
type AmountType interface {
	postgres.AmountType
}
type Reagent interface {
	postgres.Reagent
}
type Spending interface {
	postgres.Spending
}
type Extending interface {
	postgres.Extending
}
type Notes interface {
	postgres.Notes
}

type Repository struct {
	Role
	MenuItem
	Menu

	ReagentType
	AmountType
	Reagent
	Spending
	Extending
	Notes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Role:     postgres.NewRoleRepo(db),
		MenuItem: postgres.NewMenuItemRepo(db),
		Menu:     postgres.NewMenuRepo(db),

		ReagentType: postgres.NewReagentTypeRepo(db),
		AmountType:  postgres.NewAmountTypeRepo(db),
		Reagent:     postgres.NewReagentRepo(db),
		Spending:    postgres.NewSpendingRepo(db),
		Extending:   postgres.NewExtendingRepo(db),
		Notes:       postgres.NewNotesRepo(db),
	}
}
