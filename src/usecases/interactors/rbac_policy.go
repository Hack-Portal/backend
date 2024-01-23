package interactors

import (
	"context"
	"strconv"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type RbacPolicyInteractor struct {
	policyreRepo dai.RBACPolicyDai
	roleRepo     dai.RoleDai
	output       ports.RbacPolicyOutputBoundary
}

func NewRbacPolicyInteractor(policyRepo dai.RBACPolicyDai, roleRepo dai.RoleDai, output ports.RbacPolicyOutputBoundary) ports.RbacPolicyInputBoundary {
	return &RbacPolicyInteractor{
		policyreRepo: policyRepo,
		roleRepo:     roleRepo,
		output:       output,
	}
}

func (r *RbacPolicyInteractor) CreateRbacPolicy(ctx context.Context, in *request.CreateRbacPolicy) (int, *response.CreateRbacPolicy) {
	defer newrelic.FromContext(ctx).StartSegment("CreateRbacPolicy-interactor").End()
	if len(in.Policies) == 0 {
		return r.output.PresentCreateRbacPolicy(ctx, ports.NewOutput[*response.CreateRbacPolicy](
			hperror.ErrFieldRequired,
			nil,
		))
	}

	// []CasbinPolicy -> []RBACPolicy
	var policies []*models.RbacPolicy
	for _, policy := range in.Policies {
		v0, err := strconv.Atoi(policy.V0)
		if err != nil {
			return r.output.PresentCreateRbacPolicy(ctx, ports.NewOutput[*response.CreateRbacPolicy](
				err,
				nil,
			))
		}

		policies = append(policies, &models.RbacPolicy{
			PType: policy.PType,
			V0:    v0,
			V1:    policy.V1,
			V2:    policy.V2,
			V3:    policy.V3,
		})
	}

	result, err := r.policyreRepo.Create(ctx, policies)
	if err != nil {
		return r.output.PresentCreateRbacPolicy(ctx, ports.NewOutput[*response.CreateRbacPolicy](
			err,
			nil,
		))
	}

	return r.output.PresentCreateRbacPolicy(ctx, ports.NewOutput[*response.CreateRbacPolicy](
		nil,
		&response.CreateRbacPolicy{
			Id: result,
		},
	))
}

func (r *RbacPolicyInteractor) ListRbacPolicies(ctx context.Context, in *request.ListRbacPolicies) (int, *response.ListRbacPolicies) {
	defer newrelic.FromContext(ctx).StartSegment("ListRbacPolicies-interactor").End()
	result, err := r.policyreRepo.FindAll(ctx, in)
	if err != nil {
		return r.output.PresentListRbacPolicies(ctx, ports.NewOutput[*response.ListRbacPolicies](
			err,
			nil,
		))
	}

	return r.output.PresentListRbacPolicies(ctx, ports.NewOutput[*response.ListRbacPolicies](
		nil,
		&response.ListRbacPolicies{
			Policies: result,
		},
	))
}

func (r *RbacPolicyInteractor) DeleteRbacPolicy(ctx context.Context, in *request.DeleteRbacPolicy) (int, *response.DeleteRbacPolicy) {
	defer newrelic.FromContext(ctx).StartSegment("DeleteRbacPolicy-interactor").End()

	if err := r.policyreRepo.DeleteByID(ctx, in.PolicyID); err != nil {
		return r.output.PresentDeleteRbacPolicy(ctx, ports.NewOutput[*response.DeleteRbacPolicy](
			err,
			nil,
		))
	}

	return r.output.PresentDeleteRbacPolicy(ctx, ports.NewOutput[*response.DeleteRbacPolicy](
		nil,
		&response.DeleteRbacPolicy{
			PolicyID: in.PolicyID,
		},
	))
}

func (r *RbacPolicyInteractor) DeleteAllRbacPolicies(ctx context.Context) (int, *response.DeleteAllRbacPolicies) {
	defer newrelic.FromContext(ctx).StartSegment("DeleteAllRbacPolicies-interactor").End()

	if err := r.policyreRepo.DeleteAll(ctx); err != nil {
		return r.output.PresentDeleteAllRbacPolicies(ctx, ports.NewOutput[*response.DeleteAllRbacPolicies](
			err,
			nil,
		))
	}

	return r.output.PresentDeleteAllRbacPolicies(ctx, ports.NewOutput[*response.DeleteAllRbacPolicies](
		nil,
		&response.DeleteAllRbacPolicies{
			Message: "All policies deleted",
		},
	))
}
