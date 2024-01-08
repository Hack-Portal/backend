package response

type User struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
