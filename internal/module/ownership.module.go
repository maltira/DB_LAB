package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitOwnershipModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipOwnershipRepository(db)
	sc := service.NewOwnershipService(repo)
	h := http.NewOwnershipHandler(sc)

	ownershipGroup := r.Group("/ownership")
	{
		ownershipGroup.GET("/all", h.GetAll)
	}
}
