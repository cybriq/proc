package cmds

// Command is a specification for a command and can include any number of
// subcommands
type Command struct {
	Name        string
	Title       string
	Description string
	Entrypoint  func(c interface{}) error
	Commands    Commands
	Colorizer   func(a ...interface{}) string
	AppText     string
	Parent      *Command
}

// Commands are a slice of Command entries
type Commands []Command

func (c Commands) PopulateParents(parent *Command) {
	if parent != nil {
		log.T.Ln("backlinking children of", parent.Name)
	}
	for i := range c {
		c[i].Parent = parent
		c[i].Commands.PopulateParents(&c[i])
	}
}

const tabs = "\t\t\t\t\t"

// Foreach runs a closure on every node in the Commands tree, stopping if the
// closure returns false
func (c Commands) Foreach(cl func(*Command) bool) {
	c.foreach(cl, 0, 0, nil)
}

func (c Commands) foreach(cl func(*Command) bool, hereDepth, hereDist int,
	cmd *Command) (ocl func(*Command) bool, depth, dist int, cm *Command) {

	ocl = cl
	cm = cmd
	if c == nil {
		dist = hereDist
		depth = hereDepth
		return
	}
	depth = hereDepth + 1
	log.T.Ln(tabs[:depth]+"->", depth)
	dist = hereDist
	for i := range c {
		log.T.Ln(tabs[:depth]+"walking", c[i].Name, depth, dist)
		if cl(&c[i]) {
			return
		}
		dist++
		ocl, depth, dist, cm = c[i].Commands.foreach(
			cl,
			depth,
			dist,
			cm,
		)
	}
	log.T.Ln(tabs[:hereDepth]+"<-", hereDepth)
	depth--
	return
}
