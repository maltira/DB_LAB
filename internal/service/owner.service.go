package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type OwnerService interface {
	GetAll() ([]entity.ShipOwner, error)
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
