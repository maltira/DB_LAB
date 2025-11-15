package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type OwnershipService interface {
	GetAll() ([]entity.ShipOwnership, error)
}

type ownershipService struct {
	repo repository.ShipOwnershipRepository
}

func NewOwnershipService(repo repository.ShipOwnershipRepository) OwnershipService {
	return &ownershipService{repo: repo}
}

func (s *ownershipService) GetAll() ([]entity.ShipOwnership, error) {
	return s.repo.GetAll()
}
