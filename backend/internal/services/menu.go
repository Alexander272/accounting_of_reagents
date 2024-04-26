package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type MenuService struct {
	repo repository.Menu
	item MenuItem
}

func NewMenuService(repo repository.Menu, item MenuItem) *MenuService {
	return &MenuService{
		repo: repo,
		item: item,
	}
}

type Menu interface {
	GetAll(context.Context) ([]*models.MenuFull, error)
	Create(context.Context, *models.MenuDTO) error
	Update(context.Context, *models.MenuDTO) error
	Delete(context.Context, string) error
}

func (s *MenuService) GetAll(ctx context.Context) ([]*models.MenuFull, error) {
	menu, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all menu. error: %w", err)
	}

	items, err := s.item.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	menuFull := []*models.MenuFull{}

	for i, m := range menu {
		menuItem := &models.MenuItem{}
		for _, item := range items {
			if m.MenuItemId == item.Id {
				menuItem = item
				break
			}
		}

		if i == 0 || menuFull[len(menuFull)-1].Id != m.RoleId {
			menuFull = append(menuFull, &models.MenuFull{
				Id: m.RoleId,
				Role: models.RoleFull{
					Id:      m.RoleId,
					Name:    m.RoleName,
					Level:   m.RoleLevel,
					Extends: m.RoleExtends,
				},
				MenuItems: []*models.MenuItem{menuItem},
			})
		} else {
			menuFull[len(menuFull)-1].MenuItems = append(menuFull[len(menuFull)-1].MenuItems, menuItem)
		}

	}

	return menuFull, nil
}

func (s *MenuService) Create(ctx context.Context, menu *models.MenuDTO) error {
	if err := s.repo.Create(ctx, menu); err != nil {
		return fmt.Errorf("failed to create menu. error: %w", err)
	}
	return nil
}

func (s *MenuService) Update(ctx context.Context, menu *models.MenuDTO) error {
	if err := s.repo.Update(ctx, menu); err != nil {
		return fmt.Errorf("failed to update menu. error: %w", err)
	}
	return nil
}

func (s *MenuService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete menu. error: %w", err)
	}
	return nil
}
