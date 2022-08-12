package types

import (
	"strings"
	"sync"
)

// Metadata stores the information about the types.Item for documentation
// purposes
type Metadata struct {
	sync.Mutex
	name, group, description, documentation, def string
	typ                                          Type
	aliases                                      []string
}

// Desc is the named field form of Metadata for generating a Metadata
type Desc struct {
	Name, Group, Description, Documentation, Default string
	Type                                             Type
	Aliases                                          []string
}

// New allows you to create a Metadata with a sparsely filled, named field
// struct literal.
//
// name, type, group and tags all will be canonicalized to lower case.
func New(args Desc) *Metadata {
	// name, type and group should also be lower case
	return &Metadata{
		name:          strings.ToLower(args.Name),
		typ:           args.Type,
		aliases:       args.Aliases,
		group:         strings.ToLower(args.Group),
		description:   args.Description,
		documentation: args.Documentation,
		def:           args.Default,
	}
}

// Metadata accessors

func (m *Metadata) Aliases() []string {
	m.Lock()
	defer m.Unlock()
	return m.aliases
}
func (m *Metadata) Default() string {
	m.Lock()
	defer m.Unlock()
	return m.def
}

func (m *Metadata) Description() string {
	m.Lock()
	defer m.Unlock()
	return m.description
}
func (m *Metadata) Documentation() string {
	m.Lock()
	defer m.Unlock()
	return m.documentation
}
func (m *Metadata) Group() string {
	m.Lock()
	defer m.Unlock()
	return m.group
}
func (m *Metadata) Name() string {
	m.Lock()
	defer m.Unlock()
	return m.name
}
func (m *Metadata) Type() Type {
	m.Lock()
	defer m.Unlock()
	return m.typ
}
