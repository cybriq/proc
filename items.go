package proc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cybriq/proc/types"
	"go.uber.org/atomic"
)

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
