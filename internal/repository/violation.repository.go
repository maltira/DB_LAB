package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipViolationRepository interface {
	GetAll() ([]entity.Violation, error)
	GetByID(id uuid.UUID) (*entity.Violation, error)

	Update(v *entity.Violation) error
	Create(v *entity.Violation) error
	Delete(id uuid.UUID) error
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

func (r *shipViolationRepository) GetByID(id uuid.UUID) (*entity.Violation, error) {
	var v *entity.Violation
	err := r.db.Preload("Inspector").Preload("Ship").Where("id = ?", id).First(&v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *shipViolationRepository) Update(v *entity.Violation) error {
	return r.db.Save(v).Error
}

func (r *shipViolationRepository) Create(v *entity.Violation) error {
	return r.db.Create(v).Error
}

func (r *shipViolationRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Violation{}, "id = ?", id).Error
}
