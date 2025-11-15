package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InspectionHandler struct {
	sc service.InspectionService
}

func NewInspectionHandler(sc service.InspectionService) *InspectionHandler {
	return &InspectionHandler{sc}
}

func (s *InspectionHandler) GetAll(c *gin.Context) {
	inspections, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, inspections)
}
