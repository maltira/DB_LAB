package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *ViolationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)

	v, err := s.sc.GetByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, v)
}

func (s *ViolationHandler) Update(c *gin.Context) {
	var req *entity.Violation
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.UpdateV(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Нарушение успешно обновлено"})
}

func (s *ViolationHandler) Create(c *gin.Context) {
	var req *dto.ViolationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.CreateV(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Нарушение успешно добавлено"})
}

func (s *ViolationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)
	err := s.sc.DeleteV(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Нарушение успешно удалено"})
}
