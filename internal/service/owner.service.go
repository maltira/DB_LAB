package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"time"

	"github.com/google/uuid"
)

type OwnerService interface {
	GetAll() ([]entity.ShipOwner, error)
	GetByID(id uuid.UUID) (*entity.ShipOwner, error)

	UpdateOwner(shipOwner *entity.ShipOwner) error
}

type ownerService struct {
	repo repository.ShipOwnerRepository
}

func NewOwnerService(repo repository.ShipOwnerRepository) OwnerService {
	return &ownerService{repo: repo}
}

func (s *ownerService) GetAll() ([]entity.ShipOwner, error) {
	return s.repo.GetAll()
}

func (s *ownerService) GetByID(id uuid.UUID) (*entity.ShipOwner, error) {
	return s.repo.GetByID(id)
}

func (s *ownerService) UpdateOwner(shipOwner *entity.ShipOwner) error {
	shipOwner.BirthDate = shipOwner.BirthDate.Add(3 * time.Hour)
	return s.repo.Update(shipOwner)
}
