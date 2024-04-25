package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type SpendingService struct {
	repo repository.Spending
}

func NewSpendingService(repo repository.Spending) *SpendingService {
	return &SpendingService{
		repo: repo,
	}
}

type Spending interface {
	GetByReagentId(context.Context, string) ([]*models.Spending, error)
	Create(context.Context, *models.SpendingDTO) (string, error)
	Update(context.Context, *models.SpendingDTO) error
	Delete(context.Context, *models.DeleteSpendingDTO) error
}

func (s *SpendingService) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Spending, error) {
	spending, err := s.repo.GetByReagentId(ctx, reagentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get spending by reagent id. error: %w", err)
	}
	return spending, nil
}

func (s *SpendingService) Create(ctx context.Context, dto *models.SpendingDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create spending. error: %w", err)
	}
	return id, nil
}

func (s *SpendingService) Update(ctx context.Context, dto *models.SpendingDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update spending. error: %w", err)
	}
	return nil
}

func (s *SpendingService) Delete(ctx context.Context, dto *models.DeleteSpendingDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete spending. error: %w", err)
	}
	return nil
}
