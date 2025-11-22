package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InspectorRepository interface {
	GetAll() ([]entity.Inspector, error)
	GetByID(id uuid.UUID) (*entity.Inspector, error)

	Create(ins *entity.Inspector) error
	Update(ins *entity.Inspector) error
	Delete(id uuid.UUID) error
}

type inspectorRepository struct {
	db *gorm.DB
}

func NewInspectorRepository(db *gorm.DB) InspectorRepository {
	return &inspectorRepository{db}
}

func (r *inspectorRepository) GetAll() ([]entity.Inspector, error) {
	var inspectors []entity.Inspector
	err := r.db.Find(&inspectors).Error
	if err != nil {
		return nil, err
	}
	return inspectors, nil
}

func (r *inspectorRepository) GetByID(id uuid.UUID) (*entity.Inspector, error) {
	var ins *entity.Inspector
	err := r.db.Where("id = ?", id).First(&ins).Error
	if err != nil {
		return &entity.Inspector{}, err
	}
	return ins, nil
}

func (r *inspectorRepository) Update(ins *entity.Inspector) error {
	return r.db.Save(ins).Error
}

func (r *inspectorRepository) Create(ins *entity.Inspector) error {
	return r.db.Create(ins).Error
}

func (r *inspectorRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Inspector{}, "id = ?", id).Error
}
