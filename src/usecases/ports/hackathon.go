package ports

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
)

type HackathonInputBoundary interface {
	Create(input.HackathonCreate, []byte) (int, *output.CreateHackathon)
	ReadAll(input.HackathonReadAll) (int, []*output.ReadAllHackathon)
	Update()
	Delete()
}
type HackathonOutputBoundary interface {
	Create(error) (int, *output.CreateHackathon)
	ReadAll([]entities.Hackathon, []entities.HackathonStatus, error) (int, []*output.ReadAllHackathon)
	Update()
	Delete()
}
