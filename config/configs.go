package config

import (
	"fmt"
	"sync"

	"github.com/cybriq/proc/types"
)

type Configs struct {
	appName string
	items   map[string]map[string]types.Item
	sync.Mutex
	persist sync.Mutex
}

// Create a new configuration from a slice of item descriptors.
func Create(appName string, items ...types.Desc) (c Configs) {
	c = Configs{appName: appName,
		items: make(map[string]map[string]types.Item)}
	c.Lock()
	defer c.Unlock()
	for i := range items {
		name := items[i].Name
		if _, ok := c.items[items[i].Group]; !ok {
			c.items[items[i].Group] = make(map[string]types.Item)
		}
		if _, ok := c.items[items[i].Group][name]; ok {
			panic("configs contains a duplicate named item: '" +
				name + "' in group '" + items[i].Group + "'")
		}
		c.items[items[i].Group][name] = Item(types.New(items[i]))
	}
	return
}

// Item takes a Metadata and creates the appropriate item type for it.
func Item(m *types.Metadata) (t types.Item) {
	switch m.Type() {
	case types.Bool:
		t = NewBool(m)
	case types.Duration:
		t = NewDuration(m)
	case types.Float:
		t = NewFloat(m)
	case types.Int:
		t = NewInt(m)
	case types.List:
		t = NewList(m)
	case types.String:
		t = NewString(m)
	case types.Uint:
		t = NewUint(m)
	default:
		panic("invalid type: '" + fmt.Sprint(m.Type()) + "'")
	}
	return
}

// List is a helper that uses variadic parameters to return a slice of strings.
func List(items ...string) []string {
	return items
}

// IfErrNotNilPanic is a helper for functions that should never error as
// errors are purely programmer errors and not error conditions.
func IfErrNotNilPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// NewBool creates a new boolean types.Item
func NewBool(m *types.Metadata) (b *types.BooT) {
	b = &types.BooT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewDuration creates a new duration types.Item
func NewDuration(m *types.Metadata) (b *types.DurT) {
	b = &types.DurT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewFloat creates a new floating point types.Item
func NewFloat(m *types.Metadata) (b *types.FltT) {
	b = &types.FltT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewInt creates a new integer types.Item
func NewInt(m *types.Metadata) (b *types.IntT) {
	b = &types.IntT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewList creates a new list of strings types.Item
func NewList(m *types.Metadata) (b *types.LstT) {
	b = &types.LstT{Mutex: &sync.Mutex{}}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewString creates a new string types.Item
func NewString(m *types.Metadata) (b *types.StrT) {
	b = &types.StrT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}

// NewUint creates a new unsigned integer types.Item
func NewUint(m *types.Metadata) (b *types.UinT) {
	b = &types.UinT{}
	err := b.FromString(m.Default())
	IfErrNotNilPanic(err)
	b.Metadata = m
	return
}
