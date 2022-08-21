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
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.List
	o = &Opt{m: meta.New(m)}
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, ",")
	o.v.Store(split)
	return nil
}

func (o *Opt) String() (s string) {
	return strings.Join(o.v.Load().([]string), ",")
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.List = func() []string { return o.v.Load().([]string) }
	return
}
