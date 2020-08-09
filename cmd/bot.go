package main

import (
	"log"
	"os"

	"github.com/abstract300/theeye/config"
	"github.com/pkg/errors"
)

// BotStarter  wraps the discordgo.New method.
type BotStarter interface {
	New(args ...interface{}) (BotSession, error)
}

// BotCloser wraps discordgo's (*Session).Close method.
type BotCloser interface {
	Close()
}

// BotSession embed interfaces that are specific to a particular discordgo's session.
type BotSession interface {
	BotEventHandler
}

// BotEventHandler wraps discordgo's (*Session).AddHandler method.
type BotEventHandler interface {
	AddHandler(handler interface{}) func()
}

// Bot is a discord bot.
type Bot struct {
	Session BotSession
	Logger  *log.Logger
}

func NewBot(tokenFile string, fr config.FileReader, bs BotStarter) (*Bot, error) {
	tokenData, err := config.NewConfig(tokenFile, fr)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] parsing token from "+tokenFile)
	}

	dg, err := bs.New("Bot " + tokenData)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] discord session unavailable.")
	}

	logger := log.New(os.Stdout, "[TheEye]", log.Lshortfile)

	bot := &Bot{
		Session: dg,
		Logger:  logger,
	}
	return bot, nil
}
