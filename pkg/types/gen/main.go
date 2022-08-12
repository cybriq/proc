package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/cybriq/proc"
	log2 "github.com/cybriq/proc/log"
)

// Pair is a set of type names and their actual implementation type
type Pair struct{ Name, Type string }

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

var tmpl = `package types
{{$t := .}}
// This file is generated: DO NOT EDIT.
` + `//go:generate go run ./gen/main.go

import "time"

// Type is an identifier code for a type of configuration item.
type Type int

// The list of types.Item supported by proc
const (
{{range $i, $v := $t}}	{{$v.Name}}{{if $i}}{{else}} Type = iota{{end}}
{{end}})

// Names provides the string associated with the Concrete type.
var Names = []string{
{{range $t}}	"{{.Name}}",
{{end}}}

// Concrete should return a value for the correct concrete type and panic
// otherwise, except for String which should always yield a value.
type Concrete interface {
{{range $t}}	{{.Name}}() {{.Type}}
{{end}}}
`

var log = log2.GetLogger(proc.PathBase)

func main() {
	tm, _ := template.New("help").Parse(tmpl)
	var buf []byte
	b := bytes.NewBuffer(buf)
	err := tm.Execute(b, Types)
	if log.E.Chk(err) {
		panic(err)
	}
	log.I.Ln("\n" + b.String())
	err = ioutil.WriteFile("names.go", b.Bytes(), 0600)
	if log.E.Chk(err) {
		panic(err)
	}
}
