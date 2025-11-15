package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkipperHandler struct {
	sc service.SkipperService
}

func NewSkipperHandler(sc service.SkipperService) *SkipperHandler {
	return &SkipperHandler{sc}
}

func (s *SkipperHandler) GetAll(c *gin.Context) {
	skippers, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, skippers)
}
