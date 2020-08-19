package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Runner func(ctx *Context) error

// Command is a discord command.
type Command struct {
	Name  string
	Desc  string
	Usage string
	Run   Runner
}

// Permissions are for authenticating users/channels.
type Permissions struct {
	UserID    []string
	ChannelID []string
}

// Context holds the messge's context.
type Context struct {
	ChannelID    string
	AuthorID     string
	MessageID    string
	Args         []string
	CommandsHelp string
	Session      *discordgo.Session
}

// Route holds the route for a command based on the prefix.
type Route struct {
	Prefix     string
	Commands   map[string]*Command
	Permission Permissions
}

// NewRoute creates and returns a new Route with the given prefix.
func NewRoute(prefix string, perms Permissions) *Route {
	var route Route
	route.Prefix = prefix
	route.Permission = perms
	route.Commands = make(map[string]*Command, 1)

	route.NewCommand("help", "", prefix+"help", Help)

	return &route
}

// NewCommand adds a command to a route.
func (r *Route) NewCommand(name, desc, usage string, run Runner) {
	cmd := &Command{
		Name:  name,
		Desc:  desc,
		Usage: usage,
		Run:   run,
	}
	r.Commands[cmd.Name] = cmd

	helpCommand := r.Commands["help"]
	helpCommand.Desc += "- " + cmd.Name + " : " + cmd.Desc + ".\n" + "\tUsage: `" + cmd.Usage + "`\n"
}

// Execute runs the command's execution routine.
func (r *Route) Execute(run Runner, ctx *Context) error {
	//check permissions
	var approved int
	for _, userId := range r.Permission.UserID {
		if userId == ctx.AuthorID {
			approved++
			break
		}
	}

	for _, channelId := range r.Permission.ChannelID {
		if channelId == ctx.ChannelID {
			approved++
			break
		}
	}

	// not enough permissions, fail to execute command.
	if approved != 2 {
		err := ctx.Session.MessageReactionAdd(ctx.ChannelID, ctx.MessageID, "❌")
		if err != nil {
			return errors.Wrap(err, "permission denied's emoji reaction failed for "+ctx.MessageID)
		}
		return errors.New("permission denied for " + ctx.AuthorID)
	}

	return run(ctx)
}

// FindCommand finds if a command exists in a route.
func (r *Route) FindCommand(name string, ctx *Context) *Command {
	cmd := r.Commands[name]
	if cmd == nil {
		// return a dummy command.
		return &Command{Run: func(ctx *Context) error { return nil }}
	}

	if cmd.Name == "help" {
		ctx.CommandsHelp = cmd.Desc
	}

	return cmd
}

// CommandHandler is injected into the CreateMessage event handler to automatically invoke a command.
func (r *Route) CommandHandler(msg discordgo.Message, session *discordgo.Session) error {
	content := msg.Content

	cmdName, args := parseMessageContent(content, r.Prefix)

	ctx := &Context{
		ChannelID: msg.ChannelID,
		AuthorID:  msg.Author.ID,
		MessageID: msg.ID,
		Args:      args,
		Session:   session,
	}

	cmd := r.FindCommand(cmdName, ctx)

	// Not command
	if cmd.Name == "" {
		return nil
	}

	// command found
	err := r.Execute(cmd.Run, ctx)
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
