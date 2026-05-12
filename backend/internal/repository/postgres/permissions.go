package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PermissionRepo struct {
	db *sqlx.DB
	Transaction
}

func NewPermissionRepo(db *sqlx.DB, tr Transaction) *PermissionRepo {
	return &PermissionRepo{
		db:          db,
		Transaction: tr,
	}
}

type Permissions interface {
	LoadPolicy(ctx context.Context) ([]*models.Permission, error)
	Sync(ctx context.Context, tx Tx, dto []*models.PermissionDTO) error
	GetById(ctx context.Context, id string) (*models.Permission, error)
	GetAll(ctx context.Context) ([]*models.Permission, error)
	GetByRole(ctx context.Context, req *models.GetPermsByRoleDTO) ([]*models.Permission, error)
	GetInheritedByRole(ctx context.Context, roleId string) (map[string]struct{}, error)
	GetRolePermissionsMap(ctx context.Context, roleId string) (map[string]bool, error)
	CountForAll(ctx context.Context, roleToDescendants map[string][]string) (map[string]models.PermsWithCount, error)
	ReplacePermissions(ctx context.Context, tx Tx, roleId string, permissionIDs []string) error
	Create(ctx context.Context, tx Tx, dto *models.PermissionDTO) error
	Delete(ctx context.Context, tx Tx, dto *models.DeletePermissionDTO) error
	DeleteByKeys(ctx context.Context, tx Tx, dto []*models.PermissionDTO) error
}

func (r *PermissionRepo) LoadPolicy(ctx context.Context) ([]*models.Permission, error) {
	query := fmt.Sprintf(`SELECT r.slug, d.id, p.object, p.action
		FROM %s rp
		JOIN %s r ON r.id = rp.role_id
		INNER JOIN %s d ON true
		JOIN %s p ON p.id = rp.permission_id`,
		Tables.RolePermissions, Tables.Roles, Tables.Realms, Tables.Permissions,
	)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	permissions := make([]*models.Permission, 0, 50)
	for rows.Next() {
		item := &models.Permission{}
		if err := rows.Scan(&item.Role, &item.Realm, &item.Object, &item.Action); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		permissions = append(permissions, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return permissions, nil
}

func (r *PermissionRepo) Sync(ctx context.Context, tx Tx, dto []*models.PermissionDTO) error {
	if len(dto) == 0 {
		return nil
	}
	values := []string{}
	args := []interface{}{}

	for _, v := range dto {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
			len(args)+1, len(args)+2, len(args)+3, len(args)+4, len(args)+5,
		))
		args = append(args, uuid.New(), v.Object, v.Action, v.Name, v.Description)
	}

	query := fmt.Sprintf(`INSERT INTO %s (id, object, action, name, description)
			VALUES %s
			ON CONFLICT (object, action) 
			DO UPDATE SET description = EXCLUDED.description, name = EXCLUDED.name`,
		Tables.Permissions, strings.Join(values, ", "),
	)

	_, err := r.getExec(tx).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *PermissionRepo) GetById(ctx context.Context, id string) (*models.Permission, error) {
	query := fmt.Sprintf(`SELECT id, p.object, p.action
		FROM %s p WHERE id=$1`,
		Tables.Permissions,
	)
	data := &models.Permission{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&data.Id, &data.Object, &data.Action)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return data, nil
}

func (r *PermissionRepo) GetByRole(ctx context.Context, req *models.GetPermsByRoleDTO) ([]*models.Permission, error) {
	query := fmt.Sprintf(`SELECT r.slug, p.object, p.action
		FROM %s rp
		JOIN %s r ON r.id = rp.role_id
		JOIN %s p ON p.id = rp.permission_id
		WHERE r.slug = $1`,
		Tables.RolePermissions, Tables.Roles, Tables.Permissions,
	)

	data := make([]*models.Permission, 0, 50)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := &models.Permission{}
		if err := rows.Scan(&item.Id, &item.Role, &item.Object, &item.Action); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		data = append(data, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return data, nil
}

func (r *PermissionRepo) GetAll(ctx context.Context) ([]*models.Permission, error) {
	query := fmt.Sprintf(`SELECT id, object, action, name, description FROM %s ORDER BY object, action`, Tables.Permissions)

	data := make([]*models.Permission, 0, 50)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := &models.Permission{}
		if err := rows.Scan(&item.Id, &item.Object, &item.Action, &item.Name, &item.Description); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		data = append(data, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return data, nil
}

func (r *PermissionRepo) GetRolePermissionsMap(ctx context.Context, roleId string) (map[string]bool, error) {
	query := fmt.Sprintf(`SELECT permission_id FROM %s WHERE role_id = $1`, Tables.RolePermissions)

	rows, err := r.db.QueryContext(ctx, query, roleId)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions: %w", err)
	}
	defer rows.Close()

	result := make(map[string]bool)
	for rows.Next() {
		var permID string
		if err := rows.Scan(&permID); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		result[permID] = true
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return result, nil
}

func (r *PermissionRepo) GetInheritedByRole(ctx context.Context, roleID string) (map[string]struct{}, error) {
	query := fmt.Sprintf(`WITH RECURSIVE sub_roles AS (
			-- Базовый случай: берем только ПРЯМЫХ детей роли $1
			SELECT role_id 
			FROM %s 
			WHERE parent_role_id = $1
			
			UNION ALL
			
			-- Рекурсия: спускаемся дальше ко всем внукам, правнукам и т.д.
			SELECT rh.role_id 
			FROM %s rh
			JOIN sub_roles sr ON rh.parent_role_id = sr.role_id
		)
		SELECT DISTINCT rp.permission_id 
		FROM %s rp 
		WHERE rp.role_id IN (SELECT role_id FROM sub_roles)`,
		Tables.RoleHierarchy, Tables.RoleHierarchy, Tables.RolePermissions,
	)

	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	result := make(map[string]struct{})
	for rows.Next() {
		var permID string
		if err := rows.Scan(&permID); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		result[permID] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return result, nil
}

func (r *PermissionRepo) CountForAll(ctx context.Context, roleToDescendants map[string][]string) (map[string]models.PermsWithCount, error) {
	if len(roleToDescendants) == 0 {
		return make(map[string]models.PermsWithCount), nil
	}

	res := make(map[string]models.PermsWithCount)

	// Для каждой роли считаем её собственные permissions
	for roleSlug := range roleToDescendants {
		c := models.PermsWithCount{}

		// Считаем собственные permissions роли
		ownQuery := fmt.Sprintf(`
			SELECT array_agg(rp.permission_id)
			FROM %s rp
			JOIN %s r ON rp.role_id = r.id
			WHERE r.slug = $1`,
			Tables.RolePermissions, Tables.Roles,
		)
		items := pq.StringArray{}

		err := r.db.QueryRowContext(ctx, ownQuery, roleSlug).Scan(&items)
		if err != nil {
			return nil, fmt.Errorf("failed to count own perms for role %s: %w", roleSlug, err)
		}
		if items != nil {
			c.Own.Items = items
		}
		c.Own.Count = len(c.Own.Items)

		c.Total = c.Own
		res[roleSlug] = c
	}

	// Собираем все уникальные descendant slug'и
	allDescendants := make([]string, 0, len(roleToDescendants))
	descendantSet := make(map[string]struct{})
	for _, descendants := range roleToDescendants {
		for _, d := range descendants {
			if _, exists := descendantSet[d]; !exists {
				descendantSet[d] = struct{}{}
				allDescendants = append(allDescendants, d)
			}
		}
	}

	// Считаем permissions для каждого descendant
	descendantPerms := make(map[string]models.Perm)
	if len(allDescendants) > 0 {
		descQuery := fmt.Sprintf(`
			SELECT r.slug, array_agg(rp.permission_id)
			FROM %s rp
			JOIN %s r ON rp.role_id = r.id
			WHERE r.slug = ANY($1)
			GROUP BY r.slug`,
			Tables.RolePermissions, Tables.Roles,
		)

		rows, err := r.db.QueryContext(ctx, descQuery, pq.Array(allDescendants))
		if err != nil {
			return nil, fmt.Errorf("failed to count descendant perms: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var slug sql.NullString
			var perms pq.StringArray
			if err := rows.Scan(&slug, &perms); err != nil {
				return nil, err
			}
			descendantPerms[slug.String] = models.Perm{Items: perms, Count: len(perms)}
		}
	}

	// Для каждой роли суммируем permissions всех её descendants
	for roleSlug, descendants := range roleToDescendants {
		c := res[roleSlug]
		for _, d := range descendants {
			c.Inherited.Items = append(c.Inherited.Items, descendantPerms[d].Items...)
			c.Inherited.Count += descendantPerms[d].Count
		}
		c.Total = models.Perm{
			Items: append(c.Own.Items, c.Inherited.Items...),
			Count: c.Own.Count + c.Inherited.Count,
		}
		res[roleSlug] = c
	}

	return res, nil
}

func (r *PermissionRepo) Create(ctx context.Context, tx Tx, dto *models.PermissionDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, object, action) VALUES ($1, $2, $3)`,
		Tables.Permissions,
	)
	dto.Id = uuid.NewString()

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.Id, dto.Object, dto.Action)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *PermissionRepo) Delete(ctx context.Context, tx Tx, dto *models.DeletePermissionDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, Tables.Permissions)

	_, err := r.getExec(tx).ExecContext(ctx, query, dto.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *PermissionRepo) DeleteByKeys(ctx context.Context, tx Tx, dto []*models.PermissionDTO) error {
	if len(dto) == 0 {
		return nil
	}

	placeholders := make([]string, 0, len(dto)*2)
	args := make([]interface{}, 0, len(dto)*2)
	for _, v := range dto {
		placeholders = append(placeholders, fmt.Sprintf("($%d::text, $%d::text)", len(args)+1, len(args)+2))
		args = append(args, v.Object, v.Action)
	}

	// НО проще и надежнее использовать расширение unnest или values:
	query := fmt.Sprintf(`DELETE FROM %s 
        WHERE (object, action) NOT IN (
            SELECT * FROM (VALUES %s) AS t(obj, act)
        )`,
		Tables.Permissions,
		strings.Join(placeholders, ","),
	)

	_, err := r.getExec(tx).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *PermissionRepo) ReplacePermissions(ctx context.Context, tx Tx, roleId string, permissionIDs []string) error {
	exec := r.getExec(tx)

	query := fmt.Sprintf(`DELETE FROM %s WHERE role_id = $1`, Tables.RolePermissions)
	_, err := exec.ExecContext(ctx, query, roleId)
	if err != nil {
		return fmt.Errorf("failed to delete old permissions: %w", err)
	}

	if len(permissionIDs) == 0 {
		return nil
	}

	values := make([]string, 0, len(permissionIDs))
	args := make([]interface{}, 0, len(permissionIDs)*2)
	for _, id := range permissionIDs {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d)", len(args)+1, len(args)+2, len(args)+3))
		args = append(args, uuid.New(), roleId, id)
	}

	query = fmt.Sprintf(`INSERT INTO %s (id, role_id, permission_id) VALUES %s`, Tables.RolePermissions, strings.Join(values, ", "))

	_, err = exec.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert permissions: %w", err)
	}

	return nil
}
