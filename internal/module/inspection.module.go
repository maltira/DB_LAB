package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitInspectionModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipInspectionRepository(db)
	sc := service.NewInspectionService(repo)
	h := http.NewInspectionHandler(sc)

	inspectionGroup := r.Group("/inspection")
	{
		inspectionGroup.GET("/all", h.GetAll)
	}
}
