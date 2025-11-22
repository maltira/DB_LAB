package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"

	"github.com/google/uuid"
)

type InspectorService interface {
	GetAll() ([]entity.Inspector, error)
	GetByID(id uuid.UUID) (*entity.Inspector, error)

	CreateInspector(ins *dto.InspectorCreateRequest) error
	UpdateInspector(ins *entity.Inspector) error
	DeleteInspector(id uuid.UUID) error
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

func (s *inspectorService) GetByID(id uuid.UUID) (*entity.Inspector, error) {
	return s.repo.GetByID(id)
}

func (s *inspectorService) UpdateInspector(ins *entity.Inspector) error {
	return s.repo.Update(ins)
}

func (s *inspectorService) CreateInspector(ins *dto.InspectorCreateRequest) error {
	i := &entity.Inspector{
		Name:       ins.Name,
		Surname:    ins.Surname,
		Patronymic: &ins.Patronymic,
		Phone:      ins.Phone,
		Post:       ins.Post,
	}
	return s.repo.Create(i)
}

func (s *inspectorService) DeleteInspector(id uuid.UUID) error {
	return s.repo.Delete(id)
}
