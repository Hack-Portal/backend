package ports

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
)

type HackathonInputBoundary interface {
	Create(input.HackathonCreate, []byte) (int, *output.CreateHackathon)
	Read()
	Update()
	Delete()
}
type HackathonOutputBoundary interface {
	Create(error) (int, *output.CreateHackathon)
	Read()
	Update()
	Delete()
}
