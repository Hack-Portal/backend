package mock

import (
	"temp/src/datastructs/entities"
	"temp/src/datastructs/params"
	"temp/src/usecases/dai"
)

// ここでは、daiに対するmockを定義する

type MockHackathonRepository struct {
	hackathon           map[string]entities.Hackathon
	hackathonStatusTags map[string][]int32
}

func NewMockHackathonRepository() dai.HackathonRepository {
	return &MockHackathonRepository{
		hackathon:           map[string]entities.Hackathon{},
		hackathonStatusTags: map[string][]int32{},
	}
}

func (m *MockHackathonRepository) Create(arg params.HackathonCreate) error {
	h := map[string]entities.Hackathon{}
	hs := map[string][]int32{}

	h[arg.Hackathon.HackathonID] = arg.Hackathon
	m.hackathon = h

	hs[arg.Hackathon.HackathonID] = arg.Statuses
	m.hackathonStatusTags = hs
	return nil
}

func (m *MockHackathonRepository) Read()   {}
func (m *MockHackathonRepository) Update() {}
func (m *MockHackathonRepository) Delete() {}
