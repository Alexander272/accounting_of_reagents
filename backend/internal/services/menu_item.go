package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type MenuItemService struct {
	repo repository.MenuItem
}

func NewMenuItemService(repo repository.MenuItem) *MenuItemService {
	return &MenuItemService{
		repo: repo,
	}
}

type MenuItem interface {
	GetAll(context.Context) ([]*models.MenuItem, error)
	Create(context.Context, *models.MenuItemDTO) error
	Update(context.Context, *models.MenuItemDTO) error
	Delete(context.Context, string) error
}

func (s *MenuItemService) GetAll(ctx context.Context) ([]*models.MenuItem, error) {
	menuItems, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all menu items. error: %w", err)
	}
	return menuItems, nil
}

func (s *MenuItemService) Create(ctx context.Context, menu *models.MenuItemDTO) error {
	if err := s.repo.Create(ctx, menu); err != nil {
		return fmt.Errorf("failed to create menu item. error: %w", err)
	}
	return nil
}

func (s *MenuItemService) Update(ctx context.Context, menu *models.MenuItemDTO) error {
	if err := s.repo.Update(ctx, menu); err != nil {
		return fmt.Errorf("failed to update menu item. error: %w", err)
	}
	return nil
}

func (s *MenuItemService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete menu item. error: %w", err)
	}
	return nil
}
