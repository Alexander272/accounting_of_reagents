package user

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/access"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service services.Users
}

func NewUserHandlers(service services.Users) *UserHandlers {
	return &UserHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Users, middleware *middleware.Middleware) {
	handlers := NewUserHandlers(service)

	users := api.Group("/users", middleware.CheckPermissions(access.Reg.R(access.ResourceUsers).Read()))
	{
		users.GET("", handlers.getAll)

		write := users.Group("", middleware.CheckPermissions(access.Reg.R(access.ResourceUsers).Write()))
		{
			write.POST("/sync", handlers.sync)
			write.PUT("/:id", handlers.update)
		}
	}
}

func (h *UserHandlers) getAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: users})
}

func (h *UserHandlers) sync(c *gin.Context) {
	if err := h.service.SyncFromKeycloak(c); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Пользователи синхронизированы"})
}

func (h *UserHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id пользователя не задан")
		return
	}

	dto := &models.UserDataDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}
	dto.Id = id

	if err := h.service.UpdateRole(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Роль пользователя обновлена"})
}
