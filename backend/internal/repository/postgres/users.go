package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres/pq_models"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type User interface {
	GetByUsername(ctx context.Context, username string) (*models.UserData, error)
	GetById(ctx context.Context, id string) (*models.UserData, error)
	GetAll(ctx context.Context) ([]*models.UserData, error)
	Create(ctx context.Context, user *models.UserDTO) error
	SetActive(ctx context.Context, tx Tx, dto *models.UserDataDTO) error
	Update(ctx context.Context, user *models.UserDTO) error
	Delete(ctx context.Context, id string) error
	SyncFromKeycloak(ctx context.Context, keycloakUsers []*models.KeycloakUser) error
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*models.UserData, error) {
	query := fmt.Sprintf(`SELECT u.id, u.username, u.email, u.first_name, u.last_name, u.created_at, ur.id AS ur_id, ur.is_active,
			r.id AS role_id, r.name AS role_name, r.description AS role_description, r.level AS role_level,
			r.is_active AS role_is_active, r.is_editable AS role_is_editable, r.slug AS role_slug,
			rl.id AS realm_id, rl.name AS realm_name, rl.description AS realm_description, rl.created_at AS realm_created_at
		FROM %s u
		LEFT JOIN %s ur ON u.id = ur.user_id
		LEFT JOIN %s r ON ur.role_id = r.id
		LEFT JOIN %s rl ON ur.realm_id = rl.id
		WHERE username = $1`,
		Tables.Users, Tables.UserRealms, Tables.Roles, Tables.Realms,
	)

	var user []*pq_models.User
	if err := r.db.GetContext(ctx, &user, query, username); err != nil {
		return nil, fmt.Errorf("failed to get user by username. error: %w", err)
	}

	data := mapUsersData(user)
	return data[0], nil
}

func (r *UserRepo) GetById(ctx context.Context, id string) (*models.UserData, error) {
	query := fmt.Sprintf(`SELECT u.id, u.username, u.email, u.first_name, u.last_name, u.created_at, ur.id AS ur_id, ur.is_active,
			r.id AS role_id, r.name AS role_name, r.description AS role_description, r.level AS role_level,
			r.is_active AS role_is_active, r.is_editable AS role_is_editable, r.slug AS role_slug,
			rl.id AS realm_id, rl.name AS realm_name, rl.description AS realm_description, rl.created_at AS realm_created_at
		FROM %s u
		LEFT JOIN %s ur ON u.id = ur.user_id
		LEFT JOIN %s r ON ur.role_id = r.id
		LEFT JOIN %s rl ON ur.realm_id = rl.id
		WHERE u.id = $1`,
		Tables.Users, Tables.UserRealms, Tables.Roles, Tables.Realms,
	)

	var user []*pq_models.User
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("failed to get user by id. error: %w", err)
	}

	data := mapUsersData(user)
	return data[0], nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]*models.UserData, error) {
	query := fmt.Sprintf(`SELECT u.id, u.username, u.email, u.first_name, u.last_name, u.created_at, ur.id AS ur_id, ur.is_active,
			r.id AS role_id, r.name AS role_name, r.description AS role_description, r.level AS role_level,
			r.is_active AS role_is_active, r.is_editable AS role_is_editable, r.slug AS role_slug,
			rl.id AS realm_id, rl.name AS realm_name, rl.description AS realm_description, rl.created_at AS realm_created_at
		FROM %s u
		LEFT JOIN %s ur ON u.id = ur.user_id
		LEFT JOIN %s r ON ur.role_id = r.id
		LEFT JOIN %s rl ON ur.realm_id = rl.id
		ORDER BY u.first_name, u.last_name, realm_name`,
		Tables.Users, Tables.UserRealms, Tables.Roles, Tables.Realms,
	)

	var users []*pq_models.User
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, fmt.Errorf("failed to get all users. error: %w", err)
	}

	data := mapUsersData(users)
	return data, nil
}

func mapUsersData(rows []*pq_models.User) []*models.UserData {
	result := make([]*models.UserData, 0, 10)
	userIndex := make(map[string]int)

	for _, u := range rows {
		var role *models.Role
		if u.RoleId.Valid {
			role = &models.Role{
				Id:          u.RoleId.String,
				Slug:        u.RoleSlug.String,
				Name:        u.RoleName.String,
				Description: u.RoleDescription.String,
				Level:       int(u.RoleLevel.Int64),
				IsActive:    u.RoleIsActive.Bool,
				IsEditable:  u.RoleIsEditable.Bool,
			}
		}

		var realm *models.Realm
		if u.RealmId.Valid {
			realm = &models.Realm{
				Id:          u.RealmId.String,
				Name:        u.RealmName.String,
				Description: u.RealmDescription.String,
			}
		}

		userRealm := &models.UserRealm{
			Id:        u.UserRealmId.String,
			IsActive:  u.IsActive.Bool,
			CreatedAt: u.RealmCreatedAt.Time,
			RealmId:   u.RealmId.String,
			RoleId:    u.RoleId.String,
			Realm:     realm,
			Role:      role,
		}

		if idx, ok := userIndex[u.Id]; ok {
			result[idx].Realms = append(result[idx].Realms, userRealm)
		} else {
			newUser := &models.UserData{
				Id:        u.Id,
				Username:  u.Username,
				Email:     u.Email,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				CreatedAt: u.CreatedAt,
				Realms:    []*models.UserRealm{userRealm},
			}
			userIndex[u.Id] = len(result)
			result = append(result, newUser)
		}
	}

	return result
}

func (r *UserRepo) Create(ctx context.Context, user *models.UserDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, username, email, first_name, last_name)
		VALUES (:id, :username, :email, :first_name, :last_name)`,
		Tables.Users,
	)

	if _, err := r.db.NamedExecContext(ctx, query, user); err != nil {
		return fmt.Errorf("failed to create user. error: %w", err)
	}
	return nil
}

func (r *UserRepo) SetActive(ctx context.Context, tx Tx, dto *models.UserDataDTO) error {
	query := fmt.Sprintf(`UPDATE %s	SET is_active = $1
		WHERE id = $2`,
		Tables.Users,
	)

	_, err := r.db.ExecContext(ctx, query, dto.IsActive, dto.Id)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.UserDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET username = :username, email = :email, first_name = :first_name,
		last_name = :last_name, updated_at = now() WHERE id = :id`,
		Tables.Users,
	)

	if _, err := r.db.NamedExecContext(ctx, query, user); err != nil {
		return fmt.Errorf("failed to update user. error: %w", err)
	}
	return nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, Tables.Users)

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete user. error: %w", err)
	}
	return nil
}

// SyncFromKeycloak syncs users from Keycloak to local database
func (r *UserRepo) SyncFromKeycloak(ctx context.Context, keycloakUsers []*models.KeycloakUser) error {
	// Get existing users
	existingUsers := map[string]*models.UserDTO{}
	query := fmt.Sprintf(`SELECT id, username, email, first_name, last_name FROM %s`, Tables.Users)
	var users []*models.UserDTO
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return fmt.Errorf("failed to get existing users. error: %w", err)
	}
	for _, u := range users {
		existingUsers[u.Username] = u
	}

	// Sync users
	for _, ku := range keycloakUsers {
		username := ""
		if ku.Username != nil {
			username = *ku.Username
		}
		email := ""
		if ku.Email != nil {
			email = *ku.Email
		}
		firstName := ""
		if ku.FirstName != nil {
			firstName = *ku.FirstName
		}
		lastName := ""
		if ku.LastName != nil {
			lastName = *ku.LastName
		}

		existing, exists := existingUsers[username]
		if exists {
			// Update if needed
			if existing.Email != email || existing.FirstName != firstName || existing.LastName != lastName {
				r.Update(ctx, &models.UserDTO{
					Id:        existing.Id,
					Username:  username,
					Email:     email,
					FirstName: firstName,
					LastName:  lastName,
				})
			}
		} else {
			// Create new user
			id := ""
			if ku.Id != nil {
				id = *ku.Id
			}
			r.Create(ctx, &models.UserDTO{
				Id:        id,
				Username:  username,
				Email:     email,
				FirstName: firstName,
				LastName:  lastName,
			})
		}
	}

	return nil
}
