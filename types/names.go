package types

//go:generate go run ./gen/main.go

import "time"

type Type int

const (
	Bool Type = iota
	Int
	Uint
	Duration
	Float
	String
	List
)

// Names provides the string associated with the Concrete type.
var Names = []string{
	"Bool",
	"Int",
	"Uint",
	"Duration",
	"Float",
	"String",
	"List",
}

// Concrete should return a value for the correct concrete type and panic
// otherwise, except for String which should always yield a value.
type Concrete interface {
	Bool() bool
	Int() int64
	Uint() uint64
	Duration() time.Duration
	Float() float64
	String() string
	List() []string
	FromString(string) error
}
