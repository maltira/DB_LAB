package http

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/service"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OwnershipHandler struct {
	sc     service.OwnershipService
	shipSc service.ShipService
	db     *gorm.DB
}

func NewOwnershipHandler(sc service.OwnershipService, db *gorm.DB, shipSc service.ShipService) *OwnershipHandler {
	return &OwnershipHandler{sc, shipSc, db}
}

func (s *OwnershipHandler) GetAll(c *gin.Context) {
	ownerships, err := s.sc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ownerships)
}

func (s *OwnershipHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)

	own, err := s.sc.GetByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, own)
}

func (s *OwnershipHandler) Update(c *gin.Context) {
	var req *entity.ShipOwnership
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	tx := s.db.Begin()
	err := s.sc.UpdateOwnership(req, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		tx.Rollback()
		return
	}
	ship, err := s.shipSc.GetShipByID(req.ShipID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		tx.Rollback()
		return
	}
	if req.NewOwner != ship.OwnerID {
		err = s.shipSc.UpdateShip(&entity.Ship{
			ID:                 ship.ID,
			TypeID:             ship.TypeID,
			OwnerID:            req.NewOwner,
			SkipperID:          ship.SkipperID,
			ShipNumber:         ship.ShipNumber,
			RegistrationDate:   ship.RegistrationDate.Add(-3 * time.Hour),
			RegistrationStatus: ship.RegistrationStatus,
		}, tx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Запись успешно обновлена"})
}

func (s *OwnershipHandler) Create(c *gin.Context) {
	var req *dto.OwnershipCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: http.StatusBadRequest, Error: "Некорректные данные в BODY"})
		return
	}
	tx := s.db.Begin()
	err := s.sc.CreateOwnership(req, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		tx.Rollback()
		return
	}
	ship, err := s.shipSc.GetShipByID(req.ShipID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		tx.Rollback()
		return
	}
	err = s.shipSc.UpdateShip(&entity.Ship{
		ID:                 ship.ID,
		TypeID:             ship.TypeID,
		OwnerID:            req.NewOwner,
		SkipperID:          ship.SkipperID,
		ShipNumber:         ship.ShipNumber,
		RegistrationDate:   ship.RegistrationDate.Add(-3 * time.Hour),
		RegistrationStatus: ship.RegistrationStatus,
	}, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		tx.Rollback()
		return
	}
	tx.Commit()
	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Запись успешно добавлена"})
}

func (s *OwnershipHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ID := uuid.MustParse(id)
	err := s.sc.DeleteOwnership(ID)
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
