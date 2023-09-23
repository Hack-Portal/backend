package request

type CreateLike struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}

type RemoveLike struct {
	Opus int32 `form:"opus" binding:"required"`
}
