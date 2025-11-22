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

func (s *InspectorHandler) GetByID(c *gin.Context) {
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

func (s *InspectorHandler) Update(c *gin.Context) {
	var req *entity.Inspector
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.UpdateInspector(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Инспектор успешно обновлен"})
}

func (s *InspectorHandler) Create(c *gin.Context) {
	var req *dto.InspectorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := s.sc.CreateInspector(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Инспектор успешно добавлен"})
}

func (s *InspectorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)
	err := s.sc.DeleteInspector(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Инспектор успешно удалён"})
}
