package gocord

import (
	"github.com/bwmarrin/discordgo"
)

func BotCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "presence",
			Description: "Demander aux utilisateurs de mettre leur présence",
			GuildID:     *GuildID,
		},
	}
}
