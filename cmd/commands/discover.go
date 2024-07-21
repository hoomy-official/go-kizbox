package commands

import (
	"context"
	"github.com/merlindorin/go-kizbox/pkg/discover"
	"github.com/vanyda-official/go-shared/pkg/cmd"
	"log"
	"time"

	"github.com/hoomy-official/exporter-unifi-protect-cli/filter"

	"go.uber.org/zap"
)

type DiscoverCmd struct {
	filter.Filter

	Timeout time.Duration `default:"5s" help:"timeout for discovering (ns, ms, s & m)"`
}

func (d *DiscoverCmd) Run(common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	logger = logger.With(zap.Duration("timeout", d.Timeout))
	ctx := context.Background()

	ch := make(chan *discover.Gateway, 1)

	go func() {
		for gateway := range ch {
			log.Println(gateway)
		}
	}()

	return discover.NewKizboxDiscover(d.Timeout).Discover(ctx, ch)
}
