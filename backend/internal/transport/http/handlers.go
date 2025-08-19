package http

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	httpV1 "github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/limiter"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// permissions casbin.Casbin
	keycloak *auth.KeycloakClient
	services *services.Services
}

func NewHandler(services *services.Services, keycloak *auth.KeycloakClient) *Handler {
	return &Handler{
		services: services,
		keycloak: keycloak,
		// permissions: permissions,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		limiter.Limit(conf.Limiter.RPS, conf.Limiter.Burst, conf.Limiter.TTL),
		gin.CustomRecovery(h.ErrorHandler),
	)

	// Init router
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router, conf)

	return router
}

func (h *Handler) ErrorHandler(c *gin.Context, origErr any) {
	err := fmt.Errorf("unexpected error: %v", origErr)
	error_bot.Send(c, err.Error(), gin.H{"Stack trace": string(debug.Stack())})

	response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла непредвиденная ошибка: "+err.Error())
}

func (h *Handler) initAPI(router *gin.Engine, conf *config.Config) {
	middleware := middleware.NewMiddleware(h.services, conf.Auth, h.keycloak)
	handlerV1 := httpV1.NewHandler(httpV1.Deps{Services: h.services, Conf: conf, Middleware: middleware})

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
