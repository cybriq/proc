package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cybriq/proc/pkg/types"
)

func (c *Configs) Save(filename string) (err error) {
	var b []byte
	c.persist.Lock()
	defer c.persist.Unlock()
	b, err = c.MarshalJSON()
	if log.E.Chk(err) {
		return
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if log.E.Chk(err) {
		return
	}
	return
}

func (c *Configs) Load(filename string) (err error) {
	var b []byte
	c.persist.Lock()
	defer c.persist.Unlock()
	b, err = ioutil.ReadFile(filename)
	if log.E.Chk(err) {
		return
	}
	err = c.UnmarshalJSON(b)
	if log.E.Chk(err) {
		return
	}
	return
}

// MarshalJSON returns the JSON for the current state of a Configs
func (c *Configs) MarshalJSON() ([]byte, error) {
	out := make(map[string]map[string]interface{}, len(c.items))
	c.Lock()
	for i := range c.items {
		for j := range c.items[i] {
			if _, ok := out[i]; !ok {
				out[i] = make(map[string]interface{})
			}
			switch c.items[i][j].Type() {
			case types.Bool:
				out[i][j] = c.items[i][j].Bool()

			case types.Int:
				out[i][j] = c.items[i][j].Int()

			case types.Uint:
				out[i][j] = c.items[i][j].Uint()

			case types.Duration:
				out[i][j] = c.items[i][j].String()

			case types.Float:
				out[i][j] = c.items[i][j].Float()

			case types.String:
				out[i][j] = c.items[i][j].String()

			case types.List:
				out[i][j] = c.items[i][j].List()
			}
		}
	}
	c.Unlock()
	b, err := json.MarshalIndent(out, "", "\t")
	return b, err
}

// UnmarshalJSON loads a Configs with the values in a JSON string
func (c *Configs) UnmarshalJSON(bytes []byte) error {
	in := make(map[string]map[string]interface{})
	err := json.Unmarshal(bytes, &in)
	for i := range in {
		for j := range in[i] {
			v, ok := c.items[i][j]
			if !ok {
				return fmt.Errorf(
					"configuration does not contain item with"+
						" name %s", i)
			}
			switch v.Type() {
			case types.Bool:
				t := c.items[i][j].(*types.BooT)
				ta := &t
				(*ta).Set(v.Bool())

			case types.Int:
				t := c.items[i][j].(*types.IntT)
				ta := &t
				(*ta).Set(v.Int())

			case types.Uint:
				t := c.items[i][j].(*types.UinT)
				ta := &t
				(*ta).Set(v.Uint())

			case types.Duration:
				t := c.items[i][j].(*types.DurT)
				ta := &t
				(*ta).Set(v.Duration())

			case types.Float:
				t := c.items[i][j].(*types.FltT)
				ta := &t
				(*ta).Set(v.Float())

			case types.String:
				t := c.items[i][j].(*types.StrT)
				ta := &t
				(*ta).Set(v.String())

			case types.List:
				t := c.items[i][j].(*types.LstT)
				ta := &t
				(*ta).Set(v.List()...)
			}
		}
	}
	return err
}
