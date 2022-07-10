package proc_test

import (
	"testing"

	// Normally dots are bad but for a spec this makes sense
	. "gitlab.com/cybriqsystems/proc"
	// T is a common symbol for Type in Go
	T "gitlab.com/cybriqsystems/proc/types"
)

func TestCreate(t *testing.T) {
	cfgs := Create(
		Desc{
			Name:        "boolflag",
			Type:        T.Bool,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "false",
			Tags:    List("tag1", "tag2"),
			Aliases: List("BF"),
		},
	)
	_ = cfgs
}
