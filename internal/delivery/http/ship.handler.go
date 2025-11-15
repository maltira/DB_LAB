package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShipHandler struct {
	sc service.ShipService
}

func NewShipHandler(sc service.ShipService) *ShipHandler {
	return &ShipHandler{sc}
}

func (s *ShipHandler) GetAllShips(c *gin.Context) {
	ships, err := s.sc.GetAllShips()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ships)
}

func (s *ShipHandler) GetAllTypes(c *gin.Context) {
	types, err := s.sc.GetAllTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}
