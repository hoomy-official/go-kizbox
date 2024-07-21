package v1

import (
	"context"
	"net/url"

	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

type APIDevices struct {
	cl rest.Requester
}

func NewAPIDevices(cl rest.Requester) *APIDevices {
	return &APIDevices{cl: cl}
}

func (receiver *APIDevices) List(ctx context.Context, v *[]Device) error {
	return receiver.cl.GET(ctx, do.WithPath("/setup/devices"), do.WithUnmarshalBody(v))
}

func (receiver *APIDevices) Get(ctx context.Context, deviceURL string, v *Device) error {
	return receiver.cl.GET(ctx, do.WithPath("/setup/devices/%s", url.PathEscape(deviceURL)), do.WithUnmarshalBody(v))
}

func (receiver *APIDevices) States(ctx context.Context, deviceURL string, v *[]State) error {
	return receiver.cl.GET(
		ctx,
		do.WithPath("/setup/devices/%s/states", url.PathEscape(deviceURL)),
		do.WithUnmarshalBody(v),
	)
}

func (receiver *APIDevices) State(ctx context.Context, deviceURL string, stateName string, v *State) error {
	return receiver.cl.GET(
		ctx,
		do.WithPath("/setup/devices/%s/states/%s", url.PathEscape(deviceURL), stateName),
		do.WithUnmarshalBody(v),
	)
}

func (receiver *APIDevices) Controllables(ctx context.Context, controllableName string, v *[]string) error {
	return receiver.cl.GET(
		ctx,
		do.WithPath("/setup/devices/controllables/%s", url.PathEscape(controllableName)),
		do.WithUnmarshalBody(v),
	)
}
