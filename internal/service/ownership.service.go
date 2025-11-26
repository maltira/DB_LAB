package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"

	"github.com/google/uuid"
)

type OwnershipService interface {
	GetAll() ([]entity.ShipOwnership, error)
	GetByID(id uuid.UUID) (*entity.ShipOwnership, error)

	UpdateOwnership(sk *entity.ShipOwnership) error
	CreateOwnership(sk *dto.OwnershipCreateRequest) error
	DeleteOwnership(id uuid.UUID) error
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

func (s *ownershipService) GetByID(id uuid.UUID) (*entity.ShipOwnership, error) {
	return s.repo.GetByID(id)
}

func (s *ownershipService) UpdateOwnership(sk *entity.ShipOwnership) error {
	return s.repo.Update(sk)
}

func (s *ownershipService) CreateOwnership(sk *dto.OwnershipCreateRequest) error {
	o := &entity.ShipOwnership{
		OldOwner:     sk.OldOwner,
		NewOwner:     sk.NewOwner,
		TransferDate: sk.TransferDate,
		ShipID:       sk.ShipID,
	}
	return s.repo.Create(o)
}

func (s *ownershipService) DeleteOwnership(id uuid.UUID) error {
	return s.repo.Delete(id)
}
