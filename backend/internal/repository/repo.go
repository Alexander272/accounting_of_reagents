package repository

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	postgres.Transaction
}

type Permissions interface {
	postgres.Permissions
}
type Roles interface {
	postgres.Roles
}
type RoleHierarchy interface {
	postgres.RoleHierarchy
}
type Realm interface {
	postgres.Realm
}
type UserRealm interface {
	postgres.UserRealm
}
type User interface {
	postgres.User
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
	Transaction

	Permissions
	Roles
	RoleHierarchy
	Realm
	UserRealm
	User

	ReagentType
	AmountType
	Reagent
	Spending
	Extending
	Notes
}

func NewRepository(db *sqlx.DB) *Repository {
	transaction := postgres.NewTransactionRepo(db)

	return &Repository{
		Transaction: transaction,

		Permissions:   postgres.NewPermissionRepo(db, transaction),
		Roles:         postgres.NewRoleRepo(db, transaction),
		RoleHierarchy: postgres.NewRoleHierarchyRepo(db, transaction),
		Realm:         postgres.NewRealmRepo(db),
		UserRealm:     postgres.NewUserRealmRepo(db, transaction),
		User:          postgres.NewUserRepo(db),

		ReagentType: postgres.NewReagentTypeRepo(db),
		AmountType:  postgres.NewAmountTypeRepo(db),
		Reagent:     postgres.NewReagentRepo(db),
		Spending:    postgres.NewSpendingRepo(db),
		Extending:   postgres.NewExtendingRepo(db),
		Notes:       postgres.NewNotesRepo(db),
	}
}
