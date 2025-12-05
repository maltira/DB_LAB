package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryHandler struct {
	sc service.QueryService
}

func NewQueryHandler(sc service.QueryService) *QueryHandler {
	return &QueryHandler{sc: sc}
}

func (h *QueryHandler) GetAll(c *gin.Context) {
	queries, err := h.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, queries)
}

func (h *QueryHandler) Update(c *gin.Context) {
	var req *entity.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	err := h.sc.Update(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "success update"})
}
