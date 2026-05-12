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
	Get(context.Context, *models.GetReagentTypeDTO) ([]*models.ReagentType, error)
	Create(context.Context, *models.ReagentTypeDTO) (string, error)
	Update(context.Context, *models.ReagentTypeDTO) error
	Delete(context.Context, *models.DeleteReagentTypeDTO) error
}

func (r *ReagentTypeRepo) Get(ctx context.Context, req *models.GetReagentTypeDTO) ([]*models.ReagentType, error) {
	condition := ""
	args := []interface{}{}

	if req.IsPublic {
		args = append(args, req.IsPublic)
		condition = fmt.Sprintf("WHERE is_public = $%d", len(args))
	}

	query := fmt.Sprintf(`SELECT id, name, number, description, is_public 
		FROM %s %s ORDER BY number`,
		Tables.ReagentTypes, condition,
	)
	data := []*models.ReagentType{}

	if err := r.db.SelectContext(ctx, &data, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	return data, nil
}

func (r *ReagentTypeRepo) Create(ctx context.Context, dto *models.ReagentTypeDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, number, description, is_public) VALUES (:id, :name, :number, :description, :is_public)`,
		Tables.ReagentTypes,
	)
	id := uuid.New()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *ReagentTypeRepo) Update(ctx context.Context, dto *models.ReagentTypeDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=:name, is_public=:is_public, number=:number, description=:description WHERE id=:id`,
		Tables.ReagentTypes,
	)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ReagentTypeRepo) Delete(ctx context.Context, dto *models.DeleteReagentTypeDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, Tables.ReagentTypes)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
