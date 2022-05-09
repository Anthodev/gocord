package cmds

import (
	"github.com/anthodev/gocord/internal/api/response"
	"github.com/bwmarrin/discordgo"
)

func BotHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var additionalData []string
			options := i.ApplicationCommandData().Options

			if len(options) > 0 {
				optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
				for _, option := range options {
					optionMap[option.Name] = option
				}

				msgformat := "A command with options:"

				additionalData = append(additionalData, optionMap["message"].StringValue())

				response.SendResponse(msgformat, additionalData, s, i)
			} else {
				response.SendResponse("Pong!", additionalData, s, i)
			}
		},
	}
}
