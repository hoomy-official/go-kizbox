package devices

type Cmd struct {
	List  DevicesListCmd `cmd:"list" help:"list devices available"`
	Get   DevicesGetCmd  `cmd:"get" help:"get device by URL"`
	Apply ApplyCmd       `cmd:"apply" help:"apply"`
	RPC   ApplyCmd       `cmd:"rpc" help:"rpc"`
}
