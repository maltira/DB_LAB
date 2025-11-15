package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipOwnerRepository interface {
	GetAll() ([]entity.ShipOwner, error)
}

type shipOwnerRepository struct {
	db *gorm.DB
}

func NewShipOwnerRepository(db *gorm.DB) ShipOwnerRepository {
	return &shipOwnerRepository{db}
}

func (r *shipOwnerRepository) GetAll() ([]entity.ShipOwner, error) {
	var shipOwners []entity.ShipOwner
	err := r.db.Find(&shipOwners).Error
	if err != nil {
		return nil, err
	}
	return shipOwners, nil
}
