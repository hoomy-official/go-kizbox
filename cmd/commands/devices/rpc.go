package devices

import (
	"context"
	v1 "github.com/merlindorin/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"

	"github.com/hoomy-official/exporter-unifi-protect-cli/commands"
	"github.com/hoomy-official/exporter-unifi-protect-cli/filter"
	"github.com/hoomy-official/exporter-unifi-protect-cli/globals"

	"go.uber.org/zap"
)

type RPC struct {
}

// RpcCmd execute actions to a resources
type RpcCmd struct {
	filter.Filter
	Controllables []string
	Action        string `arg:"action"`
	Args          []string
}

func (s RpcCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	devices, err := commands.DeviceList(ctx, api, s.Controllables, s.Filter)
	if err != nil {
		logger.Error("cannot list device")
		return err
	}

	var actions []v1.Action
	for _, device := range devices {
		logger.Debug("open device", zap.Any("device", device))
		action := v1.Action{
			Commands: []v1.Command{
				{
					Name:       "ping",
					Parameters: []interface{}{},
				},
			},
			DeviceURL: device.DeviceURL,
		}

		actions = append(actions, action)
	}

	return api.V1.Execution.Apply(ctx, v1.Execute{Label: "cli command test", Actions: actions}, nil)
	return err
}
