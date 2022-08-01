package proc

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cybriq/proc/types"
	"go.uber.org/atomic"
)

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
