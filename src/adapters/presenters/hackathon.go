package presenters

import (
	"net/http"

	"github.com/hackhack-Geek-vol6/backend/src/datastructs/cerror"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/outputboundary"
)

type HackathonPresenter struct {
}

func NewHackathonOutputBoundary() outputboundary.HackathonOutputPort {
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

func (h *HackathonPresenter) Read()   {}
func (h *HackathonPresenter) Update() {}
func (h *HackathonPresenter) Delete() {}
