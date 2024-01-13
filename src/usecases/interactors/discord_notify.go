package interactors

import (
	"sync"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
)

type discordNotify struct {
	discordChannelRepo dai.DiscordChannelDai
	discordServerRepo  dai.DiscordServerRegistry
	discordNotifyRepo  dai.DiscordNotify
}

// DiscordNotify はDiscordに関するユースケースインターフェース
type DiscordNotify interface {
	PushNewForum(arg *models.Hackathon) error
	DeleteForums(hackathonID string) error
	CreateNewForumTag(tags []*models.StatusTag) error
}

// NewDiscordNotifyInteractor はDiscordNotifyの生成
func NewDiscordNotifyInteractor(discordChannelRepo dai.DiscordChannelDai, discordServerRepo dai.DiscordServerRegistry, discordNotifyRepo dai.DiscordNotify) DiscordNotify {
	return &discordNotify{
		discordChannelRepo: discordChannelRepo,
		discordServerRepo:  discordServerRepo,
		discordNotifyRepo:  discordNotifyRepo,
	}
}

// PushNewForum は新しいフォーラムを作成する
func (d *discordNotify) PushNewForum(arg *models.Hackathon) error {
	channels, err := d.discordServerRepo.FindAllServers()
	if err != nil {
		return err
	}

	var (
		discordForumIDs = make(chan string, len(channels))
		errChan         = make(chan error, len(channels))
		wg              sync.WaitGroup
	)

	for _, channel := range channels {
		wg.Add(1)
		go func(ch string) {
			defer wg.Done()
			forumID, err := d.discordNotifyRepo.CreateNewForum(ch, arg)
			if err != nil {
				errChan <- err
				return
			}
			discordForumIDs <- forumID
		}(channel.ChannelID)
	}

	wg.Wait()

	var forumIDs []*models.HackathonDiscordChannel

	for i := 0; i < len(channels); i++ {
		select {
		case err := <-errChan:
			return err
		case forumID := <-discordForumIDs:
			forumIDs = append(forumIDs, &models.HackathonDiscordChannel{
				HackathonID: arg.HackathonID,
				ChannelID:   forumID,
			})
		}
	}

	return d.discordChannelRepo.AddChannel(forumIDs)
}

// DeleteForums はフォーラムを削除する
func (d *discordNotify) DeleteForums(hackathonID string) error {
	channels, err := d.discordChannelRepo.GetChannelIDs(hackathonID)
	if err != nil {
		return err
	}

	var (
		errChan = make(chan error, len(channels))
		wg      sync.WaitGroup
	)
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(ch string) {
			defer wg.Done()
			err := d.discordNotifyRepo.DeleteChannel(ch)
			if err != nil {
				errChan <- err
			}
		}(channel.ChannelID)
	}

	wg.Wait()

	for i := 0; i < len(channels); i++ {
		select {
		case err := <-errChan:
			return err
		}
	}

	return d.discordChannelRepo.RemoveChannel(hackathonID)
}

// CreateNewForumTag は新しいフォーラムタグを作成する
func (d *discordNotify) CreateNewForumTag(tags []*models.StatusTag) error {
	channels, err := d.discordServerRepo.FindAllServers()
	if err != nil {
		return err
	}

	var statusTag []string

	for _, tag := range tags {
		statusTag = append(statusTag, tag.Status)
	}

	var (
		errChan   = make(chan error, len(channels))
		guildTags = make(chan []string, len(channels))
		wg        sync.WaitGroup
	)

	for _, channel := range channels {
		wg.Add(1)
		go func(ch string) {
			defer wg.Done()
			tag, err := d.discordNotifyRepo.AddAvailableTags(ch, statusTag)
			if err != nil {
				errChan <- err
				return
			}
			guildTags <- tag
		}(channel.ChannelID)
	}

	wg.Wait()
	if len(errChan) != 0 {
		return <-errChan
	}

	// TODO: ここでタグをDBに保存する

	return nil
}
