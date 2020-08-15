package main

import (
	"github.com/abstract300/theeye/command"
	"github.com/pkg/errors"
)

func SayHello(ctx *command.Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, "Hello :)")
	return errors.Wrap(err, "hello command failed")
}

func SayPing(ctx *command.Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, "pong")
	return errors.Wrap(err, "ping command failed")
}
