package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
)

type UserPresenter struct{}

func NewUserPresenter() ports.UserOutputBoundary {
	return &UserPresenter{}
}

func (up *UserPresenter) PresentInitAdmin(ctx context.Context, out *ports.OutputInitAdminData) (int, *response.User) {
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
