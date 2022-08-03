package proc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cybriq/proc/types"
	"go.uber.org/atomic"
)

// Configs are a concurrent safe key value store for configuration items
type Configs struct {
	items map[string]types.Item
	sync.Mutex
	persistenceLock sync.Mutex
}

type (
	// BooT is a boolean item type
	BooT struct {
		value atomic.Bool
		*types.Metadata
	}
	// DurT is a duration item type
	DurT struct {
		value atomic.Duration
		*types.Metadata
	}
	// FltT is a float64 item type
	FltT struct {
		value atomic.Float64
		*types.Metadata
	}
	// IntT is an int64 item type
	IntT struct {
		value atomic.Int64
		*types.Metadata
	}
	// LstT is a []string item type
	LstT struct {
		value []string
		*sync.Mutex
		*types.Metadata
	}
	// StrT is a string item type
	StrT struct {
		value atomic.String
		*types.Metadata
	}
	// UinT is an uint64 item type
	UinT struct {
		value atomic.Uint64
		*types.Metadata
	}
)

// Assert that all of the above types satisfy the types.Item interface
var _ = []types.Item{
	&BooT{},
	&DurT{},
	&FltT{},
	&IntT{},
	&LstT{},
	&StrT{},
	&UinT{},
}

// Boolean item implementation

func (b *BooT) FromString(s string) error {
	asRunes := []rune(s)
	first := string(asRunes[0])
	first = strings.ToLower(first)
	if first != "t" && first != "f" {
		return errors.New("string form of bool must start with 't' or 'f'")
	}
	if first == "t" {
		b.value.Store(true)
	} else {
		b.value.Store(false)
	}
	return nil
}
func (b *BooT) Set(bo bool)            { b.value.Store(bo) }
func (b BooT) Bool() bool              { return b.value.Load() }
func (b BooT) Duration() time.Duration { panic("type error") }
func (b BooT) Float() float64          { panic("type error") }
func (b BooT) Int() int64              { panic("type error") }
func (b BooT) List() []string          { panic("type error") }
func (b BooT) String() string          { return fmt.Sprint(b.value.Load()) }
func (b BooT) Uint() uint64            { panic("type error") }
func (b BooT) Meta() *types.Metadata   { return b.Metadata }

// Get returns a named item from the Configs
func (c *Configs) Get(name string) (t types.Item, err error) {
	var ok bool
	c.Lock()
	t, ok = c.items[name]
	c.Unlock()
	if !ok {
		err = fmt.Errorf("item '%s' not found", name)
	}
	return
}

// GetAllNames returns a lexicographically sorted list of all item names in
// the Configs.
func (c *Configs) GetAllNames() (items []string) {

	items = make([]string, len(c.items))
	counter := 0
	for i := range c.items {
		// the Create function uses the Configs.Name field as map keys,
		// so we don't need to interrogate further.
		items[counter] = i
	}
	sort.Strings(items)
	return
}

type help struct {
	types.Metadata
	detailed, markup bool
}

func toHelp() (h help) {
	return
}

var helpTemplate = `
{{.name}} ({{.type}})
Aliases: <>
Description: <>
Documentation:
<>
Default: <>
Tags: <>

`

func (c *Configs) GetHelp(name, group string, detailed,
	markdown bool) (text string,
	err error) {

	if name == "" && group == "" {
		allNames := c.GetAllNames()
		var o string
		for i := range allNames {
			o, err = c.GetHelp(allNames[i], "", detailed, markdown)
			if log.E.Chk(err) {
				panic(err)
			}
			text += o
		}
	}
	var item types.Item
	item, err = c.Get(name)
	if err != nil {
		return
	}
	// construct the text
	_ = item
	return
}

func (c *Configs) Save(filename string) (err error) {
	var b []byte
	c.persistenceLock.Lock()
	defer c.persistenceLock.Unlock()
	b, err = c.MarshalJSON()
	if log.E.Chk(err) {
		return
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if log.E.Chk(err) {
		return
	}
	return
}

func (c *Configs) Load(filename string) (err error) {
	var b []byte
	c.persistenceLock.Lock()
	defer c.persistenceLock.Unlock()
	b, err = ioutil.ReadFile(filename)
	if log.E.Chk(err) {
		return
	}
	err = c.UnmarshalJSON(b)
	if log.E.Chk(err) {
		return
	}
	return
}

// MarshalJSON returns the JSON for the current state of a Configs
func (c *Configs) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{}, len(c.items))
	c.Lock()
	for i := range c.items {
		switch c.items[i].Type() {
		case types.Bool:
			out[i] = c.items[i].Bool()

		case types.Int:
			out[i] = c.items[i].Int()

		case types.Uint:
			out[i] = c.items[i].Uint()

		case types.Duration:
			out[i] = c.items[i].Duration()

		case types.Float:
			out[i] = c.items[i].Float()

		case types.String:
			out[i] = c.items[i].String()

		case types.List:
			out[i] = c.items[i].List()
		}
	}
	c.Unlock()
	b, err := json.MarshalIndent(out, "", "\t")
	return b, err
}

// UnmarshalJSON loads a Configs with the values in a JSON string
func (c *Configs) UnmarshalJSON(bytes []byte) error {
	in := make(map[string]interface{})
	err := json.Unmarshal(bytes, &in)
	for i := range in {
		v, ok := c.items[i]
		if !ok {
			return fmt.Errorf(
				"configuration does not contain item with"+
					" name %s", i)
		}
		switch v.Type() {
		case types.Bool:
			t := c.items[i].(*BooT)
			ta := &t
			(*ta).Set(v.Bool())

		case types.Int:
			t := c.items[i].(*IntT)
			ta := &t
			(*ta).Set(v.Int())

		case types.Uint:
			t := c.items[i].(*UinT)
			ta := &t
			(*ta).Set(v.Uint())

		case types.Duration:
			t := c.items[i].(*DurT)
			ta := &t
			(*ta).Set(v.Duration())

		case types.Float:
			t := c.items[i].(*FltT)
			ta := &t
			(*ta).Set(v.Float())

		case types.String:
			t := c.items[i].(*StrT)
			ta := &t
			(*ta).Set(v.String())

		case types.List:
			t := c.items[i].(*LstT)
			ta := &t
			(*ta).Set(v.List()...)
		}
	}
	return err
}

// Duration item implementation

func (d *DurT) FromString(s string) error {
	i, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.value.Store(i)
	return nil
}
func (d *DurT) Set(du time.Duration)   { d.value.Store(du) }
func (d DurT) Bool() bool              { panic("type error") }
func (d DurT) Duration() time.Duration { return d.value.Load() }
func (d DurT) Float() float64          { panic("type error") }
func (d DurT) Int() int64              { return int64(d.value.Load()) }
func (d DurT) List() []string          { panic("type error") }
func (d DurT) String() string          { return fmt.Sprint(d.value.Load()) }
func (d DurT) Uint() uint64            { panic("type error") }
func (d DurT) Meta() *types.Metadata   { return d.Metadata }

// Floating point item implementation

func (f *FltT) FromString(s string) error {
	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.value.Store(fl)
	return nil
}
func (f *FltT) Set(fl float64)         { f.value.Store(fl) }
func (f FltT) Bool() bool              { panic("type error") }
func (f FltT) Duration() time.Duration { panic("type error") }
func (f FltT) Float() float64          { return f.value.Load() }
func (f FltT) Int() int64              { panic("type error") }
func (f FltT) List() []string          { panic("type error") }
func (f FltT) String() string          { return fmt.Sprint(f.value.Load()) }
func (f FltT) Uint() uint64            { panic("type error") }
func (f FltT) Meta() *types.Metadata   { return f.Metadata }

func (in *IntT) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	in.value.Store(i)
	return nil
}
func (in *IntT) Set(i int64)            { in.value.Store(int64(i)) }
func (in IntT) Bool() bool              { panic("type error") }
func (in IntT) Duration() time.Duration { panic("type error") }
func (in IntT) Float() float64          { panic("type error") }
func (in IntT) Int() int64              { return in.value.Load() }
func (in IntT) List() []string          { panic("type error") }
func (in IntT) String() string          { return fmt.Sprint(in.value.Load()) }
func (in IntT) Uint() uint64            { panic("type error") }
func (in IntT) Meta() *types.Metadata   { return in.Metadata }

// List of strings item implementation

func (l *LstT) FromString(s string) error {
	split := strings.Split(s, ",")
	for i := range split {
		if !strings.HasPrefix(split[i], "\"") || !strings.HasPrefix(
			split[i], "\"") {
			return errors.New(
				"list items must be enclosed in double quotes" +
					" and cannot contain commas")
		}
		split[i] = split[i][1 : len(split[i])-1]
	}
	l.Set(split...)
	return nil
}
func (l *LstT) Set(li ...string) {
	l.Mutex.Lock()
	l.value = make([]string, len(li))
	copy(l.value, li)
	l.Mutex.Unlock()
}
func (l LstT) Bool() bool              { panic("type error") }
func (l LstT) Duration() time.Duration { panic("type error") }
func (l LstT) Float() float64          { panic("type error") }
func (l LstT) Int() int64              { panic("type error") }
func (l LstT) List() (li []string) {
	l.Mutex.Lock()
	li = make([]string, len(l.value))
	copy(li, l.value)
	l.Mutex.Unlock()
	return
}
func (l LstT) String() (o string) {
	lo := l.List()
	for i := range lo {
		o += "\"" + lo[i]
		if i != len(lo)-1 {
			o += "\","
		}
	}
	return
}
func (l LstT) Uint() uint64          { panic("type error") }
func (l LstT) Meta() *types.Metadata { return l.Metadata }

// String item implementation

func (s *StrT) FromString(st string) error {
	s.value.Store(st)
	return nil
}
func (s *StrT) Set(st string)          { s.value.Store(st) }
func (s StrT) Bool() bool              { panic("type error") }
func (s StrT) Duration() time.Duration { panic("type error") }
func (s StrT) Float() float64          { panic("type error") }
func (s StrT) Int() int64              { panic("type error") }
func (s StrT) List() []string          { panic("type error") }
func (s StrT) String() string          { return s.value.Load() }
func (s StrT) Uint() uint64            { panic("type error") }
func (s StrT) Meta() *types.Metadata   { return s.Metadata }

func (u *UinT) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	u.value.Store(uint64(i))
	return nil
}
func (u *UinT) Set(ui uint64)          { u.value.Store(ui) }
func (u UinT) Bool() bool              { panic("type error") }
func (u UinT) Duration() time.Duration { panic("type error") }
func (u UinT) Float() float64          { panic("type error") }
func (u UinT) Int() int64              { panic("type error") }
func (u UinT) List() []string          { panic("type error") }
func (u UinT) String() string          { return fmt.Sprint(u.value.Load()) }
func (u UinT) Uint() uint64            { return u.value.Load() }
func (u UinT) Meta() *types.Metadata   { return u.Metadata }

// Create a new configuration from a slice of item Desc riptors.
func Create(items ...types.Desc) (c Configs) {
	c = Configs{items: make(map[string]types.Item)}
	c.Lock()
	defer c.Unlock()
	for i := range items {
		name := items[i].Name
		if _, ok := c.items[name]; ok {
			panic("configs contains a duplicate named item: '" + name + "'")
		}
		c.items[name] = Item(types.New(items[i]))
	}
	return
}

// Item takes a Metadata and creates the appropriate item type for it.
func Item(m *types.Metadata) (t types.Item) {
	switch m.Type() {
	case types.Bool:
		t = NewBool(m)
	case types.Duration:
		t = NewDuration(m)
	case types.Float:
		t = NewFloat(m)
	case types.Int:
		t = NewInt(m)
	case types.List:
		t = NewList(m)
	case types.String:
		t = NewString(m)
	case types.Uint:
		t = NewUint(m)
	default:
		panic("invalid type: '" + fmt.Sprint(m.Type()) + "'")
	}
	return
}

// List is a helper that uses variadic parameters to return a slice of strings.
func List(items ...string) []string {
	return items
}

// IfErrNotNilPanic is a helper for functions that should never error as
// errors are purely programmer errors and not error conditions.
func IfErrNotNilPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// NewBool creates a new boolean types.Item
func NewBool(m *types.Metadata) (b *BooT) {
	b = &BooT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewDuration creates a new duration types.Item
func NewDuration(m *types.Metadata) (b *DurT) {
	b = &DurT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewFloat creates a new floating point types.Item
func NewFloat(m *types.Metadata) (b *FltT) {
	b = &FltT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewInt creates a new integer types.Item
func NewInt(m *types.Metadata) (b *IntT) {
	b = &IntT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewList creates a new list of strings types.Item
func NewList(m *types.Metadata) (b *LstT) {
	b = &LstT{Mutex: &sync.Mutex{}}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewString creates a new string types.Item
func NewString(m *types.Metadata) (b *StrT) {
	b = &StrT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewUint creates a new unsigned integer types.Item
func NewUint(m *types.Metadata) (b *UinT) {
	b = &UinT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}
