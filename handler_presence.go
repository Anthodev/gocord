package gocord

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func PresenceButtonsHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"d1": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			updatePresenceHandler(s, i, 0)
		},
		"d2": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			updatePresenceHandler(s, i, 1)
		},
		"d3": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			updatePresenceHandler(s, i, 2)
		},
		"d4": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			updatePresenceHandler(s, i, 3)
		},
		"d5": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			updatePresenceHandler(s, i, 4)
		},
	}
}

func updatePresenceHandler(s *discordgo.Session, i *discordgo.InteractionCreate, dayNb int) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	updatePresenceMessageHandler(s, i, dayNb)

	if err != nil {
		return
	}
}

func updatePresenceMessageHandler(s *discordgo.Session, i *discordgo.InteractionCreate, dayNb int) {
	srcMsgEmbed := i.Interaction.Message.Embeds[0]

	nbValues := srcMsgEmbed.Fields[dayNb].Value
	nbValuesArray := strings.Split(nbValues, ",")

	updatePresenceMessage(s, i, dayNb, nbValuesArray)
}

func updatePresenceMessage(s *discordgo.Session, i *discordgo.InteractionCreate, dayNb int, dayValues []string) {
	srcMsgEmbeds := i.Interaction.Message.Embeds

	var dayUserList []string
	var userIdList []string
	userIdList = retrieveUserListFromMessageEmbed(dayValues)

	if dayValues[0] == "---" {
		dayUserList = append(dayUserList, i.Interaction.Member.User.Mention())
	} else {
		if ContainsUserID(userIdList, i.Interaction.Member.User.ID) {
			for _, v := range dayValues {
				if v != i.Interaction.Member.User.Mention() {
					if v != "---" {
						dayUserList = append(dayUserList, v)
					}
				}
			}

			if len(dayUserList) == 0 {
				dayUserList = append(dayUserList, "---")
			}
		} else {
			dayUserList = append(dayValues, i.Interaction.Member.User.Mention())
		}
	}

	srcMsgEmbeds[0].Fields[dayNb].Value = strings.Join(dayUserList, ",")

	_, err := s.FollowupMessageEdit(i.Interaction, i.Message.ID, &discordgo.WebhookEdit{
		Embeds: &srcMsgEmbeds,
	})

	if err != nil {
		return
	}
}

func retrieveUserListFromMessageEmbed(userIdList []string) []string {
	var userList []string

	for _, u := range userIdList {
		input := u
		start := strings.Index(input, "<@")
		end := strings.Index(input, ">")

		if start != -1 && end != -1 && end > start {
			output := input[start+2 : end]
			output = strings.Trim(output, "<@>")
			userList = append(userList, output)
		}

	}

	return userList
}
