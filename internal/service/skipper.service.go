package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"

	"github.com/google/uuid"
)

type SkipperService interface {
	GetAll() ([]entity.ShipSkipper, error)
	GetByID(id uuid.UUID) (*entity.ShipSkipper, error)

	UpdateSkipper(sk *entity.ShipSkipper) error
	CreateSkipper(sk *dto.SkipperCreateRequest) error
	DeleteSkipper(id uuid.UUID) error
}

type skipperService struct {
	repo repository.ShipSkipperRepository
}

func NewSkipperService(repo repository.ShipSkipperRepository) SkipperService {
	return &skipperService{repo: repo}
}

func (s *skipperService) GetAll() ([]entity.ShipSkipper, error) {
	return s.repo.GetAll()
}

func (s *skipperService) GetByID(id uuid.UUID) (*entity.ShipSkipper, error) {
	return s.repo.GetByID(id)
}

func (s *skipperService) UpdateSkipper(sk *entity.ShipSkipper) error {
	return s.repo.Update(sk)
}

func (s *skipperService) CreateSkipper(sk *dto.SkipperCreateRequest) error {
	o := &entity.ShipSkipper{
		Name:       sk.Name,
		Surname:    sk.Surname,
		Patronymic: &sk.Patronymic,
		IDNumber:   sk.IDNumber,
	}
	return s.repo.Create(o)
}

func (s *skipperService) DeleteSkipper(id uuid.UUID) error {
	return s.repo.Delete(id)
}
