package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitShipModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipRepository(db)
	sc := service.NewShipService(repo)
	h := http.NewShipHandler(sc)

	shipGroup := r.Group("/ship")
	{
		shipGroup.GET("/all", h.GetAllShips)
		shipGroup.GET("/type/all", h.GetAllTypes)
	}
}
