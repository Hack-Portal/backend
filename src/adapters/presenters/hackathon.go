package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type HackathonPresenter struct {
}

func NewHackathonPresenter() ports.HackathonOutputBoundary {
	return &HackathonPresenter{}
}

func (s *HackathonPresenter) PresentCreateHackathon(ctx context.Context, out *ports.OutputCreateHackathonData) (int, *response.CreateHackathon) {
	defer newrelic.FromContext(ctx).StartSegment("PresentCreateHackathon-presenter").End()

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
	defer newrelic.FromContext(ctx).StartSegment("PresentGetHackathon-presenter").End()

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
	defer newrelic.FromContext(ctx).StartSegment("PresentListHackathon-presenter").End()

	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	if len(out.Response) == 0 {
		out.Response = []*response.GetHackathon{}
	}

	return http.StatusOK, out.Response
}

func (s *HackathonPresenter) PresentDeleteHackathon(ctx context.Context, out *ports.OutputDeleteHackathonData) (int, *response.DeleteHackathon) {
	defer newrelic.FromContext(ctx).StartSegment("PresentDeleteHackathon-presenter").End()

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
