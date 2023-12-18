package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type StatusTagGateway struct {
	db *gorm.DB
}

func NewStatusTagGateway(db *gorm.DB) dai.StatusTagDai {
	return &StatusTagGateway{
		db: db,
	}
}

func (stg *StatusTagGateway) Create(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	result := stg.db.Select("status").Create(&statusTag)
	if result.Error != nil {
		return 0, result.Error
	}

	// get last insert id from psql
	var statusTagID int64
	err = stg.db.Raw("SELECT currval(pg_get_serial_sequence('status_tags', 'status_id'))").Scan(&statusTagID).Error
	if err != nil {
		return 0, err
	}

	return statusTagID, nil
}

func (stg *StatusTagGateway) FindAll(ctx context.Context) (statusTags []*models.StatusTag, err error) {
	result := stg.db.Find(&statusTags)
	if result.Error != nil {
		return nil, result.Error
	}

	return statusTags, nil
}

func (stg *StatusTagGateway) FindById(ctx context.Context, id int64) (statusTag *models.StatusTag, err error) {
	result := stg.db.First(&statusTag, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return statusTag, nil
}

func (stg *StatusTagGateway) Update(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	result := stg.db.Model(statusTag).Where("status_id = ?", statusTag.StatusID).Updates(statusTag)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return statusTag.StatusID, nil
}
