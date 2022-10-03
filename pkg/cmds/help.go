package cmds

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cybriq/proc/pkg/opts/config"
)

// Help is a default top level command that subsequent
func Help() (h *Command) {
	h = &Command{
		Name: "help",
		Description: "Print help information, optionally multiple keywords" +
			" can be given and will be searched to generate output",
		Documentation: "Uses partial matching, if result is ambiguous, prints" +
			" general top level help and the list of partial match terms",
		Entrypoint: func(c *Command, args []string) (err error) {
			if args == nil {
				// no args given, just print top level general help
				return
			}
			log.D.Ln("printing help")
			foundCommands := &[]*Command{}
			fops := make(map[string]config.Option)
			foundOptions := &fops
			c.ForEach(func(cm *Command, depth int) bool {
				// search:
				for i := range args {
					// check for match of current command name
					if strings.Contains(normalise(cm.Name), normalise(args[i])) {
						*foundCommands = append(*foundCommands, cm)
					}
					// check for matches on configs
					for ops := range cm.Configs {
						if strings.Contains(normalise(ops), normalise(args[i])) {
							(*foundOptions)[ops] = cm.Configs[ops]
						}
					}
				}
				return true
			}, 0, 0, c)
			var out string
			if len(*foundCommands)+len(*foundOptions) > 0 {
				for i := range *foundCommands {
					cm := (*foundCommands)[i]
					log.I.F("command: [%v] %s - %s", cm.Path, cm.Name, cm.Description)
					log.I.Ln(cm.Documentation)
				}
				for i := range *foundOptions {
					op := (*foundOptions)[i]
					om := op.Meta()
					path := op.Path().TrimPrefix().String()
					if len(path) > 0 {
						path = path + " "
					}
					log.I.F("option: %v%s - %s", path, i,
						om.Description())
					log.I.Ln(om.Documentation())
				}
			} else {
				var outs []string
				for i := range c.Commands {
					outs = append(outs, c.Commands[i].Name)
				}
				sort.Strings(outs)
				out += fmt.Sprintf("\n%s - %s\n\n", c.Name, c.Description)
				out += fmt.Sprintf("Available subcommands\n\n")
				for i := range outs {
					def := ""
					if outs[i] == c.Default[len(c.Default)-1] {
						def = " *"
					}
					out += fmt.Sprintf("\t%s%s\n", outs[i], def)
				}
				out += "\n* indicates default if no subcommand given\n\n"
				out += fmt.Sprintf("for more information:\n\n")
				out += fmt.Sprintf("\t%s help <subcommand>\n\n", os.Args[0])
				out += "Available configuration options at top level:\n\n"
				out += fmt.Sprintf("\t-%s %v - %s (default: '%s')\n",
					"flag", "[alias1 alias2]", "description", "default")
				out += "\t\t(prefix '-' can also be '--')\n\n"
				var opts []string
				for i := range c.Configs {
					opts = append(opts, i)
				}
				sort.Strings(opts)
				for i := range opts {
					aliases := c.Configs[opts[i]].Meta().Aliases()
					for j := range aliases {
						aliases[j] = strings.ToLower(aliases[j])
					}
					var al string
					if len(aliases) > 0 {
						al = fmt.Sprint(aliases, " ")
					}
					out += fmt.Sprintf("\t-%s %v- %s (default: '%s')\n", strings.ToLower(opts[i]),
						al,
						c.Configs[opts[i]].Meta().Description(),
						c.Configs[opts[i]].Meta().Default())
				}

			}
			fmt.Println(out)
			return
		},
		Parent: nil,
	}
	return
}
