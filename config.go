package proc

import (
	"encoding/json"
	"errors"
	"fmt"
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
}

// Desc is the named field form of metadata for generating a metadata
type Desc struct {
	Name, Group, Description, Documentation, Default string
	Type                                             types.Type
	Tags, Aliases                                    []string
}

type (
	// BooT is a boolean item type
	BooT struct {
		value atomic.Bool
		*metadata
	}
	// DurT is a duration item type
	DurT struct {
		value atomic.Duration
		*metadata
	}
	// FltT is a float64 item type
	FltT struct {
		value atomic.Float64
		*metadata
	}
	// IntT is an int64 item type
	IntT struct {
		value atomic.Int64
		*metadata
	}
	// LstT is a []string item type
	LstT struct {
		value []string
		*sync.Mutex
		*metadata
	}
	// StrT is a string item type
	StrT struct {
		value atomic.String
		*metadata
	}
	// UinT is an uint64 item type
	UinT struct {
		value atomic.Uint64
		*metadata
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

// metadata stores the information about the types.Item for documentation
// purposes
type metadata struct {
	sync.Mutex
	name, group, description, documentation, def string
	typ                                          types.Type
	tags, aliases                                []string
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

// Get returns a named item from the Configs
func (c *Configs) Get(name string) (t types.Item, err error) {
	var ok bool
	c.Lock()
	t, ok = c.items[name]
	c.Unlock()
	if !ok {
		err = fmt.Errorf("type '%s' not found", name)
	}
	return
}

func (c *Configs) Save(filename string) error {
	return nil
}

func (c *Configs) Load(filename string) error {
	return nil
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
func (l LstT) Uint() uint64 { panic("type error") }

// metadata accessors

func (m *metadata) Aliases() []string {
	m.Lock()
	defer m.Unlock()
	return m.aliases
}
func (m *metadata) Default() string {
	m.Lock()
	defer m.Unlock()
	return m.def
}

func (m *metadata) Description() string {
	m.Lock()
	defer m.Unlock()
	return m.description
}
func (m *metadata) Documentation() string {
	m.Lock()
	defer m.Unlock()
	return m.documentation
}
func (m *metadata) Group() string {
	m.Lock()
	defer m.Unlock()
	return m.group
}
func (m *metadata) Name() string {
	m.Lock()
	defer m.Unlock()
	return m.name
}
func (m *metadata) Tags() []string {
	m.Lock()
	defer m.Unlock()
	return m.tags
}
func (m *metadata) Type() types.Type {
	m.Lock()
	defer m.Unlock()
	return m.typ
}

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

// Create a new configuration from a slice of item Desc riptors.
func Create(items ...Desc) (c Configs) {
	c = Configs{items: make(map[string]types.Item)}
	c.Lock()
	defer c.Unlock()
	for i := range items {
		name := items[i].Name
		if _, ok := c.items[name]; ok {
			panic("configs contains a duplicate named item: '" + name + "'")
		}
		c.items[name] = Item(New(items[i]))
	}
	return
}

// Item takes a metadata and creates the appropriate item type for it.
func Item(m *metadata) (t types.Item) {
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

// New allows you to create a metadata with a sparsely filled, named field
// struct literal.
//
// name, type, group and tags all will be canonicalized to lower case.
func New(args Desc) *metadata {
	// tags should be all lower case
	for i := range args.Tags {
		args.Tags[i] = strings.ToLower(args.Tags[i])
	}
	// name, type and group should also be lower case
	return &metadata{
		name:          strings.ToLower(args.Name),
		typ:           args.Type,
		aliases:       args.Aliases,
		group:         strings.ToLower(args.Group),
		tags:          args.Tags,
		description:   args.Description,
		documentation: args.Documentation,
		def:           args.Default,
	}
}

// IfErrNotNilPanic is a helper for functions that should never error as
// errors are purely programmer errors and not error conditions.
func IfErrNotNilPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// NewBool creates a new boolean types.Item
func NewBool(m *metadata) (b *BooT) {
	b = &BooT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewDuration creates a new duration types.Item
func NewDuration(m *metadata) (b *DurT) {
	b = &DurT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewFloat creates a new floating point types.Item
func NewFloat(m *metadata) (b *FltT) {
	b = &FltT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewInt creates a new integer types.Item
func NewInt(m *metadata) (b *IntT) {
	b = &IntT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewList creates a new list of strings types.Item
func NewList(m *metadata) (b *LstT) {
	b = &LstT{Mutex: &sync.Mutex{}}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewString creates a new string types.Item
func NewString(m *metadata) (b *StrT) {
	b = &StrT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}

// NewUint creates a new unsigned integer types.Item
func NewUint(m *metadata) (b *UinT) {
	b = &UinT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.metadata = m
	return
}
