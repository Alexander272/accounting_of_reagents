package roles

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
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
	// handlers := NewRoleHandlers(service)

	roles := api.Group("/roles")
	{
		roles.GET("")
		// roles.GET("", handlers.GetAll)
		// roles.GET("/:name", handlers.Get)
		// roles.POST("", handlers.Create)
		// roles.PUT("/:id", handlers.Update)
		// roles.DELETE("/:id", handlers.Delete)
	}
}

// func (h *RoleHandlers) GetAll(c *gin.Context) {
// 	roles, err := h.service.GetAll(c, models.GetRolesDTO{})
// 	if err != nil {
// 		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
// 		h.errBot.Send(c, err.Error(), nil)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.DataResponse{Data: roles})
// }

// func (h *RoleHandlers) Get(c *gin.Context) {
// 	roleName := c.Param("name")

// 	role, err := h.service.Get(c, roleName)
// 	if err != nil {
// 		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
// 		h.errBot.Send(c, err.Error(), roleName)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.DataResponse{Data: role})
// }

// func (h *RoleHandlers) Create(c *gin.Context) {
// 	var dto models.RoleDTO
// 	if err := c.BindJSON(&dto); err != nil {
// 		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
// 		return
// 	}

// 	if err := h.service.Create(c, dto); err != nil {
// 		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
// 		h.errBot.Send(c, err.Error(), dto)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response.IdResponse{Message: "Роль создана"})
// }

// func (h *RoleHandlers) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
// 		return
// 	}

// 	var dto models.RoleDTO
// 	if err := c.BindJSON(&dto); err != nil {
// 		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
// 		return
// 	}
// 	dto.Id = id

// 	if err := h.service.Update(c, dto); err != nil {
// 		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
// 		h.errBot.Send(c, err.Error(), dto)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.IdResponse{Message: "Роль обновлена"})
// }

// func (h *RoleHandlers) Delete(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id роли не задан")
// 		return
// 	}

// 	if err := h.service.Delete(c, id); err != nil {
// 		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
// 		h.errBot.Send(c, err.Error(), id)
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, response.IdResponse{})
// }
