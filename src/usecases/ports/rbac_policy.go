package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

// RbacPolicyInputBoundary はRbacPolicyのInputBoundary
type RbacPolicyInputBoundary interface {
	CreateRbacPolicy(ctx context.Context, in *request.CreateRbacPolicy) (int, *response.CreateRbacPolicy)
	ListRbacPolicies(ctx context.Context, in *request.ListRbacPolicies) (int, *response.ListRbacPolicies)
	DeleteRbacPolicy(ctx context.Context, in *request.DeleteRbacPolicy) (int, *response.DeleteRbacPolicy)
	DeleteAllRbacPolicies(ctx context.Context) (int, *response.DeleteAllRbacPolicies)
}

// RbacPolicyOutputBoundary はRbacPolicyのOutputBoundary
type RbacPolicyOutputBoundary interface {
	PresentCreateRbacPolicy(ctx context.Context, out OutputBoundary[*response.CreateRbacPolicy]) (int, *response.CreateRbacPolicy)
	PresentListRbacPolicies(ctx context.Context, out OutputBoundary[*response.ListRbacPolicies]) (int, *response.ListRbacPolicies)
	PresentDeleteRbacPolicy(ctx context.Context, out OutputBoundary[*response.DeleteRbacPolicy]) (int, *response.DeleteRbacPolicy)
	PresentDeleteAllRbacPolicies(ctx context.Context, out OutputBoundary[*response.DeleteAllRbacPolicies]) (int, *response.DeleteAllRbacPolicies)
}
