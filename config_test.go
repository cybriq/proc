package proc_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	// Normally dots are bad but for a spec this makes sense
	"github.com/cybriq/proc"
	"github.com/cybriq/proc/types"
)

var log = proc.GetLogger(proc.PathBase)

func TestCreate(t *testing.T) {
	_ = createAndMarshalUnmarshal(t, &proc.Configs{})
}

func TestConcurrency(t *testing.T) {
	cfgs := createAndMarshalUnmarshal(t, &proc.Configs{})
	nameList := make([]string, len(descs))
	for i := range descs {
		nameList[i] = descs[i].Name
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		for j := range nameList {
			go func(name string) {
				wg.Add(1)
				item, err := cfgs.Get(name)
				if log.E.Chk(err) {
					t.Fail()
				}
				switch item.Type() {
				case types.Bool:
					v := item.Bool()
					item.(*proc.BoolT).Set(!v)
					err := item.(*proc.BoolT).
						FromString(fmt.Sprint(!v))
					if log.E.Chk(err) {
						t.Fail()
					}

				case types.Int:
					v := item.Int()
					item.(*proc.IntT).Set(v + 1)
					err := item.(*proc.IntT).
						FromString(fmt.Sprint(v + 1))
					if log.E.Chk(err) {
						t.Fail()
					}

				case types.Uint:
					v := item.Uint()
					item.(*proc.UinT).Set(v + 1)
					err := item.(*proc.UinT).
						FromString(fmt.Sprint(v + 1))
					if log.E.Chk(err) {
						t.Fail()
					}

				case types.Duration:
					v := item.Duration()
					item.(*proc.DurT).Set(v + time.Second)
					err := item.(*proc.DurT).FromString(
						fmt.Sprint(v + time.Second),
					)
					if log.E.Chk(err) {
						t.Fail()
					}

				case types.Float:
					v := item.Float()
					item.(*proc.FltT).Set(v + 1)
					err := item.(*proc.FltT).FromString(
						fmt.Sprint(v + 1),
					)
					if log.E.Chk(err) {
						t.Fail()
					}

				case types.String:
					v := item.String()
					item.(*proc.StrT).Set(v + "a")
					err := item.(*proc.StrT).FromString(
						v + " ",
					)
					if err != nil {
						t.Fail()
					}

				case types.List:
					v := item.List()
					item.(*proc.LstT).Set(append(v, v[0])...)
					err := item.(*proc.LstT).FromString(
						item.String() + `,"zzz"`,
					)
					if err != nil {
						t.Fail()
					}
				}
				wg.Done()
			}(nameList[j])
		}
	}
	wg.Wait()
}

var descs = []proc.Desc{
	{
		Name:        "boolflag",
		Type:        types.Bool,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: "false",
		Tags:    proc.List("tag1", "tag2"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "intflag",
		Type:        types.Int,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: "-42",
		Tags:    proc.List("tag1", "tag2"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "uintflag",
		Type:        types.Uint,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.
		
		With many lines of textypes.
		
		And several paragraphs
		
		- even some sort of markup
		`,
		Default: "322",
		Tags:    proc.List("tag1", "tag2"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "durationflag",
		Type:        types.Duration,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: "1h2m3s",
		Tags:    proc.List("tag1", "tag2"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "floatflag",
		Type:        types.Float,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: "3.1415927",
		Tags:    proc.List("tag1", "tag2", "tag3"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "stringflag",
		Type:        types.String,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: "itsame",
		Tags:    proc.List("tag1"),
		Aliases: proc.List("BF"),
	},
	{
		Name:        "listflag",
		Type:        types.List,
		Group:       "group",
		Description: "this is a description",
		Documentation: `This is documentation.

With many lines of textypes.

And several paragraphs

- even some sort of markup
`,
		Default: `"links","two","three","four"`,
		Tags:    proc.List("tag1", "tag2"),
		Aliases: proc.List("BF"),
	},
}

func createAndMarshalUnmarshal(t *testing.T, cfgs *proc.Configs) *proc.Configs {
	*cfgs = proc.Create(descs...)
	j, err := cfgs.MarshalJSON()
	if err != nil {
		t.Fail()
	}
	log.I.Ln("\n", string(j))
	err = cfgs.UnmarshalJSON(j)
	if err != nil {
		t.Fail()
	}
	return cfgs
}
