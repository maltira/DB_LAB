package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipSkipperRepository interface {
	GetAll() ([]entity.ShipSkipper, error)
}

type shipSkipperRepository struct {
	db *gorm.DB
}

func NewShipSkipperRepository(db *gorm.DB) ShipSkipperRepository {
	return &shipSkipperRepository{db}
}

func (r *shipSkipperRepository) GetAll() ([]entity.ShipSkipper, error) {
	var shipSkippers []entity.ShipSkipper
	err := r.db.Find(&shipSkippers).Error
	if err != nil {
		return nil, err
	}
	return shipSkippers, nil
}
