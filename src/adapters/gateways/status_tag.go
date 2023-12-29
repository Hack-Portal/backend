package gateways

import (
	"context"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type StatusTagGateway struct {
	db          *gorm.DB
	cacheClient dai.Cache[[]*models.StatusTag]
}

func NewStatusTagGateway(db *gorm.DB, cache *redis.Client) dai.StatusTagDai {
	return &StatusTagGateway{
		db:          db,
		cacheClient: NewCache[[]*models.StatusTag](cache, time.Duration(30)*time.Minute),
	}
}

func (stg *StatusTagGateway) Create(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	defer newrelic.FromContext(ctx).StartSegment("CreateStatusTag-gateway").End()

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

	return statusTagID, stg.cacheClient.Reset(ctx, "status_tags")
}

func (stg *StatusTagGateway) FindAll(ctx context.Context) (statusTags []*models.StatusTag, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllStatusTag-gateway").End()

	tags, err := stg.cacheClient.Get(ctx, "status_tags", func(ctx context.Context) ([]*models.StatusTag, error) {
		result := stg.db.Find(&statusTags)
		if result.Error != nil {
			return nil, result.Error
		}
		return statusTags, nil
	})

	return tags, nil
}

func (stg *StatusTagGateway) FindById(ctx context.Context, id int64) (statusTag *models.StatusTag, err error) {
	defer newrelic.FromContext(ctx).StartSegment("FindByIdStatusTag-gateway").End()

	result := stg.db.First(&statusTag, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return statusTag, nil
}

func (stg *StatusTagGateway) Update(ctx context.Context, statusTag *models.StatusTag) (id int64, err error) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateStatusTag-gateway").End()

	result := stg.db.Model(statusTag).Where("status_id = ?", statusTag.StatusID).Updates(statusTag)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return statusTag.StatusID, stg.cacheClient.Reset(ctx, "status_tags")
}
