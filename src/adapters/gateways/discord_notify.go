package gateways

import (
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/bwmarrin/discordgo"
)

const (
	NotifyTemplate = `
		【Title】		%s
		【応募リンク】	%s
		【応募締切】	%s
		【開催日時】	%s
		【タグ】
	`
)

type DiscordNotifyGateway struct {
	s *discordgo.Session
}

func NewDiscordNotifyGateway(s *discordgo.Session) dai.DiscordNotify {
	return &DiscordNotifyGateway{s: s}
}

func (d *DiscordNotifyGateway) CreateNewForum(channelID string, arg *models.Hackathon) (forumID string, err error) {
	forum, err := d.s.ForumThreadStartComplex(channelID, &discordgo.ThreadStart{
		Name:        arg.Name,
		AppliedTags: []string{},
	}, &discordgo.MessageSend{
		Content: arg.Name + arg.Link,
	})
	if err != nil {
		return "", err
	}

	return forum.ID, nil
}

func (d *DiscordNotifyGateway) DeleteChannel(channlID string) error {
	_, err := d.s.ChannelDelete(channlID)
	return err
}

func (d *DiscordNotifyGateway) AddAvailableTags(channelID string, tags []string) ([]string, error) {
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
