package duration

import (
	"fmt"
	"strings"
	"time"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
	"go.uber.org/atomic"
)

type Opt struct {
	m meta.Metadata
	v atomic.Duration
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.Duration
	o = &Opt{m: meta.New(m)}
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	var d time.Duration
	d, e = time.ParseDuration(s)
	if e != nil {
		return e
	}
	o.v.Store(d)
	return nil
}

func (o *Opt) String() (s string) {
	return fmt.Sprint(o.v.Load())
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Duration = func() time.Duration { return o.v.Load() }
	return
}
