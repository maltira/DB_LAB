package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type ViolationService interface {
	GetAll() ([]entity.Violation, error)
}

type violationService struct {
	repo repository.ShipViolationRepository
}

func NewViolationService(repo repository.ShipViolationRepository) ViolationService {
	return &violationService{repo: repo}
}

func (s *violationService) GetAll() ([]entity.Violation, error) {
	return s.repo.GetAll()
}
