package proc

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/cybriq/proc/types"
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
		err = fmt.Errorf("type '%s' not found", name)
	}
	return
}

func (c *Configs) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{}, len(c.items))
	for i := range c.items {
		switch c.items[i].Type() {
		case types.Bool:
			out[i] = c.items[i].Bool()

		case types.Int:
			out[i] = c.items[i].Int()

		case types.Uint:
			out[i] = c.items[i].Uint()

		case types.Duration:
			out[i] = c.items[i].Duration()

		case types.Float:
			out[i] = c.items[i].Float()

		case types.String:
			out[i] = c.items[i].String()

		case types.List:
			out[i] = c.items[i].List()
		}
	}
	b, err := json.MarshalIndent(out, "", "\t")
	return b, err
}

func (c *Configs) UnmarshalJSON(bytes []byte) error {
	in := make(map[string]interface{})
	err := json.Unmarshal(bytes, &in)
	for i := range in {
		v, ok := c.items[i]
		if !ok {
			return fmt.Errorf(
				"configuration does not contain item with"+
					" name %s", i)
		}
		switch v.Type() {
		case types.Bool:
			t := c.items[i].(*_bool)
			ta := &t
			(*ta).Set(v.Bool())

		case types.Int:
			t := c.items[i].(*_int)
			ta := &t
			(*ta).Set(v.Int())

		case types.Uint:
			t := c.items[i].(*_uin)
			ta := &t
			(*ta).Set(v.Uint())

		case types.Duration:
			t := c.items[i].(*_dur)
			ta := &t
			(*ta).Set(v.Duration())

		case types.Float:
			t := c.items[i].(*_flt)
			ta := &t
			(*ta).Set(v.Float())

		case types.String:
			t := c.items[i].(*_str)
			ta := &t
			(*ta).Set(v.String())

		case types.List:
			t := c.items[i].(*_lst)
			ta := &t
			(*ta).Set(v.List()...)
		}
	}
	return err
}
