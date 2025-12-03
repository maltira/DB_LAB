package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipRepository interface {
	GetAllShips() ([]entity.Ship, error)
	GetAllTypes() ([]entity.ShipType, error)
	GetShipByID(id uuid.UUID) (*entity.Ship, error)

	UpdateShip(ship *entity.Ship, tx *gorm.DB) error
	CreateShip(ship *entity.Ship) error
	DeleteShip(id uuid.UUID) error
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

func (r *shipRepository) GetShipByID(id uuid.UUID) (*entity.Ship, error) {
	var ship *entity.Ship
	err := r.db.Preload("Type").Preload("Owner").Preload("Skipper").Where("id = ?", id).First(&ship).Error
	if err != nil {
		return nil, err
	}
	return ship, nil
}

func (r *shipRepository) UpdateShip(ship *entity.Ship, tx *gorm.DB) error {
	if tx != nil {
		return tx.Save(ship).Error
	}
	return r.db.Save(ship).Error
}

func (r *shipRepository) CreateShip(ship *entity.Ship) error {
	return r.db.Create(ship).Error
}

func (r *shipRepository) DeleteShip(id uuid.UUID) error {
	return r.db.Delete(&entity.Ship{}, "id = ?", id).Error
}
