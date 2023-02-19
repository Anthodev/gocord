package gocord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SendResponse(message string, additionalData []string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	var v []string

	if len(additionalData) > 0 {
		v = additionalData
	}

	if len(v) > 0 {
		for _, vv := range v {
			message += "\n" + vv
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s", message),
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}

func SendEmbed(embed *discordgo.MessageEmbed, actionRow discordgo.ActionsRow, s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		fmt.Println(err)
	}

	components := []discordgo.MessageComponent{
		actionRow,
	}

	_, errFollowUp := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			embed,
		},
		Components: components,
	})
	if errFollowUp != nil {
		return
	}
}
