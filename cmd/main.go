package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/abstract300/theeye/command"
	"github.com/abstract300/theeye/token"
	"github.com/bwmarrin/discordgo"
)

var (
	router *command.Route
	dg     *discordgo.Session
)

func main() {
	Token, err := token.NewToken("token.json", token.Noop{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Discord session using the provided bot token.
	dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)

		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)

		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	/*
		err := router.CommandHandler(*m.Message, dg)
		if err != nil {
			log.Println(err)
		}
	*/
	err := router.ListenAndServe(m.Message, dg)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
