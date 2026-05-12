package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type RealmRepo struct {
	db *sqlx.DB
}

func NewRealmRepo(db *sqlx.DB) *RealmRepo {
	return &RealmRepo{
		db: db,
	}
}

type Realm interface {
	GetAll(ctx context.Context) ([]*models.Realm, error)
	GetById(ctx context.Context, id string) (*models.Realm, error)
	Create(ctx context.Context, realm *models.RealmDTO) error
	Update(ctx context.Context, realm *models.RealmDTO) error
	Delete(ctx context.Context, id string) error
}

func (r *RealmRepo) GetAll(ctx context.Context) ([]*models.Realm, error) {
	query := fmt.Sprintf(`SELECT id, name, slug, description, is_active, created_at FROM %s ORDER BY name`, Tables.Realms)

	var realms []*models.Realm
	if err := r.db.SelectContext(ctx, &realms, query); err != nil {
		return nil, fmt.Errorf("failed to get all realms. error: %w", err)
	}
	return realms, nil
}

func (r *RealmRepo) GetById(ctx context.Context, id string) (*models.Realm, error) {
	query := fmt.Sprintf(`SELECT id, name, slug, description, is_active, created_at FROM %s WHERE id = $1`, Tables.Realms)

	var realm models.Realm
	if err := r.db.GetContext(ctx, &realm, query, id); err != nil {
		return nil, fmt.Errorf("failed to get realm by id. error: %w", err)
	}
	return &realm, nil
}

func (r *RealmRepo) Create(ctx context.Context, realm *models.RealmDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, slug, description, is_active) VALUES ($1, $2, $3, $4, $5)`, Tables.Realms)

	if _, err := r.db.ExecContext(ctx, query, realm.Id, realm.Name, realm.Slug, realm.Description, realm.IsActive); err != nil {
		return fmt.Errorf("failed to create realm. error: %w", err)
	}
	return nil
}

func (r *RealmRepo) Update(ctx context.Context, realm *models.RealmDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1, slug = $2, description = $3, is_active = $4  WHERE id = $5`, Tables.Realms)

	if _, err := r.db.ExecContext(ctx, query, realm.Name, realm.Slug, realm.Description, realm.IsActive, realm.Id); err != nil {
		return fmt.Errorf("failed to update realm. error: %w", err)
	}
	return nil
}

func (r *RealmRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, Tables.Realms)

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete realm. error: %w", err)
	}
	return nil
}
