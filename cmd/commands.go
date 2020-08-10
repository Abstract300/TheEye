package main

import "github.com/abstract300/theeye/command"

func SayHello(ctx *command.Context) {
	ctx.Session.ChannelMessageSend(ctx.ChannelID, "Hello :)")
}

func SayPing(ctx *command.Context) {
	ctx.Session.ChannelMessageSend(ctx.ChannelID, "pong")
}
