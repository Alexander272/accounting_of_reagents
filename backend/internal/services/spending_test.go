package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
)

// ---- фейки: Go stdlib, ни одного мок-фреймворка ----
//
// Встраиваем сами интерфейсы (Reagent, repository.Spending, Most) в фейки.
// Это даёт всем фейкам "реализовать" интерфейс бесплатно, а мы переопределяем
// только те методы, что реально вызываются в Create. Расширение интерфейса
// в будущем не сломает тест — это та самая идиома embed-in-fake.

type fakeReagentSvc struct {
	Reagent

	group    []*models.ReagentWithRemainder
	groupErr error

	notified   []*models.ReagentNotificationDTO
	notifErr error
}

func (f *fakeReagentSvc) GetGroupRemainders(ctx context.Context, id string) ([]*models.ReagentWithRemainder, error) {
	return f.group, f.groupErr
}

func (f *fakeReagentSvc) SetGroupNotifications(ctx context.Context, dto []*models.ReagentNotificationDTO) error {
	f.notified = dto
	return f.notifErr
}

type fakeSpendingRepo struct {
	repository.Spending

	createErr error
}

func (f *fakeSpendingRepo) Create(ctx context.Context, dto *models.SpendingDTO) (string, error) {
	if f.createErr != nil {
		return "", f.createErr
	}
	return "spending-id", nil
}

type fakeMostSvc struct {
	Most

	sent []*models.Notification
}

func (f *fakeMostSvc) Send(ctx context.Context, notification *models.Notification) error {
	f.sent = append(f.sent, notification)
	return nil
}

// ---- помощники построения позиции группы ----

const (
	testReagentId   = "11111111-1111-1111-1111-111111111111"
	runOutReagentId = "22222222-2222-2222-2222-222222222222"
)

func groupItem(id string, amount, remainder float64, hasNotification bool) *models.ReagentWithRemainder {
	return &models.ReagentWithRemainder{
		Id:              id,
		Name:            "Тест-реактив",
		Document:        "док",
		Purity:          "чда",
		Manufacturer:    "произв",
		Amount:          amount,
		Remainder:       remainder,
		HasNotification: hasNotification,
	}
}

// notifiedIds — id из dto, у которых HasNotification=true.
func notifiedIds(dto []*models.ReagentNotificationDTO) map[string]bool {
	out := map[string]bool{}
	for _, v := range dto {
		if v.HasNotification {
			out[v.Id] = true
		}
	}
	return out
}

// runOutIds — id из dto, у которых HasRunOut=true.
func runOutIds(dto []*models.ReagentNotificationDTO) map[string]bool {
	out := map[string]bool{}
	for _, v := range dto {
		if v.HasRunOut {
			out[v.Id] = true
		}
	}
	return out
}

// ---- тест: has_run_out больше не привязан к has_notification ----

func TestSpendingService_Create_HasRunOutDecoupled(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name         string
		group        []*models.ReagentWithRemainder
		amount       float64
		wantNotified map[string]bool // id → должны получить HasNotification=true
		wantRunOut   map[string]bool // id → должны получить HasRunOut=true
		wantSend     bool
		wantErr      error
	}{
		{
			name: "уже уведомлённый, но в группе закончился → HasRunOut=true, повторно не уведомляем",
			group: []*models.ReagentWithRemainder{
				groupItem(testReagentId, 100, 1, true), // остаток 1 на массу 100 → мало, уже уведомлён
			},
			amount:       1, // списываем последний → остаток 0 → закончился
			wantNotified: map[string]bool{}, // НЕ переуведомляем: флаг уже true, репо не переставит
			wantRunOut:   map[string]bool{testReagentId: true},
			wantSend:     true, // все закончились → отправить "заканчивается" нет, но см. ниже
		},
		{
			name: "не уведомлённый закончился → оба флага + отправка",
			group: []*models.ReagentWithRemainder{
				groupItem(testReagentId, 100, 2, false),
			},
			amount:       2,
			wantNotified: map[string]bool{testReagentId: true},
			wantRunOut:   map[string]bool{testReagentId: true},
			wantSend:     true,
		},
		{
			name: "сосед уже закончился, целевой не тронут → только HasRunOut у соседа",
			group: []*models.ReagentWithRemainder{
				groupItem(testReagentId, 100, 88, false),
				groupItem(runOutReagentId, 10, 0, true), // сосед: остаток 0, уже уведомлён
			},
			amount:       5, // целевой остаётся 83 — не закончился
			wantNotified: map[string]bool{},           // целевой не маленький
			wantRunOut:   map[string]bool{runOutReagentId: true}, // сосед закончился
			wantSend:     false,
		},
		{
			name:    "остатка не хватает на целевой → ErrBadValue",
			group: []*models.ReagentWithRemainder{
				groupItem(testReagentId, 100, 15, false),
			},
			amount: 20, // > остатка → ошибка до любых флагов
			wantErr: models.ErrBadValue,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reagent := &fakeReagentSvc{group: tc.group}
			spending := &fakeSpendingRepo{}
			most := &fakeMostSvc{}

			svc := NewSpendingService(spending, reagent, most)

			id, err := svc.Create(ctx, &models.SpendingDTO{
				ReagentId: testReagentId,
				Amount:    tc.amount,
			})

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("ошибка: хочу %v, получил %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if id == "" {
				t.Fatal("пустой id")
			}

			gotNotified := notifiedIds(reagent.notified)
			gotRunOut := runOutIds(reagent.notified)

			for rid := range tc.wantNotified {
				if !gotNotified[rid] {
					t.Errorf("реактив %q: жду HasNotification=true, а его нет (all dto: %+v)", rid, reagent.notified)
				}
			}
			for rid := range tc.wantRunOut {
				if !gotRunOut[rid] {
					t.Errorf("реактив %q: жду HasRunOut=true, а его нет (all dto: %+v)", rid, reagent.notified)
				}
			}

			hasSend := len(most.sent) > 0
			if hasSend != tc.wantSend {
				t.Errorf("отправка уведомления: хочу %v, получил %v (отправлено %d)", tc.wantSend, hasSend, len(most.sent))
			}
		})
	}
}
