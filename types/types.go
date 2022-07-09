package types

const (
	Bool int = iota
	Int
	Uint
	Float
	String
	List
)

// Names provides the string associated with the Concrete type.
var Names = []string{"Bool", "Int", "Uint", "Float", "String", "List"}

// Concrete should return a value for the correct concrete type and panic
// otherwise, except for String which should always yield a value
type Concrete interface {
	Bool() bool
	Int() int64
	Uint() uint64
	Float() float64
	String() string
	List() []string
}

// Type represents a variable
type Type interface {
	Name() string
	Type() string
	Aliases() []string
	Group() string
	Tags() []string
	Description() string
	Documentation() string
	Default() string
	FromString(string) error
	Concrete
}
