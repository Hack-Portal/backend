package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type UserGateway struct {
	db *gorm.DB
}

func NewUserGateway(db *gorm.DB) dai.UsersDai {
	return &UserGateway{
		db: db,
	}
}

func (ug *UserGateway) Create(ctx context.Context, user *models.User) (id string, err error) {
	defer newrelic.FromContext(ctx).StartSegment("CreateUser-gateway").End()
	result := ug.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.UserID, nil
}

func (ug *UserGateway) FindAll(ctx context.Context) (users []*models.User, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllUser-gateway").End()
	result := ug.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ug *UserGateway) FindById(ctx context.Context, id string) (user *models.User, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindByIdUser-gateway").End()

	result := ug.db.First(&user, "user_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ug *UserGateway) Update(ctx context.Context, user *models.User) (id string, err error) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateUser-gateway").End()

	result := ug.db.Save(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.UserID, nil
}

func (ug *UserGateway) Delete(ctx context.Context, id string) (err error) {
	defer newrelic.FromContext(ctx).StartSegment("DeleteUser-gateway").End()

	result := ug.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
