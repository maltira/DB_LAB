package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"time"

	"github.com/google/uuid"
)

type OwnerService interface {
	GetAll() ([]entity.ShipOwner, error)
	GetByID(id uuid.UUID) (*entity.ShipOwner, error)

	CreateOwner(shipOwner *dto.OwnerCreateRequest) error
	UpdateOwner(shipOwner *entity.ShipOwner) error
	DeleteOwner(id uuid.UUID) error
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

func (s *ownerService) CreateOwner(shipOwner *dto.OwnerCreateRequest) error {
	o := &entity.ShipOwner{
		Name:         shipOwner.Name,
		Surname:      shipOwner.Surname,
		Patronymic:   &shipOwner.Patronymic,
		Address:      shipOwner.Address,
		BirthDate:    shipOwner.BirthDate,
		Phone:        shipOwner.Phone,
		TypeOfPerson: shipOwner.TypeOfPerson,
	}
	return s.repo.Create(o)
}

func (s *ownerService) DeleteOwner(id uuid.UUID) error {
	return s.repo.Delete(id)
}
