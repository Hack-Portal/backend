package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type roleGateway struct {
	db *gorm.DB
}

// NewRoleGateway はroleGatewayのインスタンスを生成する
func NewRoleGateway(db *gorm.DB) dai.RoleDai {
	return &roleGateway{
		db: db,
	}
}

// Create はRoleを作成する
func (rg *roleGateway) Create(ctx context.Context, role *models.Role) (id int64, err error) {
	result := rg.db.Create(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}

// FindAll は全てのRoleを取得する
func (rg *roleGateway) FindAll(ctx context.Context) (roles []*models.Role, err error) {
	result := rg.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

// FindById は指定したIDのRoleを取得する
func (rg *roleGateway) FindByID(ctx context.Context, id int64) (role *models.Role, err error) {
	result := rg.db.First(&role, "role_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

// Update はRoleを更新する
func (rg *roleGateway) Update(ctx context.Context, role *models.Role) (id int64, err error) {
	result := rg.db.Save(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.RoleID, nil
}
