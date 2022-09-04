package opts

import (
	"encoding"
	"sync"

	"github.com/cybriq/proc/pkg/opts/config"
)

type Op func(c interface{}) error

var NoOp = func(c interface{}) error { return nil }
var Tags = func(s ...string) []string {
	return s
}

// Command is a specification for a command and can include any number of
// subcommands, and for each Command a list of options
type Command struct {
	Name          string
	Description   string
	Documentation string
	Entrypoint    Op
	Parent        *Command
	Commands      Commands
	Configs       config.Opts
	sync.Mutex
}

// Commands are a slice of Command entries
type Commands []*Command

func Init(c *Command) *Command {
	if c.Parent != nil {
		log.T.Ln("backlinking children of", c.Parent.Name)
	}
	if c.Entrypoint == nil {
		c.Entrypoint = NoOp
	}
	for i := range c.Commands {
		c.Commands[i].Parent = c
		Init(c.Commands[i])
	}
	return c
}

func Cmd(name, desc, doc string, op Op, cfg map[string]config.Option,
	cmds ...*Command) (c *Command) {

	c = &Command{
		Name:          name,
		Description:   desc,
		Documentation: doc,
		Entrypoint:    op,
		Commands:      cmds,
		Configs:       cfg,
	}
	return
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
		current := cmd.Parent
		for current != nil {
			if current.Name != "" {
				cmdPath = current.Name + "."
			}
			current = current.Parent
		}
		text = append(text,
			[]byte("["+cmdPath+cmd.Name+"]"+"\n")...)
		return true
	})
	return
}
