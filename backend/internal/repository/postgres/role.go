package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RoleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

type Role interface {
	GetAll(context.Context, models.GetRolesDTO) ([]models.RoleFull, error)
	Get(context.Context, string) (*models.Role, error)
	Create(context.Context, models.RoleDTO) error
	Update(context.Context, models.RoleDTO) error
	Delete(context.Context, string) error
}

func (r *RoleRepo) GetAll(ctx context.Context, req models.GetRolesDTO) (roles []models.RoleFull, err error) {
	var data []models.RoleFullDTO
	query := fmt.Sprintf(`SELECT id, name, level, description, COALESCE(extends, '{}') AS extends 
		FROM %s WHERE is_show=true ORDER BY level, name`,
		RoleTable,
	)

	if err := r.db.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	for _, rfd := range data {
		roles = append(roles, models.RoleFull{
			Id:          rfd.Id,
			Name:        rfd.Name,
			Level:       rfd.Level,
			Extends:     rfd.Extends,
			Description: rfd.Description,
		})
	}

	return roles, nil
}

func (r *RoleRepo) Get(ctx context.Context, roleName string) (*models.Role, error) {
	var data []models.RoleWithMenuDTO
	query := fmt.Sprintf(`SELECT r.id, r.name, COALESCE(extends, '{}') AS extends, i.name AS menu
		FROM %s AS r 
		LEFT JOIN %s AS m ON r.id=role_id 
		LEFT JOIN %s AS i ON menu_item_id=i.id
		WHERE i.is_show=true ORDER BY level`,
		RoleTable, MenuTable, MenuItemTable,
	)

	if err := r.db.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("failed to get roles with menu. error: %w", err)
	}

	role := &models.Role{}
	menu := make(map[string][]string, 0)
	extends := make(map[string]struct{})

	//EDIT Возможно можно это как-то покрасивее написать
	for _, r := range data {
		m, exist := menu[r.Id]
		if !exist {
			menu[r.Id] = []string{r.Menu}

			if r.Name == roleName {
				role.Id = r.Id
				role.Name = r.Name
				extends[r.Id] = struct{}{}
				for _, v := range r.Extends {
					extends[v] = struct{}{}
				}
			}
		} else {
			m = append(m, r.Menu)
			menu[r.Id] = m
		}
	}

	for i := 1; i < len(extends); i++ {
		for _, r := range data {
			_, exist := extends[r.Id]
			if exist {
				for _, v := range r.Extends {
					extends[v] = struct{}{}
				}
				break
			}
		}
	}

	for k := range extends {
		role.Menu = append(role.Menu, menu[k]...)
	}

	return role, nil
}

func (r *RoleRepo) Create(ctx context.Context, role models.RoleDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, name, level, extends, description) VALUES ($1, $2, $3, $4, $5)`, RoleTable)
	id := uuid.New()

	_, err := r.db.ExecContext(ctx, query, id, role.Name, role.Level, pq.Array(role.Extends))
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *RoleRepo) Update(ctx context.Context, role models.RoleDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name=$1, level=$2, extends=$3, description=$4 WHERE id=$5`, RoleTable)

	_, err := r.db.ExecContext(ctx, query, role.Name, role.Level, pq.Array(role.Extends), role.Description, role.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *RoleRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, RoleTable)

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
