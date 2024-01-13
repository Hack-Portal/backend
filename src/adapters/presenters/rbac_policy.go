package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
)

type rbacPolicyPresenter struct{}

// NewRbacPolicyPresenter はRbacPolicyPresenterを返す
func NewRbacPolicyPresenter() ports.RbacPolicyOutputBoundary {
	return &rbacPolicyPresenter{}
}

// PresentCreateRbacPolicy はRbacPolicyの作成をpresenterする
func (r *rbacPolicyPresenter) PresentCreateRbacPolicy(ctx context.Context, out ports.OutputBoundary[*response.CreateRbacPolicy]) (int, *response.CreateRbacPolicy) {
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

// PresentListRbacPolicies はRbacPolicyの取得をpresenterする
func (r *rbacPolicyPresenter) PresentListRbacPolicies(ctx context.Context, out ports.OutputBoundary[*response.ListRbacPolicies]) (int, *response.ListRbacPolicies) {
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

// PresentDeleteRbacPolicy はRbacPolicyの削除をpresenterする
func (r *rbacPolicyPresenter) PresentDeleteRbacPolicy(ctx context.Context, out ports.OutputBoundary[*response.DeleteRbacPolicy]) (int, *response.DeleteRbacPolicy) {
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

// PresentDeleteAllRbacPolicies はRbacPolicyの全削除をpresenterする
func (r *rbacPolicyPresenter) PresentDeleteAllRbacPolicies(ctx context.Context, out ports.OutputBoundary[*response.DeleteAllRbacPolicies]) (int, *response.DeleteAllRbacPolicies) {
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
