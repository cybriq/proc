package cmds

import "os"

// Help is a default top level command that subsequent
func Help() (h *Command) {
	h = &Command{
		Name:          "Help",
		Description:   "Print help information, optionally multiple keywords can be given and will be searched to generate output",
		Documentation: "Uses partial matching, if result is ambiguous, prints general top level help and the list of partial match terms",
		Entrypoint: func(_ interface{}) (err error) {
			// First we need to parse the CLI args for the content after the
			// help keyword.
			args := make([]string, len(os.Args))
			for i := range os.Args {
				args[i] = os.Args[i]
			}
			var cursor int
			for {
				if cursor < len(args) {
					if normalise(args[cursor]) == "help" {
						args = args[cursor+1:]
						break
					}
					cursor++
				} else {
					break
				}
			}

			return
		},
		Parent: nil,
	}
	return
}
