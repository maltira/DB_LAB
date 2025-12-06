package dto

import (
	"DB_LAB/internal/entity"
	"time"

	"github.com/google/uuid"
)

type OwnerCreateRequest struct {
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Patronymic   string    `json:"patronymic"`
	Address      string    `json:"address"`
	BirthDate    time.Time `json:"birth_date"`
	Phone        string    `json:"phone"`
	TypeOfPerson string    `json:"type_of_person"`
}

type InspectorCreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
	Post       string `json:"post"`
}

type ShipCreateRequest struct {
	TypeID             uuid.UUID `json:"type_id"`
	OwnerID            uuid.UUID `json:"owner_id"`
	SkipperID          uuid.UUID `json:"skipper_id"`
	ShipNumber         string    `json:"ship_number"`
	RegistrationDate   time.Time `json:"registration_date"`
	RegistrationStatus string    `json:"registration_status"`
}

type InspectionCreateRequest struct {
	InspectorID        uuid.UUID `json:"inspector_id"`
	ShipID             uuid.UUID `json:"ship_id"`
	InspectionDate     time.Time `json:"inspection_date"`
	Result             string    `json:"result"`
	NextInspectionDate time.Time `json:"next_inspection_date"`
}

type ViolationCreateRequest struct {
	InspectorID   uuid.UUID `json:"inspector_id"`
	ShipID        uuid.UUID `json:"ship_id"`
	ViolationDate time.Time `json:"violation_date"`
	Amount        string    `json:"amount"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
}

type SkipperCreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	IDNumber   string `json:"id_number"`
}

type OwnershipCreateRequest struct {
	OldOwner     uuid.UUID `json:"old_owner"`
	NewOwner     uuid.UUID `json:"new_owner"`
	TransferDate time.Time `json:"transfer_date"`
	ShipID       uuid.UUID `json:"ship_id"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SuccessfulAuthResponse struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}
