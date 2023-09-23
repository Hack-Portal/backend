package params

import (
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type CreateHackathon struct {
	Hackathon  repository.CreateHackathonsParams
	StatusTags []int32 `json:"status_tags"`
}
