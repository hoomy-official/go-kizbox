package v1

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

type APIExecution struct {
	cl rest.Requester
}

func NewAPIExecution(cl rest.Requester) *APIExecution {
	return &APIExecution{cl: cl}
}

type Apply struct {
	ID string `json:"id"`
}

func (receiver *APIExecution) Apply(ctx context.Context, execute Execute, s *Apply) error {
	// patch resource to ensure it is acceptable
	for i, action := range execute.Actions {
		for j, command := range action.Commands {
			if command.Parameters == nil {
				// null parameters is not supported, an empty array must be provided
				execute.Actions[i].Commands[j].Parameters = make([]interface{}, 0)
			}
		}
	}

	ex, err := json.Marshal(execute)
	if err != nil {
		return err
	}

	return receiver.cl.POST(ctx, do.WithPath("/exec/apply"), do.WithBody(bytes.NewReader(ex)), do.WithUnmarshalBody(s))
}

func (receiver *APIExecution) Current(ctx context.Context, s *[]Execution) error {
	return receiver.cl.GET(ctx, do.WithPath("/exec/current"), do.WithUnmarshalBody(s))
}

func (receiver *APIExecution) Get(ctx context.Context, execID string, s *Execution) error {
	return receiver.cl.GET(ctx, do.WithPath("/exec/%s", execID), do.WithUnmarshalBody(s))
}
