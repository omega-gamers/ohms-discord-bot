package main

// InviteMetadata object for invite data that's tracked.
type InviteMetadata struct {
	Uses    int    `datastore:"uses,noindex"`
	Code    string `datastore:"code,noindex"`
	Guild   string `datastore:"guild,noindex"`
	Channel string `datastore:"channel,noindex"`
}

// Configuration object for application.
type Configuration struct {
	GuildID       string   `json:"guildID"`
	ProjectID     string   `json:"projectID"`
	Token         string   `json:"token"`
	InviteID      string   `json:"inviteID"`
	RoleIDs       []string `json:"roleIDs"`
	DumpRoleIDs   bool     `json:"dumpRoleIDs"`
	NotifyChannel string   `json:"notifyChannel"`
}
