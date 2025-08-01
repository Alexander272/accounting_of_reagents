package v1

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/amount_type"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/auth"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/extending"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/note"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/reagent"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/reagent_type"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/roles"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1/spending"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *services.Services
	conf       *config.Config
	middleware *middleware.Middleware
}

type Deps struct {
	Services   *services.Services
	Conf       *config.Config
	Middleware *middleware.Middleware
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		services:   deps.Services,
		conf:       deps.Conf,
		middleware: deps.Middleware,
	}
}

func (h *Handler) Init(group *gin.RouterGroup) {
	v1 := group.Group("/v1")

	auth.Register(v1, auth.Deps{Service: h.services.Session, Auth: h.conf.Auth})

	roles.Register(v1, h.services.Role, h.middleware)

	secure := v1.Group("", h.middleware.VerifyToken)
	reagent_type.Register(secure, h.services.ReagentType, h.middleware)
	amount_type.Register(secure, h.services.AmountType, h.middleware)
	reagent.Register(secure, h.services.Reagent, h.middleware)
	spending.Register(secure, h.services.Spending, h.middleware)
	extending.Register(secure, h.services.Extending, h.middleware)
	note.Register(secure, h.services.Note, h.middleware)

	//TODO можно попробовать ограничивать не только по разделам, но и по видам реактивов

	/* что осталось
	+ возможно я что-то забыл
	*/
}
