package request

type CreateFollow struct {
	ToAccountID string `json:"to_account_id" binding:"required"`
}

type RemoveFollow struct {
	AccountID string `form:"account_id" binding:"required"`
}

type GetFollow struct {
	Mode     bool   `form:"mode"`
	PageSize string `form:"page_size"`
	PageID   string `form:"page_id"`
}
