package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InspectorHandler struct {
	sc service.InspectorService
}

func NewInspectorHandler(sc service.InspectorService) *InspectorHandler {
	return &InspectorHandler{sc}
}

func (s *InspectorHandler) GetAll(c *gin.Context) {
	inspectors, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, inspectors)
}
