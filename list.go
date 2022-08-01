package proc

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/cybriq/proc/types"
)

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
