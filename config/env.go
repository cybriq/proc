package config

import (
	"fmt"
	"os"
	"strings"
)

func (c *Configs) LoadAllFromEnv() {
	for i := range c.items {
		for j := range c.items[i] {
			_ = c.LoadFromEnv(i, j)
		}
	}
}

func (c *Configs) LoadFromEnv(group, name string) (err error) {
	var envName string
	envName, err = c.GetEnvString(group, name)
	if log.E.Chk(err) {
		return
	}
	val := os.Getenv(envName)
	if val != "" {
		err = c.items[group][name].FromString(val)
	}
	return
}

func (c *Configs) GetEnvString(group, name string) (envName string, err error) {
	var ok bool
	_, ok = c.items[group]
	if !ok {
		return "", fmt.Errorf("no group '%s' known", group)
	}
	_, ok = c.items[group][name]
	if !ok {
		return "", fmt.Errorf("no item '%s' known in group '%s'",
			name, group)
	}
	envName = ComposeEnvName(c.appName, group, name)
	return
}

func ComposeEnvName(appName, group, name string) string {
	return fmt.Sprintf("%s_%s_%s",
		strings.ToUpper(appName),
		strings.ToUpper(group),
		strings.ToUpper(name),
	)
}
