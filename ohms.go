package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/omega-gamers/discordgo"
	"google.golang.org/appengine"
)

// Sets up the config reader and assigns application values from the config.
func init() {
	configure()
}

const configFile = "./config.json"

var (
	config *Configuration
)

func configure() {
	if content, err := ioutil.ReadFile(configFile); err != nil {
		log.Fatalf("Failed to load configuration file: %v. Terminating...", configFile)
	} else {
		c := Configuration{}
		json.Unmarshal(content, &c)
		config = &c
	}
}

func startOhms() {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
		return
	}

	if config.DumpRoleIDs {
		roles, err := bot.GuildRoles(config.GuildID)
		if err != nil {
			log.Fatalf("Failed to retrieve guild roles: %v", err)
		}
		log.Println("Roles found are: ", roles)
		return
	}

	im := getInviteMetadata(config.InviteID, bot)

	projectID := config.ProjectID
	guildID := config.GuildID

	createClient(projectID)
	updateInviteMetadata(projectID, guildID, im)

	bot.AddHandler(processNewUsers)
}

func main() {
	appengine.Main()
	startOhms()
}

// Handler function that checks if a user joined from the watched invite link.
// If the new user did join from the watched invite link, then the stored metadata
// is updated, and the new user is assigned the configured roles.
func processNewUsers(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	log.Println("Processing a new user...")
	// im := getInviteMetadata(inviteID, s)
	// if hasInviteMetadataChanged(im) {
	// 	updateInviteMetadata(projectID, guildID, im)
	//user := event.Member.User
	//userID := user.ID
	//for _, roleID := range roleIDs {
	//s.GuildMemberRoleAdd(guildID, userID, roleID)
	//}
	//}

}
