package models

type User struct {
	// TODO: Implement User model
	ID string `json:"id"`
	Username string `json:"username" validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=5,max=100"`
}
