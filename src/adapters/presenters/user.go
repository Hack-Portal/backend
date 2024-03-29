package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type UserPresenter struct{}

func NewUserPresenter() ports.UserOutputBoundary {
	return &UserPresenter{}
}

func (up *UserPresenter) PresentInitAdmin(ctx context.Context, out *ports.OutputInitAdminData) (int, *response.User) {
	defer newrelic.FromContext(ctx).StartSegment("PresentInitAdmin-presenter").End()
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

func (up *UserPresenter) PresentLogin(ctx context.Context, out ports.OutputBoundary[*response.Login]) (int, *response.Login) {
	defer newrelic.FromContext(ctx).StartSegment("PresentLogin-presenter").End()

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
