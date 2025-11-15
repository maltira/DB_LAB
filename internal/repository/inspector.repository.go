package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type InspectorRepository interface {
	GetAll() ([]entity.Inspector, error)
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
