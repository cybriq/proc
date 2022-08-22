package float

import (
	"strconv"
	"strings"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
	"go.uber.org/atomic"
)

type Opt struct {
	m meta.Metadata
	v atomic.Float64
	h []Hook
}

type Hook func(*Opt) error

func New(m meta.Data, h ...Hook) (o *Opt) {
	m.Type = meta.Float
	o = &Opt{m: meta.New(m), h: h}
	_ = o.FromString(m.Default)
	return
}

func (o *Opt) Meta() meta.Metadata   { return o.m }
func (o *Opt) Type() meta.Type       { return o.m.Typ }
func (o *Opt) ToOption() opts.Option { return o }

func (o *Opt) RunHooks() (e error) {
	for i := range o.h {
		e = o.h[i](o)
		if e != nil {
			return
		}
	}
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	var p float64
	p, e = strconv.ParseFloat(s, 64)
	if e != nil {
		return e
	}
	o.v.Store(p)
	e = o.RunHooks()
	return
}

func (o *Opt) String() (s string) {
	return strconv.FormatFloat(o.v.Load(), 'f', -1, 64)
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Float = func() float64 { return o.v.Load() }
	return
}
