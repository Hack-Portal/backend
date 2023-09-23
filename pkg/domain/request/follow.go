package request

type CreateFollowRequestBody struct {
	ToAccountID string `json:"to_account_id" binding:"required"`
}

type RemoveFollowRequestQueries struct {
	AccountID string `form:"account_id" binding:"required"`
}

type GetFollowRequestQueries struct {
	Mode     bool   `form:"mode"`
	PageSize string `form:"page_size"`
	PageID   string `form:"page_id"`
}
