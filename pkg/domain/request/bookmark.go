package request

type CreateBookmarkRequest struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}

type BookmarkRequestWildCard struct {
	AccountID string `uri:"account_id"`
}

type RemoveBookmarkRequestQueries struct {
	Opus int32 `query:"opus" binding:"required"`
}
