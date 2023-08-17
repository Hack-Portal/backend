package domain

import "time"

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	UserID    string    `json:"userID"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
