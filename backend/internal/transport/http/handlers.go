package http

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	httpV1 "github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
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
	)

	// Init router
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router, conf)

	return router
}

func (h *Handler) initAPI(router *gin.Engine, conf *config.Config) {
	middleware := middleware.NewMiddleware(h.services, conf.Auth, h.keycloak)
	handlerV1 := httpV1.NewHandler(httpV1.Deps{Services: h.services, Conf: conf, Middleware: middleware})

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	// allPaths := make(map[string]models.ApiDTO)
	// routes := router.Routes()
	// for _, v := range routes {
	// 	if v.Path != "/api/ping" {
	// 		allPaths[fmt.Sprintf("%s:%s", v.Path, v.Method)] = models.ApiDTO{
	// 			Path:   v.Path,
	// 			Method: v.Method,
	// 		}
	// 	}
	// }
	// if err := h.services.ApiPaths.Load(context.Background(), allPaths); err != nil {
	// 	logger.Error("failed to load api paths. error: ", logger.ErrAttr(err))
	// }
}
