package gateways

import (
	"context"
	"strconv"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type rbacPolicyGateway struct {
	db *gorm.DB
}

// NewRbacPolicyGateway はrbacPolicyGatewayのインスタンスを生成する
func NewRbacPolicyGateway(db *gorm.DB) dai.RBACPolicyDai {
	return &rbacPolicyGateway{
		db: db,
	}
}

// FindRoleByRole はroleに紐づくpolicyを取得する
func (r *rbacPolicyGateway) FindRoleByRole(ctx context.Context, role int) ([]*models.CasbinPolicy, error) {
	var policies []*models.RbacPolicy
	if err := r.db.Where("v0 = ?", role).Find(&policies).Error; err != nil {
		return nil, err
	}

	var casbinPolicies []*models.CasbinPolicy

	for _, p := range policies {
		casbinPolicies = append(casbinPolicies, &models.CasbinPolicy{
			PType: p.PType,
			V0:    strconv.Itoa(p.V0),
			V1:    p.V1,
			V2:    p.V2,
			V3:    p.V3,
		})
	}
	return casbinPolicies, nil
}

// FindRoleByPath はpathに紐づくpolicyを取得する
func (r *rbacPolicyGateway) FindRoleByPath(ctx context.Context, path string) ([]*models.CasbinPolicy, error) {
	var policies []*models.RbacPolicy
	if err := r.db.Where("v1 = ?", path).Find(&policies).Error; err != nil {
		return nil, err
	}

	var casbinPolicies []*models.CasbinPolicy

	for _, p := range policies {
		casbinPolicies = append(casbinPolicies, &models.CasbinPolicy{
			PType: p.PType,
			V0:    strconv.Itoa(p.V0),
			V1:    p.V1,
			V2:    p.V2,
			V3:    p.V3,
		})
	}
	return casbinPolicies, nil
}

// FindRoleByPathAndMethod はpathとmethodに紐づくpolicyを取得する
func (r *rbacPolicyGateway) FindRoleByPathAndMethod(ctx context.Context, path, method string) ([]*models.CasbinPolicy, error) {
	var policies []*models.RbacPolicy
	if err := r.db.Where("v1 = ? AND v2 = ?", path, method).Find(&policies).Error; err != nil {
		return nil, err
	}
	var casbinPolicies []*models.CasbinPolicy

	for _, p := range policies {
		casbinPolicies = append(casbinPolicies, &models.CasbinPolicy{
			PType: p.PType,
			V0:    strconv.Itoa(p.V0),
			V1:    p.V1,
			V2:    p.V2,
			V3:    p.V3,
		})
	}
	return casbinPolicies, nil
}

// FindRoleByPathAndMethodAndRole はpathとmethodとroleに紐づくpolicyを取得する
func (r *rbacPolicyGateway) Create(ctx context.Context, policy []*models.RbacPolicy) ([]int, error) {
	if err := r.db.Create(&policy).Error; err != nil {
		return nil, err
	}
	var ids []int
	for _, p := range policy {
		ids = append(ids, p.PolicyID)
	}
	return ids, nil
}

// FindRoleByPathAndMethodAndRole はpathとmethodとroleに紐づくpolicyを取得する
func (r *rbacPolicyGateway) FindAll(ctx context.Context, arg *request.ListRbacPolicies) ([]*models.RbacPolicy, error) {
	var policies []*models.RbacPolicy
	db := r.db
	if len(arg.Sub) > 0 {
		db = db.Where("v0 IN (?)", arg.Sub)
	}

	if len(arg.Obj) > 0 {
		db = db.Where("v1 IN (?)", arg.Obj)
	}

	if len(arg.Act) > 0 {
		db = db.Where("v2 IN (?)", arg.Act)
	}

	if len(arg.Eft) > 0 {
		db = db.Where("v3 IN (?)", arg.Eft)
	}

	if err := db.Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}

// FindRoleByPathAndMethodAndRole はpathとmethodとroleに紐づくpolicyを取得する
func (r *rbacPolicyGateway) DeleteByID(ctx context.Context, id int64) error {
	if err := r.db.Where("policy_id = ?", id).Delete(&models.RbacPolicy{}).Error; err != nil {
		return err
	}
	return nil
}

// FindRoleByPathAndMethodAndRole はpathとmethodとroleに紐づくpolicyを取得する
func (r *rbacPolicyGateway) DeleteAll(ctx context.Context) error {
	if err := r.db.Delete(&models.RbacPolicy{}).Error; err != nil {
		return err
	}
	return nil
}
