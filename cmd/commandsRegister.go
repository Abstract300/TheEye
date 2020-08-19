package main

import (
	"log"

	"github.com/abstract300/theeye/command"
)

func init() {
	perms := command.Permissions{
		UserID: []string{"440548249812729867", "628489151582306315"},
		//	UserID:    []string{},
		ChannelID: []string{"734465180393668608"},
	}
	//router = command.NewRoute("?", perms)
	route, err := command.NewRoute("?", perms)
	if err != nil {
		log.Printf("%+v", err)

		return
	}

	router = route

	router.NewCommand("fort", "give fortification list", Fortification)
	router.NewCommand("siege", "give siege list", Siege)
	router.NewCommand("covert", "give covert list", Covert)
	router.NewCommand("sentry", "give sentry list", Sentry)
	router.NewCommand("cons", "give conscription list", Conscription)
	router.NewCommand("tech", "give tech list", Tech)
	router.NewCommand("econ", "give econ list", Econ)
	router.NewCommand("exp", "give exp list", Exp)
	router.NewCommand("safe", "give safe list", Safe)
}
