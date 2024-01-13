package gateways

import (
	"fmt"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/bwmarrin/discordgo"
)

const (
	// DiscordNotifyTemplate はdiscordに通知する際のテンプレート
	DiscordNotifyTemplate = "【Title】%s\n【応募締切】%s\n【開催日時】%s\n【応募リンク】%s\n"
)

type discordNotifyGateway struct {
	s *discordgo.Session
}

// NewDiscordNotifyGateway はdiscordNotifyGatewayのインスタンスを生成する
func NewDiscordNotifyGateway(s *discordgo.Session) dai.DiscordNotify {
	return &discordNotifyGateway{s: s}
}

// CreateNewForum はdiscordサーバーに紐づくforumを作成する
func (d *discordNotifyGateway) CreateNewForum(channelID string, arg *models.Hackathon) (forumID string, err error) {
	forum, err := d.s.ForumThreadStartComplex(channelID, &discordgo.ThreadStart{
		Name:        arg.Name,
		AppliedTags: []string{},
	}, &discordgo.MessageSend{
		Content: fmt.Sprintf(DiscordNotifyTemplate, arg.Name, arg.Expired.Format("2006-01-02"), arg.StartDate.Format("2006-01-02"), arg.Link),
	})
	if err != nil {
		return "", err
	}

	return forum.ID, nil
}

// DeleteForum はdiscordサーバーに紐づくforumを削除する
func (d *discordNotifyGateway) DeleteChannel(channlID string) error {
	_, err := d.s.ChannelDelete(channlID)
	return err
}

// AddAvailableTags はdiscordサーバーに紐づくforum tagを作成する
func (d *discordNotifyGateway) AddAvailableTags(channelID string, tags []string) ([]string, error) {
	var arg []discordgo.ForumTag
	for _, tag := range tags {
		arg = append(arg, discordgo.ForumTag{
			Name: tag,
		})
	}

	channel, err := d.s.ChannelEditComplex(channelID, &discordgo.ChannelEdit{
		AvailableTags: &arg,
	})

	var tagIds []string
	for _, tag := range channel.AvailableTags {
		tagIds = append(tagIds, tag.ID)
	}

	channel, err = d.s.ChannelEditComplex(channelID, &discordgo.ChannelEdit{
		AppliedTags: &tagIds,
	})

	return tagIds, err
}
