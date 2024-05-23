package spending

import (
	"errors"
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/gin-gonic/gin"
)

type SpendingHandlers struct {
	service services.Spending
}

func NewSpendingHandlers(service services.Spending) *SpendingHandlers {
	return &SpendingHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Spending, middleware *middleware.Middleware) {
	handlers := NewSpendingHandlers(service)

	spending := api.Group("/spending", middleware.VerifyToken)
	{
		spending.GET(":reagentId", middleware.CheckPermissions(constants.Reagent, constants.Read), handlers.getByReagentId)
		spending.POST("", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.create)
		spending.PUT("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.update)
		spending.DELETE("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.delete)
	}
}

func (h *SpendingHandlers) getByReagentId(c *gin.Context) {
	reagentId := c.Param("reagentId")
	if reagentId == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id реагента не задан")
		return
	}

	spending, err := h.service.GetByReagentId(c, reagentId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: spending})
}

func (h *SpendingHandlers) create(c *gin.Context) {
	dto := &models.SpendingDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	id, err := h.service.Create(c, dto)
	if err != nil {
		if errors.Is(err, models.ErrBadValue) {
			response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Попытка списать большее количество реактива, чем осталось")
			return
		}

		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Расход добавлен"})
}

func (h *SpendingHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.SpendingDTO{}
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

	c.JSON(http.StatusOK, response.IdResponse{Message: "Расход обновлен"})
}

func (h *SpendingHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	reagentId := c.Query("reagentId")
	if reagentId == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id реактива не задан")
		return
	}

	dto := &models.DeleteSpendingDTO{
		Id:        id,
		ReagentId: reagentId,
	}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
