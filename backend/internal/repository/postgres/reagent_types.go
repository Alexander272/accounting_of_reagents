package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReagentTypeRepo struct {
	db *sqlx.DB
}

func NewReagentTypeRepo(db *sqlx.DB) *ReagentTypeRepo {
	return &ReagentTypeRepo{
		db: db,
	}
}

type ReagentType interface {
	GetAll(ctx context.Context) ([]*models.ReagentType, error)
	Create(context.Context, *models.ReagentTypeDTO) (string, error)
	Update(context.Context, *models.ReagentTypeDTO) error
	Delete(context.Context, *models.DeleteReagentTypeDTO) error
}

func (r *ReagentTypeRepo) GetAll(ctx context.Context) ([]*models.ReagentType, error) {
	query := fmt.Sprintf(`SELECT id, name, number, role_id FROM %s ORDER BY role_id, number`, ReagentTypesTable)
	reagentTypes := []*models.ReagentType{}

	if err := r.db.SelectContext(ctx, &reagentTypes, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return reagentTypes, nil
}

func (r *ReagentTypeRepo) Create(ctx context.Context, dto *models.ReagentTypeDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, number, role_id) VALUES (:id, :name, :number, :role_id)`, ReagentTypesTable)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *ReagentTypeRepo) Update(ctx context.Context, dto *models.ReagentTypeDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=:name, role_id=:role_id, number=:number WHERE id=:id`, ReagentTypesTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ReagentTypeRepo) Delete(ctx context.Context, dto *models.DeleteReagentTypeDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, ReagentTypesTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
