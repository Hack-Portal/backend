package response

type FollowResponse struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Icon      string `json:"icon"`
}
