package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Sets up the config reader and assigns application values from the config.
func init() {
	configure()
}

// Global config object
var (
	config      *Configuration
	debug       bool
	dumpInvite  bool
	dumpRoleIDs bool
)

// configure is a helper function that reads in the external config file.
func configure() {
	var configFile string
	flag.StringVar(&configFile, "c", "/usr/local/etc/ohms-discord-bot.d/config.json", "Config file")
	flag.BoolVar(&debug, "d", false, "Debug mode")
	flag.BoolVar(&dumpInvite, "I", false, "Dumps the invite metadata from discord")
	flag.BoolVar(&dumpRoleIDs, "R", false, "Dumps all role ids from the server")
	flag.Parse()

	if content, err := ioutil.ReadFile(configFile); err != nil {
		log.Fatalf("Failed to load configuration file: %v. Terminating...", configFile)
	} else {
		c := Configuration{}
		json.Unmarshal(content, &c)
		config = &c
	}
}

// Entry point
func main() {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)

		return
	}

	if dumpRoleIDs {
		getRoleIDs(bot)

		return
	}

	if dumpInvite {
		dumpInviteMetadata(config.InviteID, config.ChannelID, bot)

		return
	}

	im := getInviteMetadata(config.InviteID, config.ChannelID, bot)

	projectID := config.ProjectID
	guildID := config.GuildID

	// Create the datastore client
	createClient(projectID)
	// Update the datastore with the latest state.
	updateInviteMetadata(projectID, guildID, im)

	// Add handler methods
	bot.AddHandler(reportReady)
	bot.AddHandler(processNewUsers)

	bot.Open()

	// Wait for user input to terminate
	log.Println("Ohms is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	informServerOfShutdown(bot)
	bot.Close()
}

// Temporary event handler to report bot status
func reportReady(s *discordgo.Session, event *discordgo.Ready) {
	// If debug mode is not enabled, do nothing.
	if !debug {
		return
	}
	channelID := config.NotifyChannel
	_, err := s.ChannelMessageSend(channelID, "Ohms bot is online")
	if err != nil {
		log.Printf("Failed to write to channel %v, due to %v", channelID, err)
	}
}

func informServerOfShutdown(discord *discordgo.Session) {
	// If debug mode is not enabled, do nothing.
	if !debug {
		return
	}
	channelID := config.NotifyChannel
	_, err := discord.ChannelMessageSend(channelID, "Ohms bot is going offline. Goodbye...")
	if err != nil {
		log.Printf("Failed to write to notify channel: %v, due to %v", channelID, err)
	}
}

// Handler function that checks if a user joined from the watched invite link.
// If the new user did join from the watched invite link, then the stored metadata
// is updated, and the new user is assigned the configured roles.
func processNewUsers(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	log.Printf("Processing new user: %v", event.User.Username)

	projectID := config.ProjectID
	guildID := config.GuildID
	im := getInviteMetadata(config.InviteID, config.ChannelID, s)
	sim := getStoredInviteMetadata(projectID, guildID)
	if hasInviteMetadataChanged(sim, im) {
		user := event.Member.User
		userID := user.ID
		// If debug is enable, notify the admins of a user being promoted
		if debug {
			channelID := config.NotifyChannel
			_, err := s.ChannelMessageSend(channelID, "New clan member detected. Promoting "+user.Username)
			if err != nil {
				log.Printf("Failed to notify channel %v due to %v", channelID, err)
			}
		}
		updateInviteMetadata(config.ProjectID, guildID, im)

		for _, roleID := range config.RoleIDs {
			err := s.GuildMemberRoleAdd(guildID, userID, roleID)
			if err != nil {
				log.Printf("Failed to promote user %v, due to: %v", user.Username, err)
			}
		}
	}

}

// getRoleIDs is a helper function that returns the
// role name and their ids for a specific guild.
func getRoleIDs(discord *discordgo.Session) {
	roleIDMap := make(map[string]string)
	roles, err := discord.GuildRoles(config.GuildID)
	if err != nil {
		log.Fatalf("Failed to retrieve guild roles: %v", err)
	}
	for _, role := range roles {
		roleIDMap[role.Name] = role.ID
	}
	log.Printf("Roles found are: %v", roleIDMap)
}
