package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitQueryModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewQueryRepository(db)
	sc := service.NewQueryService(repo)
	h := http.NewQueryHandler(sc)

	queryGroup := r.Group("query")
	{
		queryGroup.GET("/all", h.GetAll)
		queryGroup.PUT("", h.Update)
	}
}
