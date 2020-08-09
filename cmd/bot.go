package main

import (
	"log"
	"os"

	"github.com/abstract300/theeye/config"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

// Bot is a discord bot.
type Bot struct {
	Session *discordgo.Session
	Logger  *log.Logger
}

// NewBot creates a Bot or returns error upon failure.
func NewBot(tokenFileName string, tokenFile config.FileReader) (*Bot, error) {
	tokenData, err := config.NewConfig(tokenFileName, tokenFile)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] parsing token from "+tokenFileName)
	}

	dg, err := discordgo.New("Bot " + tokenData)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] discord session unavailable.")
	}

	logger := log.New(os.Stdout, "<-:[Bot]:-> ", log.Lshortfile)

	bot := &Bot{
		Session: dg,
		Logger:  logger,
	}

	return bot, nil
}
