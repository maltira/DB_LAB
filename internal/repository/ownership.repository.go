package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type ShipOwnershipRepository interface {
	GetAll() ([]entity.ShipOwnership, error)
}

type shipOwnershipRepository struct {
	db *gorm.DB
}

func NewShipOwnershipRepository(db *gorm.DB) ShipOwnershipRepository {
	return &shipOwnershipRepository{db}
}

func (r *shipOwnershipRepository) GetAll() ([]entity.ShipOwnership, error) {
	var shipOwnerships []entity.ShipOwnership
	err := r.db.Preload("Ship").Find(&shipOwnerships).Error
	if err != nil {
		return nil, err
	}
	return shipOwnerships, nil
}
