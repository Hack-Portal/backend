package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

type DiscordNotify interface {
	CreateNewForum(channelID string, arg *models.Hackathon) (forumID string, err error)

	DeleteChannel(channlID string) error

	AddAvailableTags(channelID string, tags []string) ([]string, error)
}