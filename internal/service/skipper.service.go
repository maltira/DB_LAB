package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type SkipperService interface {
	GetAll() ([]entity.ShipSkipper, error)
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
