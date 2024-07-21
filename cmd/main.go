package main

import (
	_ "embed"

	"github.com/hoomy-official/exporter-unifi-protect-cli/commands"
	"github.com/hoomy-official/exporter-unifi-protect-cli/commands/devices"
	"github.com/hoomy-official/exporter-unifi-protect-cli/globals"
	"github.com/vanyda-official/go-shared/pkg/cmd"

	"github.com/alecthomas/kong"
)

var (
	name    = "hoomy"
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

type CLI struct {
	*cmd.Commons
	*globals.Globals

	Venitian commands.VenitianCmd `cmd:"venitians"`
	Devices  devices.Cmd          `cmd:"devices" help:"list devices availables in the current system"`
	Listen   commands.ListenCmd   `cmd:"listen" help:"listen events in the current system"`
	Discover commands.DiscoverCmd `cmd:"discover" help:"list for systems available"`
}

func main() {
	cli := CLI{
		Commons: &cmd.Commons{
			Version: cmd.NewVersion(name, version, commit, buildSource, date),
			Licence: cmd.NewLicence(license),
		},
		Globals: &globals.Globals{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description("Simple cli for managing my home automation"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Globals, cli.Commons))
}
