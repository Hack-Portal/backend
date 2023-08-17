package domain

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type CreateUserResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
