package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/middleware"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserModule(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	sc := service.NewUserService(repo)
	h := http.NewUserHandler(sc)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", h.GetUserByID, middleware.ValidateUUID(), middleware.AuthMiddleware())
		userGroup.GET("", h.GetCurrentUser, middleware.AuthMiddleware())
		userGroup.GET("/email/:email", h.GetUserByEmail)

		userGroup.POST("", h.CreateUser)
	}

}
