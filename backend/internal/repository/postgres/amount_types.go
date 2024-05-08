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
	CreateSeveral(context.Context, []*models.AmountTypeDTO) error
	Update(context.Context, *models.AmountTypeDTO) error
	UpdateSeveral(context.Context, []*models.AmountTypeDTO) error
	Delete(context.Context, *models.DeleteAmountTypeDTO) error
}

func (r *AmountTypeRepo) GetAll(ctx context.Context) ([]*models.AmountType, error) {
	query := fmt.Sprintf(`SELECT id, name, description, number FROM %s ORDER BY number`, AmountTypeTable)
	amountTypes := []*models.AmountType{}

	if err := r.db.SelectContext(ctx, &amountTypes, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return amountTypes, nil
}

func (r *AmountTypeRepo) Create(ctx context.Context, dto *models.AmountTypeDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, description, number) VALUES (:id, :name, :description, :number)`, AmountTypeTable)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *AmountTypeRepo) CreateSeveral(ctx context.Context, dto []*models.AmountTypeDTO) error {
	values := []string{}
	args := []interface{}{}
	c := 4
	for i, v := range dto {
		id := uuid.New()
		args = append(args, id, v.Name, v.Description, v.Number)
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", c*i+1, c*i+2, c*i+3, c*i+4))
	}
	query := fmt.Sprintf(`INSERT INTO %s (id, name, number) VALUES %s`, AmountTypeTable, strings.Join(values, ","))

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AmountTypeRepo) Update(ctx context.Context, dto *models.AmountTypeDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=:name, description=:description WHERE id=:id`, AmountTypeTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AmountTypeRepo) UpdateSeveral(ctx context.Context, dto []*models.AmountTypeDTO) error {
	values := []string{}
	args := []interface{}{}
	for i, v := range dto {
		args = append(args, v.Id, v.Number)
		values = append(values, fmt.Sprintf("($%d, $%d::integer)", (2*i)+1, (2*i)+2))
	}

	query := fmt.Sprintf(`UPDATE %s AS t SET number=s.number FROM (VALUES %s) AS s(id, number) WHERE t.id=s.id::uuid`,
		AmountTypeTable, strings.Join(values, ","),
	)

	// не работает (NamedExecContext похоже не проставляет значения)
	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
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
