package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ShipService interface {
	GetAllShips() ([]entity.Ship, error)
	GetAllTypes() ([]entity.ShipType, error)
	GetShipByID(id uuid.UUID) (*entity.Ship, error)

	UpdateShip(ship *entity.Ship) error
	CreateShip(ship *dto.ShipCreateRequest) error
	DeleteShip(id uuid.UUID) error
}

type shipService struct {
	repo repository.ShipRepository
}

func NewShipService(repo repository.ShipRepository) ShipService {
	return &shipService{repo: repo}
}

func (s *shipService) GetAllShips() ([]entity.Ship, error) {
	return s.repo.GetAllShips()
}

func (s *shipService) GetAllTypes() ([]entity.ShipType, error) {
	return s.repo.GetAllTypes()
}

func (s *shipService) GetShipByID(id uuid.UUID) (*entity.Ship, error) {
	return s.repo.GetShipByID(id)
}

func (s *shipService) UpdateShip(ship *entity.Ship) error {
	ship.RegistrationDate = ship.RegistrationDate.Add(3 * time.Hour)
	return s.repo.UpdateShip(ship)
}

func (s *shipService) CreateShip(ship *dto.ShipCreateRequest) error {
	sh := &entity.Ship{
		TypeID:             ship.TypeID,
		OwnerID:            ship.OwnerID,
		SkipperID:          ship.SkipperID,
		ShipNumber:         ship.ShipNumber,
		RegistrationDate:   ship.RegistrationDate.Add(3 * time.Hour),
		RegistrationStatus: ship.RegistrationStatus,
	}
	return s.repo.CreateShip(sh)
}

func (s *shipService) DeleteShip(id uuid.UUID) error {
	return s.repo.DeleteShip(id)
}
