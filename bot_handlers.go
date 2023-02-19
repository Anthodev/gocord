package gocord

import (
	"github.com/bwmarrin/discordgo"
)

func PresenceHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"presence": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			CreatePresenceEmbedMessage(s, i, false)
		},
	}
}
