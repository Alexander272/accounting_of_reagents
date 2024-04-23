package middleware

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) CheckPermissions(c *gin.Context) {
	// if strings.Contains(excludedPaths, fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())) {
	// 	c.Next()
	// 	return
	// }

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}

	user := u.(models.User)

	access, err := m.services.Permission.Enforce(user.Role, c.FullPath(), c.Request.Method)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		return
	}
	logger.Debug("permissions", logger.StringAttr("path", c.FullPath()), logger.StringAttr("method", c.Request.Method), logger.StringAttr("access", access))
	//logger.Debug("permissions path - ", c.FullPath(), " method - ", c.Request.Method, ". permission access - ", access)

	if !access {
		response.NewErrorResponse(c, http.StatusForbidden, "access denied", "нет доступа к данному разделу")
		return
	}

	c.Next()
}
