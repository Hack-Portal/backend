package presenters

import (
	"net/http"

	"github.com/hackhack-Geek-vol6/backend/src/datastructs/cerror"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
)

type HackathonPresenter struct {
}

func NewHackathonOutputBoundary() ports.HackathonOutputBoundary {
	return &HackathonPresenter{}
}

func (h *HackathonPresenter) Create(err error) (int, *output.CreateHackathon) {
	if err != nil {
		switch err {
		case cerror.ImageNotFound:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &output.CreateHackathon{
		Message: "Hackathon created successfully",
	}
}

func (h *HackathonPresenter) ReadAll(hackathons []entities.Hackathon, statuses []entities.HackathonStatus, err error) (int, []*output.ReadAllHackathon) {
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	var readAllHackathons []*output.ReadAllHackathon

	for _, hackathon := range hackathons {
		var statusTagsForHackathon []entities.StatusTag

		for _, hackathonStatus := range statuses {
			if hackathonStatus.HackathonID == hackathon.HackathonID {
				statusTagsForHackathon = append(statusTagsForHackathon, entities.StatusTag{
					StatusID: hackathonStatus.StatusID,
					Status:   hackathonStatus.Status,
				})
			}
		}

		readAllHackathon := &output.ReadAllHackathon{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  statusTagsForHackathon,
		}

		readAllHackathons = append(readAllHackathons, readAllHackathon)
	}
	return http.StatusOK, readAllHackathons
}

func (h *HackathonPresenter) Update() {}
func (h *HackathonPresenter) Delete() {}
