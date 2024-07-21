package discover

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/vanyda-official/go-shared/pkg/discover"
	"github.com/vanyda-official/go-shared/pkg/resolvers/mdns"
	"golang.org/x/sync/errgroup"

	"github.com/grandcat/zeroconf"
)

const (
	DefaultService = "_kizboxdev._tcp"
	DefaultDomain  = "local."
)

type KizboxDiscover struct {
	timeout time.Duration
}

func NewKizboxDiscover(timeout time.Duration) *KizboxDiscover {
	return &KizboxDiscover{timeout: timeout}
}

func (k KizboxDiscover) Discover(ctx context.Context, discovered chan<- *Gateway) error {
	defer close(discovered)

	g, ctx := errgroup.WithContext(ctx)
	ch := make(chan interface{}, 1)

	g.Go(func() error {
		resolver := mdns.New(
			mdns.WithIPv4AndIPv6(),
			mdns.WithDomain(DefaultDomain),
			mdns.WithService(DefaultService),
			mdns.WithTransformer(transformer),
		)

		return discover.NewDiscover(resolver, discover.WithTimeout(k.timeout)).Discover(ctx, ch)
	})

	g.Go(func() error {
		for i := range ch {
			if gateway, ok := i.(*Gateway); ok {
				discovered <- gateway
			}
		}

		return nil
	})

	return g.Wait()
}

// Gateway represents a gateway that hold details about the service such as its name,
// service type, host information, IP addresses, port, and other metadata like API version,
// gateway PIN, and firmware version.
type Gateway struct {
	Name            string
	Service         string
	Host            string
	AddrV4          []net.IP
	AddrV6          []net.IP
	Port            int
	APIVersion      string
	GatewayPin      string
	FirmwareVersion string
}

func transformer(entry *zeroconf.ServiceEntry) (interface{}, error) {
	s := &Gateway{
		Name:    entry.Instance,
		Service: entry.Service,
		Host:    entry.HostName,
		AddrV4:  entry.AddrIPv4,
		AddrV6:  entry.AddrIPv6,
		Port:    entry.Port,
	}

	for _, txt := range entry.Text {
		k, v, _ := strings.Cut(txt, "=")
		switch k {
		case "api_version":
			s.APIVersion = v
		case "gateway_pin":
			s.GatewayPin = v
		case "fw_version":
			s.FirmwareVersion = v
		}
	}

	return s, nil
}
