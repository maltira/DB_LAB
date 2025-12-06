package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
	IsAdmin  bool      `json:"is_admin" gorm:"not null;default:false"`
}
