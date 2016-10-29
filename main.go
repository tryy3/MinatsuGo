package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// BOT is a global variable containing
	// basic bot information like startime, version, ID etc.
	BOT *Bot

	// CONFIG is a global variable containing
	// basic config information like token, debug etc.
	CONFIG *Config

	// SMITE is a global variable containing
	// the smite related functions
	SMITE *Smite

	// PERMISSION is a global instance of
	// the permission related functions
	// like hasPermission()
	PERMISSION *PermissionHandler
)

func init() {
	BOT = &Bot{
		StartTime:   time.Now(),
		Version:     "0.0.1",
		Description: "Basic discord bot made in Golang",
		Author:      "tryy3",
		Website:     "pornhub.com",
	}

	CONFIG = &Config{
		Token:        "Token here",
		Debug:        false,
		DebugChannel: "",
		Prefix:       "!",
	}

	SMITE = &Smite{}

	PERMISSION = &PermissionHandler{}
}

func main() {
	// Load new config (create if not existed)
	err := loadConfiguration("config.json", CONFIG)
	if err != nil {
		fmt.Println("Error reading configuration,", err)
		return
	}

	// Load all the permissions, such as group and user permissions.
	err = PERMISSION.load()
	if err != nil {
		fmt.Println("Error loading permissions,", err)
		return
	}

	// Create a new Discord session
	discord, err := discordgo.New(CONFIG.Token)
	if err != nil {
		fmt.Println("Error creating a Discord session,", err)
		return
	}

	// Get the bot user
	user, err := discord.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
		return
	}

	// Save the bots ID for later use
	BOT.ID = user.ID

	// Add a discord chat handler
	discord.AddHandler(readChat)

	// Open/start a discord connection
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	BOT.Session = discord

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}

func sendMessage(session *discordgo.Session, channel string, message string) {
	_, err := session.ChannelMessageSend(channel, message)

	if err != nil {
		fmt.Println("Error when sending a message,", err)
	}
}

func readChat(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BOT.ID {
		return
	}

	if CONFIG.Debug && message.ChannelID != CONFIG.DebugChannel {
		return
	}

	if !strings.HasPrefix(message.Content, CONFIG.Prefix) {
		return
	}

	args := strings.Split(message.Content, " ")

	if !PERMISSION.hasPermission(message.Message, "minatsugo.command."+strings.TrimPrefix(args[0], CONFIG.Prefix)) {
		sendMessage(session, message.ChannelID, "You do not have permission to use this command.")
		return
	}

	command := findCommand(strings.TrimPrefix(args[0], CONFIG.Prefix))

	if command == nil {
		sendMessage(session, message.ChannelID, fmt.Sprintf("The command %s is not a recognizeable command, use %shelp.", message.Content, CONFIG.Prefix))
		return
	}

	command.Run(session, message, args)
}
