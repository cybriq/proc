package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/atomic"
)

// Item provides accessors for a Config item type's metadata and current
// contents, including a generic string format setter.
type Item interface {
	Name() string
	Type() Type
	Aliases() []string
	Group() string
	Description() string
	Documentation() string
	Default() string
	FromString(string) error
	Meta() *Metadata
	Concrete
}

// Name is a helper that returns the name associated with a Type.
func Name(T Type) string {
	return Names[T]
}

type (
	// BooT is a boolean item type
	BooT struct {
		value atomic.Bool
		*Metadata
	}
	// DurT is a duration item type
	DurT struct {
		value atomic.Duration
		*Metadata
	}
	// FltT is a float64 item type
	FltT struct {
		value atomic.Float64
		*Metadata
	}
	// IntT is an int64 item type
	IntT struct {
		value atomic.Int64
		*Metadata
	}
	// LstT is a []string item type
	LstT struct {
		value []string
		*sync.Mutex
		*Metadata
	}
	// StrT is a string item type
	StrT struct {
		value atomic.String
		*Metadata
	}
	// UinT is an uint64 item type
	UinT struct {
		value atomic.Uint64
		*Metadata
	}
)

// Assert that all of the above types satisfy the types.Item interface
var _ = []Item{
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
func (b BooT) Meta() *Metadata         { return b.Metadata }

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
func (d DurT) Meta() *Metadata         { return d.Metadata }

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
func (f FltT) Meta() *Metadata         { return f.Metadata }

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
func (in IntT) Meta() *Metadata         { return in.Metadata }

// List of strings item implementation

func (l *LstT) FromString(s string) error {
	split := strings.Split(s, ",")
	for i := range split {
		split[i] = split[i]
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
		o += lo[i]
		if i != len(lo)-1 {
			o += ","
		}
	}
	return
}
func (l LstT) Uint() uint64    { panic("type error") }
func (l LstT) Meta() *Metadata { return l.Metadata }

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
func (s StrT) Meta() *Metadata         { return s.Metadata }

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
func (u UinT) Meta() *Metadata         { return u.Metadata }
