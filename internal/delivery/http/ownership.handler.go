package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OwnershipHandler struct {
	sc service.OwnershipService
}

func NewOwnershipHandler(sc service.OwnershipService) *OwnershipHandler {
	return &OwnershipHandler{sc}
}

func (s *OwnershipHandler) GetAll(c *gin.Context) {
	ownerships, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ownerships)
}
