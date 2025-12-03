package module

import (
	"DB_LAB/internal/delivery/http"
	"DB_LAB/internal/middleware"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitOwnershipModule(db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewShipOwnershipRepository(db)
	shipRepo := repository.NewShipRepository(db)
	shipSc := service.NewShipService(shipRepo)
	sc := service.NewOwnershipService(repo)
	h := http.NewOwnershipHandler(sc, db, shipSc)

	ownershipGroup := r.Group("/ownership")
	{
		ownershipGroup.GET("/:id", h.GetByID, middleware.ValidateUUID())
		ownershipGroup.GET("/all", h.GetAll)

		ownershipGroup.POST("", h.Create)
		ownershipGroup.PUT("", h.Update)
		ownershipGroup.DELETE("/:id", h.Delete, middleware.ValidateUUID())
	}
}
