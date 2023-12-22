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
