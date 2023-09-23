package params

import (
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type CreateHackathonParams struct {
	Hackathon  repository.CreateHackathonsParams
	StatusTags []int32 `json:"status_tags"`
}
