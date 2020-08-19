package command

import "github.com/pkg/errors"

func Help(ctx *Context) error {
	helpPage := "TheEye's Help Page\n" + ctx.CommandsHelp
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, helpPage)
	return errors.Wrap(err, "couldn't reply")
}
