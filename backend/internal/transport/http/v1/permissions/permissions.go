package permissions

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Permissions
}

func NewHandler(service services.Permissions) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Permissions, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	// permissions := api.Group("/permissions", middleware.CheckPermissions(access.Reg.R(access.ResourcePerm).Read()))
	permissions := api.Group("/permissions")
	{
		permissions.GET("", handler.getAll)
	}
}

func (h *Handler) getAll(c *gin.Context) {
	data, err := h.service.GetGrouped(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data})
}
