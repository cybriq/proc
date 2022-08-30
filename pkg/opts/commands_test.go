package opts

import (
	"strings"
	"testing"

	log2 "github.com/cybriq/proc/pkg/log"
)

func TestCommand_Foreach(t *testing.T) {
	cm := GetCommands()
	log.I.Ln("spewing only droptxindex")
	cm.Foreach(func(cmd *Command, _ int) bool {
		if cmd.Name == "droptxindex" {
			log.I.S(cmd)
		}
		return true
	})
	log.I.Ln("printing name of all commands found on search")
	cm.Foreach(func(cmd *Command, depth int) bool {
		log.I.Ln(strings.Repeat("\t", depth) + cmd.Name)
		return true
	})
}

func TestCommand_MarshalText(t *testing.T) {

	log2.SetLogLevel(log2.Info)
	o := New(GetCommands())
	// log.I.S(o)
	conf, err := o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	log.I.Ln("\n" + string(conf))
}

// GetCommands returns available subcommands in Parallelcoin Pod
func GetCommands() (c *Command) {
	c = &Command{
		Commands: Commands{
			{
				Name:        "gui",
				Description: "ParallelCoin GUI Wallet/Miner/Explorer",
			},
			{
				Name:        "version",
				Description: "print version and exit",
			},
			{
				Name:        "ctl",
				Description: "command line wallet and chain RPC client",
			},
			{
				Name:        "node",
				Description: "ParallelCoin blockchain node",
				Commands: []*Command{
					{
						Name:        "dropaddrindex",
						Description: "drop the address database index",
					},
					{
						Name:        "droptxindex",
						Description: "drop the transaction database index",
					},
					{
						Name:        "dropcfindex",
						Description: "drop the cfilter database index",
					},
					{
						Name:        "dropindexes",
						Description: "drop all of the indexes",
					},
					{
						Name:        "resetchain",
						Description: "deletes the current blockchain cache to force redownload",
					},
				},
			},
			{
				Name:        "wallet",
				Description: "run the wallet server (requires a chain node to function)",
				Entrypoint:  func(c interface{}) error { return nil },
				Commands: []*Command{
					{
						Name:        "drophistory",
						Description: "reset the wallet transaction history",
					},
				},
			},
			{
				Name:        "kopach",
				Description: "standalone multicast miner for easy mining farm deployment",
			},
			{
				Name:        "worker",
				Description: "single thread worker process, normally started by kopach",
			},
		},
	}
	return
}
