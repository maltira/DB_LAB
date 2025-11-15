package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitInspectorModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewInspectorRepository(db)
	sc := service.NewInspectorService(repo)
	h := http.NewInspectorHandler(sc)

	inspectorGroup := r.Group("/inspector")
	{
		inspectorGroup.GET("/all", h.GetAll)
	}
}
