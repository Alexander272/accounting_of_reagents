package amount_type

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

type AmountTypeHandlers struct {
	service services.AmountType
}

func NewAmountTypeHandlers(services services.AmountType) *AmountTypeHandlers {
	return &AmountTypeHandlers{
		service: services,
	}
}

func Register(api *gin.RouterGroup, service services.AmountType, middleware *middleware.Middleware) {
	handlers := NewAmountTypeHandlers(service)

	amountTypes := api.Group("/amount-types", middleware.VerifyToken)
	{
		amountTypes.GET("", middleware.CheckPermissions(constants.Types, constants.Read), handlers.getAll)
		amountTypes.POST("", middleware.CheckPermissions(constants.Types, constants.Write), handlers.create)
		amountTypes.POST("/few", middleware.CheckPermissions(constants.Types, constants.Write), handlers.createSeveral)
		amountTypes.PUT("/:id", middleware.CheckPermissions(constants.Types, constants.Write), handlers.update)
		amountTypes.PUT("", middleware.CheckPermissions(constants.Types, constants.Write), handlers.updateSeveral)
		amountTypes.DELETE("/:id", middleware.CheckPermissions(constants.Types, constants.Write), handlers.delete)
	}
}

func (h *AmountTypeHandlers) getAll(c *gin.Context) {
	amountTypes, err := h.service.GetAll(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: amountTypes})
}

func (h *AmountTypeHandlers) create(c *gin.Context) {
	dto := &models.AmountTypeDTO{}
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

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Единица измерения создана"})
}

func (h *AmountTypeHandlers) createSeveral(c *gin.Context) {
	dto := []*models.AmountTypeDTO{}
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.CreateSeveral(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Единицы измерения созданы"})
}

func (h *AmountTypeHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id единиц измерения не задан")
		return
	}

	dto := &models.AmountTypeDTO{}
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

	c.JSON(http.StatusOK, response.IdResponse{Message: "Единица измерения обновлена"})
}

func (h *AmountTypeHandlers) updateSeveral(c *gin.Context) {
	dto := []*models.AmountTypeDTO{}
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.UpdateSeveral(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Единицы измерения обновлены"})
}

func (h *AmountTypeHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id единиц измерения не задан")
		return
	}
	dto := &models.DeleteAmountTypeDTO{Id: id}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
