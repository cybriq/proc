package text

import (
	"strings"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
	"go.uber.org/atomic"
)

type Opt struct {
	m meta.Metadata
	v atomic.String
	h []Hook
}

type Hook func(*Opt) error

func New(m meta.Data, h ...Hook) (o *Opt) {
	m.Type = meta.Text
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
	o.v.Store(s)
	e = o.RunHooks()
	return
}

func (o *Opt) String() (s string) {
	return o.v.Load()
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Text = func() string { return o.v.Load() }
	return
}
