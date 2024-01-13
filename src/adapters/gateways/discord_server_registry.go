package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type discordServerRegistryGateways struct {
	db *gorm.DB
}

// NewDiscordServerRegistryGateways はdiscordServerRegistryGatewaysのインスタンスを生成する
func NewDiscordServerRegistryGateways(db *gorm.DB) dai.DiscordServerRegistry {
	return &discordServerRegistryGateways{db: db}
}

// AddServer はdiscordサーバーを登録する
func (d *discordServerRegistryGateways) AddServer(arg *models.DiscordServerRegistry) error {
	return d.db.Create(arg).Error
}

// FindServer はdiscordサーバーを取得する
func (d *discordServerRegistryGateways) FindAllServers() ([]*models.DiscordServerRegistry, error) {
	var servers []*models.DiscordServerRegistry
	err := d.db.Find(&servers).Error
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// FindServer はdiscordサーバーを取得する
func (d *discordServerRegistryGateways) DeleteServer(arg *models.DiscordServerRegistry) error {
	return d.db.Delete(arg).Error
}
