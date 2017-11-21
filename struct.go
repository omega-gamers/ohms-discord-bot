package main

// InviteMetadata object for invite data that's tracked.
type InviteMetadata struct {
	Uses    int    `datastore:"uses,noindex" json:"uses"`
	Code    string `datastore:"code,noindex" json:"code"`
	Guild   string `datastore:"guild,noindex" json:"guild"`
	Channel string `datastore:"channel,noindex" json:"channel"`
}

// Configuration object for application.
type Configuration struct {
	GuildID       string   `json:"guildID"`
	ProjectID     string   `json:"projectID"`
	Token         string   `json:"token"`
	InviteID      string   `json:"inviteID"`
	RoleIDs       []string `json:"roleIDs"`
	NotifyChannel string   `json:"notifyChannel"`
	ChannelID     string   `json:"channelID"`
}
