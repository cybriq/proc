package proc

import (
	"fmt"
	"strconv"
	"time"

	"go.uber.org/atomic"
)

type Int struct {
	value atomic.Int64
	Meta
}

func (in *Int) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	in.value.Store(i)
	return nil
}

func (in *Int) Bool() bool              { panic("type error") }
func (in *Int) Int() int64              { return in.value.Load() }
func (in *Int) Duration() time.Duration { panic("type error") }
func (in *Int) Uint() uint64            { panic("type error") }
func (in *Int) Float() float64          { panic("type error") }
func (in *Int) String() string          { return fmt.Sprint(in.value.Load()) }
func (in *Int) List() []string          { panic("type error") }

func (in *Int) Set(i int) { in.value.Store(int64(i)) }
