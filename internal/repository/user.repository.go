package repository

import (
	"DB_LAB/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)

	GetByID(ID uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
