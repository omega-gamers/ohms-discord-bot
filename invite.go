package main

import (
	"fmt"

	"github.com/omega-gamers/discordgo"
)

func getInviteMetadata(inviteID string, discord *discordgo.Session) (im *InviteMetadata) {
	invite, err := discord.Invite(inviteID)
	if err != nil {
		fmt.Println("Failed to retrieve guild invite information.", err)
		return
	}

	im = &InviteMetadata{
		Uses:    invite.Uses,
		Code:    invite.Code,
		Guild:   invite.Guild.ID,
		Channel: invite.Channel.ID,
	}

	return
}
