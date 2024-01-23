package request

type InitAdmin struct {
	InitAdminToken string `json:"init_admin_token" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type Login struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
