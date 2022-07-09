package proc

import (
	"fmt"
	"time"

	"go.uber.org/atomic"
)

type Duration struct {
	value atomic.Duration
	Meta
}

func (d *Duration) FromString(s string) error {
	i, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.value.Store(i)
	return nil
}

func (d *Duration) Bool() bool              { panic("type error") }
func (d *Duration) Int() int64              { return int64(d.value.Load()) }
func (d *Duration) Duration() time.Duration { return d.value.Load() }
func (d *Duration) Uint() uint64            { panic("type error") }
func (d *Duration) Float() float64          { panic("type error") }
func (d *Duration) String() string          { return fmt.Sprint(d.value.Load()) }
func (d *Duration) List() []string          { panic("type error") }

func (d *Duration) Set(du time.Duration) { d.value.Store(du) }
