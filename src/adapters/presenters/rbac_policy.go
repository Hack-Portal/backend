package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type RbacPolicyPresenter struct{}

func NewRbacPolicyPresenter() ports.RbacPolicyOutputBoundary {
	return &RbacPolicyPresenter{}
}

func (r *RbacPolicyPresenter) PresentCreateRbacPolicy(ctx context.Context, out ports.OutputBoundary[*response.CreateRbacPolicy]) (int, *response.CreateRbacPolicy) {
	defer newrelic.FromContext(ctx).StartSegment("PresentCreateRbacPolicy").End()
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

func (r *RbacPolicyPresenter) PresentListRbacPolicies(ctx context.Context, out ports.OutputBoundary[*response.ListRbacPolicies]) (int, *response.ListRbacPolicies) {
	defer newrelic.FromContext(ctx).StartSegment("PresentListRbacPolicies").End()
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

func (r *RbacPolicyPresenter) PresentDeleteRbacPolicy(ctx context.Context, out ports.OutputBoundary[*response.DeleteRbacPolicy]) (int, *response.DeleteRbacPolicy) {
	defer newrelic.FromContext(ctx).StartSegment("PresentDeleteRbacPolicy").End()
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

func (r *RbacPolicyPresenter) PresentDeleteAllRbacPolicies(ctx context.Context, out ports.OutputBoundary[*response.DeleteAllRbacPolicies]) (int, *response.DeleteAllRbacPolicies) {
	defer newrelic.FromContext(ctx).StartSegment("PresentDeleteAllRbacPolicies").End()
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
