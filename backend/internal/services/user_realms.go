package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/events"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres"
	"github.com/google/uuid"
)

type UserRealmService struct {
	repo     repository.UserRealm
	realm    Realm
	roles    Roles
	eventBus *events.PolicyEventManager
}

func NewUserRealmService(repo repository.UserRealm, realm Realm, roles Roles, eventBus *events.PolicyEventManager) *UserRealmService {
	return &UserRealmService{
		repo:     repo,
		realm:    realm,
		roles:    roles,
		eventBus: eventBus,
	}
}

type UserRealm interface {
	GetAll(ctx context.Context) ([]*models.UserRealm, error)
	GetByUserId(ctx context.Context, userId string) ([]*models.UserRealm, error)
	GetByUserAndRealm(ctx context.Context, userId, realmId string) (*models.UserRealm, error)
	Create(ctx context.Context, userRealm *models.UserRealmDTO) error
	CreateSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error
	Update(ctx context.Context, userRealm *models.UserRealmDTO) error
	UpdateSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error
	Delete(ctx context.Context, id string) error
	DeleteByUserAndRealm(ctx context.Context, userId, realmId string) error
	DeleteSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error
}

func (s *UserRealmService) GetAll(ctx context.Context) ([]*models.UserRealm, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user realms. error: %w", err)
	}
	return data, nil
}

func (s *UserRealmService) GetByUserId(ctx context.Context, userId string) ([]*models.UserRealm, error) {
	data, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user realms by user id. error: %w", err)
	}
	return data, nil
}

func (s *UserRealmService) GetByUserAndRealm(ctx context.Context, userId, realmId string) (*models.UserRealm, error) {
	data, err := s.repo.GetByUserAndRealm(ctx, userId, realmId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user realm by user id and realm id. error: %w", err)
	}
	return data, nil
}

func (s *UserRealmService) Create(ctx context.Context, userRealm *models.UserRealmDTO) error {
	if userRealm.Id == "" {
		userRealm.Id = uuid.NewString()
	}

	// Verify realm exists
	_, err := s.realm.GetById(ctx, userRealm.RealmId)
	if err != nil {
		return fmt.Errorf("realm not found: %w", err)
	}

	// Verify role exists (assuming Role interface has Get method)
	_, err = s.roles.GetOne(ctx, &models.GetRoleDTO{Id: userRealm.RoleId})
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}

	if err := s.repo.Create(ctx, userRealm); err != nil {
		return fmt.Errorf("failed to create user realm. error: %w", err)
	}

	s.eventBus.Notify(events.PolicyEvent{})
	return nil
}

func (s *UserRealmService) CreateSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error {
	if len(dto) == 0 {
		return nil
	}
	if err := s.repo.CreateSeveral(ctx, tx, dto); err != nil {
		return fmt.Errorf("failed to create several user realms. error: %w", err)
	}
	return nil
}

func (s *UserRealmService) Update(ctx context.Context, userRealm *models.UserRealmDTO) error {
	if err := s.repo.Update(ctx, userRealm); err != nil {
		return fmt.Errorf("failed to update user realm. error: %w", err)
	}
	s.eventBus.Notify(events.PolicyEvent{})
	return nil
}

func (s *UserRealmService) UpdateSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error {
	if len(dto) == 0 {
		return nil
	}
	if err := s.repo.UpdateSeveral(ctx, tx, dto); err != nil {
		return fmt.Errorf("failed to update several user realms. error: %w", err)
	}
	return nil
}

func (s *UserRealmService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user realm. error: %w", err)
	}
	s.eventBus.Notify(events.PolicyEvent{})
	return nil
}

func (s *UserRealmService) DeleteByUserAndRealm(ctx context.Context, userId, realmId string) error {
	if err := s.repo.DeleteByUserAndRealm(ctx, userId, realmId); err != nil {
		return fmt.Errorf("failed to delete user realm by user and realm. error: %w", err)
	}
	s.eventBus.Notify(events.PolicyEvent{})
	return nil
}

func (s *UserRealmService) DeleteSeveral(ctx context.Context, tx postgres.Tx, dto []*models.UserRealmDTO) error {
	if len(dto) == 0 {
		return nil
	}
	if err := s.repo.DeleteSeveral(ctx, tx, dto); err != nil {
		return fmt.Errorf("failed to delete several user realms. error: %w", err)
	}
	return nil
}
