package outputboundary

import "github.com/hackhack-Geek-vol6/backend/src/datastructs/output"

// ここでは主にPresenterを汎化している

// TODO:ここもどうようにするか考える
type HackathonOutputPort interface {
	Create(error) (int, *output.CreateHackathon)
	Read()
	Update()
	Delete()
}
