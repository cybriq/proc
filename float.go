package proc

import (
	"fmt"
	"strconv"

	"go.uber.org/atomic"
)

type Float struct {
	value atomic.Float64
	Meta
}

func (f *Float) FromString(s string) error {
	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.value.Store(fl)
	return nil
}

func (f *Float) Bool() bool     { panic("type error") }
func (f *Float) Int() int64     { panic("type error") }
func (f *Float) Uint() uint64   { panic("type error") }
func (f *Float) Float() float64 { return f.value.Load() }
func (f *Float) String() string { return fmt.Sprint(f.value.Load()) }
func (f *Float) List() []string { panic("type error") }

func (f *Float) Set(bo float64) { f.value.Store(bo) }
