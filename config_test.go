package proc_test

import (
	"testing"

	// Normally dots are bad but for a spec this makes sense
	. "github.com/cybriq/proc"
	// T is a common symbol for Type in Go
	T "github.com/cybriq/proc/types"
)

var log = GetLogger(PathBase)

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
		Desc{
			Name:        "intflag",
			Type:        T.Int,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "-42",
			Tags:    List("tag1", "tag2"),
			Aliases: List("BF"),
		},
		Desc{
			Name:        "uintflag",
			Type:        T.Uint,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "322",
			Tags:    List("tag1", "tag2"),
			Aliases: List("BF"),
		},
		Desc{
			Name:        "durationflag",
			Type:        T.Duration,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "1h2m3s",
			Tags:    List("tag1", "tag2"),
			Aliases: List("BF"),
		},
		Desc{
			Name:        "floatflag",
			Type:        T.Float,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "3.1415927",
			Tags:    List("tag1", "tag2", "tag3"),
			Aliases: List("BF"),
		},
		Desc{
			Name:        "stringflag",
			Type:        T.String,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: "itsame",
			Tags:    List("tag1"),
			Aliases: List("BF"),
		},
		Desc{
			Name:        "listflag",
			Type:        T.List,
			Group:       "group",
			Description: "this is a description",
			Documentation: `This is documentation.

With many lines of text.

And several paragraphs

- even some sort of markup
`,
			Default: `"links","two","three","four"`,
			Tags:    List("tag1", "tag2"),
			Aliases: List("BF"),
		},
	)
	j, err := cfgs.MarshalJSON()
	if err != nil {
		t.Fail()
	}
	log.I.Ln("\n", string(j))
	err = cfgs.UnmarshalJSON(j)
	if err != nil {
		t.Fail()
	}
}
