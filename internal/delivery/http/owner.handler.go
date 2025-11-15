package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OwnerHandler struct {
	sc service.OwnerService
}

func NewOwnerHandler(sc service.OwnerService) *OwnerHandler {
	return &OwnerHandler{sc}
}

func (s *OwnerHandler) GetAll(c *gin.Context) {
	owners, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, owners)
}
