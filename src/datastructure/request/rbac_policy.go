package request

import "github.com/Hack-Portal/backend/src/datastructure/models"

type CreateRbacPolicy struct {
	Policies []models.CasbinPolicy `json:"policies"`
}

type ListRbacPolicies struct {
	Sub []string `query:"sub"`
	Obj []string `query:"obj"`
	Act []string `query:"act"`
	Eft []string `query:"eft"`
}

type DeleteRbacPolicy struct {
	PolicyID int64 `param:"policy_id" validate:"required"`
}
