package proc

import (
	"strings"
	"sync"
	"time"
)

type List struct {
	value []string
	sync.Mutex
	Meta
}

// FromString converts a comma separated list of strings into a List
func (l *List) FromString(s string) error {
	split := strings.Split(s, ",")
	l.Set(split)
	return nil
}
func (l *List) Bool() bool              { panic("type error") }
func (l *List) Int() int64              { panic("type error") }
func (l *List) Duration() time.Duration { panic("type error") }
func (l *List) Uint() uint64            { panic("type error") }
func (l *List) Float() float64          { panic("type error") }

func (l *List) String() (o string) {
	o = "["
	lo := l.List()
	for i := range lo {
		o += "\"" + lo[i] + "\","
	}
	o += "]"
	return
}

func (l *List) List() (li []string) {
	l.Mutex.Lock()
	li = make([]string, len(l.value))
	copy(li, l.value)
	l.Mutex.Unlock()
	return
}

func (l *List) Set(li []string) {
	l.Mutex.Lock()
	l.value = make([]string, len(li))
	copy(l.value, li)
	l.Mutex.Unlock()
}
