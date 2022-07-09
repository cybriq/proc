package proc

import (
	"fmt"
	"strconv"
	"time"

	"go.uber.org/atomic"
)

type _uin struct {
	value atomic.Uint64
	*metadata
}

func NewUint(m *metadata) (b *_uin) {
	b = &_uin{}
	err := b.FromString(m.Default())
	if err != nil {
		panic(err)
	}
	b.metadata = m
	return
}

func (u *_uin) FromString(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	u.value.Store(uint64(i))
	return nil
}

func (u *_uin) Bool() bool              { panic("type error") }
func (u *_uin) Int() int64              { panic("type error") }
func (u *_uin) Duration() time.Duration { panic("type error") }
func (u *_uin) Uint() uint64            { return u.value.Load() }
func (u *_uin) Float() float64          { panic("type error") }
func (u *_uin) String() string          { return fmt.Sprint(u.value.Load()) }
func (u *_uin) List() []string          { panic("type error") }

func (u *_uin) Set(ui int) { u.value.Store(uint64(ui)) }
