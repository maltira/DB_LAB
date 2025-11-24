package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"time"

	"github.com/google/uuid"
)

type InspectionService interface {
	GetAll() ([]entity.TechnicalInspection, error)
	GetByID(id uuid.UUID) (*entity.TechnicalInspection, error)

	Update(ins *entity.TechnicalInspection) error
	Create(ins *dto.InspectionCreateRequest) error
	Delete(id uuid.UUID) error
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

func (s *inspectionService) GetByID(id uuid.UUID) (*entity.TechnicalInspection, error) {
	return s.repo.GetByID(id)
}

func (s *inspectionService) Update(ins *entity.TechnicalInspection) error {
	ins.InspectionDate = ins.InspectionDate.Add(3 * time.Hour)
	ins.NextInspectionDate = ins.NextInspectionDate.Add(3 * time.Hour)
	return s.repo.Update(ins)
}

func (s *inspectionService) Create(ins *dto.InspectionCreateRequest) error {
	sh := &entity.TechnicalInspection{
		InspectorID:        ins.InspectorID,
		ShipID:             ins.ShipID,
		InspectionDate:     ins.InspectionDate.Add(3 * time.Hour),
		Result:             ins.Result,
		NextInspectionDate: ins.NextInspectionDate.Add(3 * time.Hour),
	}
	return s.repo.Create(sh)
}

func (s *inspectionService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
