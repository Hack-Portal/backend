package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
)

type userPresenter struct{}

// NewUserPresenter はUserPresenterを返す
func NewUserPresenter() ports.UserOutputBoundary {
	return &userPresenter{}
}

// PresentInitAdmin はUserの作成をpresenterする
func (up *userPresenter) PresentInitAdmin(ctx context.Context, out ports.OutputBoundary[*response.User]) (int, *response.User) {
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

// PresentLogin はUserのログインをpresenterする
func (up *userPresenter) PresentLogin(ctx context.Context, out ports.OutputBoundary[*response.Login]) (int, *response.Login) {
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
