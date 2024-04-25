package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres/pq_models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReagentRepo struct {
	db *sqlx.DB
}

func NewReagentRepo(db *sqlx.DB) *ReagentRepo {
	return &ReagentRepo{
		db: db,
	}
}

type Reagent interface {
	Get(context.Context, *models.Params) (*models.ReagentList, error)
	Create(context.Context, *models.ReagentDTO) (string, error)
	Update(context.Context, *models.ReagentDTO) error
}

func (r *ReagentRepo) Get(ctx context.Context, req *models.Params) (*models.ReagentList, error) {
	params := []interface{}{}
	count := 1

	order := " ORDER BY "
	for _, s := range req.Sort {
		order += fmt.Sprintf("%s %s, ", s.Field, s.Type)
	}
	order += "created_at, id"

	filter := ""
	if len(req.Filters) > 0 {
		filter = "WHERE "
		filters := []string{}

		for _, ns := range req.Filters {
			for _, sv := range ns.Values {
				filters = append(filters, getFilterLine(sv.CompareType, ns.Field, count))
				if sv.CompareType == "in" {
					sv.Value = strings.ReplaceAll(sv.Value, ",", "|")
				}
				params = append(params, sv.Value)
				count++
			}
		}
		filter += strings.Join(filters, " AND ")
	}

	params = append(params, req.Page.Limit, req.Page.Offset)

	//TODO решить как получать comments и notes
	query := fmt.Sprintf(`SELECT r.id, t.name AS type, r.name, uname, document, purity, date_of_manufacture, consignment, manufacturer, shelf_life, closet, shelf,
		receipt_date, (amount || ' ' || a.name) AS amount, control_date, protocol, result, e.date_of_extending, e.period_of_extending, seizure, disposal,
		(SELECT SUM(amount) FROM %s WHERE reagent_id=r.id GROUP BY reagent_id) AS spending,
		COUNT(*) OVER() as total_count
		FROM %s AS r
		LEFT JOIN %s AS t ON r.type_id=t.id
		LEFT JOIN LATERAL (SELECT date_of_extending, period_of_extending FROM %s WHERE reagent_id=r.id ORDER BY date_of_extending 
			DESC LIMIT 1) AS e ON true
		LEFT JOIN %s AS a ON r.amount_type_id=a.id
		%s%s LIMIT $%d OFFSET $%d`,
		SpendingTable, ReagentsTable, ReagentTypesTable, ExtendingTable, AmountTypeTable,
		filter, order, count, count+1,
	)
	reagents := []*pq_models.ReagentDTO{}

	if err := r.db.SelectContext(ctx, &reagents, query, params...); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	list := &models.ReagentList{
		List: []*models.Reagent{},
	}

	if len(reagents) > 0 {
		list.Total = reagents[0].Total

		for _, r := range reagents {
			list.List = append(list.List, &models.Reagent{
				Id:                r.Id,
				Type:              r.Type,
				Name:              r.Name,
				Uname:             r.Uname,
				Document:          r.Document,
				Purity:            r.Purity,
				DateOfManufacture: r.DateOfManufacture,
				Consignment:       r.Consignment,
				Manufacturer:      r.Manufacturer,
				ShelfLife:         r.ShelfLife,
				Closet:            r.Closet,
				Shelf:             r.Shelf,
				ReceiptDate:       r.ReceiptDate,
				Amount:            r.Amount,
				ControlDate:       r.ControlDate,
				Protocol:          r.Protocol,
				Result:            r.Result,
				Spending:          r.Spending,
				DateOfExtending:   r.DateOfExtending,
				Period:            r.Period,
				Seizure:           r.Seizure,
				Disposal:          r.Disposal,
				Comments:          r.Comments,
				Notes:             r.Notes,
			})
		}
	}

	return list, nil
}

func (r *ReagentRepo) GetById(ctx context.Context, id string) (*models.Reagent, error) {

	return nil, fmt.Errorf("not implemented")
}

func (r *ReagentRepo) Create(ctx context.Context, dto *models.ReagentDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, type_id, name, uname, document, purity, date_of_manufacture, consignment, manufacturer, shelf_life, closet, 
		shelf, receipt_date, amount, amount_type_id, control_date, protocol, result, seizure, disposal) VALUES (:id, :type_id, :name, :uname, :document, :purity,
		:date_of_manufacture, :consignment, :manufacturer, :shelf_life, :closet, :shelf, :receipt_date, :amount, :amount_type_id, :control_date, :protocol,
		:result, :seizure, :disposal)`,
		ReagentsTable,
	)
	id := uuid.New()
	dto.Id = id.String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return id.String(), nil
}

func (r *ReagentRepo) Update(ctx context.Context, dto *models.ReagentDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET type_id=:type_id, name=:name, uname=:uname, document=:document, purity=:purity, date_of_manufacture=:date_of_manufacture,
		consignment=:consignment, manufacturer=:manufacturer, shelf_life=:shelf_life, closet=:closet, shelf=:shelf, receipt_date=:receipt_date, amount=:amount,
		amount_type_id=:amount_type_id, control_date=:control_date, protocol=:protocol, result=:result, seizure=:seizure, disposal=:disposal
		WHERE id=:id`, ReagentsTable,
	)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}