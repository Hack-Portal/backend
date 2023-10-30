package param

import "temp/src/entities"

type CreateHackathon struct {
	Hackathon  *entities.Hackathon
	StatusTags []int
}
