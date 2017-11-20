package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

const (
	metadata string = "metadata"
)

var (
	ctx    context.Context
	client *datastore.Client
)

func createClient(projectID string) {
	log.Println("creating client...")
	var err error

	ctx = context.Background()
	client, err = datastore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func updateInviteMetadata(projectID string, guildID string, data *InviteMetadata) {
	isReady()

	metadataKey := getInviteMetadataKey(projectID, guildID)

	if _, err := client.Put(ctx, metadataKey, data); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
}

func getStoredInviteMetadata(projectID string, guildID string) *InviteMetadata {
	isReady()
	im := &InviteMetadata{}

	metadataKey := getInviteMetadataKey(projectID, guildID)
	if err := client.Get(ctx, metadataKey, im); err != nil {
		log.Fatalf("Failed to locate invite metadata for provided key: %s", err)
	}

	return im
}

func hasInviteMetadataChanged(sim *InviteMetadata, im *InviteMetadata) bool {
	if sim.Uses != im.Uses {
		log.Println("Invite has been used, stored data has changed")
		return true
	}

	return false
}

// isReady is a Helper function that checks if the
// client variable has been initialized.
func isReady() bool {
	if client == nil {
		log.Fatalf("Client wasn't initialized prior to invoking updateInviteMetadata")
		return false
	}

	return true
}

func getInviteMetadataKey(projectID string, guildID string) *datastore.Key {
	name := projectID + guildID

	metadataKey := datastore.NameKey(metadata, name, nil)

	return metadataKey
}
