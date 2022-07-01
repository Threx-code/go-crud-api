package users

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated"`
}

type CreateRequest struct {
	Email     string `validate:"required,email"`
	Firstname string `validate:"required,alpha"`
	Lastname  string `validate:"required,alpha"`
	Password  string `validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	User        UserResponse
	AccessToken string `json:"access_token"`
}

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func NewUserResponse(user *Users) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type ResError struct {
	Error string `json:"error"`
}
