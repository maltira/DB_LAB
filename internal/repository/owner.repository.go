package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipOwnerRepository interface {
	GetAll() ([]entity.ShipOwner, error)
	GetByID(id uuid.UUID) (*entity.ShipOwner, error)

	Create(shipOwner *entity.ShipOwner) error
	Update(shipOwner *entity.ShipOwner) error
	Delete(id uuid.UUID) error
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

func (r *shipOwnerRepository) GetByID(id uuid.UUID) (*entity.ShipOwner, error) {
	var shipOwner *entity.ShipOwner
	err := r.db.Where("id = ?", id).First(&shipOwner).Error
	if err != nil {
		return nil, err
	}
	return shipOwner, nil
}

func (r *shipOwnerRepository) Update(shipOwner *entity.ShipOwner) error {
	return r.db.Save(shipOwner).Error
}

func (r *shipOwnerRepository) Create(shipOwner *entity.ShipOwner) error {
	return r.db.Create(shipOwner).Error
}

func (r *shipOwnerRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.ShipOwner{}, "id = ?", id).Error
}
