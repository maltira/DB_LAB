package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *InspectionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)

	ins, err := s.sc.GetByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ins)
}

func (s *InspectionHandler) Update(c *gin.Context) {
	var req *entity.TechnicalInspection
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.Update(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Осмотр успешно обновлен"})
}

func (s *InspectionHandler) Create(c *gin.Context) {
	var req *dto.InspectionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Осмотр успешно добавлен"})
}

func (s *InspectionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)
	err := s.sc.Delete(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Осмотр успешно удален"})
}
