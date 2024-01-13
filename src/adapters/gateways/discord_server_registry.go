package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type DiscordServerRegistryGateways struct {
	db *gorm.DB
}

func NewDiscordServerRegistryGateways(db *gorm.DB) dai.DiscordServerRegistry {
	return &DiscordServerRegistryGateways{db: db}
}

func (d *DiscordServerRegistryGateways) AddServer(arg *models.DiscordServerRegistry) error {
	return d.db.Create(arg).Error
}

func (d *DiscordServerRegistryGateways) FindAllServers() ([]*models.DiscordServerRegistry, error) {
	var servers []*models.DiscordServerRegistry
	err := d.db.Find(&servers).Error
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (d *DiscordServerRegistryGateways) DeleteServer(arg *models.DiscordServerRegistry) error {
	return d.db.Delete(arg).Error
}
