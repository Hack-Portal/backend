package domain

type CreateRateRequestBody struct {
	Rate int32 `json:"rate"`
}

type ListRateParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
}
