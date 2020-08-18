# Help page generator

1. The view of the page:

```

~ TheEye ~ Aids you in your koc endeavours

Command Help:
(generate list as an ordered list, with subcommands.)
1. ?command : description
2. ?command2: description of command2
 ....
(generate every command registered at compile-time in this order)
```

1. State machine:
```
type Runner func(ctx *Context) error

type Command struct {
	Name 	string
	Desc 	string
	Help 	string
	Run 	Runner
}

type Permissions struct {
	UserID 		[]string
	ChannelID 	[]string
}

type Context struct {
	ChannelID string
	AuthorID  string
	MessageID string
	Args      []string
	Response  string
	Session   *discordgo.Session
}

type Route struct {
	Prefix 		string
	Commands 	map[string]*Command
	Permission 	Permissions
}

func NewRoute(prefix string, perms Permission) *Route {
	var route &Route
	route.Prefix = prefix
	route.Permission = perms
	
	route.NewCommand("help", Help)
}

func (r *Route) HelpGenerator() {
	var helpPage string
	
	for name, cmd := range r.Commands {
		helpPage += "- "  + cmd.Name + " : " + cmd.Desc + ".\n"
	}
	
	helpCmd := r.Commands["help"]
	helpCmd.Help = helpPage 
}

```


(b) Help command
```
func Help(ctx *Context) error {
	Reply("```\n" + ctx.Response + "\n```")
}
```
