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

func (s *SkipperHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)

	skipper, err := s.sc.GetByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, skipper)
}

func (s *SkipperHandler) Update(c *gin.Context) {
	var req *entity.ShipSkipper
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.UpdateSkipper(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Запись успешно обновлена"})
}

func (s *SkipperHandler) Create(c *gin.Context) {
	var req *dto.SkipperCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.CreateSkipper(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Запись успешно добавлена"})
}

func (s *SkipperHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)
	err := s.sc.DeleteSkipper(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Запись успешно удалена"})
}
