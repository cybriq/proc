package proc

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/atomic"
)

type Bool struct {
	value atomic.Bool
	Meta
}

func (b *Bool) FromString(s string) error {
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

func (b *Bool) Bool() bool     { return b.value.Load() }
func (b *Bool) Int() int64     { panic("type error") }
func (b *Bool) Uint() uint64   { panic("type error") }
func (b *Bool) Float() float64 { panic("type error") }
func (b *Bool) String() string { return fmt.Sprint(b.value.Load()) }
func (b *Bool) List() []string { panic("type error") }

func (b *Bool) Set(bo bool) { b.value.Store(bo) }
