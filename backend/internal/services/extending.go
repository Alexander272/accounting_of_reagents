package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type ExtendingService struct {
	repo    repository.Extending
	reagent Reagent
}

func NewExtendingService(repo repository.Extending, reagent Reagent) *ExtendingService {
	return &ExtendingService{
		repo:    repo,
		reagent: reagent,
	}
}

type Extending interface {
	GetByReagentId(context.Context, string) ([]*models.Extending, error)
	Create(context.Context, *models.ExtendingDTO) (string, error)
	Update(context.Context, *models.ExtendingDTO) error
	Delete(context.Context, *models.DeleteExtendingDTO) error
}

func (s *ExtendingService) GetByReagentId(ctx context.Context, reagentId string) ([]*models.Extending, error) {
	extending, err := s.repo.GetByReagentId(ctx, reagentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get extending by reagent id. error: %w", err)
	}
	return extending, nil
}

func (s *ExtendingService) Create(ctx context.Context, dto *models.ExtendingDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create extending. error: %w", err)
	}

	if err := s.reagent.ClearIsOverdue(ctx, dto.ReagentId); err != nil {
		return id, err
	}

	return id, nil
}

func (s *ExtendingService) Update(ctx context.Context, dto *models.ExtendingDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update extending. error: %w", err)
	}
	return nil
}

func (s *ExtendingService) Delete(ctx context.Context, dto *models.DeleteExtendingDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete extending. error: %w", err)
	}
	return nil
}
