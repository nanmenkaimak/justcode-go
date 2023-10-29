package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
