package gateways

import (
	"context"
	"fmt"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type HackathonGateway struct {
	db *gorm.DB
}

func NewHackathonGateway(db *gorm.DB) dai.HackathonDai {
	return &HackathonGateway{
		db: db,
	}
}

func (h *HackathonGateway) Create(ctx context.Context, hackathon *models.Hackathon, hackathonStatus []int64) error {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathon-gateway").End()

	return h.db.Transaction(func(tx *gorm.DB) error {
		result := h.db.Create(hackathon)
		if result.Error != nil {
			return result.Error
		}

		if len(hackathonStatus) > 0 {
			if err := h.createStatusTag(ctx, hackathon.HackathonID, hackathonStatus); err != nil {
				return err
			}
		}

		return nil
	})
}

func (h *HackathonGateway) createStatusTag(ctx context.Context, HackathonID string, hackathonStatus []int64) error {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathonStatus-gateway").End()

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

func (h *HackathonGateway) Find(ctx context.Context, hackathonID string) (*models.Hackathon, error) {
	defer newrelic.FromContext(ctx).StartSegment("FindHackathon-gateway").End()

	var hackathon models.Hackathon
	result := h.db.First(&hackathon, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hackathon, nil
}

func (h *HackathonGateway) FindAll(ctx context.Context, arg dai.FindAllParams) ([]*models.Hackathon, error) {
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

	var hackathons []*models.Hackathon
	result := chain.Select("DISTINCT (hackathons.hackathon_id)", "name", "icon", "link", "expired", "start_date", "term", "created_at").Where("expired > ?", time.Now()).Find(&hackathons)
	if result.Error != nil {
		return nil, result.Error
	}
	return hackathons, nil
}

func (h *HackathonGateway) Delete(ctx context.Context, hackathonID string) error {
	defer newrelic.FromContext(ctx).StartSegment("DeleteHackathon-gateway").End()

	result := h.db.Delete(&models.Hackathon{}, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
