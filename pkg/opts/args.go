package opts

import "strings"

// ParseCLIArgs reads a command line argument slice (presumably from
// os.Args), identifies the command to run and returns a list of Entry
// values that should be overlaid last after default,
// config and environment variables are parsed.
//
// Rules for constructing CLI args:
//
// - Commands are identified by name, and must appear in their hierarchic
//   order to invoke subcommands. They are matched as normalised to lower
//   case.
//
// - Options can be preceded by "--" or "-", and the full name, or the
//   alias, normalised to lower case for matching, and if there is an "="
//   after it, the value is after this, otherwise, the next element in the
//   args is the value, except booleans, which default to true unless set
//   to false. For these, the prefix "no" or similar indicates the
//   semantics of the true value.
//
// - Options only match when preceded by their relevant Command, except for
//   the root Command, and these options must precede any other command
//   options.
//
// - If no command is selected, the root Command.Default is selected. This
//   can optionally be used for subcommands as well, though it is unlikely
//   needed, if found, the Default of the tip of the Command branch
//   selected by the CLI if there is one, otherwise the Command itself.
func (c *Command) ParseCLIArgs(args []string) (run *Command,
	entries []Entry, err error) {

	log.I.Ln(args)

	var segments [][]string
	var depth, last, cursor int
	var done bool
	current := c
	// First pass matches Command names in order to slice up the sections
	// where relevant config items will be found.
	for !done {
		for i := range current.Commands {
			// log.I.Ln(current.Commands[i])
			if normalise(args[cursor]) == normalise(current.Commands[i].Name) {
				log.I.Ln(args[cursor], normalise(current.Commands[i].Name))
				depth++
				current = current.Commands[i]
				segments = append(segments, args[last:cursor])
				log.I.Ln(segments, cursor)
				last = cursor
				break
			}
		}
		cursor++
		if cursor >= len(args) {
			log.I.Ln(len(args), cursor)
			segments = append(segments, args[last:cursor])
			done = true
		}
	}
	// The segments that have been cut from args will now provide the root level
	// command name, and all subsequent items until the next segment should be
	// names found in the configs map.
	for i := range segments {

		log.I.Ln(i, segments[i])
	}
	return
}

func normalise(s string) string {
	return strings.ToLower(s)
}
