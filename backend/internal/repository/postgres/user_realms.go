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

type UserRealmRepo struct {
	db *sqlx.DB
	Transaction
}

func NewUserRealmRepo(db *sqlx.DB, tr Transaction) *UserRealmRepo {
	return &UserRealmRepo{
		db:          db,
		Transaction: tr,
	}
}

type UserRealm interface {
	GetAll(ctx context.Context) ([]*models.UserRealm, error)
	GetByUserId(ctx context.Context, userId string) ([]*models.UserRealm, error)
	GetByUserAndRealm(ctx context.Context, userId, realmId string) (*models.UserRealm, error)
	Create(ctx context.Context, userRealm *models.UserRealmDTO) error
	CreateSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error
	Update(ctx context.Context, userRealm *models.UserRealmDTO) error
	UpdateSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error
	Delete(ctx context.Context, id string) error
	DeleteByUserAndRealm(ctx context.Context, userId, realmId string) error
	DeleteSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error
}

func (r *UserRealmRepo) GetAll(ctx context.Context) ([]*models.UserRealm, error) {
	query := fmt.Sprintf(`SELECT ur.id, ur.user_id, ur.realm_id, ur.role_id, ur.is_active,
		    r.id as role_id, r.slug as role_slug, r.name as role_name, r.level as role_level,
		    rl.id as realm_id, rl.name as realm_name, rl.description as realm_description
		FROM %s ur
		LEFT JOIN %s r ON ur.role_id = r.id
		LEFT JOIN %s rl ON ur.realm_id = rl.id`,
		Tables.UserRealms, Tables.Roles, Tables.Realms,
	)

	var userRealms []*pq_models.UserRealm
	if err := r.db.SelectContext(ctx, &userRealms, query); err != nil {
		return nil, fmt.Errorf("failed to get all user realms. error: %w", err)
	}
	var data []*models.UserRealm

	for _, ur := range userRealms {
		data = append(data, &models.UserRealm{
			Id:       ur.Id,
			UserId:   ur.UserId,
			RealmId:  ur.RealmId,
			IsActive: ur.IsActive,
			Role: &models.Role{
				Id:    ur.RoleId,
				Slug:  ur.RoleSlug,
				Name:  ur.RoleName,
				Level: ur.RoleLevel,
			},
			Realm: &models.Realm{
				Id:          ur.RealmId,
				Name:        ur.RealmName,
				Description: ur.RealmDescription,
			},
		})
	}
	return data, nil
}

func (r *UserRealmRepo) GetByUserId(ctx context.Context, userId string) ([]*models.UserRealm, error) {
	query := fmt.Sprintf(`SELECT ur.id, ur.user_id, ur.realm_id, ur.role_id,
		    r.id as role_id, r.name as role_name, r.level as role_level,
		    rl.id as realm_id, rl.name as realm_name, rl.description as realm_description
		FROM %s ur
		LEFT JOIN %s r ON ur.role_id = r.id
		LEFT JOIN %s rl ON ur.realm_id = rl.id
		WHERE ur.user_id = $1`,
		Tables.UserRealms, Tables.Roles, Tables.Realms,
	)

	var userRealms []*pq_models.UserRealm
	if err := r.db.SelectContext(ctx, &userRealms, query, userId); err != nil {
		return nil, fmt.Errorf("failed to get user realms by user id. error: %w", err)
	}
	var data []*models.UserRealm

	for _, ur := range userRealms {
		data = append(data, &models.UserRealm{
			Id:       ur.Id,
			UserId:   ur.UserId,
			RealmId:  ur.RealmId,
			IsActive: ur.IsActive,
			Role: &models.Role{
				Id:    ur.RoleId,
				Name:  ur.RoleName,
				Level: ur.RoleLevel,
			},
			Realm: &models.Realm{
				Id:          ur.RealmId,
				Name:        ur.RealmName,
				Description: ur.RealmDescription,
			},
		})
	}
	return data, nil
}

func (r *UserRealmRepo) GetByUserAndRealm(ctx context.Context, userId, realmId string) (*models.UserRealm, error) {
	query := fmt.Sprintf(`
		SELECT ur.id, ur.user_id, ur.realm_id, ur.role_id, ur.is_active
		FROM %s ur
		WHERE ur.user_id = $1 AND ur.realm_id = $2`, Tables.UserRealms)

	var data models.UserRealm
	if err := r.db.GetContext(ctx, &data, query, userId, realmId); err != nil {
		return nil, fmt.Errorf("failed to get user realm. error: %w", err)
	}
	return &data, nil
}

func (r *UserRealmRepo) Create(ctx context.Context, dto *models.UserRealmDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, user_id, realm_id, role_id) VALUES ($1, $2, $3, $4)`, Tables.UserRealms)

	if _, err := r.db.ExecContext(ctx, query, dto.Id, dto.UserId, dto.RealmId, dto.RoleId); err != nil {
		return fmt.Errorf("failed to create user realm. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) CreateSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, user_id, realm_id, role_id) VALUES `, Tables.UserRealms)

	args := []interface{}{}
	values := []string{}
	for _, ur := range dto {
		ur.Id = uuid.NewString()
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", len(args)+1, len(args)+2, len(args)+3, len(args)+4))
		args = append(args, ur.Id, ur.UserId, ur.RealmId, ur.RoleId)
	}

	query += strings.Join(values, ", ")

	if _, err := r.getExec(tx).ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to create user realms. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) Update(ctx context.Context, dto *models.UserRealmDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET role_id = $1, is_active = $2 WHERE user_id = $3 AND realm_id = $4`, Tables.UserRealms)

	if _, err := r.db.ExecContext(ctx, query, dto.RoleId, dto.IsActive, dto.UserId, dto.RealmId); err != nil {
		return fmt.Errorf("failed to update user realm. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) UpdateSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error {
	values := []string{}
	args := []interface{}{}
	for i, v := range dto {
		tmp := []interface{}{v.Id, v.UserId, v.RealmId, v.RoleId, v.IsActive}
		args = append(args, tmp...)
		numbers := []string{}
		for j := range tmp {
			numbers = append(numbers, fmt.Sprintf("$%d", i*len(tmp)+j+1))
		}
		values = append(values, fmt.Sprintf("(%s)", strings.Join(numbers, ",")))
	}

	query := fmt.Sprintf(`UPDATE %s AS t SET role_id=s.role_id::uuid, is_active=s.is_active::boolean 
		FROM (VALUES %s) AS s(id, user_id, realm_id, role_id, is_active) 
		WHERE t.user_id=s.user_id::uuid AND t.realm_id=s.realm_id::uuid`,
		Tables.UserRealms, strings.Join(values, ","),
	)

	if _, err := r.getExec(tx).ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, Tables.UserRealms)

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete user realm. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) DeleteByUserAndRealm(ctx context.Context, userId, realmId string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1 AND realm_id = $2`, Tables.UserRealms)

	if _, err := r.db.ExecContext(ctx, query, userId, realmId); err != nil {
		return fmt.Errorf("failed to delete user realm by user and realm. error: %w", err)
	}
	return nil
}

func (r *UserRealmRepo) DeleteSeveral(ctx context.Context, tx Tx, dto []*models.UserRealmDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=ANY($1)`, Tables.UserRealms)

	ids := []interface{}{}
	for _, d := range dto {
		ids = append(ids, d.Id)
	}

	if _, err := r.getExec(tx).ExecContext(ctx, query, ids); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
