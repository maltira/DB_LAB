package middleware

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token") // проверка существования токена
		if err != nil {                    // токена нет
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
			return
		}

		// токен есть - валидация токена
		userID, isAdmin, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
			return
		}
		c.Set("userID", userID)
		c.Set("isAdmin", isAdmin)
		c.Next()
	}
}
