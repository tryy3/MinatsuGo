package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
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
			output += strings.ToLower(CONFIG.Prefix + cmd.Name)
		}

		sendMessage(session, message.ChannelID, output)
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

func infoCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	var output string
	m, _ := mem.VirtualMemory()
	h, _ := host.Info()
	l, _ := load.Avg()
	ls, _ := load.Misc()
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", h)
	fmt.Printf("%#v\n", l)
	fmt.Printf("%#v\n", ls)
	output += "***BOT INFO***\n"
	output += "**Uptime:** " + string(time.Since(BOT.StartTime)) + "\n"
	output += "***SERVER INFO***\n"
	output += "**Uptime:** " + string(h.Uptime) + "\n"
	output += "**OS:** " + h.OS + "\n"
	output += "**Total Ram:** " + strconv.FormatFloat(float64(m.Total)/1024/1024/1024, 'f', 2, 32) + "GB\n"
	output += "**Available Ram:** " + strconv.FormatFloat(float64(m.Available)/1024/1024/1024, 'f', 2, 32) + "GB\n"
	output += "**Used Ram:** " + strconv.FormatFloat(float64(m.Used)/1024/1024/1024, 'f', 2, 32) + "GB (*" + strconv.FormatFloat(m.UsedPercent, 'f', -1, 32) + "%*)\n"

	sendMessage(session, message.ChannelID, output)
}

func shutdownCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	sendMessage(session, message.ChannelID, "***Shutting down, good bye cruel world!***")
	close(runChan)
}

func restartCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
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
