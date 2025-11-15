package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type InspectorService interface {
	GetAll() ([]entity.Inspector, error)
}

type inspectorService struct {
	repo repository.InspectorRepository
}

func NewInspectorService(repo repository.InspectorRepository) InspectorService {
	return &inspectorService{repo: repo}
}

func (s *inspectorService) GetAll() ([]entity.Inspector, error) {
	return s.repo.GetAll()
}
