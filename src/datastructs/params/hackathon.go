package params

import "temp/src/datastructs/entities"

// ここでは、gatewaysで扱う構造体を定義する

type HackathonCreate struct {
	Hackathon entities.Hackathon `json:"hackathon"`
	Statuses  []int32            `json:"statuses"`
}
