package main

import "github.com/abstract300/theeye/command"

func init() {
	router = command.NewRoute("?")
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
