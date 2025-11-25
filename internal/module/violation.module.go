package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/middleware"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitViolationModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipViolationRepository(db)
	sc := service.NewViolationService(repo)
	h := http.NewViolationHandler(sc)

	violationGroup := r.Group("/violation")
	{
		violationGroup.GET("/:id", h.GetByID, middleware.ValidateUUID())
		violationGroup.GET("/all", h.GetAll)

		violationGroup.POST("", h.Create)
		violationGroup.PUT("", h.Update)
		violationGroup.DELETE("/:id", h.Delete, middleware.ValidateUUID())
	}
}
