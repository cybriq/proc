package types

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
