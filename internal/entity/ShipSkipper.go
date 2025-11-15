package entity

import "github.com/google/uuid"

type ShipSkipper struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name       string    `json:"name" gorm:"not null"`
	Surname    string    `json:"surname" gorm:"not null"`
	Patronymic *string   `json:"patronymic"`
	IDNumber   string    `json:"id_number" gorm:"not null"`
}
