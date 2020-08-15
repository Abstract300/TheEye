package main

import "github.com/abstract300/theeye/command"

func init() {
	perms := command.Permission{
		UserID: []string{"440548249812729867", "628489151582306315"},
		//	UserID:    []string{},
		ChannelID: []string{"734465180393668608"},
	}
	router = command.NewRoute("?", perms)
	router.NewCommand("fort", Fortification)
	router.NewCommand("siege", Siege)
	router.NewCommand("siege", Siege)
	router.NewCommand("covert", Covert)
	router.NewCommand("sentry", Sentry)
	router.NewCommand("cons", Conscription)
	router.NewCommand("tech", Tech)
	router.NewCommand("econ", Econ)
	router.NewCommand("exp", Exp)
	router.NewCommand("safe", Safe)
}
