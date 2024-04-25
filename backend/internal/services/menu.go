package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type MenuService struct {
	repo repository.Menu
}

func NewMenuService(repo repository.Menu) *MenuService {
	return &MenuService{
		repo: repo,
	}
}

type Menu interface {
	GetAll(context.Context) ([]*models.MenuFull, error)
	Create(context.Context, *models.MenuDTO) error
	Update(context.Context, *models.MenuDTO) error
	Delete(context.Context, string) error
}

func (s *MenuService) GetAll(ctx context.Context) ([]*models.MenuFull, error) {
	// m, err := s.repo.GetAll(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get all menu. error: %w", err)
	// }

	// api, err := s.api.GetAll(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// for i, m2 := range m {
	// 	item := models.MenuItem{}
	// 	for _, mi := range api {
	// 		if mi.Id == m2.MenuItemId {
	// 			item = mi
	// 			break
	// 		}
	// 	}

	// 	if i == 0 || menu[len(menu)-1].Id != m2.RoleId {
	// 		menu = append(menu, models.MenuFull{
	// 			Id: m2.RoleId,
	// 			Role: models.RoleFull{
	// 				Id:      m2.RoleId,
	// 				Name:    m2.RoleName,
	// 				Level:   m2.RoleLevel,
	// 				Extends: m2.RoleExtends,
	// 			},
	// 			MenuItems: []models.MenuItem{item},
	// 		})
	// 	} else {
	// 		menu[len(menu)-1].MenuItems = append(menu[len(menu)-1].MenuItems, item)
	// 	}
	// }

	// return menu, nil
	return nil, fmt.Errorf("not implemented")
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
