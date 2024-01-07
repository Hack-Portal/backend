package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
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
	result := rg.db.Create(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}

func (rg *RoleGateway) FindAll(ctx context.Context) (roles []*models.Role, err error) {
	result := rg.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (rg *RoleGateway) FindById(ctx context.Context, id int64) (role *models.Role, err error) {
	result := rg.db.First(role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

func (rg *RoleGateway) Update(ctx context.Context, role *models.Role) (id int64, err error) {
	result := rg.db.Save(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}
