package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/middleware"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSkipperModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipSkipperRepository(db)
	sc := service.NewSkipperService(repo)
	h := http.NewSkipperHandler(sc)

	skipperGroup := r.Group("/skipper")
	{
		skipperGroup.GET("/:id", h.GetByID, middleware.ValidateUUID())
		skipperGroup.GET("/all", h.GetAll)

		skipperGroup.POST("", h.Create)
		skipperGroup.PUT("", h.Update)
		skipperGroup.DELETE("/:id", h.Delete, middleware.ValidateUUID())
	}
}
