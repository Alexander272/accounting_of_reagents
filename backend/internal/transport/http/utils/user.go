package utils

import (
	"net/http"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/constants"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/gin-gonic/gin"
)

// func GetActor(c *gin.Context) *models.Actor {
// 	u, exists := c.Get(constants.CtxUser)
// 	if !exists {
// 		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "Сессия не найдена")
// 		return nil
// 	}
// 	user := u.(models.User)

// 	actor := &models.Actor{
// 		ID:   user.ID,
// 		Name: user.Name,
// 	}
// 	return actor
// }

func GetUser(c *gin.Context) *models.User {
	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "Сессия не найдена")
		return nil
	}
	user := u.(models.User)
	return &user
}
