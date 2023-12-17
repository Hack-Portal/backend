package request

type CreateStatusTag struct {
	Status string `json:"status"`
}

type GetStatusTagByID struct {
	ID int64 `json:"id"`
}

type UpdateStatusTag struct {
	ID int64 `json:"id"`
	CreateStatusTag
}
