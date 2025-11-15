package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViolationHandler struct {
	sc service.ViolationService
}

func NewViolationHandler(sc service.ViolationService) *ViolationHandler {
	return &ViolationHandler{sc}
}

func (s *ViolationHandler) GetAll(c *gin.Context) {
	violations, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, violations)
}
