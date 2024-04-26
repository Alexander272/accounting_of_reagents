package roles

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/gin-gonic/gin"
)

type RoleHandlers struct {
	service services.Role
}

func NewRoleHandlers(service services.Role) *RoleHandlers {
	return &RoleHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Role, middleware *middleware.Middleware) {
	handlers := NewRoleHandlers(service)

	roles := api.Group("/roles", middleware.VerifyToken)
	{
		roles.GET("", middleware.CheckPermissions(constants.Roles, constants.Read), handlers.getAll)
		roles.GET("/:name", middleware.CheckPermissions(constants.Roles, constants.Write), handlers.get)
		roles.POST("", middleware.CheckPermissions(constants.Roles, constants.Write), handlers.create)
		roles.PUT("/:id", middleware.CheckPermissions(constants.Roles, constants.Write), handlers.update)
		roles.DELETE("/:id", middleware.CheckPermissions(constants.Roles, constants.Write), handlers.delete)
	}
}

func (h *RoleHandlers) getAll(c *gin.Context) {
	roles, err := h.service.GetAll(c, &models.GetRolesDTO{})
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: roles})
}

func (h *RoleHandlers) get(c *gin.Context) {
	roleName := c.Param("name")

	role, err := h.service.Get(c, roleName)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), roleName)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: role})
}

func (h *RoleHandlers) create(c *gin.Context) {
	dto := &models.RoleDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.Create(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Message: "Роль создана"})
}

func (h *RoleHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
		return
	}

	dto := &models.RoleDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}
	dto.Id = id

	if err := h.service.Update(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Роль обновлена"})
}

func (h *RoleHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
