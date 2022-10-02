package path

import (
	"strings"
)

type Op func(c interface{}) error

type Path []string

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

func GetIndent(d int) string {
	return strings.Repeat("\t", d)
}
