package repository

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Role interface {
	postgres.Role
}

type Repository struct {
	Role
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Role: postgres.NewRoleRepo(db),
	}
}
