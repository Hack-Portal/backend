package response

import "github.com/Hack-Portal/backend/src/datastructure/models"

type CreateRbacPolicy struct {
	Id []int `json:"id"`
}

type ListRbacPolicies struct {
	Policies []*models.RbacPolicy `json:"policies"`
}

type DeleteRbacPolicy struct {
	PolicyID int64 `json:"policy_id"`
}

type DeleteAllRbacPolicies struct {
	Message string `json:"message"`
}
