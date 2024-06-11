package note

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

type NoteHandlers struct {
	service services.Note
}

func NewNoteHandlers(service services.Note) *NoteHandlers {
	return &NoteHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Note, middleware *middleware.Middleware) {
	handlers := NewNoteHandlers(service)

	notes := api.Group("/notes", middleware.VerifyToken)
	{
		notes.GET(":reagentId", middleware.CheckPermissions(constants.Reagent, constants.Read), handlers.getByReagentId)
		notes.POST("", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.create)
		notes.PUT("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.update)
		notes.DELETE("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.delete)
	}
}

func (h *NoteHandlers) getByReagentId(c *gin.Context) {
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

func (h *NoteHandlers) create(c *gin.Context) {
	dto := &models.NoteDTO{}
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
	logger.Info("Добавлены комментарии", logger.StringAttr("reagentId", dto.ReagentId))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Примечание добавлено"})
}

func (h *NoteHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.NoteDTO{}
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
	logger.Info("Обновлены комментарии", logger.StringAttr("reagentId", dto.ReagentId))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Примечание обновлено"})
}

func (h *NoteHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}
	dto := &models.DeleteNoteDTO{Id: id}

	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удалены комментарии", logger.StringAttr("id", dto.Id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
