package devices

import (
	"context"
	"fmt"
	v1 "github.com/merlindorin/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"
	"slices"

	"github.com/hoomy-official/exporter-unifi-protect-cli/filter"
	"github.com/hoomy-official/exporter-unifi-protect-cli/globals"
)

type DevicesListCmd struct {
	filter.Filter
}

func (d *DevicesListCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := global.Client()
	ctx := context.Background()

	var devices []v1.Device
	err = cl.V1.Devices.List(ctx, &devices)
	if err != nil {
		logger.Error("cannot list devices")
		return err
	}

	for _, device := range devices {
		if (len(d.URLS) == 0 && len(d.Labels) == 0) ||
			slices.Contains(d.URLS, device.DeviceURL) ||
			slices.Contains(d.Labels, device.Label) {
			logger.Info(fmt.Sprintf("Device %s (%s)", device.Label, device.DeviceURL))
		}
	}

	return nil
}
