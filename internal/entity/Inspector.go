package entity

import "github.com/google/uuid"

type Inspector struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name       string    `json:"name" gorm:"not null"`
	Surname    string    `json:"surname" gorm:"not null"`
	Patronymic *string   `json:"patronymic"`
	Post       string    `json:"post" gorm:"not null"`
	Phone      string    `json:"phone" gorm:"not null"`
}
