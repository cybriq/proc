package proc

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/cybriq/proc/types"
)

func TestCreate(t *testing.T) {
	_ = createAndMarshalUnmarshal(t, &Configs{})
}

func TestConcurrency(t *testing.T) {
	cfgs := createAndMarshalUnmarshal(t, &Configs{})
	nameList := make([]string, len(descs))
	for i := range descs {
		nameList[i] = descs[i].Name
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		for j := range nameList {
			go readwrite(nameList[j], wg, cfgs, t)
		}
	}
	wg.Wait()
	j, err := cfgs.MarshalJSON()
	if err != nil {
		t.Fail()
	}
	log.I.Ln("\n", string(j))
}

func readwrite(name string, wg sync.WaitGroup, cfgs *Configs,
	t *testing.T) {

	wg.Add(1)
	item, err := cfgs.Get(name)
	if log.E.Chk(err) {
		t.Fail()
	}
	switch item.Type() {
	case types.Bool:
		v := item.Bool()
		item.(*BooT).Set(!v)
		err = item.FromString(fmt.Sprint(!v))
		if log.E.Chk(err) {
			t.Fail()
		}

	case types.Int:
		v := item.Int()
		item.(*IntT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Fail()
		}

	case types.Uint:
		v := item.Uint()
		item.(*UinT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Fail()
		}

	case types.Duration:
		v := item.Duration()
		item.(*DurT).Set(v + time.Second)
		err = item.FromString(fmt.Sprint(v + time.Second))
		if log.E.Chk(err) {
			t.Fail()
		}

	case types.Float:
		v := item.Float()
		item.(*FltT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Fail()
		}

	case types.String:
		v := item.String()
		item.(*StrT).Set(v + "a")
		err = item.FromString(v + " ")
		if err != nil {
			t.Fail()
		}

	case types.List:
		v := item.List()
		item.(*LstT).Set(append(v, v[0])...)
		err = item.FromString(item.String() + `,"zzz"`)
		if err != nil {
			t.Fail()
		}
	}
	wg.Done()
}

var descs = []Desc{
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
		Tags:    List("tag1", "tag2"),
		Aliases: List("BF"),
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
		Tags:    List("tag1", "tag2"),
		Aliases: List("BF"),
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
		Tags:    List("tag1", "tag2"),
		Aliases: List("BF"),
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
		Tags:    List("tag1", "tag2"),
		Aliases: List("BF"),
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
		Tags:    List("tag1", "tag2", "tag3"),
		Aliases: List("BF"),
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
		Tags:    List("tag1"),
		Aliases: List("BF"),
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
		Tags:    List("tag1", "tag2"),
		Aliases: List("BF"),
	},
}

func createAndMarshalUnmarshal(t *testing.T, cfgs *Configs) *Configs {
	*cfgs = Create(descs...)
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
