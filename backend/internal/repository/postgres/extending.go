package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ExtendingRepo struct {
	db *sqlx.DB
}

func NewExtendingRepo(db *sqlx.DB) *ExtendingRepo {
	return &ExtendingRepo{
		db: db,
	}
}

type Extending interface {
	GetByReagentId(context.Context, string) ([]*models.Extending, error)
	Create(context.Context, *models.ExtendingDTO) (string, error)
	Update(context.Context, *models.ExtendingDTO) error
	Delete(context.Context, *models.DeleteExtendingDTO) error
}

func (r *ExtendingRepo) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Extending, error) {
	query := fmt.Sprintf(`SELECT id, date_of_extending, period_of_extending FROM %s WHERE reagent_id=$1`, ExtendingTable)
	extending := []*models.Extending{}

	if err := r.db.SelectContext(ctx, &extending, query, reagentId); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return extending, nil
}

func (r *ExtendingRepo) Create(ctx context.Context, dto *models.ExtendingDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, reagent_id, date_of_extending, period_of_extending) VALUES 
		(:id, :reagent_id, :date_of_extending, :period_of_extending)`, ExtendingTable,
	)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *ExtendingRepo) Update(ctx context.Context, dto *models.ExtendingDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET date_of_extending=:date_of_extending, period_of_extending=:period_of_extending WHERE id=:id`, ExtendingTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ExtendingRepo) Delete(ctx context.Context, dto *models.DeleteExtendingDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, ExtendingTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
