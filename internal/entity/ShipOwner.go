package entity

import (
	"time"

	"github.com/google/uuid"
)

type ShipOwner struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name         string    `json:"name" gorm:"not null"`
	Surname      string    `json:"surname" gorm:"not null"`
	Patronymic   *string   `json:"patronymic"`
	Address      string    `json:"address" gorm:"not null"`
	BirthDate    time.Time `json:"birth_date" gorm:"not null"`
	Phone        string    `json:"phone" gorm:"not null"`
	TypeOfPerson string    `json:"type_of_person" gorm:"not null;check:type_of_person IN ('legal', 'private')"`

	// Связи
}
