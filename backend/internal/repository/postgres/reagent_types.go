package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres/pq_models"
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
	query := fmt.Sprintf(`SELECT id, name, number, description, role_id FROM %s ORDER BY role_id, number`, ReagentTypesTable)
	data := []*pq_models.ReagentTypeDTO{}
	reagentTypes := []*models.ReagentType{}

	if err := r.db.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	for _, d := range data {
		reagentTypes = append(reagentTypes, &models.ReagentType{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
			Number:      d.Number,
			RoleId:      d.RoleId,
		})
	}

	return reagentTypes, nil
}

func (r *ReagentTypeRepo) Create(ctx context.Context, dto *models.ReagentTypeDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, number, description, role_id) VALUES (:id, :name, :number, :description, :role_id)`, ReagentTypesTable)
	id := uuid.New()

	data := &pq_models.ReagentTypeDTO{
		Id:          id.String(),
		Name:        dto.Name,
		Description: dto.Description,
		Number:      dto.Number,
		RoleId:      dto.RoleId,
	}

	if _, err := r.db.NamedExecContext(ctx, query, data); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *ReagentTypeRepo) Update(ctx context.Context, dto *models.ReagentTypeDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=:name, role_id=:role_id, number=:number, description=:description WHERE id=:id`, ReagentTypesTable)

	data := &pq_models.ReagentTypeDTO{
		Id:          dto.Id,
		Name:        dto.Name,
		Description: dto.Description,
		Number:      dto.Number,
		RoleId:      dto.RoleId,
	}

	if _, err := r.db.NamedExecContext(ctx, query, data); err != nil {
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
