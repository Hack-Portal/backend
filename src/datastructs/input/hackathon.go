package input

import "time"

// ここでは Controllerでバインドしたい構造体を定義する
// 所謂リクエストボディ等の構造体を定義する

type HackathonCreate struct {
	Name       string    `form:"name"`
	Link       string    `form:"link"`
	Expired    time.Time `form:"expired"`
	StartDate  time.Time `form:"start_date"`
	Term       int32     `form:"term"`
	StatusTags string    `form:"status_tags"`
}

type HackathonReadAll struct {
	PageSize int     `form:"page_size"`
	PageID   int     `form:"page_id"`
	SortTag  []int32 `form:"size_tag[]"`
}

type HackathonUpdate struct{}
type HackathonDelete struct{}
