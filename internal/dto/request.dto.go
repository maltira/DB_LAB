package dto

import "time"

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
