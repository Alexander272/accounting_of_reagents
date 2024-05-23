package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SpendingRepo struct {
	db *sqlx.DB
}

func NewSpendingRepo(db *sqlx.DB) *SpendingRepo {
	return &SpendingRepo{
		db: db,
	}
}

type Spending interface {
	GetByReagentId(context.Context, string) ([]*models.Spending, error)
	Create(context.Context, *models.SpendingDTO) (string, error)
	Update(context.Context, *models.SpendingDTO) error
	Delete(context.Context, *models.DeleteSpendingDTO) (float64, error)
}

func (r *SpendingRepo) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Spending, error) {
	query := fmt.Sprintf(`SELECT id, reagent_id, date_of_spending, amount FROM %s WHERE reagent_id=$1 ORDER BY date_of_spending DESC`, SpendingTable)
	spending := []*models.Spending{}

	if err := r.db.SelectContext(ctx, &spending, query, reagentId); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return spending, nil
}

func (r *SpendingRepo) Create(ctx context.Context, dto *models.SpendingDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, reagent_id, date_of_spending, amount) VALUES (:id, :reagent_id, :date_of_spending, :amount)`, SpendingTable)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *SpendingRepo) Update(ctx context.Context, dto *models.SpendingDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET date_of_spending=:date_of_spending, amount=:amount WHERE id=:id`, SpendingTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *SpendingRepo) Delete(ctx context.Context, dto *models.DeleteSpendingDTO) (float64, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING amount`, SpendingTable)

	row := r.db.QueryRowxContext(ctx, query, dto.Id)
	if row.Err() != nil {
		return 0, fmt.Errorf("failed to execute query. error: %w", row.Err())
	}

	var amount float64
	if err := row.Scan(&amount); err != nil {
		return 0, fmt.Errorf("failed to scan result. error: %w", err)
	}
	return amount, nil

	// if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
	// 	return 0, fmt.Errorf("failed to execute query. error: %w", err)
	// }
	// return nil
}
