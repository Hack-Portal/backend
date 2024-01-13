package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type hackathonPresenter struct{}

// NewHackathonPresenter はHackathonPresenterを返す
func NewHackathonPresenter() ports.HackathonOutputBoundary {
	return &hackathonPresenter{}
}

// PresentCreateHackathon はHackathonの作成をpresenterする
func (s *hackathonPresenter) PresentCreateHackathon(ctx context.Context, out ports.OutputBoundary[*response.CreateHackathon]) (int, *response.CreateHackathon) {
	defer newrelic.FromContext(ctx).StartSegment("PresentCreateHackathon-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusCreated, out.Response()
}

// PresentGetHackathon はHackathonの取得をpresenterする
func (s *hackathonPresenter) PresentListHackathon(ctx context.Context, out ports.OutputBoundary[[]*response.GetHackathon]) (int, []*response.GetHackathon) {
	defer newrelic.FromContext(ctx).StartSegment("PresentListHackathon-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	respose := out.Response()
	if len(respose) == 0 {
		respose = []*response.GetHackathon{}
	}

	return http.StatusOK, respose
}

// PresentDeleteHackathon はHackathonの削除をpresenterする
func (s *hackathonPresenter) PresentDeleteHackathon(ctx context.Context, out ports.OutputBoundary[*response.DeleteHackathon]) (int, *response.DeleteHackathon) {
	defer newrelic.FromContext(ctx).StartSegment("PresentDeleteHackathon-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, out.Response()
}
