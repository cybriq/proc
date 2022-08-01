package proc

import (
	"strings"
	"sync"

	"github.com/cybriq/proc/types"
)

// metadata automatically implements everything except the inputs and outputs
type metadata struct {
	sync.Mutex
	name, group, description, documentation, def string
	typ                                          types.Type
	tags, aliases                                []string
}

// Desc is the named field form of metadata for generating a metadata
type Desc struct {
	Name, Group, Description, Documentation, Default string
	Type                                             types.Type
	Tags, Aliases                                    []string
}

func isType(s string) (is bool) {
	for i := range types.Names {
		if s == types.Names[i] {
			is = true
		}
	}
	return
}

// New allows you to create a metadata with a sparsely filled, named field
// struct literal.
//
// name, type, group and tags all will be canonicalized to lower case.
func New(args Desc) *metadata {
	// tags should be all lower case
	for i := range args.Tags {
		args.Tags[i] = strings.ToLower(args.Tags[i])
	}
	// name, type and group should also be lower case
	return &metadata{
		name:          strings.ToLower(args.Name),
		typ:           args.Type,
		aliases:       args.Aliases,
		group:         strings.ToLower(args.Group),
		tags:          args.Tags,
		description:   args.Description,
		documentation: args.Documentation,
		def:           args.Default,
	}
}

func (m *metadata) Name() string {
	m.Lock()
	defer m.Unlock()
	return m.name
}
func (m *metadata) Type() types.Type {
	m.Lock()
	defer m.Unlock()
	return m.typ
}
func (m *metadata) Aliases() []string {
	m.Lock()
	defer m.Unlock()
	return m.aliases
}
func (m *metadata) Group() string {
	m.Lock()
	defer m.Unlock()
	return m.group
}
func (m *metadata) Tags() []string {
	m.Lock()
	defer m.Unlock()
	return m.tags
}
func (m *metadata) Description() string {
	m.Lock()
	defer m.Unlock()
	return m.description
}
func (m *metadata) Documentation() string {
	m.Lock()
	defer m.Unlock()
	return m.documentation
}
func (m *metadata) Default() string {
	m.Lock()
	defer m.Unlock()
	return m.def
}
