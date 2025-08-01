package extending

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type ExtendingHandlers struct {
	service services.Extending
}

func NewExtendingHandlers(service services.Extending) *ExtendingHandlers {
	return &ExtendingHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Extending, middleware *middleware.Middleware) {
	handlers := NewExtendingHandlers(service)

	extending := api.Group("/extending", middleware.CheckPermissions(constants.Reagent, constants.Read))
	{
		extending.GET(":reagentId", handlers.getByReagentId)

		write := extending.Group("", middleware.CheckPermissions(constants.Reagent, constants.Write))
		{
			write.POST("", handlers.create)
			write.PUT("/:id", handlers.update)
			write.DELETE("/:id", handlers.delete)
		}
	}
}

func (h *ExtendingHandlers) getByReagentId(c *gin.Context) {
	reagentId := c.Param("reagentId")
	if reagentId == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id реагента не задан")
		return
	}

	extending, err := h.service.GetByReagentId(c, reagentId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: extending})
}

func (h *ExtendingHandlers) create(c *gin.Context) {
	dto := &models.ExtendingDTO{}
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
	logger.Info("Добавлено продление срока", logger.StringAttr("reagentId", dto.ReagentId))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Продление срока годности создано"})
}

func (h *ExtendingHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.ExtendingDTO{}
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
	logger.Info("Обновлено продление срока", logger.StringAttr("reagentId", dto.ReagentId))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Продление срока годности обновлено"})
}

func (h *ExtendingHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}
	dto := &models.DeleteExtendingDTO{Id: id}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удалено продление срока", logger.StringAttr("id", dto.Id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
