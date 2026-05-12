package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type ReagentTypeService struct {
	repo repository.ReagentType
}

func NewReagentTypeService(repo repository.ReagentType) *ReagentTypeService {
	return &ReagentTypeService{
		repo: repo,
	}
}

type ReagentType interface {
	Get(context.Context, *models.GetReagentTypeDTO) ([]*models.ReagentType, error)
	Create(context.Context, *models.ReagentTypeDTO) (string, error)
	Update(context.Context, *models.ReagentTypeDTO) error
	Delete(context.Context, *models.DeleteReagentTypeDTO) error
}

func (s *ReagentTypeService) Get(ctx context.Context, req *models.GetReagentTypeDTO) ([]*models.ReagentType, error) {
	types, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all reagent types. error: %w", err)
	}
	return types, nil
}

func (s *ReagentTypeService) Create(ctx context.Context, dto *models.ReagentTypeDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create reagent type. error: %w", err)
	}
	return id, nil
}

func (s *ReagentTypeService) Update(ctx context.Context, dto *models.ReagentTypeDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update reagent type. error: %w", err)
	}
	return nil
}

func (s *ReagentTypeService) Delete(ctx context.Context, dto *models.DeleteReagentTypeDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete reagent type. error: %w", err)
	}
	return nil
}
