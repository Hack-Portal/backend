package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

type DiscordChannelDai interface {
	AddChannel(arg []*models.HackathonDiscordChannel) error
	GetChannelIDs(hackathonID string) ([]*models.HackathonDiscordChannel, error)
	RemoveChannel(hackathonID string) error
}
