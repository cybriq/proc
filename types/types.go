package types

// Item represents a variable
type Item interface {
	Name() string
	Type() Type
	Aliases() []string
	Group() string
	Tags() []string
	Description() string
	Documentation() string
	Default() string
	FromString(string) error
	Concrete
}

func Name(T Type) string {
	return Names[T]
}
