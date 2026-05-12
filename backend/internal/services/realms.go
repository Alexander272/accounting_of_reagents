package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/google/uuid"
)

type RealmService struct {
	repo repository.Realm
}

func NewRealmService(repo repository.Realm) *RealmService {
	return &RealmService{
		repo: repo,
	}
}

type Realm interface {
	GetAll(ctx context.Context) ([]*models.Realm, error)
	GetById(ctx context.Context, id string) (*models.Realm, error)
	Create(ctx context.Context, realm *models.RealmDTO) error
	Update(ctx context.Context, realm *models.RealmDTO) error
	Delete(ctx context.Context, id string) error
}

func (s *RealmService) GetAll(ctx context.Context) ([]*models.Realm, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get realms. error: %w", err)
	}
	return data, nil
}

func (s *RealmService) GetById(ctx context.Context, id string) (*models.Realm, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get realm by id. error: %w", err)
	}
	return data, nil
}

func (s *RealmService) Create(ctx context.Context, realm *models.RealmDTO) error {
	if realm.Id == "" {
		realm.Id = uuid.NewString()
	}

	if err := s.repo.Create(ctx, realm); err != nil {
		return fmt.Errorf("failed to create realm. error: %w", err)
	}
	return nil
}

func (s *RealmService) Update(ctx context.Context, realm *models.RealmDTO) error {
	if err := s.repo.Update(ctx, realm); err != nil {
		return fmt.Errorf("failed to update realm. error: %w", err)
	}
	return nil
}

func (s *RealmService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete realm. error: %w", err)
	}
	return nil
}
