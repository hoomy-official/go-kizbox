package devices

import (
	"context"
	"fmt"
	"github.com/hoomy-official/exporter-unifi-protect-cli/globals"
	v1 "github.com/merlindorin/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"
)

type DevicesGetCmd struct {
	URL string `arg:"URL"`
}

func (d *DevicesGetCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := global.Client()
	ctx := context.Background()

	var device v1.Device

	err = cl.V1.Devices.Get(ctx, d.URL, &device)
	if err != nil {
		logger.Error("cannot get device")
		return err
	}

	fmt.Printf("%s", device)

	return nil
}
