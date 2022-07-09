package proc

// Meta automatically implements everything except the inputs and outputs
type Meta struct {
	name, typ, group, description, documentation, def string
	tags, aliases                                     []string
}

type Struct struct {
	Name, Type, Group, Description, Documentation, Default string
	Tags, Aliases                                          []string
}

// NewMeta allows you to create a Meta with a sparsely filled, named field
// struct literal
func NewMeta(args Struct) Meta {
	return Meta{
		name:          args.Name,
		typ:           args.Type,
		aliases:       args.Aliases,
		group:         args.Group,
		tags:          args.Tags,
		description:   args.Description,
		documentation: args.Documentation,
		def:           args.Default,
	}
}

func (m *Meta) Name() string          { return m.name }
func (m *Meta) Type() string          { return m.typ }
func (m *Meta) Aliases() []string     { return m.aliases }
func (m *Meta) Group() string         { return m.group }
func (m *Meta) Tags() []string        { return m.tags }
func (m *Meta) Description() string   { return m.description }
func (m *Meta) Documentation() string { return m.documentation }
func (m *Meta) Default() string       { return m.def }
