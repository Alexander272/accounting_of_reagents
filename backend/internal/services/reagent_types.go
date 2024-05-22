package services

import (
	"context"
	"fmt"
	"sort"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type ReagentTypeService struct {
	repo repository.ReagentType
	role Role
}

func NewReagentTypeService(repo repository.ReagentType, role Role) *ReagentTypeService {
	return &ReagentTypeService{
		repo: repo,
		role: role,
	}
}

type ReagentType interface {
	GetByRole(context.Context, string) ([]*models.ReagentType, error)
	Create(context.Context, *models.ReagentTypeDTO) (string, error)
	Update(context.Context, *models.ReagentTypeDTO) error
	Delete(context.Context, *models.DeleteReagentTypeDTO) error
}

func (s *ReagentTypeService) GetByRole(ctx context.Context, role string) ([]*models.ReagentType, error) {
	roles, err := s.role.GetAll(ctx, &models.GetRolesDTO{})
	if err != nil {
		return nil, err
	}

	reagentTypes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all reagent types. error: %w", err)
	}

	index := -1
	foundedRole := &models.RoleFull{}
	for i, r := range roles {
		if r.Name == role {
			index = i
			foundedRole = r
			break
		}
	}

	types := []*models.ReagentType{}
	for _, rt := range reagentTypes {
		for _, rId := range rt.RoleId {
			if rId == foundedRole.Id {
				types = append(types, rt)
			}
		}

	}

	extends := foundedRole.Extends
	for len(extends) > 0 {
		for i := index; i >= 0; i-- {
			if roles[i].Id == extends[0] {
				foundedRole = roles[i]
				index = i
				extends = append(extends[:0], extends[1:]...)
				extends = append(extends, foundedRole.Extends...)
				break
			}
		}

		for _, rt := range reagentTypes {
			for _, rId := range rt.RoleId {
				if rId == foundedRole.Id {
					types = append(types, rt)
				}
			}
			// if rt.RoleId == foundedRole.Id {
			// 	types = append(types, rt)
			// }
		}

	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].Number < types[j].Number
	})

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
