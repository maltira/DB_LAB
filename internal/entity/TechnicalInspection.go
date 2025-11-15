package entity

import (
	"time"

	"github.com/google/uuid"
)

type TechnicalInspection struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	InspectorID        uuid.UUID `json:"inspector_id" gorm:"type:uuid;not null"`
	ShipID             uuid.UUID `json:"ship_id" gorm:"type:uuid;not null"`
	InspectionDate     time.Time `json:"inspection_date" gorm:"not null"`
	Result             string    `json:"result" gorm:"not null; check: result IN ('Годно к эксплутации', 'Годно с замечаниями', 'Ограниченно годно', 'Не годно к эксплутации')"`
	NextInspectionDate time.Time `json:"next_inspection_date" gorm:"not null"`

	Inspector Inspector `gorm:"foreignkey:InspectorID"`
	Ship      Ship      `gorm:"foreignkey:ShipID"`
}
