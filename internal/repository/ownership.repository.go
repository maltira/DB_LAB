package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipOwnershipRepository interface {
	GetAll() ([]entity.ShipOwnership, error)
	GetByID(id uuid.UUID) (*entity.ShipOwnership, error)

	Create(s *entity.ShipOwnership) error
	Update(s *entity.ShipOwnership) error
	Delete(id uuid.UUID) error
}

type shipOwnershipRepository struct {
	db *gorm.DB
}

func NewShipOwnershipRepository(db *gorm.DB) ShipOwnershipRepository {
	return &shipOwnershipRepository{db}
}

func (r *shipOwnershipRepository) GetAll() ([]entity.ShipOwnership, error) {
	var shipOwnerships []entity.ShipOwnership
	err := r.db.Preload("Ship").Preload("ShipOldOwner").Preload("ShipNewOwner").Find(&shipOwnerships).Error
	if err != nil {
		return nil, err
	}
	return shipOwnerships, nil
}

func (r *shipOwnershipRepository) GetByID(id uuid.UUID) (*entity.ShipOwnership, error) {
	var s *entity.ShipOwnership
	err := r.db.
		Preload("Ship").
		Preload("ShipOldOwner").
		Preload("ShipNewOwner").
		Where("id = ?", id).First(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *shipOwnershipRepository) Update(s *entity.ShipOwnership) error {
	return r.db.Save(s).Error
}

func (r *shipOwnershipRepository) Create(s *entity.ShipOwnership) error {
	return r.db.Create(s).Error
}

func (r *shipOwnershipRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.ShipOwnership{}, "id = ?", id).Error
}
