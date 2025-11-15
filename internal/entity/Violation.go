package entity

import (
	"time"

	"github.com/google/uuid"
)

type Violation struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	ShipID        uuid.UUID `json:"ship_id" gorm:"type:uuid;not null"`
	InspectorID   uuid.UUID `json:"inspector_id" gorm:"type:uuid;not null"`
	ViolationDate time.Time `json:"violation_date" gorm:"not null"`
	Amount        int       `json:"amount" gorm:"not null"`
	Description   string    `json:"description" gorm:"not null"`
	Status        string    `json:"status" gorm:"not null; check: status IN ('Исполнено', 'Не исполонено'); default: 'Не исполонено'"`

	Ship      Ship      `gorm:"foreignKey:ShipID"`
	Inspector Inspector `gorm:"foreignKey:InspectorID"`
}
