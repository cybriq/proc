package toggle

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/atomic"

	"github.com/cybriq/proc/pkg/opts"
	"github.com/cybriq/proc/pkg/opts/meta"
)

type Opt struct {
	m meta.Metadata
	v atomic.Bool
}

func New(m meta.Data) (o *Opt) {
	m.Type = meta.Bool
	o = &Opt{m: meta.New(m)}
	return
}

func (o *Opt) FromString(s string) (e error) {
	s = strings.TrimSpace(s)
	switch s {
	case "f", "false", "off", "-":
		o.v.Store(false)
	case "t", "true", "on", "+":
		o.v.Store(true)
	default:
		return fmt.Errorf("string '%s' does not parse to boolean", s)
	}
	return nil
}

func (o *Opt) String() (s string) {
	return strconv.FormatBool(o.v.Load())
}

func (o *Opt) Value() (c opts.Concrete) {
	c = opts.NewConcrete()
	c.Bool = func() bool { return o.v.Load() }
	return
}
