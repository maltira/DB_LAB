package entity

import (
	"time"

	"github.com/google/uuid"
)

type ShipType struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name string    `json:"name" gorm:"not null"`
}
type Ship struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	TypeID             uuid.UUID `json:"type_id" gorm:"type:uuid;not null"`
	OwnerID            uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	SkipperID          uuid.UUID `json:"skipper_id" gorm:"type:uuid;not null"`
	ShipNumber         string    `json:"ship_number" gorm:"not null"`
	RegistrationDate   time.Time `json:"registration_date" gorm:"not null"`
	RegistrationStatus string    `json:"registration_status" gorm:"not null;check: registration_status IN ('Активный', 'Истёкший'); default:'Активный'"`

	Type    ShipType    `gorm:"foreignKey:TypeID"`
	Owner   ShipOwner   `gorm:"foreignKey:OwnerID"`
	Skipper ShipSkipper `gorm:"foreignKey:OwnerID"`
}
