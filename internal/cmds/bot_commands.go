package cmds

import "github.com/bwmarrin/discordgo"

func BotCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Pings the bot.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "Message to send",
					Required:    false,
				},
			},
		},
	}
}
