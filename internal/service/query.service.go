package service

import (
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
)

type QueryService interface {
	GetAll() ([]entity.Query, error)
	Update(q *entity.Query) error
}

type queryService struct {
	repo repository.QueryRepository
}

func NewQueryService(repo repository.QueryRepository) QueryService {
	return &queryService{repo: repo}
}

func (qs *queryService) GetAll() ([]entity.Query, error) {
	return qs.repo.GetAll()
}

func (qs *queryService) Update(q *entity.Query) error {
	return qs.repo.Update(q)
}
