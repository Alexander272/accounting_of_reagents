package services

import (
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
)

type Services struct {
	Menu
	MenuItem
	Role
	Session
	Permission

	AmountType
	ReagentType
	Reagent
	Spending
	Extending
	Note
}

type Deps struct {
	Repos           *repository.Repository
	Keycloak        *auth.KeycloakClient
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	BotUrl          string
	ChannelId       string
	ErrorBotUrl     string
}

func NewServices(deps Deps) *Services {
	// TODO можно включить для keycloak настройку что он за прокси и запустить сервер на 80 (или на другом) порту для вывода интерфейса
	// TODO при авторизации пользователя его можно искать сразу по нескольким realm

	// TODO для чего я делаю экземпляр ботов для каждого сервиса, когда нужно запустить один и отправлять все запросы на него. тоже самое относится и сервису email, файловому (file - minio) и возможно к некоторым другим. можно в принципе сделать один сервис бота для ошибок и рассылок (стоит рассмотреть и попробовать. можно попробовать связать шаблон и формат данных, а еще бота от имени которого будет отправляться сообщение)

	// notification := NewNotificationService(si, most, errorBot)

	menuItem := NewMenuItemService(deps.Repos.MenuItem)
	menu := NewMenuService(deps.Repos.Menu, menuItem)
	role := NewRoleService(deps.Repos.Role)
	session := NewSessionService(deps.Keycloak, role)
	permission := NewPermissionService("configs/privacy.conf", menu)

	most := NewMostService(deps.BotUrl, deps.ChannelId)

	amountType := NewAmountTypeService(deps.Repos.AmountType)
	reagentType := NewReagentTypeService(deps.Repos.ReagentType, role)
	reagent := NewReagentService(deps.Repos.Reagent, reagentType, most)
	spending := NewSpendingService(deps.Repos.Spending, reagent, most)
	extending := NewExtendingService(deps.Repos.Extending)
	note := NewNoteService(deps.Repos.Notes)

	return &Services{
		MenuItem:   menuItem,
		Menu:       menu,
		Role:       role,
		Session:    session,
		Permission: permission,

		AmountType:  amountType,
		ReagentType: reagentType,
		Reagent:     reagent,
		Spending:    spending,
		Extending:   extending,
		Note:        note,
	}
}
