package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type InspectionService interface {
	GetAll() ([]entity.TechnicalInspection, error)
}

type inspectionService struct {
	repo repository.ShipInspectionRepository
}

func NewInspectionService(repo repository.ShipInspectionRepository) InspectionService {
	return &inspectionService{repo: repo}
}

func (s *inspectionService) GetAll() ([]entity.TechnicalInspection, error) {
	return s.repo.GetAll()
}
