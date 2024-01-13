package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

// DiscordServerRegistry はDiscordServerRegistryに関するデータアクセスインターフェース
type DiscordServerRegistry interface {
	AddServer(arg *models.DiscordServerRegistry) error
	FindAllServers() ([]*models.DiscordServerRegistry, error)

	DeleteServer(arg *models.DiscordServerRegistry) error
}
