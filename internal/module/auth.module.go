package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthModule(r *gin.RouterGroup, db *gorm.DB) {
	repoUser := repository.NewUserRepository(db)
	scUser := service.NewUserService(repoUser)

	h := http.NewAuthHandler(scUser)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/logout", h.Logout)
		authGroup.GET("/status", h.AuthStatus)
	}
}
