package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MenuItemRepo struct {
	db *sqlx.DB
}

func NewMenuItemRepo(db *sqlx.DB) *MenuItemRepo {
	return &MenuItemRepo{
		db: db,
	}
}

type MenuItem interface {
	GetAll(context.Context) ([]*models.MenuItem, error)
	Create(context.Context, models.MenuItemDTO) error
	Update(context.Context, models.MenuItemDTO) error
	Delete(context.Context, string) error
}

func (r *MenuItemRepo) GetAll(ctx context.Context) ([]*models.MenuItem, error) {
	query := fmt.Sprintf(`SELECT id, name, description, is_show FROM %s`, MenuItemTable)
	items := []*models.MenuItem{}

	if err := r.db.SelectContext(ctx, items, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return items, nil
}

func (r *MenuItemRepo) Create(ctx context.Context, menu models.MenuItemDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, name, description, is_show) VALUES ($1, $2, $3, $4)`, MenuItemTable)
	id := uuid.New()

	_, err := r.db.ExecContext(ctx, query, id, menu.Name, menu.Description, menu.IsShow)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *MenuItemRepo) Update(ctx context.Context, menu models.MenuItemDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=$1, description=$2, is_show=$3 WHERE id=$4`, MenuItemTable)

	_, err := r.db.ExecContext(ctx, query, menu.Name, menu.Description, menu.IsShow, menu.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *MenuItemRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, MenuItemTable)

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
