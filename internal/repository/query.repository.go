package repository

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

type QueryRepository interface {
	GetAll() ([]entity.Query, error)
	Update(query *entity.Query) error
}

type queryRepository struct {
	db *gorm.DB
}

func NewQueryRepository(db *gorm.DB) QueryRepository {
	return &queryRepository{db: db}
}

func (r *queryRepository) GetAll() ([]entity.Query, error) {
	var queries []entity.Query
	err := r.db.Find(&queries).Error
	if err != nil {
		return nil, err
	}
	return queries, nil
}

func (r *queryRepository) Update(query *entity.Query) error {
	return r.db.Save(&query).Error
}
