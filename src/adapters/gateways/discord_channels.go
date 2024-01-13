package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type discordChannelGateway struct {
	db *gorm.DB
}

// NewDiscordChannelGateway はdiscordChannelGatewayのインスタンスを生成する
func NewDiscordChannelGateway(db *gorm.DB) dai.DiscordChannelDai {
	return &discordChannelGateway{db: db}
}

// AddChannel はdiscordサーバーに紐づくchannelを作成する
func (d *discordChannelGateway) AddChannel(arg []*models.HackathonDiscordChannel) error {
	return d.db.Create(arg).Error
}

// GetChannelIDs はdiscordサーバーに紐づくchannelを取得する
func (d *discordChannelGateway) GetChannelIDs(hackathonID string) ([]*models.HackathonDiscordChannel, error) {
	var channels []*models.HackathonDiscordChannel
	err := d.db.Where("hackathon_id = ?", hackathonID).Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

// RemoveChannel はdiscordサーバーに紐づくchannelを削除する
func (d *discordChannelGateway) RemoveChannel(hackathonID string) error {
	return d.db.Where("hackathon_id = ?", hackathonID).Delete(&models.HackathonDiscordChannel{}).Error
}
