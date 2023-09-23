package request

type CreateBookmark struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}
type RemoveBookmark struct {
	Opus int32 `query:"opus" binding:"required"`
}
