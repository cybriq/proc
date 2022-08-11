package config

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/cybriq/proc/types"
)

func TestCreate(t *testing.T) {
	cfgs := createConfig(t, &Configs{})
	_ = marshalUnmarshal(t, cfgs)
}

func TestConcurrency(t *testing.T) {
	cfgs := createConfig(t, &Configs{})
	nameList := make(map[string]string, len(descs))
	for i := range descs {
		nameList[descs[i].Group] = descs[i].Name
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		for j := range nameList {
			go readwrite(j, nameList[j], &wg, cfgs, t)
		}
	}
	wg.Wait()
	_ = marshalUnmarshal(t, cfgs)
}

func TestPersistence(t *testing.T) {
	cfgs := createConfig(t, &Configs{})
	b, err := cfgs.MarshalJSON()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	s := string(b)
	filename := "./test" + fmt.Sprint(rand.Int()) + ".json"
	err = cfgs.Save(filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	err = cfgs.Load(filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	b, err = cfgs.MarshalJSON()
	err = os.Remove(filename)
	if err != nil {
		t.Log("file '" + filename + "' not deleted")
	}
	if string(b) != s {
		t.Log("written file does not match after being re-read")
		t.Fail()
	}
}

func TestConfigs_GetHelp(t *testing.T) {
	cfgs := createConfig(t, &Configs{})
	text, _ := cfgs.GetHelp("", "", false)
	// if err != nil {
	// 	panic(err)
	// }
	t.Log("\n" + text)
	text, _ = cfgs.GetHelp("", "", true)
	// if err != nil {
	// 	panic(err)
	// }
	t.Log("\n" + text)
	text, _ = cfgs.GetHelp("", "group2", true)
	// if err != nil {
	// 	panic(err)
	// }
	t.Log("\n" + text)
	text, _ = cfgs.GetHelp("stringflag", "", true)
	// if err != nil {
	// 	panic(err)
	// }
	t.Log("\n" + text)
}

func TestConfigs_LoadAllFromEnv(t *testing.T) {
	c := createConfig(t, &Configs{})
	for i := range c.items {
		for j := range c.items[i] {
			t.Log(c.GetEnvString(i, j))
		}
	}
	c.LoadAllFromEnv()
}

func readwrite(group, name string, wg *sync.WaitGroup, cfgs *Configs,
	t *testing.T) {

	wg.Add(1)
	item, err := cfgs.Get(group, name)
	if log.E.Chk(err) {
		t.Log(err)
		t.Fail()
	}
	switch item.Type() {
	case types.Bool:
		v := item.Bool()
		item.(*types.BooT).Set(!v)
		err = item.FromString(fmt.Sprint(!v))
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.Int:
		v := item.Int()
		item.(*types.IntT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.Uint:
		v := item.Uint()
		item.(*types.UinT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.Duration:
		v := item.Duration()
		item.(*types.DurT).Set(v + time.Second)
		err = item.FromString(fmt.Sprint(v + time.Second))
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.Float:
		v := item.Float()
		item.(*types.FltT).Set(v + 1)
		err = item.FromString(fmt.Sprint(v + 1))
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.String:
		v := item.String()
		item.(*types.StrT).Set(v + "a")
		err = item.FromString(v + " ")
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}

	case types.List:
		v := item.List()
		item.(*types.LstT).Set(append(v, v[0])...)
		err = item.FromString(item.String() + `,zzz`)
		if log.E.Chk(err) {
			t.Log(err)
			t.Fail()
		}
	}
	wg.Done()
}

var descs = []types.Desc{
	{
		Name:        "boolflag",
		Type:        types.Bool,
		Group:       "group2",
		Description: "this is a boolean flag",
		Documentation: `
		This is documentation.
		
		With many lines of text.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "false",
		Aliases: List("BF"),
	},
	{
		Name:        "intflag",
		Type:        types.Int,
		Group:       "group2",
		Description: "this is an integer flag, stored as a 64 bit signed integer",
		Documentation: `
		This is documentation.
		
		With many lines of text.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "-42",
		Aliases: List("IF", "Hitchhikers"),
	},
	{
		Name:        "uintflag",
		Type:        types.Uint,
		Group:       "group3",
		Description: "this is an unsigned 64 bit integer",
		Documentation: `
		This is documentation.
				
		With many lines of text.
				
		And several paragraphs
				
		- even some sort of markup
`,
		Default: "322",
		Aliases: List("UF"),
	},
	{
		Name:        "durationflag",
		Type:        types.Duration,
		Group:       "group3",
		Description: "this is a time.Duration, storing a length of time as a 64 bit integer",
		Documentation: `
		This is documentation.
		
		With many lines of textypes.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "1h2m3s",
		Aliases: List("DF"),
	},
	{
		Name:        "floatflag",
		Type:        types.Float,
		Group:       "group1",
		Description: "this is a 64 bit floating point number",
		Documentation: `
		This is documentation.
		
		With many lines of text.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "3.1415927",
		Aliases: List("FF", "floaty"),
	},
	{
		Name:        "stringflag",
		Type:        types.String,
		Group:       "",
		Description: "this is just a plain old string, use quotes if it contains spaces",
		Documentation: `
		This is documentation.
		
		With many lines of text.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "itsa me",
		Aliases: List("SF"),
	},
	{
		Name:        "listflag",
		Type:        types.List,
		Group:       "",
		Description: "This is a series of strings separated by commas - to add spaces the list must be quoted",
		Documentation: `
		This is documentation.
		
		With many lines of text.
		
		And several paragraphs
		
		- even some sort of markup
`,
		Default: "links,two,three,four",
		Aliases: List("BF"),
	},
}

func createConfig(t *testing.T, cfgs *Configs) *Configs {
	*cfgs = Create("testapp", descs...)
	return cfgs
}

func marshalUnmarshal(t *testing.T, cfgs *Configs) *Configs {
	j, err := cfgs.MarshalJSON()
	if log.E.Chk(err) {
		t.Fail()
	}
	s := string(j)
	log.I.Ln("\n", string(j))
	err = cfgs.UnmarshalJSON(j)
	if log.E.Chk(err) {
		t.Fail()
	}
	j, err = cfgs.MarshalJSON()
	if log.E.Chk(err) {
		t.Fail()
	}
	if string(j) != s {
		t.Log("marshal/unmarshal changed the content of configs")
		t.Fail()
	}
	return cfgs
}
