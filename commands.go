package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

var commands []*Command

func init() {
	commands = []*Command{
		&Command{
			Name: "Info",
			Aliases: []string{
				"i",
			},
			Description: "Bot and server information",
			Syntax: []string{
				"<command>",
			},
			Run: infoCommand,
		},
		&Command{
			Name: "Help",
			Aliases: []string{
				"h",
				"helpme",
			},
			Description: "Basic help command",
			Syntax: []string{
				"<command> [list]",
				"<command> find (command)",
			},
			Run: helpCommand,
		},
		&Command{
			Name: "Shutdown",
			Aliases: []string{
				"stop",
			},
			Description: "Command to shutdown the bot",
			Syntax: []string{
				"<command>",
			},
			Run: shutdownCommand,
		},
		&Command{
			Name: "Restart",
			Aliases: []string{
				"reboot",
			},
			Description: "Command to restart the bot",
			Syntax: []string{
				"<command>",
			},
			Run: restartCommand,
		},
		&Command{
			Name: "Smite",
			Aliases: []string{
				"god",
			},
			Description: "Command to get a random god from smite",
			Syntax: []string{
				"<command> (settings)",
			},
			Run: smiteCommand,
		},
	}
}

func helpCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	if len(args) <= 1 || strings.ToLower(args[1]) == "list" {
		var output = "Registered commands: "

		for _, cmd := range commands {
			output += strings.ToLower(CONFIG.Prefix+cmd.Name) + ", "
		}

		sendMessage(session, message.ChannelID, strings.TrimSuffix(output, ", "))
	}
	if len(args) == 3 && strings.ToLower(args[1]) == "find" {
		var cmd = findCommand(strings.ToLower(args[2]))
		var output string

		var aliasOutput = ""

		for _, alias := range cmd.Aliases {
			aliasOutput += CONFIG.Prefix + alias + ", "
		}

		output += "**Name:** " + CONFIG.Prefix + cmd.Name + "\n"
		output += "**Description:** " + cmd.Description + "\n"
		output += "**Aliases:** " + strings.TrimSuffix(aliasOutput, ", ") + "\n"

		output += "**Syntax:** "

		for _, syntax := range cmd.Syntax {
			output += strings.Replace(syntax, "<command>", CONFIG.Prefix+strings.ToLower(cmd.Name), 1) + "\n\t"
		}

		sendMessage(session, message.ChannelID, output)
	}
}

func getTime(t uint64) string {
	s := strconv.FormatUint(t, 10)
	d, _ := time.ParseDuration(s)
	o := time.Unix(time.Now().Unix()-int64(t), 0)
	fmt.Println(t)
	fmt.Println(int64(t))
	fmt.Println(s)
	fmt.Println(d)
	fmt.Println(o)
	return humanize.Time(o)
}

func infoCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	var output string
	m, _ := mem.VirtualMemory()
	h, _ := host.Info()
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	output += "***BOT INFO***\n"
	output += "**Uptime:** " + humanize.Time(BOT.StartTime) + "\n"
	output += "**Discordgo:** " + discordgo.VERSION + "\n"
	output += "**Go:** " + runtime.Version() + "\n"
	output += "**Bot:** v." + BOT.Version + "\n"
	output += "**Allocated Ram:** " + humanize.Bytes(stats.TotalAlloc) + "\n"
	output += "**Total Allocated Ram:** " + humanize.Bytes(stats.Alloc) + "\n"
	output += "**Bot Obtained Ram:** " + humanize.Bytes(stats.Sys) + "\n"
	output += "\n***SERVER INFO***\n"
	output += "**Uptime:** " + getTime(h.Uptime) + "\n"
	output += "**OS:** " + h.OS + "\n"
	output += "**Total Ram:** " + humanize.Bytes(m.Total) + "\n"
	output += "**Available Ram:** " + humanize.Bytes(m.Available) + "\n"
	output += "**Used Ram:** " + humanize.Bytes(m.Used) + " (*" + strconv.FormatFloat(m.UsedPercent, 'f', -1, 32) + "%*)\n"

	sendMessage(session, message.ChannelID, output)
}

func shutdownCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	sendMessage(session, message.ChannelID, "***Shutting down, good bye cruel world!***")
	os.Exit(1)
}

func restartCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	sendMessage(session, message.ChannelID, "***Why you rebooting me, icri!***")
	os.Exit(2)
}

func smiteCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	if len(args) >= 2 {
		if args[1] == "random" {
			types := parseTypes(args[2:])
			fmt.Printf("%#v\n", types)
			/*gods := getGods([]string{"Chinese", "Magical"})

			for _, god := range gods {
				fmt.Printf("%#v\n", god)
			}*/
		}
		if args[1] == "update" {
			SMITE.UpdateGods()

			for _, god := range SMITE.Gods {
				fmt.Printf("%#v\n", god)
			}
		}
	}
}

func findCommand(cmd string) *Command {
	cmd = strings.ToLower(cmd)
	for _, c := range commands {
		if cmd == strings.ToLower(c.Name) {
			return c
		}
		for _, alias := range c.Aliases {
			if cmd == strings.ToLower(alias) {
				return c
			}
		}
	}
	return nil
}
