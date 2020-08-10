package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

// Route holds the route for a command based on the prefix.
type Route struct {
	Prefix string
	Cmd    map[string]*Command
}

// Command is a command to be run and its attribues
type Command struct {
	Name string
	Desc string
	Run  func(ctx *Context)
}

// Context provides context for a command to execute
type Context struct {
	ChannelID string
	Args      []string
	Session   *discordgo.Session
}

// NewRoute generates a Route
func NewRoute(prefix string) *Route {
	return &Route{
		Prefix: prefix,
		Cmd:    make(map[string]*Command, 1),
	}
}

// NewCommand adds a command to a route
func (r *Route) NewCommand(name string, run func(ctx *Context)) {
	cmd := &Command{
		Name: name,
		Run:  run,
	}
	r.Cmd[cmd.Name] = cmd
}

// FindCommand finds if a command exists in a route.
func (r *Route) FindCommand(name string) (*Command, error) {
	cmd := r.Cmd[name]
	if cmd == nil {
		return &Command{}, errors.New("Command not found.")
	}
	return cmd, nil
}

// CommandHandler is injected into the CreateMessage event handler to automatically invoke a command.
func (r *Route) CommandHandler(msg discordgo.Message, session *discordgo.Session) error {
	channelID := msg.ChannelID
	content := msg.Content
	cmdName, args, err := parseMessageContent(content, r.Prefix)
	if err != nil {
		return errors.Wrap(err, "Cannot handle command.")
	}

	ctx := &Context{
		ChannelID: channelID,
		Args:      args,
		Session:   session,
	}

	cmd, err := r.FindCommand(cmdName)
	if err != nil {
		return errors.Wrap(err, "FindCommand failed.")
	}

	cmd.Run(ctx)

	return nil
}

// parseMessageContent parses message for the command name and its args
func parseMessageContent(msg, prefix string) (string, []string, error) {
	tokens := strings.Split(msg, " ")
	namePrefix := tokens[0]

	ok := strings.HasPrefix(namePrefix, prefix)
	if ok != true {
		return "", []string{}, errors.New("Illegal route")
	}

	return namePrefix[1:], tokens[1:], nil
}
