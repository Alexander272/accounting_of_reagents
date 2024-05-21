package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

type ReagentService struct {
	repo        repository.Reagent
	reagentType ReagentType
	most        Most
}

func NewReagentService(repo repository.Reagent, reagentType ReagentType, most Most) *ReagentService {
	return &ReagentService{
		repo:        repo,
		reagentType: reagentType,
		most:        most,
	}
}

type Reagent interface {
	Get(context.Context, *models.Params) (*models.ReagentList, error)
	GetById(context.Context, string) (*models.EditReagent, error)
	GetByIdList(context.Context, []string) ([]*models.Reagent, error)
	GetRemainder(context.Context, string) (*models.ReagentWithRemainder, error)
	GetOverdue(context.Context) ([]*models.Reagent, error)
	SendOverdue(context.Context) error
	PrepareOrder(context.Context, []string) error
	Create(context.Context, *models.ReagentDTO) (string, error)
	Update(context.Context, *models.ReagentDTO) error
	SetNotification(context.Context, *models.ReagentNotificationDTO) error
	SetIsOverdue(context.Context, []string) error
	ClearIsOverdue(context.Context, string) error
	SetDeleteStamp(context.Context, string) error
	DeleteOld(context.Context) error
	Delete(context.Context, *models.DeleteReagentDTO) error
}

func (s *ReagentService) Get(ctx context.Context, req *models.Params) (*models.ReagentList, error) {
	//TODO из-за того что у меня есть одинаковые типы (для разных ролей) фильтр по id типа не выдает нужный мне результат
	reagentTypes, err := s.reagentType.GetByRole(ctx, req.User.Role)
	if err != nil {
		return nil, err
	}

	isEmpty := true
	for _, f := range req.Filters {
		if f.Field == "typeId" {
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
			Field: "typeId",
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

	for _, i := range list.List {
		//TODO думаю когда реактив закончился нужно менять подсветку или делать другим цвет текста
		shelfLife := time.Unix(int64(i.DateOfManufacture), 0)
		shelfLife = shelfLife.AddDate(0, i.ShelfLife, 0)
		shelfLife = shelfLife.AddDate(0, i.SumPeriod, 0)
		now := time.Now()

		if shelfLife.Compare(time.Date(now.Year(), now.Month()+1, now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())) <= 0 {
			i.Background = constants.OrangeColor
		}
		if shelfLife.Compare(now) <= 0 {
			i.Background = constants.RedColor
		}
		if i.Seizure != "" {
			i.Background = constants.BeigeColor
		}
		// if i.HasRunOut {}
	}

	return list, nil
}

func (s *ReagentService) GetById(ctx context.Context, id string) (*models.EditReagent, error) {
	reagent, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get reagent by id. error: %w", err)
	}
	return reagent, nil
}

func (s *ReagentService) GetByIdList(ctx context.Context, list []string) ([]*models.Reagent, error) {
	reagents, err := s.repo.GetByIdList(ctx, list)
	if err != nil {
		return nil, fmt.Errorf("failed to get reagents by id list. error: %w", err)
	}
	return reagents, nil
}

func (s *ReagentService) GetRemainder(ctx context.Context, id string) (*models.ReagentWithRemainder, error) {
	remainder, err := s.repo.GetRemainder(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get remainder. error: %w", err)
	}
	return remainder, nil
}

func (s *ReagentService) GetOverdue(ctx context.Context) ([]*models.Reagent, error) {
	list, err := s.repo.Get(ctx, &models.Params{Page: &models.Page{Limit: 999999}})
	if err != nil {
		return nil, fmt.Errorf("failed to get reagents list. error: %w", err)
	}

	newList := []*models.Reagent{}
	for _, i := range list.List {
		shelfLife := time.Unix(int64(i.DateOfManufacture), 0)
		shelfLife = shelfLife.AddDate(0, i.ShelfLife, 0)
		shelfLife = shelfLife.AddDate(0, i.SumPeriod, 0)
		now := time.Now()

		if shelfLife.Compare(time.Date(now.Year(), now.Month()+1, now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())) <= 0 &&
			!i.IsOverdue && i.Seizure == "" {
			newList = append(newList, i)
		}
	}

	return newList, nil
}
func (s *ReagentService) SendOverdue(ctx context.Context) error {
	reagents, err := s.GetOverdue(ctx)
	if err != nil {
		return err
	}

	if len(reagents) == 0 {
		return nil
	}

	notification := &models.Notification{
		Message: "У следующих реактивов истекает срок годности",
		Data:    reagents,
	}
	if err := s.most.Send(ctx, notification); err != nil {
		return fmt.Errorf("failed to send notification. error: %w", err)
	}

	idList := []string{}
	for _, r := range reagents {
		idList = append(idList, r.Id)
	}
	if err := s.SetIsOverdue(ctx, idList); err != nil {
		return err
	}

	return nil
}

func (s *ReagentService) PrepareOrder(ctx context.Context, idList []string) error {
	reagents, err := s.GetByIdList(ctx, idList)
	if err != nil {
		return err
	}

	notification := &models.Notification{
		Message: "Данные для заказа",
		Data:    reagents,
	}
	if err := s.most.Send(ctx, notification); err != nil {
		return fmt.Errorf("failed to send notification. error: %w", err)
	}
	return nil
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

func (s *ReagentService) SetNotification(ctx context.Context, dto *models.ReagentNotificationDTO) error {
	if err := s.repo.SetNotification(ctx, dto); err != nil {
		return fmt.Errorf("failed to set has_notification and has_run_out. error: %w", err)
	}
	return nil
}

func (s *ReagentService) SetIsOverdue(ctx context.Context, idList []string) error {
	if err := s.repo.SetIsOverdue(ctx, idList); err != nil {
		return fmt.Errorf("failed to set is overdue. error: %w", err)
	}
	return nil
}
func (s *ReagentService) ClearIsOverdue(ctx context.Context, id string) error {
	if err := s.repo.ClearIsOverdue(ctx, id); err != nil {
		return fmt.Errorf("failed to clear is overdue. error: %w", err)
	}
	return nil
}

func (s *ReagentService) SetDeleteStamp(ctx context.Context, id string) error {
	if err := s.repo.SetDeleteStamp(ctx, id); err != nil {
		return fmt.Errorf("failed to set deleted stamp by id. error: %w", err)
	}
	return nil
}

func (s *ReagentService) DeleteOld(ctx context.Context) error {
	if err := s.repo.DeleteOld(ctx); err != nil {
		return fmt.Errorf("failed to delete old. error: %w", err)
	}
	return nil
}
func (s *ReagentService) Delete(ctx context.Context, dto *models.DeleteReagentDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete. error: %w", err)
	}
	return nil
}
