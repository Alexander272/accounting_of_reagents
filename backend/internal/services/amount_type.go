package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type AmountTypeService struct {
	repo repository.AmountType
}

func NewAmountTypeService(repo repository.AmountType) *AmountTypeService {
	return &AmountTypeService{
		repo: repo,
	}
}

type AmountType interface {
	GetAll(context.Context) ([]*models.AmountType, error)
	Create(context.Context, *models.AmountTypeDTO) (string, error)
	Update(context.Context, *models.AmountTypeDTO) error
	UpdateSeveral(context.Context, []*models.AmountTypeDTO) error
	Delete(context.Context, *models.DeleteAmountTypeDTO) error
}

func (s *AmountTypeService) GetAll(ctx context.Context) ([]*models.AmountType, error) {
	amountTypes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all amount types. error: %w", err)
	}
	return amountTypes, nil
}

func (s *AmountTypeService) Create(ctx context.Context, dto *models.AmountTypeDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create amount type. error: %w", err)
	}
	return id, nil
}

func (s *AmountTypeService) Update(ctx context.Context, dto *models.AmountTypeDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update amount type. error: %w", err)
	}
	return nil
}

func (s *AmountTypeService) UpdateSeveral(ctx context.Context, dto []*models.AmountTypeDTO) error {
	if err := s.repo.UpdateSeveral(ctx, dto); err != nil {
		return fmt.Errorf("failed to several update amount type. error: %w", err)
	}
	return nil
}

func (s *AmountTypeService) Delete(ctx context.Context, dto *models.DeleteAmountTypeDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete amount type. error: %w", err)
	}
	return nil
}
