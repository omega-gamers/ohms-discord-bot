package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func getInviteMetadata(inviteID string, discord *discordgo.Session) (im *InviteMetadata) {
	invite, err := discord.Invite(inviteID)
	if err != nil {
		log.Println("Failed to retrieve guild invite information.", err)
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

func dumpInviteMetadata(inviteID string, discord *discordgo.Session) {
	im := getInviteMetadata(inviteID, discord)

	log.Printf("Invite metadata is: %v", im)
}
