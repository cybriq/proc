package proc

import (
	"strings"
	"sync"
	"time"
)

type _lst struct {
	value []string
	sync.Mutex
	*metadata
}

func NewList(m *metadata) (b *_lst) {
	b = &_lst{}
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
	l.Set(split)
	return nil
}
func (l *_lst) Bool() bool              { panic("type error") }
func (l *_lst) Int() int64              { panic("type error") }
func (l *_lst) Duration() time.Duration { panic("type error") }
func (l *_lst) Uint() uint64            { panic("type error") }
func (l *_lst) Float() float64          { panic("type error") }

func (l *_lst) String() (o string) {
	o = "["
	lo := l.List()
	for i := range lo {
		o += "\"" + lo[i] + "\","
	}
	o += "]"
	return
}

func (l *_lst) List() (li []string) {
	l.Mutex.Lock()
	li = make([]string, len(l.value))
	copy(li, l.value)
	l.Mutex.Unlock()
	return
}

func (l *_lst) Set(li []string) {
	l.Mutex.Lock()
	l.value = make([]string, len(li))
	copy(l.value, li)
	l.Mutex.Unlock()
}
