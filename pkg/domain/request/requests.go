package request

type ListRequest struct {
	PageSize int32 `form:"page_size"`
	PageID   int32 `form:"page_id"`
}
