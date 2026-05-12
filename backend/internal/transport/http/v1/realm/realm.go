package realm

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/access"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service    services.Realm
	userRealm  services.UserRealm
	middleware *middleware.Middleware
}

func NewHandler(service services.Realm, userRealm services.UserRealm, middleware *middleware.Middleware) *Handler {
	return &Handler{
		service:    service,
		userRealm:  userRealm,
		middleware: middleware,
	}
}

type Deps struct {
	Service    services.Realm
	UserRealm  services.UserRealm
	Middleware *middleware.Middleware
}

func Register(router *gin.RouterGroup, deps Deps) {
	h := NewHandler(deps.Service, deps.UserRealm, deps.Middleware)

	realms := router.Group("/realms", deps.Middleware.CheckPermissions(access.Reg.R(access.ResourceRealm).Read()))
	{
		realms.GET("", h.GetAll)
		realms.GET("/:id", h.GetById)

		write := realms.Group("", deps.Middleware.CheckPermissions(access.Reg.R(access.ResourceRealm).Write()))
		{
			write.POST("", h.Create)
			write.PUT("/:id", h.Update)
			write.DELETE("/:id", h.Delete)
		}
	}
}

func (h *Handler) GetAll(c *gin.Context) {
	realms, err := h.service.GetAll(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to get realms")
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: realms})
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid realm id")
		return
	}

	realm, err := h.service.GetById(c, id)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to get realm")
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: realm})
}

func (h *Handler) Create(c *gin.Context) {
	var dto models.RealmDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "failed to parse request body")
		return
	}

	if dto.Id == "" {
		dto.Id = uuid.New().String()
	}

	if err := h.service.Create(c, &dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to create realm")
		return
	}
	c.JSON(http.StatusCreated, response.IdResponse{Message: "Realm created", Id: dto.Id})
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid realm id")
		return
	}

	var dto models.RealmDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "failed to parse request body")
		return
	}
	dto.Id = id

	if err := h.service.Update(c, &dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to update realm")
		return
	}
	c.JSON(http.StatusOK, response.IdResponse{Message: "Realm updated", Id: id})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid realm id")
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to delete realm")
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) GetByUserId(c *gin.Context) {
	userId := c.Param("userId")
	if _, err := uuid.Parse(userId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid user id")
		return
	}

	userRealms, err := h.userRealm.GetByUserId(c, userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to get user realms")
		return
	}
	c.JSON(http.StatusOK, userRealms)
}

func (h *Handler) CreateUserRealm(c *gin.Context) {
	var dto models.UserRealmDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "failed to parse request body")
		return
	}

	if dto.Id == "" {
		dto.Id = uuid.New().String()
	}

	if err := h.userRealm.Create(c, &dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to create user realm")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": dto.Id})
}

func (h *Handler) UpdateUserRealm(c *gin.Context) {
	userId := c.Param("userId")
	realmId := c.Param("realmId")
	if _, err := uuid.Parse(userId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid user id")
		return
	}
	if _, err := uuid.Parse(realmId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid realm id")
		return
	}

	var dto models.UserRealmDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "failed to parse request body")
		return
	}
	dto.UserId = userId
	dto.RealmId = realmId

	if err := h.userRealm.Update(c, &dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to update user realm")
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUserRealm(c *gin.Context) {
	userId := c.Param("userId")
	realmId := c.Param("realmId")
	if _, err := uuid.Parse(userId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid user id")
		return
	}
	if _, err := uuid.Parse(realmId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "invalid realm id")
		return
	}

	if err := h.userRealm.DeleteByUserAndRealm(c, userId, realmId); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "failed to delete user realm")
		return
	}
	c.Status(http.StatusOK)
}
