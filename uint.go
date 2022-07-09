package proc

import (
	"fmt"
	"strconv"

	"go.uber.org/atomic"
)

type Uint struct {
	value atomic.Uint64
	Meta
}

func (u *Uint) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	u.value.Store(uint64(i))
	return nil
}

func (u *Uint) Bool() bool     { panic("type error") }
func (u *Uint) Int() int64     { panic("type error") }
func (u *Uint) Uint() uint64   { return u.value.Load() }
func (u *Uint) Float() float64 { panic("type error") }
func (u *Uint) String() string { return fmt.Sprint(u.value.Load()) }
func (u *Uint) List() []string { panic("type error") }

func (u *Uint) Set(bo int) { u.value.Store(uint64(bo)) }
