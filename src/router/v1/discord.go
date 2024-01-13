package v1

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (s *v1router) discordHandler() {
	s.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})
}
