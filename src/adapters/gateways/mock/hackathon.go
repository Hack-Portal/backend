package mock

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
)

type MockHackathonRepository struct {
	hackathons            map[string]entities.Hackathon
	hackathon_status_tags map[string][]entities.HackathonStatusTag
}

func NewMockHackathonRepository() dai.HackathonRepository {
	return &MockHackathonRepository{
		hackathons:            map[string]entities.Hackathon{},
		hackathon_status_tags: map[string][]entities.HackathonStatusTag{},
	}
}

func (m *MockHackathonRepository) Create(arg params.HackathonCreate) error {
	h := map[string]entities.Hackathon{
		arg.Hackathon.HackathonID: arg.Hackathon,
	}
	m.hackathons = h

	var hackathonStatusTag []entities.HackathonStatusTag
	for _, tag := range arg.Statuses {
		hackathonStatusTag = append(hackathonStatusTag, entities.HackathonStatusTag{
			HackathonID: arg.Hackathon.HackathonID,
			StatusID:    tag,
		})
	}

	hst := map[string][]entities.HackathonStatusTag{
		arg.Hackathon.HackathonID: hackathonStatusTag,
	}
	m.hackathon_status_tags = hst
	return nil
}

func (m *MockHackathonRepository) ReadAll(arg params.HackathonReadAll) ([]entities.Hackathon, []entities.HackathonStatus, error) {
	return nil, nil, nil
}

func (m *MockHackathonRepository) Update() {

}

func (m *MockHackathonRepository) Delete() {

}
