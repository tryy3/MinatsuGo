package main 
 
import ( 
  "time" 
 
  "github.com/bwmarrin/discordgo" 
) 
 
// Command is an exported type that 
// contains the bots command information 
type Command struct { 
  Name        string 
  Aliases     []string 
  Description string 
  Syntax      []string 
  Run         func(*discordgo.Session, *discordgo.MessageCreate, []string) 
  SubCommand  *Command 
} 
 
// Group is an exported type that 
// contains the permission for running 
// certain commands, uses discord roles 
type Group struct { 
  Name        string 
  Members     []string 
  Permissions []string 
} 
 
// Bot is an exported type that 
// contains the basic information 
// about the bot, some autogenerated 
// while others is manually set 
type Bot struct { 
  StartTime   time.Time 
  ID          string 
  Version     string 
  Description string 
  Author      string 
  Website     string 
} 
 
// Config is an exported type that 
// contains basic settings like Token 
// TODO: Add a better config system. 
type Config struct { 
  Token        string 
  Debug        bool 
  DebugChannel string 
  Prefix       string 
} 
Initial commit of all work