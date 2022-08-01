package proc

import (
	"time"

	"gitlab.com/cybriqsystems/proc/types"
	"go.uber.org/atomic"
)

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
