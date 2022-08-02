package types

// Item provides accessors for a Config item type's metadata and current
// contents, including a generic string format setter.
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

// Name is a helper that returns the name associated with a Type.
func Name(T Type) string {
	return Names[T]
}
