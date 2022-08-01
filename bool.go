package proc

import (
	"errors"
	"fmt"
	"strings"
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
