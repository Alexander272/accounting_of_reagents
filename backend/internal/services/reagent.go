package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type ReagentService struct {
	repo        repository.Reagent
	reagentType ReagentType
}

func NewReagentService(repo repository.Reagent, reagentType ReagentType) *ReagentService {
	return &ReagentService{
		repo:        repo,
		reagentType: reagentType,
	}
}

type Reagent interface {
	Get(context.Context, *models.Params) (*models.ReagentList, error)
	Create(context.Context, *models.ReagentDTO) (string, error)
	Update(context.Context, *models.ReagentDTO) error
	Delete(context.Context, *models.DeleteReagentDTO) error
}

func (s *ReagentService) Get(ctx context.Context, req *models.Params) (*models.ReagentList, error) {
	reagentTypes, err := s.reagentType.GetByRole(ctx, req.User.Role)
	if err != nil {
		return nil, err
	}

	isEmpty := true
	for _, f := range req.Filters {
		if f.Field == "reagentTypeId" {
			isEmpty = false
			break
		}
	}
	if isEmpty {
		values := []string{}
		for _, v := range reagentTypes {
			values = append(values, v.Id)
		}
		req.Filters = append(req.Filters, &models.Filter{
			Field: "reagentTypeId",
			Values: []*models.FilterValue{{
				CompareType: "in",
				Value:       strings.Join(values, ","),
			}},
		})
	}

	list, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get reagents list. error: %w", err)
	}

	return list, nil
}

func (s *ReagentService) Create(ctx context.Context, dto *models.ReagentDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create reagent. error: %w", err)
	}
	return id, nil
}

func (s *ReagentService) Update(ctx context.Context, dto *models.ReagentDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update reagent. error: %w", err)
	}
	return nil
}

func (s *ReagentService) Delete(ctx context.Context, dto *models.DeleteReagentDTO) error {
	return fmt.Errorf("not implemented")
}
