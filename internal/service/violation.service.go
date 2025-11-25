package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ViolationService interface {
	GetAll() ([]entity.Violation, error)
	GetByID(id uuid.UUID) (*entity.Violation, error)

	DeleteV(id uuid.UUID) error
	CreateV(v *dto.ViolationCreateRequest) error
	UpdateV(v *entity.Violation) error
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

func (s *violationService) GetByID(id uuid.UUID) (*entity.Violation, error) {
	return s.repo.GetByID(id)
}

func (s *violationService) UpdateV(v *entity.Violation) error {
	v.ViolationDate = v.ViolationDate.Add(3 * time.Hour)
	return s.repo.Update(v)
}

func (s *violationService) CreateV(v *dto.ViolationCreateRequest) error {
	o := &entity.Violation{
		InspectorID:   v.InspectorID,
		ShipID:        v.ShipID,
		ViolationDate: v.ViolationDate.Add(3 * time.Hour),
		Amount:        v.Amount,
		Description:   v.Description,
		Status:        v.Status,
	}
	return s.repo.Create(o)
}

func (s *violationService) DeleteV(id uuid.UUID) error {
	return s.repo.Delete(id)
}
