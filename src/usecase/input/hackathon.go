package input

import (
	"github.com/hackhack-Geek-vol6/backend/src/entities"
	"github.com/hackhack-Geek-vol6/backend/src/entities/request"
)

type HackathonInputPort interface {
	Create(hackathon request.CreateHackathon, iamge []byte) error
	Get(request.ListRequest) ([]*entities.Hackathon, error)
	UpdateByID(hackathon *entities.Hackathon) error
	DeleteByID(hackathonID int32) error

	Approve(hackathonID int32) error
}
