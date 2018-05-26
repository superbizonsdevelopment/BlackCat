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
	Commands = make(map[string]command.Command)
	HelpCommand = command.Command{Name: "help", Description: "hehe", MinArgs: 0, MaxArgs: 0, Permission: 0}
	CommandList = []command.Command {
		command.Command {
			Name: "!help",
			Description: "help command description",
			MinArgs: 0,
			MaxArgs: 0,
			Permission: 0,
		},
		command.Command {
			Name: "!avatar",
			Description: "avatar command description",
			MinArgs: 0,
			MaxArgs: 0,
			Permission: 0,
		},
	}
)

func RegisterCommands() {
	for _, cmd := range CommandList{
		Commands[cmd.Name] = command.Command {
			cmd.Name,
			cmd.Description,
			cmd.MinArgs,
			cmd.MaxArgs,
			cmd.Permission,
		}
	}
}

func main() {
	RegisterCommands()
	log.Println("Starting BlackCat...")
	
	for _, cmd := range Commands {
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
		
		if m.Content == Commands[m.Content].Name {
			s.ChannelMessageSend(m.ChannelID, Commands[m.Content].Description)
		}
	
	}
}