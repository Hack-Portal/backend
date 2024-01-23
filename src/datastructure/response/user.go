package response

type User struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Login struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
