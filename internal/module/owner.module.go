package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitOwnerModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipOwnerRepository(db)
	sc := service.NewOwnerService(repo)
	h := http.NewOwnerHandler(sc)

	ownerGroup := r.Group("/owner")
	{
		ownerGroup.GET("/all", h.GetAll)
	}
}
