package app

import "github.com/cybriq/proc/pkg/cmds"

type App struct {
	cmds.Command
	cmds.Envs
}

func New(cmd *cmds.Command) (a App) {
	cmds.GetConfigBase(cmd.Configs, cmd.Name, false)
	_, cmd = cmds.Init(cmd)
	cfgFile := cmd.GetOpt(cmds.Path{cmd.Name, "ConfigFile"})
	// cfgFile.SetExpanded(cfgFile)
	_ = cfgFile.Expanded()
	return
}
