package proc

import (
	"fmt"
	"time"

	"go.uber.org/atomic"
)

type _dur struct {
	value atomic.Duration
	*metadata
}

func NewDuration(m *metadata) (b *_dur) {
	b = &_dur{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (d *_dur) FromString(s string) error {
	i, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.value.Store(i)
	return nil
}

func (d *_dur) Bool() bool              { panic("type error") }
func (d *_dur) Int() int64              { return int64(d.value.Load()) }
func (d *_dur) Duration() time.Duration { return d.value.Load() }
func (d *_dur) Uint() uint64            { panic("type error") }
func (d *_dur) Float() float64          { panic("type error") }
func (d *_dur) String() string          { return fmt.Sprint(d.value.Load()) }
func (d *_dur) List() []string          { panic("type error") }

func (d *_dur) Set(du time.Duration) { d.value.Store(du) }
