package service

import (
	"DB_LAB/internal/dto"
	"DB_LAB/internal/entity"
	"DB_LAB/internal/repository"
	"DB_LAB/internal/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Create(usr *dto.CreateUserRequest) (*entity.User, error)

	GetByID(ID uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) Create(usr *dto.CreateUserRequest) (*entity.User, error) {
	passwordHash, err := utils.HashPassword(usr.Password)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:     usr.Name,
		Email:    usr.Email,
		Password: passwordHash,
		IsAdmin:  usr.IsAdmin,
	}
	return u.repo.Create(user)
}

func (u *userService) GetByID(id uuid.UUID) (*entity.User, error) {
	return u.repo.GetByID(id)
}
func (u *userService) GetByEmail(email string) (*entity.User, error) {
	return u.repo.GetByEmail(email)
}
