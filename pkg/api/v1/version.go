package v1

import (
	"context"

	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

type APIVersion struct {
	cl rest.Requester
}

func NewAPIVersion(cl rest.Requester) *APIVersion {
	return &APIVersion{cl: cl}
}

func (a *APIVersion) Get(ctx context.Context, v *Version) error {
	return a.cl.GET(ctx, do.WithPath("/apiVersion"), do.WithUnmarshalBody(v))
}
