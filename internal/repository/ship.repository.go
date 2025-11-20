package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipRepository interface {
	GetAllShips() ([]entity.Ship, error)
	GetAllTypes() ([]entity.ShipType, error)
}

type shipRepository struct {
	db *gorm.DB
}

func NewShipRepository(db *gorm.DB) ShipRepository {
	return &shipRepository{db}
}

func (r *shipRepository) GetAllShips() ([]entity.Ship, error) {
	var ships []entity.Ship
	err := r.db.Preload("Type").Preload("Owner").Preload("Skipper").Find(&ships).Error
	if err != nil {
		return nil, err
	}
	return ships, nil
}

func (r *shipRepository) GetAllTypes() ([]entity.ShipType, error) {
	var types []entity.ShipType
	err := r.db.Find(&types).Error
	if err != nil {
		return nil, err
	}
	return types, nil
}
