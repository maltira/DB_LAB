package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"DB_LAB/internal/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: err.Error()})
		return
	}

	// Получаем пользоваетля по почте
	user, err := h.userService.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Указан неверный email или пароль"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	// Проверяем соответствие пароля и генерируем токен
	isPasswordCorrect := utils.CheckPasswordHash(req.Password, user.Password)
	token := ""

	if !isPasswordCorrect {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Указан неверный email или пароль"})
		return
	} else {
		token, err = utils.GenerateToken(user.ID, user.IsAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
			return
		}
		c.SetCookie("token", token, 3600*24, "/", "", false, true)
	}
	c.JSON(http.StatusOK, dto.SuccessfulAuthResponse{
		User:  *user,
		Token: token,
	})
}

func (h *AuthHandler) AuthStatus(c *gin.Context) {
	tokenStr, _ := c.Cookie("token")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessfulAuthResponse{
		Token: tokenStr,
		User:  *user,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "logged out successfully"})
}
