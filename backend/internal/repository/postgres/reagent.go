package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres/pq_models"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	GetById(context.Context, string) (*models.EditReagent, error)
	GetByIdList(context.Context, []string) ([]*models.Reagent, error)
	GetRemainder(context.Context, string) (*models.ReagentWithRemainder, error)
	GetRemainderNew(context.Context, string) ([]*models.ReagentWithRemainder, error)
	GetAllShelfLife(context.Context) (models.GroupedReagents, error)
	GetUniqueData(context.Context, *models.GetUniqueDTO) ([]string, error)
	Create(context.Context, *models.ReagentDTO) (string, error)
	Update(context.Context, *models.ReagentDTO) error
	SetNotification(context.Context, *models.ReagentNotificationDTO) error
	SetNotificationNew(context.Context, []*models.ReagentNotificationDTO) error
	SetIsOverdue(context.Context, []string) error
	ClearIsOverdue(context.Context, string) error
	SetDeleteStamp(context.Context, string) error
	DeleteOld(context.Context) error
	Delete(context.Context, *models.DeleteReagentDTO) error
}

func (r *ReagentRepo) getColumnName(field string) string {
	columns := map[string]string{
		"type":                        "t.name",
		"typeId":                      "type_id",
		"name":                        "r.name",
		"uname":                       "uname",
		"document":                    "document",
		"purity":                      "purity",
		"dateOfManufacture":           "date_of_manufacture",
		"consignment":                 "consignment",
		"manufacturer":                "manufacturer",
		"shelfLife":                   "shelf_life",
		"place_closet":                "closet",
		"place_shelf":                 "shelf",
		"incomingControl_receiptDate": "receipt_date",
		"incomingControl_amount":      "amount",
		"incomingControl_date":        "control_date",
		"incomingControl_protocol":    "protocol",
		"incomingControl_result":      "result",
		// "spending":"spending",
		"extending_date":      "date_of_extending",
		"extending_period":    "period_of_extending",
		"seizureInformation":  "seizure",
		"disposalInformation": "disposal",
	}

	return columns[field]
}

func (r *ReagentRepo) Get(ctx context.Context, req *models.Params) (*models.ReagentList, error) {
	params := []interface{}{}
	count := 1

	order := " ORDER BY "
	for _, s := range req.Sort {
		order += fmt.Sprintf("%s %s, ", r.getColumnName(s.Field), s.Type)
	}
	order += "r.created_at DESC, r.id"

	filter := "WHERE deleted IS NULL"
	if len(req.Filters) > 0 {
		filter += " AND "
		filters := []string{}

		for _, ns := range req.Filters {
			for _, sv := range ns.Values {
				filters = append(filters, getFilterLine(sv.CompareType, r.getColumnName(ns.Field), count))
				if sv.CompareType == "in" {
					sv.Value = strings.ReplaceAll(sv.Value, ",", "|")
				}
				params = append(params, sv.Value)
				count++
			}
		}
		filter += strings.Join(filters, " AND ")
	}

	search := ""
	if req.Search != nil {
		search = " AND ("

		list := []string{}
		for _, f := range req.Search.Fields {
			list = append(list, fmt.Sprintf("%s ILIKE '%%'||$%d||'%%'", r.getColumnName(f), count))
		}
		params = append(params, req.Search.Value)
		count++
		search += strings.Join(list, " OR ") + ")"
	}

	params = append(params, req.Page.Limit, req.Page.Offset)

	query := fmt.Sprintf(`SELECT r.id, t.name AS type, r.name, uname, document, purity, date_of_manufacture, consignment, manufacturer, shelf_life, closet, shelf,
		receipt_date, (amount || ' ' || a.name) AS amount, control_date, protocol, result, COALESCE(e.date_of_extending, 0) AS date_of_extending, 
		has_run_out, is_overdue,
		COALESCE(e.period_of_extending, 0) AS period_of_extending, seizure, disposal, COALESCE(comment,'') AS comment, COALESCE(note,'') AS note,
		COALESCE((SELECT SUM(amount) FROM %s WHERE reagent_id=r.id GROUP BY reagent_id), 0) AS spending, a.name AS spending_type,
		-- COALESCE((SELECT SUM(period_of_extending) FROM %s WHERE reagent_id=r.id GROUP BY reagent_id), 0) AS sum_period,
		COUNT(*) OVER() as total_count
		FROM %s AS r
		LEFT JOIN %s AS t ON r.type_id=t.id
		LEFT JOIN LATERAL (SELECT date_of_extending, period_of_extending FROM %s WHERE reagent_id=r.id ORDER BY date_of_extending 
			DESC LIMIT 1) AS e ON true
		LEFT JOIN %s AS a ON r.amount_type_id=a.id
		LEFT JOIN %s AS n ON r.id=n.reagent_id
		%s%s%s LIMIT $%d OFFSET $%d`,
		SpendingTable, ExtendingTable, ReagentsTable, ReagentTypesTable, ExtendingTable, AmountTypeTable, NotesTable,
		filter, search, order, count, count+1,
	)
	reagents := []*pq_models.ReagentDTO{}

	// logger.Debug("get", logger.StringAttr("query", query))

	if err := r.db.SelectContext(ctx, &reagents, query, params...); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	list := &models.ReagentList{
		List: []*models.Reagent{},
	}

	if len(reagents) > 0 {
		list.Total = reagents[0].Total

		for _, r := range reagents {
			sp := strconv.FormatFloat(math.Round(r.Spending*10000)/10000, 'f', -1, 64)

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
				Spending:          fmt.Sprintf("%s %s", sp, r.SpendingType),
				DateOfExtending:   r.DateOfExtending,
				Period:            r.Period,
				Seizure:           r.Seizure,
				Disposal:          r.Disposal,
				Comments:          r.Comments,
				Notes:             r.Notes,
				HasRunOut:         r.HasRunOut,
				IsOverdue:         r.IsOverdue,
				// SumPeriod:         r.SumPeriod,
			})
		}
	}

	return list, nil
}

func (r *ReagentRepo) GetById(ctx context.Context, id string) (*models.EditReagent, error) {
	query := fmt.Sprintf(`SELECT id, type_id, name, uname, document, purity, date_of_manufacture, consignment, manufacturer, shelf_life, closet, 
		shelf, receipt_date, amount, amount_type_id, control_date, protocol, result, seizure, disposal FROM %s WHERE id=$1`,
		ReagentsTable,
	)

	reagent := &models.EditReagent{}
	if err := r.db.GetContext(ctx, reagent, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	return reagent, nil
}

func (r *ReagentRepo) GetByIdList(ctx context.Context, list []string) ([]*models.Reagent, error) {
	query := fmt.Sprintf(`SELECT id, name, uname, document, purity, manufacturer FROM %s WHERE id=ANY($1::uuid[])`, ReagentsTable)

	reagents := []*models.Reagent{}
	if err := r.db.SelectContext(ctx, &reagents, query, pq.Array(list)); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return reagents, nil
}

// ! Deprecated
func (r *ReagentRepo) GetRemainder(ctx context.Context, id string) (*models.ReagentWithRemainder, error) {
	/*
	   WITH Reagent AS (
	   	SELECT name, purity
	   	FROM reagents WHERE id='4f4f55a8-e2ab-4c0a-ba9d-04c1331940f0'
	   )

	   SELECT o.name, o.purity, SUM(amount - COALESCE(s.spent,0)) AS remainder FROM reagents AS o
	   INNER JOIN Reagent AS r ON true
	   LEFT JOIN LATERAL(SELECT SUM(amount) AS spent FROM spending WHERE reagent_id=o.id GROUP BY reagent_id) AS s ON true
	   WHERE o.name=r.name AND o.purity=r.purity AND is_overdue=false AND deleted IS NULL
	   GROUP BY o.name, o.purity
	*/

	query := fmt.Sprintf(`SELECT id, amount, name, document, purity, manufacturer, has_notification,
		amount - COALESCE((SELECT SUM(amount) FROM %s WHERE reagent_id=$1 GROUP BY reagent_id), 0) AS remainder
		FROM %s WHERE id=$1`,
		SpendingTable, ReagentsTable,
	)

	remainder := &models.ReagentWithRemainder{}
	if err := r.db.GetContext(ctx, remainder, query, id); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return remainder, nil
}
func (r *ReagentRepo) GetRemainderNew(ctx context.Context, id string) ([]*models.ReagentWithRemainder, error) {
	query := fmt.Sprintf(`WITH Reagent AS (
			SELECT name, purity
			FROM %s WHERE id=$1
		)

		SELECT id, o.name, o.purity, document, manufacturer, has_notification, amount, amount - COALESCE(s.spent,0) AS remainder FROM %s AS o
		INNER JOIN Reagent AS r ON true
		LEFT JOIN LATERAL(SELECT SUM(amount) AS spent FROM %s WHERE reagent_id=o.id GROUP BY reagent_id) AS s ON true
		WHERE o.name=r.name AND o.purity=r.purity AND is_overdue=false AND deleted IS NULL AND has_notification=false`,
		ReagentsTable, ReagentsTable, SpendingTable,
	)
	data := []*models.ReagentWithRemainder{}

	if err := r.db.SelectContext(ctx, &data, query, id); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *ReagentRepo) GetAllShelfLife(ctx context.Context) (models.GroupedReagents, error) {
	query := fmt.Sprintf(`SELECT id, name, document, purity, date_of_manufacture, shelf_life, is_overdue, seizure,
		COALESCE(date_of_extending,0) AS date_of_extending, COALESCE(period_of_extending,0) AS period_of_extending
		FROM %s AS r
		LEFT JOIN LATERAL (SELECT date_of_extending, period_of_extending FROM %s WHERE reagent_id=r.id ORDER BY date_of_extending 
			DESC LIMIT 1) AS e ON true
		WHERE deleted IS NULL
		ORDER BY name, purity, date_of_manufacture DESC`,
		ReagentsTable, ExtendingTable,
	)
	reagents := []*pq_models.ReagentDTO{}

	if err := r.db.SelectContext(ctx, &reagents, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	data := make(map[string][]*models.Reagent, 0)
	for _, r := range reagents {
		data[r.Name] = append(data[r.Name], &models.Reagent{
			Id:                r.Id,
			Name:              r.Name,
			Document:          r.Document,
			Purity:            r.Purity,
			Seizure:           r.Seizure,
			IsOverdue:         r.IsOverdue,
			DateOfManufacture: r.DateOfManufacture,
			ShelfLife:         r.ShelfLife,
			DateOfExtending:   r.DateOfExtending,
			Period:            r.Period,
		})
	}
	return data, nil
}

func (r *ReagentRepo) GetUniqueData(ctx context.Context, req *models.GetUniqueDTO) ([]string, error) {
	req.Field = r.getColumnName(req.Field)

	query := fmt.Sprintf(`SELECT DISTINCT(%s) AS item FROM %s AS r WHERE %s!='' AND %s IS NOT NULL`,
		req.Field, ReagentsTable, req.Field, req.Field,
	)
	tmp := []pq_models.UniqueData{}

	if err := r.db.SelectContext(ctx, &tmp, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	data := []string{}
	for _, v := range tmp {
		data = append(data, v.Item)
	}
	return data, nil
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

// ! Deprecated
func (r *ReagentRepo) SetNotification(ctx context.Context, dto *models.ReagentNotificationDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET has_notification=:has_notification, has_run_out=:has_run_out WHERE id=:id`, ReagentsTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
func (r *ReagentRepo) SetNotificationNew(ctx context.Context, dto []*models.ReagentNotificationDTO) error {
	values := []string{}
	args := []interface{}{}
	for i, v := range dto {
		tmp := []interface{}{v.Id, v.HasNotification, v.HasRunOut}
		args = append(args, tmp...)
		numbers := []string{}
		for j := range tmp {
			numbers = append(numbers, fmt.Sprintf("$%d", i*len(tmp)+j+1))
		}
		values = append(values, fmt.Sprintf("(%s)", strings.Join(numbers, ",")))
	}

	query := fmt.Sprintf(`UPDATE %s AS t SET has_notification=s.has_notification::boolean, has_run_out=s.has_run_out::boolean
		FROM (VALUES %s) AS s(id, has_notification, has_run_out) WHERE t.id=s.id::uuid`,
		ReagentsTable, strings.Join(values, ","),
	)

	logger.Debug("query", logger.StringAttr("query", query))

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ReagentRepo) SetIsOverdue(ctx context.Context, idList []string) error {
	query := fmt.Sprintf(`UPDATE %s SET is_overdue=true WHERE id=ANY($1::uuid[])`, ReagentsTable)

	if _, err := r.db.ExecContext(ctx, query, pq.Array(idList)); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
func (r *ReagentRepo) ClearIsOverdue(ctx context.Context, id string) error {
	query := fmt.Sprintf(`UPDATE %s SET is_overdue=false WHERE id=$1`, ReagentsTable)

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ReagentRepo) SetDeleteStamp(ctx context.Context, id string) error {
	query := fmt.Sprintf(`UPDATE %s SET deleted=$1 WHERE id=$2`, ReagentsTable)

	if _, err := r.db.ExecContext(ctx, query, time.Now(), id); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
func (r *ReagentRepo) DeleteOld(ctx context.Context) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE deleted<$1`, ReagentsTable)

	now := time.Now()
	date := time.Date(now.Year(), now.Month()-1, now.Day(), now.Hour(), 0, 0, 0, now.Location())

	if _, err := r.db.ExecContext(ctx, query, date); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ReagentRepo) Delete(ctx context.Context, dto *models.DeleteReagentDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, ReagentsTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
