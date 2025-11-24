package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipInspectionRepository interface {
	GetAll() ([]entity.TechnicalInspection, error)
	GetByID(id uuid.UUID) (*entity.TechnicalInspection, error)

	Update(ins *entity.TechnicalInspection) error
	Create(ins *entity.TechnicalInspection) error
	Delete(id uuid.UUID) error
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

func (r *shipInspectionRepository) GetByID(id uuid.UUID) (*entity.TechnicalInspection, error) {
	var ins *entity.TechnicalInspection
	err := r.db.Preload("Inspector").Preload("Ship").Where("id = ?", id).First(&ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (r *shipInspectionRepository) Update(ins *entity.TechnicalInspection) error {
	return r.db.Save(ins).Error
}

func (r *shipInspectionRepository) Create(ins *entity.TechnicalInspection) error {
	return r.db.Create(ins).Error
}

func (r *shipInspectionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.TechnicalInspection{}, "id = ?", id).Error
}
