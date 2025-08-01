package reagent_type

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

type ReagentTypeHandlers struct {
	service services.ReagentType
}

func NewReagentTypeHandlers(service services.ReagentType) *ReagentTypeHandlers {
	return &ReagentTypeHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.ReagentType, middleware *middleware.Middleware) {
	handlers := NewReagentTypeHandlers(service)

	reagentTypes := api.Group("/reagent-types", middleware.CheckPermissions(constants.Types, constants.Read))
	{
		reagentTypes.GET("", handlers.get)

		write := reagentTypes.Group("", middleware.CheckPermissions(constants.Types, constants.Write))
		{
			write.POST("", handlers.create)
			write.PUT("/:id", handlers.update)
			write.DELETE("/:id", handlers.delete)
		}
	}
}

func (h *ReagentTypeHandlers) get(c *gin.Context) {
	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}

	user := u.(models.User)

	reagentTypes, err := h.service.GetByRole(c, user.Role)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), user)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: reagentTypes})
}

func (h *ReagentTypeHandlers) create(c *gin.Context) {
	dto := &models.ReagentTypeDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	id, err := h.service.Create(c, dto)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Тип реагента создан"})
}

func (h *ReagentTypeHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id типа реагента не задан")
		return
	}

	dto := &models.ReagentTypeDTO{}
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

	c.JSON(http.StatusOK, response.IdResponse{Message: "Тип реагента обновлен"})
}

func (h *ReagentTypeHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id типа реагента не задан")
		return
	}
	dto := &models.DeleteReagentTypeDTO{Id: id}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
