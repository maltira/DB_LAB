package entity

import (
	"time"

	"github.com/google/uuid"
)

type ShipOwnership struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	OldOwner     uuid.UUID `json:"old_owner" gorm:"type:uuid;not null;check: old_owner != new_owner"`
	NewOwner     uuid.UUID `json:"new_owner" gorm:"type:uuid;not null"`
	TransferDate time.Time `json:"transfer_date" gorm:"not null"`
	ShipID       uuid.UUID `json:"ship_id" gorm:"type:uuid;not null"`

	ShipOldOwner ShipOwner `gorm:"foreignkey:OldOwner"`
	ShipNewOwner ShipOwner `gorm:"foreignkey:NewOwner"`
	Ship         Ship      `gorm:"foreignkey:ShipID"`
}
