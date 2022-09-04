package opts

import (
	"strings"
	"testing"

	log2 "github.com/cybriq/proc/pkg/log"
	integer "github.com/cybriq/proc/pkg/opts/Integer"
	"github.com/cybriq/proc/pkg/opts/config"
	"github.com/cybriq/proc/pkg/opts/duration"
	"github.com/cybriq/proc/pkg/opts/float"
	"github.com/cybriq/proc/pkg/opts/list"
	"github.com/cybriq/proc/pkg/opts/meta"
	"github.com/cybriq/proc/pkg/opts/text"
	"github.com/cybriq/proc/pkg/opts/toggle"
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
	o := Init(GetCommands())
	// log.I.S(o)
	conf, err := o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	log.I.Ln("\n" + string(conf))
}

// GetCommands returns available subcommands in Parallelcoin Pod
func GetCommands() (c *Command) {
	tags := Tags
	c = &Command{
		Configs: config.Opts{
			"AutoPorts": toggle.New(meta.Data{
				Label:         "Automatic Ports",
				Tags:          tags("node", "wallet"),
				Description:   "RPC and controller ports are randomized, use with controller for automatic peer discovery",
				Documentation: "<placeholder for detailed documentation>",
			}),
			"AutoListen": toggle.New(meta.Data{
				Aliases:       []string{"AL"},
				Tags:          tags("node", "wallet"),
				Label:         "Automatic Listeners",
				Description:   "automatically update inbound addresses dynamically according to discovered network interfaces",
				Documentation: "<placeholder for detailed documentation>",
			}),
			"CAFile": text.New(
				meta.Data{
					Aliases:       []string{"CA"},
					Tags:          tags("node", "wallet", "tls"),
					Label:         "Certificate Authority File",
					Description:   "certificate authority file for TLS certificate validation",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"ConfigFile": text.New(
				meta.Data{
					Aliases:       []string{"CF"},
					Label:         "Configuration File",
					Description:   "location of configuration file, cannot actually be changed",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"CPUProfile": text.New(
				meta.Data{
					Aliases:       []string{"CPR"},
					Tags:          tags("node", "wallet", "kopach", "worker"),
					Label:         "CPU Profile",
					Description:   "write cpu profile to this file",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"DataDir": text.New(
				meta.Data{
					Aliases: []string{"DD"},
					Label:   "Data Directory",
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Description:   "root folder where application data is stored",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"DisableRPC": toggle.New(
				meta.Data{
					Aliases:       []string{"NRPC"},
					Tags:          tags("node", "wallet"),
					Label:         "Disable RPC",
					Description:   "disable rpc servers, as well as kopach controller",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Discovery": toggle.New(
				meta.Data{
					Aliases:       []string{"DI"},
					Tags:          tags("node"),
					Label:         "Disovery",
					Description:   "enable LAN peer discovery in GUI",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Hilite": list.New(
				meta.Data{
					Aliases: []string{"HL"},
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Label:         "Hilite",
					Description:   "list of packages that will print with attention getters",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Locale": text.New(
				meta.Data{
					Aliases: []string{"LC"},
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Label:         "Language",
					Description:   "user interface language i18 localization",
					Documentation: "<placeholder for detailed documentation>",
					Options:       []string{"en"},
				},
			),
			"LimitPass": text.New(
				meta.Data{
					Aliases:       []string{"LP"},
					Tags:          tags("node", "wallet"),
					Label:         "Limit Password",
					Description:   "limited user password",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"LimitUser": text.New(
				meta.Data{
					Aliases:       []string{"LU"},
					Tags:          tags("node", "wallet"),
					Label:         "Limit Username",
					Description:   "limited user name",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"LogDir": text.New(
				meta.Data{
					Aliases: []string{"LD"},
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Label:         "Log Directory",
					Description:   "folder where log files are written",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"LogFilter": list.New(
				meta.Data{
					Aliases: []string{"LF"},
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Label:         "Log Filter",
					Description:   "list of packages that will not print logs",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"LogLevel": text.New(
				meta.Data{
					Aliases: []string{"LL"},
					Tags: tags("node", "wallet", "ctl", "kopach",
						"worker"),
					Label:       "Log Level",
					Description: "maximum log level to output",
					Options: []string{
						"off",
						"fatal",
						"error",
						"info",
						"check",
						"debug",
						"trace",
					},
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"OneTimeTLSKey": toggle.New(
				meta.Data{
					Aliases:       []string{"OTK"},
					Tags:          tags("node", "wallet"),
					Label:         "One Time TLS Key",
					Description:   "generate a new TLS certificate pair at startup, but only write the certificate to disk",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Password": text.New(
				meta.Data{
					Aliases:       []string{"PW"},
					Tags:          tags("node", "wallet"),
					Label:         "Password",
					Description:   "password for client RPC connections",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"PipeLog": toggle.New(
				meta.Data{
					Aliases: []string{"PL"},
					Label:   "Pipe Logger",
					Tags: tags(
						"node",
						"wallet",
						"ctl",
						"kopach",
						"worker",
					),
					Description:   "enable pipe based logger IPC",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Profile": text.New(
				meta.Data{
					Aliases: []string{"HPR"},
					Tags: tags("node", "wallet", "ctl", "kopach",
						"worker"),
					Label:       "Profile",
					Description: "http profiling on given port (1024-40000)",
					// Type:        "",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"RPCCert": text.New(
				meta.Data{
					Aliases:       []string{"RC"},
					Tags:          tags("node", "wallet"),
					Label:         "RPC Cert",
					Description:   "location of RPC TLS certificate",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"RPCKey": text.New(
				meta.Data{
					Aliases:       []string{"RK"},
					Tags:          tags("node", "wallet"),
					Label:         "RPC Key",
					Description:   "location of rpc TLS key",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"RunAsService": toggle.New(
				meta.Data{
					Aliases:       []string{"RS"},
					Label:         "Run As Service",
					Description:   "shuts down on lock timeout",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Save": toggle.New(
				meta.Data{
					Aliases:       []string{"SV"},
					Label:         "Save Configuration",
					Description:   "save opts given on commandline",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"ServerTLS": toggle.New(
				meta.Data{
					Aliases:       []string{"ST"},
					Tags:          tags("node", "wallet"),
					Label:         "Server TLS",
					Description:   "enable TLS for the wallet connection to node RPC server",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"ClientTLS": toggle.New(
				meta.Data{
					Aliases:       []string{"CT"},
					Tags:          tags("node", "wallet"),
					Label:         "TLS",
					Description:   "enable TLS for RPC client connections",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"TLSSkipVerify": toggle.New(
				meta.Data{
					Aliases:       []string{"TSV"},
					Tags:          tags("node", "wallet"),
					Label:         "TLS Skip Verify",
					Description:   "skip TLS certificate verification (ignore CA errors)",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"Username": text.New(
				meta.Data{
					Aliases:       []string{"UN"},
					Tags:          tags("node", "wallet"),
					Label:         "Username",
					Description:   "password for client RPC connections",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"UseWallet": toggle.New(
				meta.Data{
					Aliases:       []string{"WC"},
					Tags:          tags("ctl"),
					Label:         "Connect to Wallet",
					Description:   "set ctl to connect to wallet instead of chain server",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
			"WalletOff": toggle.New(
				meta.Data{
					Aliases:       []string{"WO"},
					Tags:          tags("wallet"),
					Label:         "Wallet Off",
					Description:   "turn off the wallet backend",
					Documentation: "<placeholder for detailed documentation>",
				},
			),
		},
		Commands: Commands{
			{
				Name:        "gui",
				Description: "ParallelCoin GUI Wallet/Miner/Explorer",
				Configs: config.Opts{
					"DarkTheme": toggle.New(
						meta.Data{
							Aliases:       []string{"DT"},
							Tags:          tags("gui"),
							Label:         "Dark Theme",
							Description:   "sets dark theme for GUI",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
				},
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
				Configs: config.Opts{
					"AddCheckpoints": list.New(meta.Data{
						Aliases:       []string{"AC"},
						Tags:          tags("node"),
						Label:         "Add Checkpoints",
						Description:   "add custom checkpoints",
						Documentation: "<placeholder for detailed documentation>",
					}),
					"AddPeers": list.New(meta.Data{
						Aliases:       []string{"AP"},
						Tags:          tags("node"),
						Label:         "Add Peers",
						Description:   "manually adds addresses to try to connect to",
						Documentation: "<placeholder for detailed documentation>",
					}),
					"AddrIndex": toggle.New(meta.Data{
						Aliases:       tags("AI"),
						Tags:          tags("node"),
						Label:         "Address Index",
						Description:   "maintain a full address-based transaction index which makes the searchrawtransactions RPC available",
						Documentation: "<placeholder for detailed documentation>",
					}),
					"BanDuration": duration.New(
						meta.Data{
							Aliases:       []string{"BD"},
							Tags:          tags("node", "policy"),
							Label:         "Ban Duration",
							Description:   "how long a ban of a misbehaving peer lasts",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BanThreshold": integer.New(
						meta.Data{
							Aliases:       []string{"BT"},
							Tags:          tags("node", "policy"),
							Label:         "Ban Threshold",
							Description:   "ban score that triggers a ban (default 100)",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlockMaxSize": integer.New(
						meta.Data{
							Aliases:       []string{"BMXS"},
							Tags:          tags("node", "policy"),
							Label:         "Block Max Size",
							Description:   "maximum block size in bytes to be used when creating a block",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlockMaxWeight": integer.New(
						meta.Data{
							Aliases:       []string{"BMXW"},
							Tags:          tags("node", "policy"),
							Label:         "Block Max Weight",
							Description:   "maximum block weight to be used when creating a block",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlockMinSize": integer.New(
						meta.Data{
							Aliases:       []string{"BMS"},
							Tags:          tags("node", "policy"),
							Label:         "Block Min Size",
							Description:   "minimum block size in bytes to be used when creating a block",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlockMinWeight": integer.New(
						meta.Data{
							Aliases:       []string{"BMW"},
							Tags:          tags("node"),
							Label:         "Block Min Weight",
							Description:   "minimum block weight to be used when creating a block",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlockPrioritySize": integer.New(
						meta.Data{
							Aliases:       []string{"BPS"},
							Tags:          tags("node"),
							Label:         "Block Priority Size",
							Description:   "size in bytes for high-priority/low-fee transactions when creating a block",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"BlocksOnly": toggle.New(
						meta.Data{
							Aliases:       []string{"BO"},
							Tags:          tags("node"),
							Label:         "Blocks Only",
							Description:   "do not accept transactions from remote peers",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"ConnectPeers": list.New(meta.Data{
						Aliases:       []string{"CPS"},
						Tags:          tags("node"),
						Label:         "Connect Peers",
						Description:   "connect ONLY to these addresses (disables inbound connections)",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"Controller": toggle.New(meta.Data{
						Aliases:       []string{"CN"},
						Tags:          tags("node"),
						Label:         "Enable Controller",
						Description:   "delivers mining jobs over multicast",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"DbType": text.New(meta.Data{
						Aliases: []string{"DB"},
						Tags:    tags("node"),
						Label:   "Database Type",
						Description: "type of database storage engine to use for node (" +
							"only one right now, ffldb)",
						Documentation: "<placeholder for detailed documentation>",
						Options:       tags("ffldb"),
					},
					),
					"DisableBanning": toggle.New(meta.Data{
						Aliases:       []string{"NB"},
						Tags:          tags("node"),
						Label:         "Disable Banning",
						Description:   "disables banning of misbehaving peers",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"DisableCheckpoints": toggle.New(meta.Data{
						Aliases:       []string{"NCP"},
						Tags:          tags("node"),
						Label:         "Disable Checkpoints",
						Description:   "disables all checkpoints",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"DisableDNSSeed": toggle.New(meta.Data{
						Aliases:       []string{"NDS"},
						Tags:          tags("node"),
						Label:         "Disable DNS Seed",
						Description:   "disable seeding of addresses to peers",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"DisableListen": toggle.New(meta.Data{
						Aliases:       []string{"NL"},
						Tags:          tags("node", "wallet"),
						Label:         "Disable Listen",
						Description:   "disables inbound connections for the peer to peer network",
						Documentation: "<placeholder for detailed documentation>",
					},
					),
					"ExternalIPs": list.New(
						meta.Data{
							Aliases:       []string{"EI"},
							Tags:          tags("node"),
							Label:         "External IP Addresses",
							Description:   "extra addresses to tell peers they can connect to",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"FreeTxRelayLimit": float.New(
						meta.Data{
							Aliases:       []string{"LR"},
							Tags:          tags("node"),
							Label:         "Free Tx Relay Limit",
							Description:   "limit relay of transactions with no transaction fee to the given amount in thousands of bytes per minute",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"LAN": toggle.New(
						meta.Data{
							Tags:          tags("node"),
							Label:         "LAN Testnet Mode",
							Description:   "run without any connection to nodes on the internet (does not apply on mainnet)",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"MaxOrphanTxs": integer.New(
						meta.Data{
							Aliases:       []string{"MO"},
							Tags:          tags("node"),
							Label:         "Max Orphan Txs",
							Description:   "max number of orphan transactions to keep in memory",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"MaxPeers": integer.New(
						meta.Data{
							Aliases:       []string{"MP"},
							Tags:          tags("node"),
							Label:         "Max Peers",
							Description:   "maximum number of peers to hold connections with",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"MinRelayTxFee": float.New(
						meta.Data{
							Aliases:       []string{"MRTF"},
							Tags:          tags("node"),
							Label:         "Min Relay Transaction Fee",
							Description:   "the minimum transaction fee in DUO/kB to be considered a non-zero fee",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"Network": text.New(
						meta.Data{
							Aliases:     []string{"NW"},
							Tags:        tags("node", "wallet"),
							Label:       "Network",
							Description: "connect to this network:",
							Options: []string{
								"mainnet",
								"testnet",
								"regtestnet",
								"simnet",
							},
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"NoCFilters": toggle.New(
						meta.Data{
							Aliases:       []string{"NCF"},
							Tags:          tags("node"),
							Label:         "No CFilters",
							Description:   "disable committed filtering (CF) support",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"NodeOff": toggle.New(
						meta.Data{
							Aliases:       []string{"NO"},
							Tags:          tags("node"),
							Label:         "Node Off",
							Description:   "turn off the node backend",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"NoPeerBloomFilters": toggle.New(
						meta.Data{
							Aliases:       []string{"NPBF"},
							Tags:          tags("node"),
							Label:         "No Peer Bloom Filters",
							Description:   "disable bloom filtering support",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"NoRelayPriority": toggle.New(
						meta.Data{
							Aliases:       []string{"NRPR"},
							Tags:          tags("node"),
							Label:         "No Relay Priority",
							Description:   "do not require free or low-fee transactions to have high priority for relaying",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"OnionEnabled": toggle.New(
						meta.Data{
							Aliases:       []string{"OE"},
							Tags:          tags("node"),
							Label:         "Onion Enabled",
							Description:   "enable tor proxy",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"OnionProxyAddress": text.New(
						meta.Data{
							Aliases:       []string{"OPA"},
							Tags:          tags("node"),
							Label:         "Onion Proxy Address",
							Description:   "address of tor proxy you want to connect to",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"OnionProxyPass": text.New(
						meta.Data{
							Aliases:       []string{"OPW"},
							Tags:          tags("node"),
							Label:         "Onion Proxy Password",
							Description:   "password for tor proxy",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"OnionProxyUser": text.New(
						meta.Data{
							Aliases:       []string{"OU"},
							Tags:          tags("node"),
							Label:         "Onion Proxy Username",
							Description:   "tor proxy username",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"P2PConnect": list.New(
						meta.Data{
							Aliases:       []string{"P2P"},
							Tags:          tags("node"),
							Label:         "P2P Connect",
							Description:   "list of addresses reachable from connected networks",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"P2PListeners": list.New(
						meta.Data{
							Aliases:       []string{"LA"},
							Tags:          tags("node"),
							Label:         "P2PListeners",
							Description:   "list of addresses to bind the node listener to",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"ProxyAddress": text.New(
						meta.Data{
							Aliases:       []string{"PA"},
							Tags:          tags("node"),
							Label:         "Proxy",
							Description:   "address of proxy to connect to for outbound connections",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"ProxyPass": text.New(
						meta.Data{
							Aliases:       []string{"PPW"},
							Tags:          tags("node"),
							Label:         "Proxy Pass",
							Description:   "proxy password, if required",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"ProxyUser": text.New(
						meta.Data{
							Aliases:       []string{"PU"},
							Tags:          tags("node"),
							Label:         "ProxyUser",
							Description:   "proxy username, if required",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RejectNonStd": toggle.New(
						meta.Data{
							Aliases:       []string{"REJ"},
							Tags:          tags("node"),
							Label:         "Reject Non Std",
							Description:   "reject non-standard transactions regardless of the default settings for the active network",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RelayNonStd": toggle.New(
						meta.Data{
							Aliases:       []string{"RNS"},
							Tags:          tags("node"),
							Label:         "Relay Nonstandard Transactions",
							Description:   "relay non-standard transactions regardless of the default settings for the active network",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCConnect": text.New(
						meta.Data{
							Aliases:       []string{"RA"},
							Tags:          tags("node"),
							Label:         "RPC Connect",
							Description:   "full node RPC for wallet",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCListeners": list.New(
						meta.Data{
							Aliases:       []string{"RL"},
							Tags:          tags("node"),
							Label:         "Node RPC Listeners",
							Description:   "addresses to listen for RPC connections",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCMaxClients": integer.New(
						meta.Data{
							Aliases:       []string{"RMXC"},
							Tags:          tags("node"),
							Label:         "Maximum Node RPC Clients",
							Description:   "maximum number of clients for regular RPC",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCMaxConcurrentReqs": integer.New(
						meta.Data{
							Aliases:       []string{"RMCR"},
							Tags:          tags("node"),
							Label:         "Maximum Node RPC Concurrent Reqs",
							Description:   "maximum number of requests to process concurrently",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCMaxWebsockets": integer.New(
						meta.Data{
							Aliases:       []string{"RMWS"},
							Tags:          tags("node"),
							Label:         "Maximum Node RPC Websockets",
							Description:   "maximum number of websocket clients to allow",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"RPCQuirks": toggle.New(
						meta.Data{
							Aliases:       []string{"RQ"},
							Tags:          tags("node"),
							Label:         "Emulate Bitcoin Core RPC Quirks",
							Description:   "enable bugs that replicate bitcoin core RPC's JSON",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"SigCacheMaxSize": integer.New(
						meta.Data{
							Aliases:       []string{"SCM"},
							Tags:          tags("node"),
							Label:         "Signature Cache Max Size",
							Description:   "the maximum number of entries in the signature verification cache",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"Solo": toggle.New(
						meta.Data{
							Label:         "Solo Generate",
							Tags:          tags("node"),
							Description:   "mine even if not connected to a network",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"TorIsolation": toggle.New(
						meta.Data{
							Aliases:       []string{"TI"},
							Tags:          tags("node"),
							Label:         "Tor Isolation",
							Description:   "makes a separate proxy connection for each connection",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"TrickleInterval": duration.New(
						meta.Data{
							Aliases:       []string{"TKI"},
							Tags:          tags("node"),
							Label:         "Trickle Interval",
							Description:   "minimum time between attempts to send new inventory to a connected peer",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"TxIndex": toggle.New(
						meta.Data{
							Aliases:       []string{"TXI"},
							Tags:          tags("node"),
							Label:         "Tx Index",
							Description:   "maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"UPNP": toggle.New(
						meta.Data{
							Aliases:       []string{"UP"},
							Tags:          tags("node"),
							Label:         "UPNP",
							Description:   "enable UPNP for NAT traversal",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"UserAgentComments": list.New(
						meta.Data{
							Aliases:       []string{"UA"},
							Tags:          tags("node"),
							Label:         "User Agent Comments",
							Description:   "comment to add to the user agent -- See BIP 14 for more information",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"UUID": integer.New(
						meta.Data{
							Label:         "UUID",
							Description:   "instance unique id (32bit random value) (json mangles big 64 bit integers due to float64 numbers)",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"Whitelists": list.New(
						meta.Data{
							Aliases:       []string{"WL"},
							Tags:          tags("node"),
							Label:         "Whitelists",
							Description:   "peers that you don't want to ever ban",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
				},
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
				Commands: []*Command{
					{
						Name:        "drophistory",
						Description: "reset the wallet transaction history",
						Configs:     config.Opts{},
					},
				},
				Configs: config.Opts{
					"WalletFile": text.New(
						meta.Data{
							Aliases:       []string{"WF"},
							Tags:          tags("wallet"),
							Label:         "Wallet File",
							Description:   "wallet database file",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"WalletPass": text.New(
						meta.Data{
							Aliases: []string{"WPW"},
							Label:   "Wallet Pass",
							Tags:    tags("wallet"),
							Description: "password encrypting public data in wallet - only hash is stored" +
								" so give on command line or in environment POD_WALLETPASS",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"WalletRPCListeners": list.New(
						meta.Data{
							Aliases:       []string{"WRL"},
							Tags:          tags("wallet"),
							Label:         "Wallet RPC Listeners",
							Description:   "addresses for wallet RPC server to listen on",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"WalletRPCMaxClients": integer.New(
						meta.Data{
							Aliases:       []string{"WRMC"},
							Tags:          tags("wallet"),
							Label:         "Legacy RPC Max Clients",
							Description:   "maximum number of RPC clients allowed for wallet RPC",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"WalletRPCMaxWebsockets": integer.New(
						meta.Data{
							Aliases:       []string{"WRMWS"},
							Tags:          tags("wallet"),
							Label:         "Legacy RPC Max Websockets",
							Description:   "maximum number of websocket clients allowed for wallet RPC",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"WalletServer": text.New(
						meta.Data{
							Aliases:       []string{"WS"},
							Tags:          tags("wallet"),
							Label:         "Wallet Server",
							Description:   "node address to connect wallet server to",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
				},
			},
			{
				Name:        "kopach",
				Description: "standalone multicast miner for easy mining farm deployment",
				Configs: config.Opts{
					"Generate": toggle.New(
						meta.Data{
							Aliases:       []string{"GB"},
							Tags:          tags("node", "kopach"),
							Label:         "Generate Blocks",
							Description:   "turn on Kopach CPU miner",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"GenThreads": integer.New(
						meta.Data{
							Aliases:       []string{"GT"},
							Tags:          tags("kopach"),
							Label:         "Generate Threads",
							Description:   "number of threads to mine with",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
					"MulticastPass": text.New(
						meta.Data{
							Aliases:       []string{"PM"},
							Tags:          tags("node", "kopach"),
							Label:         "Multicast Pass",
							Description:   "password that encrypts the connection to the mining controller",
							Documentation: "<placeholder for detailed documentation>",
						},
					),
				},
			},
			{
				Name:        "worker",
				Description: "single thread worker process, normally started by kopach",
			},
		},
	}
	return
}
