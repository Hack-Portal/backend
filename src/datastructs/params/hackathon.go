package params

import "github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"

// ここでは、gatewaysで扱う構造体を定義する

type HackathonCreate struct {
	Hackathon entities.Hackathon `json:"hackathon"`
	Statuses  []int32            `json:"statuses"`
}

type HackathonReadAll struct {
	Limit   int
	Offset  int
	SortTag []int32
}
