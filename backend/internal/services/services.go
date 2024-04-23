package services

import (
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
)

type Services struct {
	Session
}

type Deps struct {
	Repos           *repository.Repository
	Keycloak        *auth.KeycloakClient
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	BotUrl          string
	ErrorBotUrl     string
}

func NewServices(deps Deps) *Services {

	// TODO можно включить для keycloak настройку что он за прокси и запустить сервер на 80 (или на другом) порту для вывода интерфейса
	// TODO при авторизации пользователя его можно искать сразу по нескольким realm
	// session := NewSessionService(deps.Keycloak, role, filter)

	// TODO для чего я делаю экземпляр ботов для каждого сервиса, когда нужно запустить один и отправлять все запросы на него. тоже самое относится и сервису email, файловому (file - minio) и возможно к некоторым другим. можно в принципе сделать один сервис бота для ошибок и рассылок (стоит рассмотреть и попробовать. можно попробовать связать шаблон и формат данных, а еще бота от имени которого будет отправляться сообщение)

	// permission := NewPermissionService("configs/privacy.conf", menu)

	// notification := NewNotificationService(si, most, errorBot)

	return &Services{}
}
