package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AmountTypeRepo struct {
	db *sqlx.DB
}

func NewAmountTypeRepo(db *sqlx.DB) *AmountTypeRepo {
	return &AmountTypeRepo{
		db: db,
	}
}

type AmountType interface {
	GetAll(context.Context) ([]*models.AmountType, error)
	Create(context.Context, *models.AmountTypeDTO) (string, error)
	Update(context.Context, *models.AmountTypeDTO) error
	UpdateSeveral(context.Context, []*models.AmountTypeDTO) error
	Delete(context.Context, *models.DeleteAmountTypeDTO) error
}

func (r *AmountTypeRepo) GetAll(ctx context.Context) ([]*models.AmountType, error) {
	query := fmt.Sprintf(`SELECT id, name, number FROM %s`, AmountTypeTable)
	amountTypes := []*models.AmountType{}

	if err := r.db.Select(&amountTypes, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return amountTypes, nil
}

func (r *AmountTypeRepo) Create(ctx context.Context, dto *models.AmountTypeDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, number) VALUES (:id, :name, :number)`, AmountTypeTable)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *AmountTypeRepo) Update(ctx context.Context, dto *models.AmountTypeDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=:name WHERE id=:id`, AmountTypeTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AmountTypeRepo) UpdateSeveral(ctx context.Context, dto []*models.AmountTypeDTO) error {
	//TODO надо это протестировать
	values := []string{}
	args := make(map[string]interface{})
	for i, v := range dto {
		args[fmt.Sprintf("id_%d", i)] = v.Id
		args[fmt.Sprintf("number_%d", i)] = v.Number

		values = append(values, fmt.Sprintf("(:id_%d, :number_%d)", i, i))
	}

	query := fmt.Sprintf(`UPDATE %s AS t SET number=s.number FROM (VALUES %s) AS s(id, number) WHERE t.id=s.id`,
		AmountTypeTable, strings.Join(values, ","),
	)

	if _, err := r.db.NamedExecContext(ctx, query, args); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AmountTypeRepo) Delete(ctx context.Context, dto *models.DeleteAmountTypeDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, AmountTypeTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
