package v1

import (
	"context"

	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

type APIGateway struct {
	cl rest.Requester
}

func NewAPIGateway(cl rest.Requester) *APIGateway {
	return &APIGateway{cl: cl}
}

func (a *APIGateway) List(ctx context.Context, v *[]Gateway) error {
	return a.cl.GET(ctx, do.WithPath("/setup/gateways"), do.WithUnmarshalBody(v))
}
