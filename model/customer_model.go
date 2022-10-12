package model

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required"`
}
