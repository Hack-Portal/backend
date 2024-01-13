package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type DiscordChannelGateway struct {
	db *gorm.DB
}

func NewDiscordChannelGateway(db *gorm.DB) dai.DiscordChannelDai {
	return &DiscordChannelGateway{db: db}
}

func (d *DiscordChannelGateway) AddChannel(arg []*models.HackathonDiscordChannel) error {
	return d.db.Create(arg).Error
}

func (d *DiscordChannelGateway) GetChannelIDs(hackathonID string) ([]*models.HackathonDiscordChannel, error) {
	var channels []*models.HackathonDiscordChannel
	err := d.db.Where("hackathon_id = ?", hackathonID).Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (d *DiscordChannelGateway) RemoveChannel(hackathonID string) error {
	return d.db.Where("hackathon_id = ?", hackathonID).Delete(&models.HackathonDiscordChannel{}).Error
}
