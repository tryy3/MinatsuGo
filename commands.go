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
Initial commit of all work