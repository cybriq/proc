package config

import (
	"fmt"
	"sort"

	"github.com/cybriq/proc/pkg/types"
)

// Get returns a named item from the Configs
func (c *Configs) Get(group, name string) (t types.Item, err error) {
	var ok bool
	c.Lock()
	t, ok = c.items[group][name]
	c.Unlock()
	if !ok {
		err = fmt.Errorf("item '%s.%s' not found", group, name)
	}
	return
}

type Types []types.Item

func (t Types) Len() int           { return len(t) }
func (t Types) Less(i, j int) bool { return t[i].Name() < t[j].Name() }
func (t Types) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type TypeGroup struct {
	Name string
	Types
}

type TypesGrouped []TypeGroup

func (t TypesGrouped) Len() int           { return len(t) }
func (t TypesGrouped) Less(i, j int) bool { return t[i].Name < t[i].Name }
func (t TypesGrouped) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (c *Configs) GetByName(name string) (t Types) {
	c.Lock()
	for i := range c.items {
		for j := range c.items[i] {
			if j == name {
				t = append(t, c.items[i][j])
			}
		}
	}
	c.Unlock()
	sort.Sort(t)
	return
}

func (c *Configs) GetGroup(group string) (t TypeGroup) {
	c.Lock()
	if _, ok := c.items[group]; ok {
		t.Name = group
		for i := range c.items[group] {
			t.Types = append(t.Types, c.items[group][i])
		}
	}
	c.Unlock()
	sort.Sort(t.Types)
	return
}

func (c *Configs) GetGroups() (tg TypesGrouped) {
	var g []string
	for i := range c.items {
		g = append(g, i)
	}
	sort.Strings(g)
	for i := range g {
		tg = append(tg, c.GetGroup(g[i]))
	}
	return
}
