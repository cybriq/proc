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
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.Text
	o = &Opt{m: meta.New(m)}
	_ = o.FromString(m.Default)
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	o.v.Store(s)
	return nil
}

func (o *Opt) String() (s string) {
	return o.v.Load()
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Text = func() string { return o.v.Load() }
	return
}
