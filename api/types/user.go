package types

import "time"

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUser struct {
	Firstname string `json:"firstName" validate:"required"`
	Lastname  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=20"`
}
