package request

type ListRequest struct {
	PageSize int `form:"page_size"`
	PageID   int `form:"page_id"`
}
