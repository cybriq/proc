package opts

import (
	"time"

	"github.com/cybriq/proc/pkg/opts/meta"
)

// Concrete is a struct of functions that return the concrete values. Only the
// intended type will return a value, the rest always return zero.
//
// Note there is no meta.Text type because this can be had using the
// Option.String method below in all cases.
type Concrete struct {
	Bool     func() bool
	Duration func() time.Duration
	Float    func() float64
	Integer  func() int64
	List     func() []string
	Text     func() string
}

// NewConcrete provides a Concrete with all functions returning zero values
func NewConcrete() Concrete {
	return Concrete{
		func() bool { return false },
		func() time.Duration { return 0 },
		func() float64 { return 0 },
		func() int64 { return 0 },
		func() []string { return nil },
		func() string { return "" },
	}
}

// Option interface reads and writes string formats for options and returns a
// Concrete value to the appropriate concrete value, with the type indicated.
type Option interface {
	FromString(s string) (e error)
	String() (s string)
	Value() (c Concrete)
	Type() (t meta.Type)
	Meta() (md meta.Metadata)
}

type Config map[string]map[string]Option

// MarshalText produces a standard TOML format document via toml.Marshal that
// can be used to toml.UnmarshalTOML back into the Config.
func (c Config) MarshalText() (b []byte, e error) {
	return
}
