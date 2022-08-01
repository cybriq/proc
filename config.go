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

type Configs struct {
	items map[string]types.Item
	sync.Mutex
}

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

func (c *Configs) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{}, len(c.items))
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
	b, err := json.MarshalIndent(out, "", "\t")
	return b, err
}

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
			t := c.items[i].(*_bool)
			ta := &t
			(*ta).Set(v.Bool())

		case types.Int:
			t := c.items[i].(*_int)
			ta := &t
			(*ta).Set(v.Int())

		case types.Uint:
			t := c.items[i].(*_uin)
			ta := &t
			(*ta).Set(v.Uint())

		case types.Duration:
			t := c.items[i].(*_dur)
			ta := &t
			(*ta).Set(v.Duration())

		case types.Float:
			t := c.items[i].(*_flt)
			ta := &t
			(*ta).Set(v.Float())

		case types.String:
			t := c.items[i].(*_str)
			ta := &t
			(*ta).Set(v.String())

		case types.List:
			t := c.items[i].(*_lst)
			ta := &t
			(*ta).Set(v.List()...)
		}
	}
	return err
}

type _bool struct {
	value atomic.Bool
	*metadata
}

var _ types.Item = &_bool{}

func NewBool(m *metadata) (b *_bool) {
	b = &_bool{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (b *_bool) FromString(s string) error {
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

func (b _bool) Bool() bool              { return b.value.Load() }
func (b _bool) Int() int64              { panic("type error") }
func (b _bool) Duration() time.Duration { panic("type error") }
func (b _bool) Uint() uint64            { panic("type error") }
func (b _bool) Float() float64          { panic("type error") }
func (b _bool) String() string          { return fmt.Sprint(b.value.Load()) }
func (b _bool) List() []string          { panic("type error") }

func (b *_bool) Set(bo bool) { b.value.Store(bo) }

type _dur struct {
	value atomic.Duration
	*metadata
}

var _ types.Item = &_dur{}

func NewDuration(m *metadata) (b *_dur) {
	b = &_dur{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (d *_dur) FromString(s string) error {
	i, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.value.Store(i)
	return nil
}

func (d _dur) Bool() bool              { panic("type error") }
func (d _dur) Int() int64              { return int64(d.value.Load()) }
func (d _dur) Duration() time.Duration { return d.value.Load() }
func (d _dur) Uint() uint64            { panic("type error") }
func (d _dur) Float() float64          { panic("type error") }
func (d _dur) String() string          { return fmt.Sprint(d.value.Load()) }
func (d _dur) List() []string          { panic("type error") }

func (d *_dur) Set(du time.Duration) { d.value.Store(du) }

type _flt struct {
	value atomic.Float64
	*metadata
}

var _ types.Item = &_flt{}

func NewFloat(m *metadata) (b *_flt) {
	b = &_flt{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (f *_flt) FromString(s string) error {
	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.value.Store(fl)
	return nil
}

func (f _flt) Bool() bool              { panic("type error") }
func (f _flt) Int() int64              { panic("type error") }
func (f _flt) Duration() time.Duration { panic("type error") }
func (f _flt) Uint() uint64            { panic("type error") }
func (f _flt) Float() float64          { return f.value.Load() }
func (f _flt) String() string          { return fmt.Sprint(f.value.Load()) }
func (f _flt) List() []string          { panic("type error") }

func (f *_flt) Set(fl float64) { f.value.Store(fl) }

type _int struct {
	value atomic.Int64
	*metadata
}

var _ types.Item = &_int{}

func NewInt(m *metadata) (b *_int) {
	b = &_int{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (in *_int) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	in.value.Store(i)
	return nil
}

func (in _int) Bool() bool              { panic("type error") }
func (in _int) Int() int64              { return in.value.Load() }
func (in _int) Duration() time.Duration { panic("type error") }
func (in _int) Uint() uint64            { panic("type error") }
func (in _int) Float() float64          { panic("type error") }
func (in _int) String() string          { return fmt.Sprint(in.value.Load()) }
func (in _int) List() []string          { panic("type error") }

func (in *_int) Set(i int64) { in.value.Store(int64(i)) }

type _lst struct {
	value []string
	*sync.Mutex
	*metadata
}

var _ types.Item = &_lst{}

func NewList(m *metadata) (b *_lst) {
	b = &_lst{Mutex: &sync.Mutex{}}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

// FromString converts a comma separated list of strings into a _lst
func (l *_lst) FromString(s string) error {
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
func (l _lst) Bool() bool              { panic("type error") }
func (l _lst) Int() int64              { panic("type error") }
func (l _lst) Duration() time.Duration { panic("type error") }
func (l _lst) Uint() uint64            { panic("type error") }
func (l _lst) Float() float64          { panic("type error") }

func (l _lst) String() (o string) {
	o = "["
	lo := l.List()
	for i := range lo {
		o += "\"" + lo[i] + "\","
	}
	o += "]"
	return
}

func (l _lst) List() (li []string) {
	l.Mutex.Lock()
	li = make([]string, len(l.value))
	copy(li, l.value)
	l.Mutex.Unlock()
	return
}

func (l *_lst) Set(li ...string) {
	l.Mutex.Lock()
	l.value = make([]string, len(li))
	copy(l.value, li)
	l.Mutex.Unlock()
}

// List is a more compact way of declaring a []string
func List(items ...string) []string {
	return items
}

type _str struct {
	value atomic.String
	*metadata
}

var _ types.Item = &_str{}

func NewString(m *metadata) (b *_str) {
	b = &_str{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (s *_str) FromString(st string) error {
	s.value.Store(st)
	return nil
}

func (s _str) Bool() bool              { panic("type error") }
func (s _str) Int() int64              { panic("type error") }
func (s _str) Duration() time.Duration { panic("type error") }
func (s _str) Uint() uint64            { panic("type error") }
func (s _str) Float() float64          { panic("type error") }
func (s _str) String() string          { return s.value.Load() }
func (s _str) List() []string          { panic("type error") }

func (s *_str) Set(st string) { s.value.Store(st) }

type _uin struct {
	value atomic.Uint64
	*metadata
}

var _ types.Item = &_uin{}

func NewUint(m *metadata) (b *_uin) {
	b = &_uin{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (u *_uin) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	u.value.Store(uint64(i))
	return nil
}

func (u _uin) Bool() bool              { panic("type error") }
func (u _uin) Int() int64              { panic("type error") }
func (u _uin) Duration() time.Duration { panic("type error") }
func (u _uin) Uint() uint64            { return u.value.Load() }
func (u _uin) Float() float64          { panic("type error") }
func (u _uin) String() string          { return fmt.Sprint(u.value.Load()) }
func (u _uin) List() []string          { panic("type error") }

func (u *_uin) Set(ui uint64) { u.value.Store(ui) }

// metadata automatically implements everything except the inputs and outputs
type metadata struct {
	sync.Mutex
	name, group, description, documentation, def string
	typ                                          types.Type
	tags, aliases                                []string
}

// Desc is the named field form of metadata for generating a metadata
type Desc struct {
	Name, Group, Description, Documentation, Default string
	Type                                             types.Type
	Tags, Aliases                                    []string
}

func isType(s string) (is bool) {
	for i := range types.Names {
		if s == types.Names[i] {
			is = true
		}
	}
	return
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

func (m *metadata) Name() string {
	m.Lock()
	defer m.Unlock()
	return m.name
}
func (m *metadata) Type() types.Type {
	m.Lock()
	defer m.Unlock()
	return m.typ
}
func (m *metadata) Aliases() []string {
	m.Lock()
	defer m.Unlock()
	return m.aliases
}
func (m *metadata) Group() string {
	m.Lock()
	defer m.Unlock()
	return m.group
}
func (m *metadata) Tags() []string {
	m.Lock()
	defer m.Unlock()
	return m.tags
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
func (m *metadata) Default() string {
	m.Lock()
	defer m.Unlock()
	return m.def
}

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
