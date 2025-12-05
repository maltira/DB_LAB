package entity

import "github.com/google/uuid"

type Query struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Route  string    `json:"route" gorm:"not null"`
	Name   string    `json:"name" gorm:"not null"`
	Access bool      `json:"access" gorm:"not null;default:true"`
}
