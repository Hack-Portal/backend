package output

import "temp/src/entities"

type HackathonOutputPort interface {
	RenderCreate() error
	RenderGet(hackathons []*entities.Hackathon) error
	RenderUpdateByID(hackathon *entities.Hackathon) error
	RenderDeleteByID(hackathonID int32) error

	RenderApprove(hackathonID int32) error

	RenderError(err error) error
}
