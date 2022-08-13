package text

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/cybriq/proc/pkg/opts/meta"
	"github.com/cybriq/proc/pkg/opts/opt"
	"github.com/cybriq/proc/pkg/opts/sanitizers"
)

// Opt stores a string configuration value
type Opt struct {
	meta.Data
	hook  []Hook
	Value *atomic.Value
	Def   string
}

type Hook func(s []byte) error

// New creates a new Opt with a given default value set
func New(m meta.Data, def string, hook ...Hook) *Opt {
	v := &atomic.Value{}
	v.Store([]byte(def))
	return &Opt{Value: v, Data: m, Def: def, hook: hook}
}

// SetName sets the name for the generator
func (x *Opt) SetName(name string) {
	x.Data.Option = strings.ToLower(name)
	x.Data.Name = name
}

// Type returns the receiver wrapped in an interface for identifying its type
func (x *Opt) Type() interface{} {
	return x
}

// GetMetadata returns the metadata of the opt type
func (x *Opt) GetMetadata() *meta.Data {
	return &x.Data
}

// ReadInput sets the value from a string
func (x *Opt) ReadInput(input string) (o opt.Option, e error) {
	if input == "" {
		e = fmt.Errorf(
			"string opt %s %v may not be empty", x.Name(),
			x.Data.Aliases,
		)
		return
	}
	if strings.HasPrefix(input, "=") {
		// the following removes leading `=` and retains any following
		// instances of `=`
		input = strings.Join(strings.Split(input, "=")[1:], "=")
	}
	if x.Data.Options != nil {
		var matched string
		e = fmt.Errorf("option value not found '%s'", input)
		for _, i := range x.Data.Options {
			op := i
			if len(i) >= len(input) {
				op = i[:len(input)]
			}
			if input == op {
				if e == nil {
					return x, fmt.Errorf(
						"ambiguous short option value"+
							" '%s' matches multiple "+
							"options: %s, %s",
						input, matched, i,
					)
				}
				matched = i
				e = nil
			} else {
				continue
			}
		}
		if log.E.Chk(e) {
			return
		}
		input = matched
	} else {
		var cleaned string
		if cleaned, e = sanitizers.StringType(
			x.Data.Type, input,
			x.Data.DefaultPort,
		); log.E.Chk(e) {
			return
		}
		if cleaned != "" {
			log.I.Ln("setting value for", x.Data.Name, cleaned)
			input = cleaned
		}
	}
	e = x.Set(input)
	return x, e
}

// LoadInput sets the value from a string
func (x *Opt) LoadInput(input string) (o opt.Option, e error) {
	return x.ReadInput(input)
}

// Name returns the name of the opt
func (x *Opt) Name() string {
	return x.Data.Option
}

// AddHooks appends callback hooks to be run when the value is changed
func (x *Opt) AddHooks(hook ...Hook) {
	x.hook = append(x.hook, hook...)
}

// SetHooks sets a new slice of hooks
func (x *Opt) SetHooks(hook ...Hook) {
	x.hook = hook
}

// V returns the stored string
func (x *Opt) V() string {
	return string(x.Value.Load().([]byte))
}

// Empty returns true if the string is empty
func (x *Opt) Empty() bool {
	return len(x.Value.Load().([]byte)) == 0
}

// Bytes returns the raw bytes in the underlying storage note that this returns
// a copy because anything done to the slice affects all accesses afterwards,
// thus there is also a zero function
//
// todo: make an option for the byte buffer to be MMU fenced to prevent
//  elevated privilege processes from accessing this memory.
func (x *Opt) Bytes() []byte {
	byt := x.Value.Load().([]byte)
	o := make([]byte, len(byt))
	copy(o, byt)
	return o
}

// Zero the bytes
func (x *Opt) Zero() {
	byt := x.Value.Load().([]byte)
	for i := range byt {
		byt[i] = 0
	}
	x.Value.Store(byt)
}

func (x *Opt) runHooks(s []byte) (e error) {
	for i := range x.hook {
		if e = x.hook[i](s); log.E.Chk(e) {
			break
		}
	}
	return
}

// Set the value stored
func (x *Opt) Set(s string) (e error) {
	if e = x.runHooks([]byte(s)); !log.E.Chk(e) {
		x.Value.Store([]byte(s))
	}
	return
}

// SetBytes sets the string from bytes
func (x *Opt) SetBytes(s []byte) (e error) {
	if e = x.runHooks(s); !log.E.Chk(e) {
		x.Value.Store(s)
	}
	return
}

// Opt returns a string representation of the value
func (x *Opt) String() string {
	return fmt.Sprintf("%s: '%s'", x.Data.Option, x.V())
}

// MarshalJSON returns the json representation
func (x *Opt) MarshalJSON() (b []byte, e error) {
	v := string(x.Value.Load().([]byte))
	return json.Marshal(&v)
}

// UnmarshalJSON decodes a JSON representation
func (x *Opt) UnmarshalJSON(data []byte) (e error) {
	v := x.Value.Load().([]byte)
	e = json.Unmarshal(data, &v)
	x.Value.Store(v)
	return
}
