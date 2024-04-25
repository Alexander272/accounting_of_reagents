package middleware

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
	// "github.com/casbin/casbin/v2"
)

type Middleware struct {
	keycloak *auth.KeycloakClient
	// TODO стоит наверное получать не все сервисы, а только те что используются
	services *services.Services
	auth     config.AuthConfig
}

func NewMiddleware(services *services.Services, auth config.AuthConfig, keycloak *auth.KeycloakClient) *Middleware {
	return &Middleware{
		// enforcer: enforcer,
		keycloak: keycloak,
		services: services,
		auth:     auth,
	}
}
