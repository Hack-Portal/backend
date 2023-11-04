package ports

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
)

// ここでは主にPresenterを汎化している

// TODO:ここもどうようにするか考える
type HackathonOutputBoundary interface {
	Create(error) (int, *output.CreateHackathon)
	Read()
	Update()
	Delete()
}

type HackathonInputBoundary interface {
	Create(input.HackathonCreate, []byte) (int, *output.CreateHackathon)
	Read()
	Update()
	Delete()
}
