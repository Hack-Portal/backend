package input

import (
	"temp/src/entities"
	"temp/src/entities/request"
)

type HackathonInputPort interface {
	Create(hackathon request.CreateHackathon, iamge []byte) error
	Get(request.ListRequest) ([]*entities.Hackathon, error)
	UpdateByID(hackathon *entities.Hackathon) error
	DeleteByID(hackathonID int32) error

	Approve(hackathonID int32) error
}
