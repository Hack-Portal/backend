package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type HackathonStatusGateway struct {
	db *gorm.DB
}

func NewHackathonStatusGateway(db *gorm.DB) dai.HackathonStatusDai {
	return &HackathonStatusGateway{
		db: db,
	}
}

func (hs *HackathonStatusGateway) FindAll(ctx context.Context, HackathonID []string) ([]*models.JoinedStatusTag, error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllHackathonStatus-gateway").End()
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
}

func (hs *HackathonStatusGateway) Delete(ctx context.Context, HackathonID string) error {
	defer newrelic.FromContext(ctx).StartSegment("DeleteHackathonStatus-gateway").End()

	result := hs.db.Delete(&models.HackathonStatusTag{}, "hackathon_id = ?", HackathonID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
