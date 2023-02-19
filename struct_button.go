package gocord

import "github.com/bwmarrin/discordgo"

type Button struct {
	*discordgo.Button
}

func NewButton() *Button {
	return &Button{&discordgo.Button{}}
}

func (b *Button) SetLabel(label string) *Button {
	b.Label = label
	return b
}

func (b *Button) SetStyle(style discordgo.ButtonStyle) *Button {
	b.Style = style
	return b
}

func (b *Button) SetCustomID(id string) *Button {
	b.CustomID = id
	return b
}
