package gocord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strconv"
	"time"
)

var (
	startDate = time.Now()
	endDate   = time.Now()
)

func CreatePresenceEmbedMessage(s *discordgo.Session, i *discordgo.InteractionCreate, isUpdated bool) {
	options := i.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	if len(options) > 0 {
		for _, option := range options {
			optionsMap[option.Name] = option
		}

		isUpdated, _ = strconv.ParseBool(optionsMap["update"].StringValue())
	}

	originalDate := time.Now()
	currentDate := time.Now()

	if isUpdated {
		for _, v := range i.Interaction.Message.Embeds {
			if v.Type == "rich" {
				title := v.Title

				re := regexp.MustCompile(`\d+`)
				match := re.FindSubmatch([]byte(title))

				if strconv.Itoa(weekStartDate(originalDate).Day()) == string(match[0]) {
					startDate = weekStartDate(currentDate)
					endDate = weekEndDate(startDate)
				} else {
					startDate = nextWeekStartDate(currentDate)
					endDate = nextWeekEndDate(startDate)
				}

				CreatePresenceMessage(s, i, startDate, endDate)
			}
		}
	} else {
		if int(originalDate.Weekday()) == 0 || int(originalDate.Weekday()) >= 2 {
			startDate = nextWeekStartDate(currentDate)
			endDate = nextWeekEndDate(startDate)
		} else {
			startDate = weekStartDate(currentDate)
			endDate = weekEndDate(startDate)
		}

		CreatePresenceMessage(s, i, startDate, endDate)
	}
}

func CreatePresenceMessage(s *discordgo.Session, i *discordgo.InteractionCreate, startDate time.Time, endDate time.Time) {
	embed := CreatePresenceEmbed()

	actionRow := CreatePresenceActionRow()

	SendEmbed(embed, actionRow, s, i)
}

func CreatePresenceEmbed() *discordgo.MessageEmbed {
	embed := NewGenericEmbed(
		fmt.Sprintf("Présence au bureau pour la semaine du %s au %s", startDate.Format("02/01"), endDate.Format("02/01")),
		"Indiquer vos jours de présence prévus au bureau pour la semaine indiquée.",
	)

	fields := []*discordgo.MessageEmbedField{
		{
			"**Lundi**",
			"---",
			false,
		},
		{
			"**Mardi**",
			"---",
			false,
		},
		{
			"**Mercredi**",
			"---",
			false,
		},
		{
			"**Jeudi**",
			"---",
			false,
		},
		{
			"**Vendredi**",
			"---",
			false,
		},
	}

	embed.Fields = fields

	return embed
}

func CreatePresenceActionRow() discordgo.ActionsRow {
	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			NewButton().
				SetLabel("Lundi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d1"),
			NewButton().
				SetLabel("Mardi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d2"),
			NewButton().
				SetLabel("Mercredi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d3"),
			NewButton().
				SetLabel("Jeudi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d4"),
			NewButton().
				SetLabel("Vendredi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d5"),
		},
	}

	return actionRow
}

func weekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}

func weekEndDate(date time.Time) time.Time {
	return date.Add(time.Duration(4*24) * time.Hour)
}

func nextWeekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) + 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}

func nextWeekEndDate(date time.Time) time.Time {
	return date.Add(time.Duration(4*24) * time.Hour)
}
