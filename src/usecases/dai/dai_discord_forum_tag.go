package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

type DiscordServerForumTagDai interface {
	CreateNewForumTag(arg []*models.DiscordServerForumTag) error
}
