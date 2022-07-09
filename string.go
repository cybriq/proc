package proc

import (
	"go.uber.org/atomic"
)

type String struct {
	value atomic.String
	Meta
}

func (s *String) FromString(st string) error {
	s.value.Store(st)
	return nil
}

func (s *String) Bool() bool     { panic("type error") }
func (s *String) Int() int64     { panic("type error") }
func (s *String) Uint() uint64   { panic("type error") }
func (s *String) Float() float64 { panic("type error") }
func (s *String) String() string { return s.value.Load() }
func (s *String) List() []string { panic("type error") }

func (s *String) Set(st string) { s.value.Store(st) }
