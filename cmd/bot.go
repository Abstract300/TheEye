package main

import (
	"log"
	"os"

	"github.com/abstract300/theeye/config"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Bot struct {
	Session *discordgo.Session
	Logger  *log.Logger
}

func NewBot(tokenFile string) (*Bot, error) {
	bot := &Bot{}
	token := &Token{}
	tokenData, err := NewToken(tokenFile, token)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] Token unavailable.")
	}

	dg, err := discordgo.New("Bot " + tokenData)
	if err != nil {
		return &Bot{}, errors.Wrap(err, "[Error] discord session unavailable.")
	}

	logger := log.New(os.Stdout, "[TheEye]", log.Lshortfile)

	bot.Session = dg
	bot.Logger = logger
	return bot, nil
}

func NewToken(tokenFile string, token *Token) (string, error) {
	tokenData, err := config.NewConfig("token.json", token)
	if err != nil {
		return "", errors.Wrap(err, "[Error] parsing token from "+tokenFile)
	}
	return tokenData, nil
}
