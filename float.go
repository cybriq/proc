package proc

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.com/cybriqsystems/proc/types"
	"go.uber.org/atomic"
)

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
