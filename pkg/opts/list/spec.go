package list

import (
	"strings"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
	"go.uber.org/atomic"
)

type Opt struct {
	m meta.Metadata
	v atomic.Value
	h []Hook
}

type Hook func(*Opt) error

func New(m meta.Data, h ...Hook) (o *Opt) {
	m.Type = meta.List
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
	split := strings.Split(s, ",")
	o.v.Store(split)
	e = o.RunHooks()
	return
}

func (o *Opt) String() (s string) {
	return strings.Join(o.v.Load().([]string), ",")
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.List = func() []string { return o.v.Load().([]string) }
	return
}
