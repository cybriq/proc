package path

import (
	"strings"
)

type Path []string

func (p Path) TrimPrefix() Path {
	if len(p) > 1 {
		return p[1:]
	}
	return p[:0]
}

func (p Path) String() string {
	return strings.Join(p, " ")
}

func (p Path) Parent() (p1 Path) {
	if len(p) > 0 {
		p1 = p[:len(p)-1]
	}
	return
}

func (p Path) Child(child string) (p1 Path) {
	p1 = append(p, child)
	// log.I.Ln(p, p1)
	return
}

func (p Path) Common(p2 Path) (o Path) {
	for i := range p {
		if len(p2) < i {
			if p[i] == p2[i] {
				o = append(o, p[i])
			}
		}
	}
	return
}

func GetIndent(d int) string {
	return strings.Repeat("\t", d)
}
