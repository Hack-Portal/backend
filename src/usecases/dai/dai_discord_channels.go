package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

// DiscordChannelDai はDiscordChannelに関するデータアクセスインターフェース
type DiscordChannelDai interface {
	AddChannel(arg []*models.HackathonDiscordChannel) error
	GetChannelIDs(hackathonID string) ([]*models.HackathonDiscordChannel, error)
	RemoveChannel(hackathonID string) error
}
