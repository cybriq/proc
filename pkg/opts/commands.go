package opts

import (
	"encoding"
	"sync"

	"github.com/cybriq/proc/pkg/opts/config"
)

type Op func(c interface{}) error

var NoOp = func(c interface{}) error { return nil }

// Command is a specification for a command and can include any number of
// subcommands, and for each Command a list of options
type Command struct {
	Name          string
	Description   string
	Documentation string
	Entrypoint    Op
	parent        *Command
	Commands      Commands
	Opts          map[string]config.Option
	sync.Mutex
}

// Commands are a slice of Command entries
type Commands []*Command

func New(c *Command) *Command {
	initCommand(c)
	return c
}

func initCommand(c *Command) {
	if c.parent != nil {
		log.T.Ln("backlinking children of", c.parent.Name)
	}
	if c.Entrypoint == nil {
		c.Entrypoint = NoOp
	}
	for i := range c.Commands {
		c.Commands[i].parent = c
		initCommand(c.Commands[i])
	}
}

const tabs = "\t\t\t\t\t\t\t\t\t\t"

// Foreach runs a closure on every node in the Commands tree, stopping if the
// closure returns false
func (c *Command) Foreach(cl func(*Command, int) bool) {
	c.foreach(cl, 0, 0, nil)
}
func (c *Command) foreach(cl func(*Command, int) bool, hereDepth, hereDist int,
	cmd *Command) (ocl func(*Command, int) bool, depth, dist int,
	cm *Command) {

	ocl = cl
	cm = cmd
	if c == nil {
		dist = hereDist
		depth = hereDepth
		return
	}
	depth = hereDepth + 1
	log.T.Ln(tabs[:depth]+"->", depth)
	dist = hereDist
	for i := range c.Commands {
		log.T.Ln(tabs[:depth]+"walking", c.Commands[i].Name, depth,
			dist)
		if !cl(c.Commands[i], hereDepth) {
			return
		}
		dist++
		ocl, depth, dist, cm = c.Commands[i].foreach(
			cl,
			depth,
			dist,
			cm,
		)
	}
	log.T.Ln(tabs[:hereDepth]+"<-", hereDepth)
	depth--
	return
}

var _ encoding.TextMarshaler = &Command{}

func (c *Command) MarshalText() (text []byte, err error) {
	c.Foreach(func(cmd *Command, depth int) bool {
		text = append(text,
			[]byte("# "+cmd.Description+"\n")...)
		var cmdPath string
		current := cmd.parent
		for current != nil {
			if current.Name != "" {
				cmdPath = current.Name + "."
			}
			current = current.parent
		}
		text = append(text,
			[]byte("["+cmdPath+cmd.Name+"]"+"\n")...)
		return true
	})
	return
}
