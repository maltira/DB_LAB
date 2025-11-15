package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type ShipService interface {
	GetAllShips() ([]entity.Ship, error)
	GetAllTypes() ([]entity.ShipType, error)
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
