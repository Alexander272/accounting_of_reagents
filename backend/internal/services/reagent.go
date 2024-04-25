package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
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
	user := ctx.Value(constants.CtxUser)
	logger.Debug("context", logger.AnyAttr("user", user))

	return nil, fmt.Errorf("not implemented")
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
