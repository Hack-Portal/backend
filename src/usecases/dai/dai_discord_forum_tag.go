package dai

import "github.com/Hack-Portal/backend/src/datastructure/models"

// DiscordServerForumTagDai はDiscordServerForumTagに関するデータアクセスインターフェース
type DiscordServerForumTagDai interface {
	CreateNewForumTag(arg []*models.DiscordServerForumTag) error
	FindByStatusID(statusID []int) ([]*models.DiscordServerForumTag, error)
}
