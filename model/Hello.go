package model

type HelloRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type HelloResponse struct {
	Message string `json:"message"`
}