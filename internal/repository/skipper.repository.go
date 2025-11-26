package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipSkipperRepository interface {
	GetAll() ([]entity.ShipSkipper, error)
	GetByID(id uuid.UUID) (*entity.ShipSkipper, error)

	Update(s *entity.ShipSkipper) error
	Create(s *entity.ShipSkipper) error
	Delete(id uuid.UUID) error
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

func (r *shipSkipperRepository) GetByID(id uuid.UUID) (*entity.ShipSkipper, error) {
	var s *entity.ShipSkipper
	err := r.db.Where("id = ?", id).First(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *shipSkipperRepository) Update(s *entity.ShipSkipper) error {
	return r.db.Save(s).Error
}

func (r *shipSkipperRepository) Create(s *entity.ShipSkipper) error {
	return r.db.Create(s).Error
}

func (r *shipSkipperRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.ShipSkipper{}, "id = ?", id).Error
}
