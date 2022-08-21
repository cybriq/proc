package integer

import (
	"strconv"
	"strings"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
	"go.uber.org/atomic"
)

type Opt struct {
	m meta.Metadata
	v atomic.Int64
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.Integer
	o = &Opt{m: meta.New(m)}
	_ = o.FromString(m.Default)
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	var p int64
	p, e = strconv.ParseInt(s, 10, 64)
	if e != nil {
		return e
	}
	o.v.Store(p)
	return nil
}

func (o *Opt) String() (s string) {
	return strconv.FormatInt(o.v.Load(), 10)
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Integer = func() int64 { return o.v.Load() }
	return
}
