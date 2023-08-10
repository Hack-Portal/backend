package domain

type CreateFollowRequestBody struct {
	ToAccountID string `json:"to_account_id" binding:"required"`
}

type RemoveFollowRequestQueries struct {
	ToAccountID string `from:"to_account_id" binding:"required"`
}

type GetFollowRequestQueries struct {
	Mode     bool   `form:"mode"`
	PageSize string `form:"page_size"`
	PageID   string `form:"page_id"`
}

type FollowResponse struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Icon      string `json:"icon"`
}
