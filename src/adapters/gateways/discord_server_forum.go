package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type discordServerForumTagGateway struct {
	db *gorm.DB
}

// NewDiscordServerForumTagGateway はdiscordServerForumTagGatewayのインスタンスを生成する
func NewDiscordServerForumTagGateway(db *gorm.DB) dai.DiscordServerForumTagDai {
	return &discordServerForumTagGateway{db: db}
}

// CreateNewForumTag はdiscordサーバーに紐づくforum tagを作成する
func (d *discordServerForumTagGateway) CreateNewForumTag(arg []*models.DiscordServerForumTag) error {
	return d.db.Create(arg).Error
}

// FindByServerID はdiscordサーバーに紐づくforum tagを取得する
func (d *discordServerForumTagGateway) FindByStatusID(statusID []int) ([]*models.DiscordServerForumTag, error) {
	var tags []*models.DiscordServerForumTag
	err := d.db.Where("status_id IN (?)", statusID).Find(&tags).Error
	return tags, err
}
