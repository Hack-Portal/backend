package presenters

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
)

type HackathonPresenter struct {
	logger *slog.Logger
}

func NewHackathonPresenter() ports.HackathonOutputBoundary {
	return &HackathonPresenter{}
}

func (s *HackathonPresenter) PresentCreateHackathon(ctx context.Context, out *ports.OutputCreateHackathonData) (int, *response.CreateHackathon) {
	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusCreated, out.Response
}

func (s *HackathonPresenter) PresentGetHackathon(ctx context.Context, out *ports.OutputGetHackathonData) (int, *response.GetHackathon) {
	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, out.Response
}

func (s *HackathonPresenter) PresentListHackathon(ctx context.Context, out *ports.OutputListHackathonData) (int, []*response.GetHackathon) {
	log.Println("out", out)
	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, out.Response
}
