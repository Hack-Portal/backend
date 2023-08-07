package domain

type CreateFollowRequestBody struct {
	ToUserID string `json:"to_user_id" binding:"required"`
}

type RemoveFollowRequestQueries struct {
	ToUserID string `json:"to_user_id" binding:"required"`
}

type GetFollowRequestQueries struct {
	Mode     bool   `form:"mode"`
	PageSize string `form:"page_size"`
	PageID   string `form:"page_id"`
}

type FollowResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}
