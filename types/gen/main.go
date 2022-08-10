package main

import (
	"fmt"
	"io/ioutil"
)

// Pair is a set of type names and their actual implementation type
type Pair struct{ name, typ string }

// Types is the list of names and concrete primitive types they represent
var Types = []Pair{
	{
		"Bool", "bool",
	},
	{
		"Int", "int64",
	},
	{
		"Uint", "uint64",
	},
	{
		"Duration", "time.Duration",
	},
	{
		"Float", "float64",
	},
	{
		"String", "string",
	},
	{
		"List", "[]string",
	},
}

func main() {

	o := `package types

// This file is generated: DO NOT EDIT.
` + `//go:generate go run ./gen/main.go

import "time"

// Type is an identifier code for a type of configuration item.
type Type int

// The list of types.Item supported by proc
const (
`
	for i := range Types {
		if i == 0 {
			o += fmt.Sprintf("\t%s Type = iota\n", Types[i].name)
		} else {
			o += fmt.Sprintf("\t%s\n", Types[i].name)
		}
	}
	o += `)

// Names provides the string associated with the Concrete type.
var Names = []string{
`
	for i := range Types {
		o += fmt.Sprintf("\t\"%s\",\n", Types[i].name)
	}
	o += `}

// Concrete should return a value for the correct concrete type and panic
// otherwise, except for String which should always yield a value.
type Concrete interface {
`
	for i := range Types {
		o += fmt.Sprintf("\t%s() %s\n", Types[i].name, Types[i].typ)
	}

	o += `}
`
	err := ioutil.WriteFile("names.go", []byte(o), 0600)
	if err != nil {
		panic(err)
	}
}
