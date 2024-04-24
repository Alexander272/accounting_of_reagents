package v1

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *services.Services
	conf       *config.Config
	middleware *middleware.Middleware
	cookieName string
}

type Deps struct {
	Services   *services.Services
	Conf       *config.Config
	Middleware *middleware.Middleware
	CookieName string
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		services:   deps.Services,
		conf:       deps.Conf,
		middleware: deps.Middleware,
		cookieName: deps.CookieName,
	}
}

func (h *Handler) Init(group *gin.RouterGroup) {
	// v1 := group.Group("/v1")

	//TODO можно попробовать ограничивать не только по разделам, но и по видам реактивов
}
