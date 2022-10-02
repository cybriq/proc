package cmds

import (
	"strings"
	"testing"

	log2 "github.com/cybriq/proc/pkg/log"
	"github.com/cybriq/proc/pkg/opts/config"
	"github.com/cybriq/proc/pkg/path"
)

func TestCommand_Foreach(t *testing.T) {
	log2.SetLogLevel(log2.Info)
	cm, _ := Init(GetExampleCommands(), nil)
	log.I.Ln("spewing only droptxindex")
	cm.ForEach(func(cmd *Command, _ int) bool {
		if cmd.Name == "droptxindex" {
			log.I.S(cmd)
		}
		return true
	}, 0, 0, cm)
	log.I.Ln("printing name of all commands found on search")
	cm.ForEach(func(cmd *Command, depth int) bool {
		log.I.Ln()
		log.I.F("%s%s #(%d)", path.GetIndent(depth), cmd.Path, depth)
		for i := range cmd.Configs {
			log.I.F("%s%s -%s %v #%v (%d)", path.GetIndent(depth),
				cmd.Configs[i].Path(), i, cmd.Configs[i].String(), cmd.Configs[i].Meta().Aliases(), depth)
		}
		return true
	}, 0, 0, cm)
}

func TestCommand_MarshalText(t *testing.T) {

	log2.SetLogLevel(log2.Info)
	o, _ := Init(GetExampleCommands(), nil)
	// log.I.S(o)
	conf, err := o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	log.I.Ln("\n" + string(conf))
}

func TestCommand_UnmarshalText(t *testing.T) {
	log2.SetLogLevel(log2.Info)
	o, _ := Init(GetExampleCommands(), nil)
	var conf []byte
	var err error
	conf, err = o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = o.UnmarshalText(conf)
	if err != nil {
		t.FailNow()
	}
}

func TestCommand_ParseCLIArgs(t *testing.T) {
	args1 := "/random/path/to/server_binary --cafile ~/some/cafile --LC=cn node -addrindex --BD 48h30s"
	args1s := strings.Split(args1, " ")
	log2.SetLogLevel(log2.Debug)
	o, _ := Init(GetExampleCommands(), nil)
	run, err := o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	_, _ = run, err
	args2 := "/random/path/to/server_binary node -addrindex --BD=48h30s  -RPCMaxConcurrentReqs -16 dropaddrindex "
	args2s := strings.Split(args2, " ")
	run, err = o.ParseCLIArgs(args2s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	args3 := "node -addrindex --BD 48h30s dropaddrindex somegarbage --autoports"
	args3s := strings.Split(args3, " ")
	run, err = o.ParseCLIArgs(args3s)
	// This one must fail, 'somegarbage' is not a command and has no -/-- prefix
	if err == nil {
		t.FailNow()
	}
	args4 := "/random/path/to/server_binary --cafile ~/some/cafile --LC=cn"
	args4s := strings.Split(args4, " ")
	run, err = o.ParseCLIArgs(args4s)
	if log.E.Chk(err) {
		t.FailNow()
	}
}

func TestCommand_GetEnvs(t *testing.T) {
	log2.SetLogLevel(log2.Info)
	o, _ := Init(GetExampleCommands(), nil)
	envs := o.GetEnvs()
	var out []string
	err := envs.ForEach(func(env string, opt config.Option) error {
		out = append(out, env)
		return nil
	})
	for i := range out { // verifying ordering groups subcommands
		log.I.Ln(out[i])
	}
	if err != nil {
		t.FailNow()
	}
}
