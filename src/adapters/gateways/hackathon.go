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

type hackathonGateway struct {
	db          *gorm.DB
	cacheClient dai.Cache[[]*models.Hackathon]
}

// NewHackathonGateway はhackathonGatewayのインスタンスを生成する
func NewHackathonGateway(db *gorm.DB, cache *redis.Client) dai.HackathonDai {
	return &hackathonGateway{
		db:          db,
		cacheClient: NewCache[[]*models.Hackathon](cache, time.Duration(5)*time.Minute),
	}
}

// Create はhackathonを作成する
func (h *hackathonGateway) Create(ctx context.Context, hackathon *models.Hackathon, hackathonStatus []int64) error {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathon-gateway").End()

	return h.db.Transaction(func(tx *gorm.DB) error {
		result := h.db.Create(hackathon)
		if result.Error != nil {
			return result.Error
		}

		if err := h.createStatusTag(ctx, hackathon.HackathonID, hackathonStatus); err != nil {
			return err
		}

		return h.cacheClient.Reset(ctx, "hackathons")
	})
}

// createStatusTag はhackathonに紐づくstatusを作成する
func (h *hackathonGateway) createStatusTag(ctx context.Context, HackathonID string, hackathonStatus []int64) error {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathonStatus-gateway").End()
	if len(hackathonStatus) == 0 {
		return nil
	}

	var hackathonStatusTags []*models.HackathonStatusTag
	for _, status := range hackathonStatus {
		hackathonStatusTags = append(hackathonStatusTags, &models.HackathonStatusTag{
			HackathonID: HackathonID,
			StatusID:    status,
		})
	}

	result := h.db.Model(&models.HackathonStatusTag{}).Create(hackathonStatusTags)

	return result.Error
}

// Find はhackathonを取得する
func (h *hackathonGateway) Find(ctx context.Context, hackathonID string) (*models.Hackathon, error) {
	defer newrelic.FromContext(ctx).StartSegment("FindHackathon-gateway").End()

	var hackathon models.Hackathon
	result := h.db.First(&hackathon, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hackathon, nil
}

// FindAll は全てのhackathonを取得する
func (h *hackathonGateway) FindAll(ctx context.Context, arg dai.FindAllParams) ([]*models.Hackathon, error) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllHackathon-gateway").End()

	var key string = "hackathons"
	chain := h.db.Limit(arg.Limit).Offset(arg.Offset)

	if len(arg.Tags) > 0 {
		chain.Joins("JOIN hackathon_status_tags ON hackathons.hackathon_id = hackathon_status_tags.hackathon_id").
			Where("hackathon_status_tags.status_id IN ?", arg.Tags)
		key = fmt.Sprintf("%s-%v", key, arg.Tags)
	}

	if arg.New {
		chain.Order("created_at DESC")
		key = fmt.Sprintf("%s-new", key)
	}

	if arg.LongTerm {
		chain.Order("term DESC")
		key = fmt.Sprintf("%s-longterm", key)
	}

	if arg.NearDeadline {
		chain.Order("expired ASC")
		key = fmt.Sprintf("%s-neardeadline", key)
	}

	hackathons, err := h.cacheClient.Get(ctx, key, func(ctx context.Context) ([]*models.Hackathon, error) {
		var hackathons []*models.Hackathon
		result := chain.Select("DISTINCT (hackathons.hackathon_id)", "name", "icon", "link", "expired", "start_date", "term", "created_at").Where("expired > ?", time.Now()).Find(&hackathons)
		if result.Error != nil {
			return nil, result.Error
		}
		return hackathons, nil
	})

	return hackathons, err
}

// Delete はhackathonを更新する
func (h *hackathonGateway) Delete(ctx context.Context, hackathonID string) error {
	defer newrelic.FromContext(ctx).StartSegment("DeleteHackathon-gateway").End()

	result := h.db.Delete(&models.Hackathon{}, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return result.Error
	}

	result = h.db.Delete(&models.HackathonStatusTag{}, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return result.Error
	}

	return h.cacheClient.Reset(ctx, "hackathons")
}
