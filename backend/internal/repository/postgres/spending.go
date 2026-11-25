package postgres

import (
	"context"
	"fmt"
	"math"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	GetReport(context.Context, *models.SpendingReportParams) ([]*models.SpendingReportItem, error)
	Create(context.Context, *models.SpendingDTO) (string, error)
	Update(context.Context, *models.SpendingDTO) error
	Delete(context.Context, *models.DeleteSpendingDTO) (float64, error)
}

func (r *SpendingRepo) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Spending, error) {
	query := fmt.Sprintf(`SELECT id, reagent_id, date_of_spending, amount FROM %s WHERE reagent_id=$1 ORDER BY date_of_spending DESC`, Tables.Spending)
	spending := []*models.Spending{}

	if err := r.db.SelectContext(ctx, &spending, query, reagentId); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	for i := range spending {
		spending[i].Amount = math.Round(spending[i].Amount*100000) / 100000
	}

	return spending, nil
}

func (r *SpendingRepo) GetReport(ctx context.Context, params *models.SpendingReportParams) ([]*models.SpendingReportItem, error) {
	report := []*models.SpendingReportItem{}
	if len(params.TypeIds) == 0 {
		return report, nil
	}

	query := fmt.Sprintf(`WITH spent AS (
			SELECT reagent_id,
				COALESCE(SUM(amount) FILTER (WHERE date_of_spending < $1), 0) AS before,
				COALESCE(SUM(amount) FILTER (WHERE date_of_spending BETWEEN $1 AND $2), 0) AS in_period,
				COALESCE(SUM(amount) FILTER (WHERE date_of_spending <= $2), 0) AS total
			FROM %s
			GROUP BY reagent_id
		)
		SELECT r.name, COALESCE(r.uname, '') AS uname, COALESCE(r.consignment, '') AS consignment,
			COALESCE(a.name, '') AS unit,
			r.amount - s.before AS amount_at_start, s.in_period AS spent, r.amount - s.total AS remainder
		FROM %s AS r
		INNER JOIN spent AS s ON s.reagent_id = r.id
		LEFT JOIN %s AS a ON r.amount_type_id = a.id
		WHERE s.in_period > 0 AND r.deleted IS NULL AND r.realm_id = $3 AND r.type_id = ANY($4::uuid[])
		ORDER BY r.name, r.consignment`,
		Tables.Spending, Tables.Reagents, Tables.AmountType,
	)

	if err := r.db.SelectContext(ctx, &report, query, params.From, params.To, params.RealmId, pq.Array(params.TypeIds)); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	for i := range report {
		report[i].Amount = math.Round(report[i].Amount*100000) / 100000
		report[i].Spent = math.Round(report[i].Spent*100000) / 100000
		report[i].Remainder = math.Round(report[i].Remainder*100000) / 100000
	}

	return report, nil
}

func (r *SpendingRepo) Create(ctx context.Context, dto *models.SpendingDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, reagent_id, date_of_spending, amount) VALUES (:id, :reagent_id, :date_of_spending, :amount)`, Tables.Spending)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *SpendingRepo) Update(ctx context.Context, dto *models.SpendingDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET date_of_spending=:date_of_spending, amount=:amount WHERE id=:id`, Tables.Spending)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *SpendingRepo) Delete(ctx context.Context, dto *models.DeleteSpendingDTO) (float64, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING amount`, Tables.Spending)

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
