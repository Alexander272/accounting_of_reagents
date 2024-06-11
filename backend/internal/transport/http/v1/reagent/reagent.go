package reagent

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type ReagentHandlers struct {
	service services.Reagent
}

func NewReagentHandlers(service services.Reagent) *ReagentHandlers {
	return &ReagentHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Reagent, middleware *middleware.Middleware) {
	handlers := NewReagentHandlers(service)

	reagents := api.Group("/reagents", middleware.VerifyToken)
	{
		reagents.GET("", middleware.CheckPermissions(constants.Reagent, constants.Read), handlers.get)
		reagents.GET("/:id", middleware.CheckPermissions(constants.Reagent, constants.Read), handlers.getById)
		reagents.POST("/order", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.prepareOrder)
		reagents.POST("", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.create)
		reagents.PUT("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.update)
		reagents.DELETE("/:id", middleware.CheckPermissions(constants.Reagent, constants.Write), handlers.delete)
	}
}

func (h *ReagentHandlers) get(c *gin.Context) {
	params := &models.Params{
		Page:    &models.Page{},
		Sort:    []*models.Sort{},
		Filters: []*models.Filter{},
		User:    &models.User{},
	}

	page := c.Query("page")
	size := c.Query("size")

	sortLine := c.Query("sort_by")
	filters := c.QueryMap("filters")

	limit, err := strconv.Atoi(size)
	if err != nil {
		params.Page.Limit = 15
	} else {
		params.Page.Limit = limit
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		params.Page.Offset = 0
	} else {
		params.Page.Offset = (p - 1) * params.Page.Limit
	}

	if sortLine != "" {
		sort := strings.Split(sortLine, ",")
		for _, v := range sort {
			field, found := strings.CutPrefix(v, "-")
			t := "ASC"
			if found {
				t = "DESC"
			}

			params.Sort = append(params.Sort, &models.Sort{
				Field: field,
				Type:  t,
			})
		}
	}

	for k, v := range filters {
		valueMap := c.QueryMap(k)

		values := []*models.FilterValue{}
		for key, value := range valueMap {
			values = append(values, &models.FilterValue{
				CompareType: key,
				Value:       value,
			})
		}

		f := &models.Filter{
			Field:     k,
			FieldType: v,
			Values:    values,
		}

		params.Filters = append(params.Filters, f)
	}

	search := c.QueryMap("search")
	for key, value := range search {
		params.Search = &models.Search{
			Value:  value,
			Fields: strings.Split(key, ","),
		}
	}

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}

	user := u.(models.User)
	params.User = &user

	list, err := h.service.Get(c, params)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), params)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: list.List, Total: list.Total})
}

func (h *ReagentHandlers) getById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	reagent, err := h.service.GetById(c, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			response.NewErrorResponse(c, http.StatusNotFound, err.Error(), err.Error())
			return
		}
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: reagent})
}

func (h *ReagentHandlers) prepareOrder(c *gin.Context) {
	dto := &models.ReagentOrderDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.PrepareOrder(c, dto.List); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	c.JSON(http.StatusOK, response.IdResponse{Message: "Заказ сформирован"})
}

func (h *ReagentHandlers) create(c *gin.Context) {
	dto := &models.ReagentDTO{}
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
	logger.Info("Добавлен реактив", logger.StringAttr("title", dto.Name))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Реактив добавлен"})
}

func (h *ReagentHandlers) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.ReagentDTO{}
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
	logger.Info("Обновлен реактив", logger.StringAttr("title", dto.Name))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Реактив обновлен"})
}

func (h *ReagentHandlers) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	if err := h.service.SetDeleteStamp(c, id); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удален реактив", logger.StringAttr("id", id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
