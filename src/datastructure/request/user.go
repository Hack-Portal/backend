package request

type InitAdmin struct {
	InitAdminToken string `json:"init_admin_token" validate:"required"`
	Name           string `json:"name" validate:"required"`
}
