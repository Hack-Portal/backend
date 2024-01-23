package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type RoleGateway struct {
	db *gorm.DB
}

func NewRoleGateway(db *gorm.DB) dai.RoleDai {
	return &RoleGateway{
		db: db,
	}
}

func (rg *RoleGateway) Create(ctx context.Context, role *models.Role) (id int64, err error) {
	defer newrelic.FromContext(ctx).StartSegment("CreateRole-gateway").End()
	result := rg.db.Create(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}

func (rg *RoleGateway) FindAll(ctx context.Context) (roles []*models.Role, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllRole-gateway").End()
	result := rg.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (rg *RoleGateway) FindById(ctx context.Context, id int64) (role *models.Role, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindByIdRole-gateway").End()
	result := rg.db.First(&role, "role_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

func (rg *RoleGateway) Update(ctx context.Context, role *models.Role) (id int64, err error) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateRole-gateway").End()
	result := rg.db.Save(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}
