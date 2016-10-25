package main 
 
import ( 
  "fmt" 
  "strings" 
  "time" 
 
  "github.com/bwmarrin/discordgo" 
) 
 
var ( 
  // BOT is an instance of the Bot struct 
  BOT *Bot 
 
  // CONFIG is an instance of the Config struct 
  CONFIG *Config 
 
  runChan chan struct{} 
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
} 
 
func main() { 
  err := loadConfiguration("config.json", CONFIG) 
 
  if err != nil { 
    fmt.Println("Error reading configuration,", err) 
    return 
  } 
 
  // create a new Discord session using the token supplied in the command line 
  discord, err := discordgo.New(CONFIG.Token) 
  if err != nil { 
    fmt.Println("Error creating a Discord session,", err) 
    return 
  } 
 
  user, err := discord.User("@me") 
  if err != nil { 
    fmt.Println("error obtaining account details,", err) 
  } 
 
  BOT.ID = user.ID 
 
  discord.AddHandler(readChat) 
 
  err = discord.Open() 
  if err != nil { 
    fmt.Println("Error opening connection,", err) 
    return 
  } 
 
  fmt.Println("Bot is now running. Precc CTRL-C to exit.") 
  runChan := make(chan struct{}) 
  <-runChan 
  return 
} 
 
func sendMessage(session *discordgo.Session, channel string, message string) { 
Initial commit of all work