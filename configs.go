package proc

import (
	"errors"
	"sync"

	"gitlab.com/cybriqsystems/proc/types"
)

type Configs struct {
	items map[string]types.Item
	sync.Mutex
}

func Create(items ...Desc) (c Configs) {
	c = Configs{items: make(map[string]types.Item)}
	c.Lock()
	defer c.Unlock()
	for i := range items {
		name := items[i].Name
		if _, ok := c.items[name]; ok {
			panic("configs contains a duplicate named item: '" + name + "'")
		}
		c.items[name] = Item(New(items[i]))
	}
	return
}

func (c *Configs) Get(name string) (t types.Item, err error) {
	var ok bool
	c.Lock()
	t, ok = c.items[name]
	c.Unlock()
	if !ok {
		err = errors.New("type '" + name + "' not found")
	}
	return
}
