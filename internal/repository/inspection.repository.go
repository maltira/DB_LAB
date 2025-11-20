package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipInspectionRepository interface {
	GetAll() ([]entity.TechnicalInspection, error)
}

type shipInspectionRepository struct {
	db *gorm.DB
}

func NewShipInspectionRepository(db *gorm.DB) ShipInspectionRepository {
	return &shipInspectionRepository{db}
}

func (r *shipInspectionRepository) GetAll() ([]entity.TechnicalInspection, error) {
	var shipInspection []entity.TechnicalInspection
	err := r.db.Preload("Inspector").Preload("Ship").Find(&shipInspection).Error
	if err != nil {
		return nil, err
	}
	return shipInspection, nil
}
