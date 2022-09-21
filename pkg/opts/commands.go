package opts

import (
	"encoding"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/cybriq/proc/pkg/opts/config"
	"github.com/cybriq/proc/pkg/opts/meta"
	"github.com/naoina/toml"
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
func (c *Command) Foreach(cl func(*Command, int) bool, hereDepth, hereDist int,
	cmd *Command) (ocl func(*Command, int) bool, depth, dist int,
	cm *Command) {
	ocl = cl
	cm = cmd
	if hereDepth == 0 {
		if !ocl(cm, hereDepth) {
			return
		}
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
		ocl, depth, dist, cm = c.Commands[i].Foreach(
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
		if cmd == nil {
			log.I.Ln("cmd empty")
			return true
		}
		cfgNames := make([]string, 0, len(cmd.Configs))
		for i := range cmd.Configs {
			cfgNames = append(cfgNames, i)
		}
		if len(cfgNames) < 1 {
			return true
		}
		if cmd.Name != "" {
			var cmdPath string
			current := cmd.Parent
			for current != nil {
				if current.Name != "" {
					cmdPath = current.Name + "."
				}
				current = current.Parent
			}
			text = append(text,
				[]byte("# "+cmdPath+cmd.Name+": "+cmd.Description+"\n")...)
			text = append(text,
				[]byte("["+cmdPath+cmd.Name+"]"+"\n\n")...)
		}
		sort.Strings(cfgNames)
		for _, i := range cfgNames {
			md := cmd.Configs[i].Meta()
			lq, rq := "", ""
			st := cmd.Configs[i].String()
			df := md.Default()
			switch cmd.Configs[i].Type() {
			case meta.Duration, meta.Text:
				lq, rq = "\"", "\""
			case meta.List:
				lq, rq = "[ \"", "\" ]"
				st = strings.ReplaceAll(st, ",", "\", \"")
				df = strings.ReplaceAll(df, ",", "\", \"")
				if st == "" {
					lq, rq = "[ ", "]"
				}
			}
			text = append(text,
				[]byte("# "+i+" - "+md.Description()+
					" - default: "+lq+df+rq+"\n")...)
			text = append(text,
				[]byte(i+" = "+lq+st+rq+"\n")...)
		}
		text = append(text, []byte("\n")...)
		return true
	}, 0, 0, c)
	return
}

func (c *Command) UnmarshalText(text []byte) (err error) {
	var out interface{}
	err = toml.Unmarshal(text, &out)
	// log.I.S(out)
	o := walk([]string{}, out, []Entry{})
	// log.I.Ln(o)
	_ = o
	return
}

type Path []string

func (p Path) String() string {
	return strings.Join(p, "/")
}

type Entry struct {
	path  Path
	name  string
	value interface{}
}

func (e Entry) String() string {
	return fmt.Sprint(e.path, "/", e.name, "=", e.value)
}

type Entries []Entry

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	iPath, jPath := e[i].String(), e[j].String()
	return iPath < jPath
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func walk(p []string, v interface{}, in Entries) (o Entries) {
	o = append(o, in...)
	var parent []string
	for i := range p {
		parent = append(parent, p[i])
	}
	// log.I.Ln(p, parent)
	switch vv := v.(type) {
	case map[string]interface{}:
		for i := range vv {
			// log.I.Ln(reflect.TypeOf(vv[i]))
			switch vvv := vv[i].(type) {
			case map[string]interface{}:
				o = walk(append(parent, i), vvv, o)
			default:
				o = append(o, Entry{
					path:  parent,
					name:  i,
					value: vv[i],
				})

			}
		}
	}
	return
}
