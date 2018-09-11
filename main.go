package main

import(
	"log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"./config"
	"./command"
)

var (
	Token string
	Prefix string
)

func main() {
	command.RegisterCommands()
	log.Println("Starting BlackCat...")
	
	for _, cmd := range command.Commands {
		log.Println("Register new command:", cmd.Name)
	}
	
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

	if strings.HasPrefix(m.Content, config.Prefix) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		
		if m.Content == command.Commands[m.Content].Name {
			s.ChannelMessageSend(m.ChannelID, command.Commands[m.Content].Description)
		}
	
	}
}