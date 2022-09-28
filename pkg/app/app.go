package app

import "github.com/cybriq/proc/pkg/cmds"

type App struct {
	cmds.Command
	cmds.Envs
}
