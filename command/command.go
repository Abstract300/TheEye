package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Runner func(ctx *Context) error

// Route holds the route for a command based on the prefix.
type Route struct {
	Prefix string
	Cmd    map[string]*Command
}

// Command is a command to be run and its attribues.
type Command struct {
	Name string
	Desc string
	Run  Runner
}

// Context provides context for a command to execute.
type Context struct {
	ChannelID string
	Args      []string
	Session   *discordgo.Session
}

// NewRoute generates a Route.
func NewRoute(prefix string) *Route {
	return &Route{
		Prefix: prefix,
		Cmd:    make(map[string]*Command, 1),
	}
}

// NewCommand adds a command to a route.
func (r *Route) NewCommand(name string, run Runner) {
	cmd := &Command{
		Name: name,
		Run:  run,
	}
	r.Cmd[cmd.Name] = cmd
}

// FindCommand finds if a command exists in a route.
func (r *Route) FindCommand(name string) *Command {
	cmd := r.Cmd[name]
	if cmd == nil {
		// return a dummy command.
		return &Command{Run: func(ctx *Context) error { return nil }}
	}

	return cmd
}

// CommandHandler is injected into the CreateMessage event handler to automatically invoke a command.
func (r *Route) CommandHandler(msg discordgo.Message, session *discordgo.Session) error {
	content := msg.Content

	cmdName, args := parseMessageContent(content, r.Prefix)

	ctx := &Context{
		ChannelID: msg.ChannelID,
		Args:      args,
		Session:   session,
	}

	cmd := r.FindCommand(cmdName)

	err := cmd.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "command failed to execute")
	}

	return nil
}

// parseMessageContent parses message for the command name and its args.
func parseMessageContent(msg, prefix string) (string, []string) {
	tokens := strings.Split(msg, " ")
	namePrefix := tokens[0]

	ok := strings.HasPrefix(namePrefix, prefix)
	if !ok {
		return "", []string{}
	}

	return namePrefix[1:], tokens[1:]
}
