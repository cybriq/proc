package cmds

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/cybriq/proc/pkg/opts/config"
	"github.com/cybriq/proc/pkg/opts/meta"
	"github.com/cybriq/proc/pkg/opts/text"
	"github.com/cybriq/proc/pkg/path"
)

type Op func(c *Command, args []string) error

var NoOp = func(c *Command, args []string) error { return nil }
var Tags = func(s ...string) []string {
	return s
}

// Command is a specification for a command and can include any number of
// subcommands, and for each Command a list of options
type Command struct {
	path.Path
	Name          string
	Description   string
	Documentation string
	Entrypoint    Op
	Parent        *Command
	Commands      Commands
	Configs       config.Opts
	Default       []string // specifies default subcommand to execute
	sync.Mutex
}

// Commands are a slice of Command entries
type Commands []*Command

func (c *Command) AddCommand(cm *Command) {
	c.Commands = append(c.Commands, cm)
}

const configFilename = "config.toml"

// GetConfigBase creates an option set that should go in the root of a
// Command specification for an application, providing a data directory path
// and config file path.
//
// This exists in order to simplify setup for application configuration
// persistence.
func GetConfigBase(in config.Opts, appName string, abs bool) {
	var defaultDataDir, defaultConfigFile string
	switch runtime.GOOS {
	case "linux", "aix", "freebsd", "netbsd", "openbsd", "dragonfly":
		defaultDataDir = fmt.Sprintf("~/.%s", appName)
		defaultConfigFile =
			fmt.Sprintf("%s/%s", defaultDataDir, configFilename)
	case "windows":
		defaultDataDir = fmt.Sprintf("%%LOCALAPPDATA%%\\%s", appName)
		defaultConfigFile =
			fmt.Sprintf("%%LOCALAPPDATA%%\\%s\\%s", defaultDataDir,
				configFilename)
	case "darwin":
		defaultDataDir = filepath.Join(
			"~", "Library",
			"Application Support", strings.ToUpper(appName),
		)
		defaultConfigFile = filepath.Join(defaultDataDir, configFilename)
	}
	options := config.Opts{
		"ConfigFile": text.New(meta.Data{
			Aliases:       []string{"CF"},
			Label:         "Configuration File",
			Description:   "location of configuration file",
			Documentation: "<placeholder for detailed documentation>",
			Default:       defaultConfigFile,
		}, text.NormalizeFilesystemPath(abs, appName)),
		"DataDir": text.New(meta.Data{
			Aliases:       []string{"DD"},
			Label:         "Data Directory",
			Description:   "root folder where application data is stored",
			Documentation: "<placeholder for detailed documentation>",
			Default:       defaultDataDir,
		}, text.NormalizeFilesystemPath(abs, appName)),
	}
	for i := range options {
		in[i] = options[i]
	}
}

// Init sets up a Command to be ready to use. Puts the reverse paths into the
// tree structure, puts sane defaults into command launchers, runs the hooks on
// all the defined configuration values, and sets the paths on each Command and
// Option so that they can be directly interrogated for their location.
func Init(c *Command, p path.Path) (cmd *Command, err error) {
	if c.Parent != nil {
		log.T.Ln("backlinking children of", c.Parent.Name)
	}
	if c.Entrypoint == nil {
		c.Entrypoint = NoOp
	}
	if p == nil {
		p = path.Path{c.Name}
	}
	c.Path = p // .Parent()
	for i := range c.Configs {
		c.Configs[i].SetPath(p)
	}
	for i := range c.Commands {
		c.Commands[i].Parent = c
		c.Commands[i].Path = p.Child(c.Commands[i].Name)
		_, _ = Init(c.Commands[i], p.Child(c.Commands[i].Name))
	}
	c.ForEach(func(cmd *Command, _ int) bool {
		for i := range cmd.Configs {
			err := cmd.Configs[i].RunHooks()
			if log.E.Chk(err) {
				return false
			}
		}
		return true
	}, 0, 0, c)
	return c, err
}

func (c *Command) GetOpt(path path.Path) (o config.Option) {
	p := make([]string, len(path))
	for i := range path {
		p[i] = path[i]
	}
	switch {
	case len(p) < 1:
		// not found
		return
	case len(p) > 2:
		// search subcommands
		for i := range c.Commands {
			if c.Commands[i].Name == p[1] {
				return c.Commands[i].GetOpt(p[1:])
			}
		}
	case len(p) == 2:
		// check name matches path, search for config item
		if c.Name == p[0] {
			for i := range c.Configs {
				if i == p[1] {
					return c.Configs[i]
				}
			}
		}
	}
	return nil
}

// Cmd is a convenience function but probably unnecessary when named field
// sparse struct literals are just as convenient.
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

// ForEach runs a closure on every node in the Commands tree, stopping if the
// closure returns false
func (c *Command) ForEach(cl func(*Command, int) bool, hereDepth,
	hereDist int, cmd *Command) (ocl func(*Command, int) bool, depth,
	dist int, cm *Command) {
	ocl = cl
	cm = cmd
	if hereDepth == 0 {
		if !ocl(cm, hereDepth) {
			return
		}
	}
	depth = hereDepth + 1
	log.T.Ln(path.GetIndent(depth)+"->", depth)
	dist = hereDist
	for i := range c.Commands {
		log.T.Ln(path.GetIndent(depth)+"walking", c.Commands[i].Name, depth,
			dist)
		if !cl(c.Commands[i], depth) {
			return
		}
		dist++
		ocl, depth, dist, cm = c.Commands[i].ForEach(
			cl,
			depth,
			dist,
			cm,
		)
	}
	log.T.Ln(path.GetIndent(hereDepth)+"<-", hereDepth)
	depth--
	return
}
