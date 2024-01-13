package gateways

import (
	"context"
	"fmt"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type hackathonStatusGateway struct {
	db          *gorm.DB
	cacheClient dai.Cache[[]*models.JoinedStatusTag]
}

// NewHackathonStatusGateway はhackathonStatusGatewayのインスタンスを生成する
func NewHackathonStatusGateway(db *gorm.DB, cache *redis.Client) dai.HackathonStatusDai {
	return &hackathonStatusGateway{
		db:          db,
		cacheClient: NewCache[[]*models.JoinedStatusTag](cache, time.Duration(5)*time.Minute),
	}
}

// FindAll は全てのhackathonに紐づくstatusを取得する
func (hs *hackathonStatusGateway) FindAll(ctx context.Context, HackathonID []string) ([]*models.JoinedStatusTag, error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllHackathonStatus-gateway").End()
	key := fmt.Sprintf("hackathon_status_%v", HackathonID)
	hackathonStatusTags, err := hs.cacheClient.Get(ctx, key, func(ctx context.Context) ([]*models.JoinedStatusTag, error) {
		var hackathonStatusTags []*models.JoinedStatusTag
		result := hs.db.Model(&models.HackathonStatusTag{}).
			Joins("JOIN status_tags ON status_tags.status_id = hackathon_status_tags.status_id").
			Where("hackathon_status_tags.hackathon_id IN (?)", HackathonID).
			Select("hackathon_id", "status_tags.status_id as status_id", "status").
			Find(&hackathonStatusTags)
		if result.Error != nil {
			return nil, result.Error
		}

		return hackathonStatusTags, nil
	})
	return hackathonStatusTags, err
}

// Delete はhackathonに紐づくstatusを削除する
func (hs *hackathonStatusGateway) Delete(ctx context.Context, HackathonID string) error {
	defer newrelic.FromContext(ctx).StartSegment("DeleteHackathonStatus-gateway").End()

	result := hs.db.Delete(&models.HackathonStatusTag{}).
		Where("hackathon_id = ?", HackathonID)
	if result.Error != nil {
		return result.Error
	}

	return hs.cacheClient.Reset(ctx, "hackathon_status")
}
