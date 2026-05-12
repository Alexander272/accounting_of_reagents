package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RoleRepo struct {
	db *sqlx.DB
	Transaction
}

func NewRoleRepo(db *sqlx.DB, tr Transaction) *RoleRepo {
	return &RoleRepo{
		db:          db,
		Transaction: tr,
	}
}

type Roles interface {
	GetOne(ctx context.Context, req *models.GetRoleDTO) (*models.Role, error)
	GetAll(ctx context.Context) ([]*models.Role, error)
	GetUserCount(ctx context.Context, req []string) (map[string]int, error)
	IsExists(ctx context.Context, roleName string) (bool, error)
	IsExistsById(ctx context.Context, Id string) (bool, error)
	GetIdsBySlugs(ctx context.Context, slugs []string) (map[string]string, error)
	Create(ctx context.Context, tx Tx, dto *models.RoleDTO) error
	Update(ctx context.Context, tx Tx, dto *models.RoleDTO) error
	Delete(ctx context.Context, tx Tx, dto *models.DeleteRoleDTO) error

	AssignPermission(ctx context.Context, tx Tx, dto *models.RolePermissionDTO) error
	AssignPermissions(ctx context.Context, tx Tx, roleId string, permissionIds []string) error
	DeletePermission(ctx context.Context, tx Tx, dto *models.RolePermissionDTO) error
}

func (r *RoleRepo) GetOne(ctx context.Context, req *models.GetRoleDTO) (*models.Role, error) {
	condition := ""
	params := []interface{}{}
	if req.Id != "" {
		params = append(params, req.Id)
		condition = fmt.Sprintf("WHERE Id = $%d", len(params))
	}
	if req.Slug != "" {
		params = append(params, req.Slug)
		condition = fmt.Sprintf("WHERE slug = $%d", len(params))
	}
	if condition == "" {
		return nil, models.ErrInvalidInput
	}

	query := fmt.Sprintf(`SELECT id, slug, name, description, level, is_active, is_system, is_editable, created_at, updated_at FROM %s %s`,
		Tables.Roles, condition,
	)
	data := &models.Role{}

	err := r.db.QueryRowContext(ctx, query, params...).Scan(
		&data.Id,
		&data.Slug,
		&data.Name,
		&data.Description,
		&data.Level,
		&data.IsActive,
		&data.IsSystem,
		&data.IsEditable,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return data, nil
}

func (r *RoleRepo) GetIdsBySlugs(ctx context.Context, slugs []string) (map[string]string, error) {
	if len(slugs) == 0 {
		return make(map[string]string), nil
	}

	placeholders := make([]string, 0, len(slugs))
	args := make([]interface{}, 0, len(slugs))
	for i, slug := range slugs {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		args = append(args, slug)
	}

	query := fmt.Sprintf(`SELECT slug, Id FROM %s WHERE slug IN (%s)`,
		Tables.Roles, strings.Join(placeholders, ", "))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	result := make(map[string]string)
	for rows.Next() {
		var slug string
		var Id string
		if err := rows.Scan(&slug, &Id); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result[slug] = Id
	}

	return result, nil
}

func (r *RoleRepo) GetAll(ctx context.Context) ([]*models.Role, error) {
	query := fmt.Sprintf(`SELECT id, slug, name, description, level, is_active, is_system, is_editable, created_at, updated_at FROM %s 
		ORDER BY level`,
		Tables.Roles,
	)
	data := []*models.Role{}

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := &models.Role{}
		if err := rows.Scan(
			&item.Id, &item.Slug, &item.Name, &item.Description, &item.Level, &item.IsActive, &item.IsSystem, &item.IsEditable,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		data = append(data, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return data, nil
}

func (r *RoleRepo) GetUserCount(ctx context.Context, req []string) (map[string]int, error) {
	// Если входящий список пуст, сразу возвращаем пустую мапу
	if len(req) == 0 {
		return make(map[string]int), nil
	}

	query := fmt.Sprintf(`SELECT role_id, COUNT(*) FROM %s 
		WHERE role_id = ANY($1)
		GROUP BY role_id`,
		Tables.UserRealms,
	)

	// Передаем слайс Ids (в зависимости от драйвера может понадобиться pq.Array)
	rows, err := r.db.QueryContext(ctx, query, pq.Array(req))
	if err != nil {
		return nil, fmt.Errorf("failed to get user counts: %w", err)
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var roleId string
		var count int
		if err := rows.Scan(&roleId, &count); err != nil {
			return nil, fmt.Errorf("scan count error: %w", err)
		}
		counts[roleId] = count
	}

	return counts, nil
}

func (r *RoleRepo) IsExists(ctx context.Context, roleName string) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE name = $1)`, Tables.Roles)
	var exists bool

	err := r.db.QueryRowContext(ctx, query, roleName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return exists, nil
}
func (r *RoleRepo) IsExistsById(ctx context.Context, id string) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1 AND is_active = true)`, Tables.Roles)
	var exists bool

	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return exists, nil
}

func (r *RoleRepo) Create(ctx context.Context, tx Tx, dto *models.RoleDTO) error {
	if dto.Slug == "root" || dto.Slug == "superadmin" {
		return models.ErrReservedRole
	}

	query := fmt.Sprintf(`INSERT INTO %s (id, slug, name, level, description, is_system)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING Id, created_at`,
		Tables.Roles,
	)

	err := r.getExec(tx).QueryRowxContext(
		ctx, query, uuid.NewString(), dto.Slug, dto.Name, dto.Level, dto.Description, dto.IsSystem,
	).Scan(&dto.Id, &dto.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *RoleRepo) Update(ctx context.Context, tx Tx, dto *models.RoleDTO) error {
	if dto.Slug == "root" || dto.Slug == "superadmin" {
		return models.ErrReservedRole
	}

	query := fmt.Sprintf(`UPDATE %s SET name=$1, level=$2, slug=$3, description=$4, updated_at=NOW() WHERE id=$5`,
		Tables.Roles,
	)

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.Name, dto.Level, dto.Slug, dto.Description, dto.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *RoleRepo) Delete(ctx context.Context, tx Tx, dto *models.DeleteRoleDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 AND NOT is_system`, Tables.Roles)

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *RoleRepo) AssignPermission(ctx context.Context, tx Tx, dto *models.RolePermissionDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (role_Id, permission_Id) VALUES ($1, $2)`, Tables.RolePermissions)

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.RoleId, dto.PermissionId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *RoleRepo) AssignPermissions(ctx context.Context, tx Tx, roleId string, permissionIds []string) error {
	if len(permissionIds) == 0 {
		return nil
	}

	values := make([]string, 0, len(permissionIds))
	args := make([]interface{}, 0, len(permissionIds)*2)
	for i, permId := range permissionIds {
		values = append(values, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		args = append(args, roleId, permId)
	}

	query := fmt.Sprintf(`INSERT INTO %s (role_Id, permission_Id) VALUES %s`, Tables.RolePermissions, strings.Join(values, ", "))

	_, err := r.getExec(tx).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *RoleRepo) DeletePermission(ctx context.Context, tx Tx, dto *models.RolePermissionDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE role_Id=$1 AND permission_Id=$2`, Tables.RolePermissions)

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.RoleId, dto.PermissionId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
