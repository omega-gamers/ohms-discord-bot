package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// getInviteMetadata returns the invite metadata object.
func getInviteMetadata(inviteID string, channelID string, discord *discordgo.Session) (im *InviteMetadata) {
	invites, err := discord.ChannelInvites(channelID)
	if err != nil {
		log.Println("Failed to retrieve guild invite information.", err)
		return
	}

	for _, invite := range invites {
		if invite.Code == inviteID {
			im = &InviteMetadata{
				Uses:    invite.Uses,
				Code:    invite.Code,
				Guild:   invite.Guild.ID,
				Channel: invite.Channel.ID,
			}
		}
	}

	return
}

// dumpInviteMetadata handles dumping the invite metadata from discord.
func dumpInviteMetadata(inviteID string, channelID string, discord *discordgo.Session) {
	im := getInviteMetadata(inviteID, channelID, discord)

	log.Printf("Invite metadata is: %v", im)
}
