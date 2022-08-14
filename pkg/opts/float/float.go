package float

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/cybriq/proc/pkg/opts/meta"
	"github.com/cybriq/proc/pkg/opts/opt"
	"github.com/cybriq/proc/pkg/opts/sanitizers"

	"go.uber.org/atomic"
)

// Opt stores an float64 configuration value
type Opt struct {
	meta.Data
	hook     []Hook
	Min, Max float64
	clamp    func(input float64) (result float64)
	value    *atomic.Float64
	Def      float64
}

type Hook func(f float64) error

// New returns a new Opt value set to a default value
func New(m meta.Data, def float64, min, max float64, hook ...Hook) *Opt {
	m.Type = fmt.Sprint(reflect.TypeOf(def))
	return &Opt{
		value: atomic.NewFloat64(def),
		Data:  m,
		Def:   def,
		Min:   min,
		Max:   max,
		hook:  hook,
		clamp: sanitizers.ClampFloat(min, max),
	}
}

// SetName sets the name for the generator
func (x *Opt) SetName(name string) {
	x.Data.Option = strings.ToLower(name)
	x.Data.Name = name
}

// Type returns the receiver wrapped in an interface for identifying its type
func (x *Opt) Type() reflect.Type {
	return reflect.TypeOf(x.value.Load())
}

func (x *Opt) Value() interface{} {
	return x.value.Load()
}

// GetMetadata returns the metadata of the opt type
func (x *Opt) GetMetadata() *meta.Data {
	return &x.Data
}

// ReadInput sets the value from a string
func (x *Opt) ReadInput(input string) (o opt.Option, e error) {
	if input == "" {
		e = fmt.Errorf(
			"floating point number opt %s %v may not be empty",
			x.Name(), x.Data.Aliases,
		)
		return
	}
	if strings.HasPrefix(input, "=") {
		// the following removes leading and trailing '='
		input = strings.Join(strings.Split(input, "=")[1:], "=")
	}
	var v float64
	if v, e = strconv.ParseFloat(input, 64); log.E.Chk(e) {
		return
	}
	if e = x.Set(v); log.E.Chk(e) {
	}
	return x, e
}

// LoadInput sets the value from a string (this is the same as the above but
// differs for Strings)
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

// V returns the value stored
func (x *Opt) V() float64 {
	return x.value.Load()
}

func (x *Opt) runHooks(f float64) (e error) {
	for i := range x.hook {
		if e = x.hook[i](f); log.E.Chk(e) {
			break
		}
	}
	return
}

// Set the value stored
func (x *Opt) Set(f float64) (e error) {
	f = x.clamp(f)
	if e = x.runHooks(f); !log.E.Chk(e) {
		x.value.Store(f)
	}
	return
}

// String returns a string representation of the value
func (x *Opt) String() string {
	return fmt.Sprintf("%s: %0.8f", x.Data.Option, x.V())
}

// MarshalJSON returns the json representation of
func (x *Opt) MarshalJSON() (b []byte, e error) {
	v := x.value.Load()
	return json.Marshal(&v)
}

// UnmarshalJSON decodes a JSON representation of
func (x *Opt) UnmarshalJSON(data []byte) (e error) {
	v := x.value.Load()
	e = json.Unmarshal(data, &v)
	e = x.Set(v)
	return
}
