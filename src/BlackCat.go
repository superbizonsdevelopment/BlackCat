package main

import(
	"log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"./config"
)

var (
	Token string
	Prefix string
)

func main() {
	log.Println("Starting BlackCat...")
	
	err := config.LoadConfig()
	
	if err != nil {
		log.Println(err.Error())
		return
	}
	
	Token := config.Token
	Prefix := config.Prefix
	
	discord, err := discordgo.New("Bot " + Token)
	
	if err != nil {
		log.Println(err.Error())
		return
	}
	
	discord.AddHandler(messageCreate)
	
	err = discord.Open()
	
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}
	
	log.Println("Bot is working!")
	log.Println("Prefix:", Prefix)
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
	
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if m.Content == "help" {
		s.ChannelMessageSend(m.ChannelID, "Actually, bot hasn't got any commands except `help`")
	}
}