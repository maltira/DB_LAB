package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipViolationRepository interface {
	GetAll() ([]entity.Violation, error)
}

type shipViolationRepository struct {
	db *gorm.DB
}

func NewShipViolationRepository(db *gorm.DB) ShipViolationRepository {
	return &shipViolationRepository{db}
}

func (r *shipViolationRepository) GetAll() ([]entity.Violation, error) {
	var shipViolations []entity.Violation
	err := r.db.Preload("Ship").Preload("Inspector").Find(&shipViolations).Error
	if err != nil {
		return nil, err
	}
	return shipViolations, nil
}
