package dto

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
