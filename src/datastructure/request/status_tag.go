package request

type CreateStatusTag struct {
	Status string `json:"status"`
}

type GetStatusTagByID struct {
	ID int64 `param:"id" json:"id"`
}

type UpdateStatusTag struct {
	ID     int64  `param:"id" json:"id"`
	Status string `json:"status"`
}
