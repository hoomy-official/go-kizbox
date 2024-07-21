package v1

import (
	"context"

	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

type APIEvent struct {
	cl rest.Requester
}

func NewAPIEvent(cl rest.Requester) *APIEvent {
	return &APIEvent{cl: cl}
}

func (receiver *APIEvent) Register(ctx context.Context, v *EventRegister) error {
	return receiver.cl.POST(ctx, do.WithPath("/events/register"), do.WithUnmarshalBody(v))
}

func (receiver *APIEvent) Fetch(ctx context.Context, eventRegister EventRegister, v *[]map[string]interface{}) error {
	return receiver.cl.POST(ctx, do.WithPath("/events/%s/fetch", eventRegister.ID), do.WithUnmarshalBody(v))
}

func (receiver *APIEvent) Unregister(ctx context.Context, listenerID string) error {
	return receiver.cl.POST(ctx, do.WithPath("/events/%s/unregister", listenerID))
}
