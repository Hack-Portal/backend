package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type userGateway struct {
	db *gorm.DB
}

// NewUserGateway はuserGatewayのインスタンスを生成する
func NewUserGateway(db *gorm.DB) dai.UsersDai {
	return &userGateway{
		db: db,
	}
}

// Create はUserを作成する
func (ug *userGateway) Create(ctx context.Context, user *models.User) (id string, err error) {
	result := ug.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.UserID, nil
}

// FindAll は全てのUserを取得する
func (ug *userGateway) FindAll(ctx context.Context) (users []*models.User, err error) {
	result := ug.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindByID は指定したIDのUserを取得する
func (ug *userGateway) FindByID(ctx context.Context, id string) (user *models.User, err error) {
	result := ug.db.First(&user, "user_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// Update はUserを更新する
func (ug *userGateway) Update(ctx context.Context, user *models.User) (id string, err error) {
	result := ug.db.Save(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.UserID, nil
}

// Delete はUserを削除する
func (ug *userGateway) Delete(ctx context.Context, id string) (err error) {
	result := ug.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
