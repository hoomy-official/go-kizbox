package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	v1 "github.com/hoomy-official/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/net/do"
	"github.com/vanyda-official/go-shared/pkg/net/rest"
)

const (
	apiPath = "/enduser-mobile-web/1/enduserAPI"
)

var (
	//nolint:gochecknoglobals // lazyness
	defaultClientOptions = []Option{WithDefaultHTTPClient()}

	//nolint:gochecknoglobals // lazyness
	defaultRequesterOptions = []do.Option{
		do.WithJSONRequest(),
		do.WithPostRequestHandler("http_error_code_handler", defaultHTTPErrorCodeHandler),
	}
)

type APIClient struct {
	HTTPClient  *http.Client
	baseURL     *url.URL
	authAdapter AuthAdapter

	rest.Requester
	V1 V1
}

type V1 struct {
	Version   *v1.APIVersion
	Gateway   *v1.APIGateway
	Devices   *v1.APIDevices
	Execution *v1.APIExecution
	Event     *v1.APIEvent
}

func NewClient(options ...Option) *APIClient {
	apiClient := &APIClient{}

	for _, option := range append(defaultClientOptions, options...) {
		option.apply(apiClient)
	}

	apiClient.Requester = rest.NewRest(
		apiClient.baseURL,
		append(
			defaultRequesterOptions,
			do.WithClient(apiClient.HTTPClient),
			do.WithPreRequestHandler("auth-adapter", do.PreRequestHandlerFunc(apiClient.authAdapter)),
		)...,
	)

	apiClient.V1 = V1{
		Version:   v1.NewAPIVersion(apiClient),
		Gateway:   v1.NewAPIGateway(apiClient),
		Devices:   v1.NewAPIDevices(apiClient),
		Execution: v1.NewAPIExecution(apiClient),
		Event:     v1.NewAPIEvent(apiClient),
	}

	return apiClient
}

func (cl *APIClient) Do(ctx context.Context, options ...do.Option) error {
	return do.Do(
		ctx,
		cl.baseURL,
		append(
			append(
				defaultRequesterOptions,
				do.WithClient(cl.HTTPClient),
				do.WithPreRequestHandler("auth-adapter", do.PreRequestHandlerFunc(cl.authAdapter)),
			),
			options...,
		)...,
	)
}

func defaultHTTPErrorCodeHandler(_ context.Context, _ *http.Request, response *http.Response) error {
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(
			"unexcpected response status code %d: %s",
			response.StatusCode,
			http.StatusText(response.StatusCode),
		)
	}

	return nil
}
