package roles

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

type Handler struct {
	service services.Roles
}

func NewHandler(service services.Roles) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Roles, middleware *middleware.Middleware) {
	handlers := NewHandler(service)

	roles := api.Group("/roles", middleware.CheckPermissions(access.Reg.R(access.ResourceRoles).Read()))
	{
		roles.GET("", handlers.getAll)
		roles.GET("/all/stats", handlers.getWithStats)

		write := roles.Group("", middleware.CheckPermissions(access.Reg.R(access.ResourceRoles).Write()))
		{
			write.GET("/item/:name", handlers.get)

			write.POST("", handlers.create)
			write.PUT("/:id", handlers.update)
		}

		delete := roles.Group("", middleware.CheckPermissions(access.Reg.R(access.ResourceRoles).Delete()))
		{
			delete.DELETE("/:id", handlers.delete)
		}

		permissions := roles.Group("", middleware.CheckPermissions(access.Reg.R(access.ResourcePerm).Read()))
		{
			permissions.GET("/:id/permissions", handlers.getWithPermissions)
		}
	}
}

func (h *Handler) getAll(c *gin.Context) {
	roles, err := h.service.GetAll(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: roles})
}

func (h *Handler) get(c *gin.Context) {
	slug := c.Param("name")
	dto := &models.GetRoleDTO{Slug: slug}

	role, err := h.service.GetOne(c, dto)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: role})
}

func (h *Handler) getWithPermissions(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
		return
	}

	role, err := h.service.GetOneWithPermissions(c, &models.GetRoleDTO{Id: id})
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: role})
}

func (h *Handler) getWithStats(c *gin.Context) {
	roles, err := h.service.GetWithStats(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: roles})
}

func (h *Handler) create(c *gin.Context) {
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

func (h *Handler) update(c *gin.Context) {
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
	if id != dto.Id {
		response.NewErrorResponse(c, http.StatusBadRequest, "id is not equal to dto.ID", "Некорректные данные")
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

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
		return
	}
	dto := &models.DeleteRoleDTO{Id: id}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
