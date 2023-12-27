package gateways

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type HackathonGateway struct {
	db *gorm.DB
}

func NewHackathonGateway(db *gorm.DB) dai.HackathonDai {
	return &HackathonGateway{db: db}
}

func (h *HackathonGateway) Create(ctx context.Context, hackathon *models.Hackathon) error {
	result := h.db.Create(hackathon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (h *HackathonGateway) Find(ctx context.Context, hackathonID string) (*models.Hackathon, error) {
	var hackathon models.Hackathon
	result := h.db.First(&hackathon, "hackathon_id = ?", hackathonID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hackathon, nil
}

func (h *HackathonGateway) FindAll(ctx context.Context, size, id int) ([]*models.Hackathon, error) {
	var hackathons []*models.Hackathon
	result := h.db.Limit(int(size)).Offset(id).Find(&hackathons)
	if result.Error != nil {
		return nil, result.Error
	}
	return hackathons, nil
}
