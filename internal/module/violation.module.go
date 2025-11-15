package module

import (
	"DB_LAB/internal/delivery/http"
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
		violationGroup.GET("/all", h.GetAll)
	}
}
