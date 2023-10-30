package param

import "github.com/hackhack-Geek-vol6/backend/src/entities"

type CreateHackathon struct {
	Hackathon  *entities.Hackathon
	StatusTags []int
}
