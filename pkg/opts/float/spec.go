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
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.Float
	o = &Opt{m: meta.New(m)}
	_ = o.FromString(m.Default)
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
	return nil
}

func (o *Opt) String() (s string) {
	return strconv.FormatFloat(o.v.Load(), 'f', -1, 64)
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Float = func() float64 { return o.v.Load() }
	return
}
