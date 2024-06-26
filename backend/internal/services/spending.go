package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type SpendingService struct {
	repo    repository.Spending
	reagent Reagent
	most    Most
}

func NewSpendingService(repo repository.Spending, reagent Reagent, most Most) *SpendingService {
	return &SpendingService{
		repo:    repo,
		reagent: reagent,
		most:    most,
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
	remainder, err := s.reagent.GetRemainder(ctx, dto.ReagentId)
	if err != nil {
		return "", err
	}

	if remainder.Remainder < dto.Amount {
		return "", models.ErrBadValue
	}

	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create spending. error: %w", err)
	}

	// проверять осталось ли больше 30 % от изначальной массы. если меньше отправлять уведомление. думаю хватит одного уведомления, а это значит нужно какой-то флаг добавить
	if (remainder.Remainder-dto.Amount)/remainder.Amount <= .3 {
		reagentNotification := &models.ReagentNotificationDTO{
			Id:              remainder.Id,
			HasNotification: true,
			HasRunOut:       (remainder.Remainder - dto.Amount) == 0,
		}
		if err := s.reagent.SetNotification(ctx, reagentNotification); err != nil {
			return "", err
		}

		if remainder.HasNotification {
			return id, nil
		}

		notification := &models.Notification{
			Message: "Заканчивается реактив.",
			Data: []*models.Reagent{{
				Id:           remainder.Id,
				Name:         remainder.Name,
				Document:     remainder.Document,
				Purity:       remainder.Purity,
				Manufacturer: remainder.Manufacturer,
			}},
		}
		if err := s.most.Send(ctx, notification); err != nil {
			return "", fmt.Errorf("failed to send notification. error: %w", err)
		}
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
	remainder, err := s.reagent.GetRemainder(ctx, dto.ReagentId)
	if err != nil {
		return err
	}

	amount, err := s.repo.Delete(ctx, dto)
	if err != nil {
		return fmt.Errorf("failed to delete spending. error: %w", err)
	}

	reagentNotification := &models.ReagentNotificationDTO{
		Id:              remainder.Id,
		HasNotification: (remainder.Remainder+amount)/remainder.Amount <= .3,
		HasRunOut:       false,
	}
	if err := s.reagent.SetNotification(ctx, reagentNotification); err != nil {
		return err
	}
	return nil
}
