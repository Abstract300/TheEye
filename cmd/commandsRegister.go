package main

import "github.com/abstract300/theeye/command"

func init() {
	perms := command.Permissions{
		UserID: []string{"440548249812729867", "628489151582306315"},
		//	UserID:    []string{},
		ChannelID: []string{"734465180393668608"},
	}
	router = command.NewRoute("?", perms)
	router.NewCommand("fort", "give fortification list", "?fort", Fortification)
	router.NewCommand("siege", "give siege list", "?siege", Siege)
	router.NewCommand("covert", "give covert list", "?covert", Covert)
	router.NewCommand("sentry", "give sentry list", "?sentry", Sentry)
	router.NewCommand("cons", "give conscription list", "?cons", Conscription)
	router.NewCommand("tech", "give tech list", "?tech", Tech)
	router.NewCommand("econ", "give econ list", "?econ", Econ)
	router.NewCommand("exp", "give exp list", "?exp", Exp)
	router.NewCommand("safe", "give safe list", "?safe", Safe)
}
