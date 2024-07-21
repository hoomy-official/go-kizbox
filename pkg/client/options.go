package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type Option func(cl *APIClient)

func (w Option) apply(cl *APIClient) {
	w(cl)
}

func WithBaseURL(baseURL *url.URL) Option {
	return func(a *APIClient) {
		a.baseURL = baseURL.JoinPath(apiPath)
	}
}

func WithHTTPClient(cl *http.Client) Option {
	return func(a *APIClient) {
		a.HTTPClient = cl
	}
}

func WithAuthAdapter(adapter AuthAdapter) Option {
	return func(a *APIClient) {
		a.authAdapter = adapter
	}
}

func WithAuthToken(token string) Option {
	return WithAuthAdapter(TokenAuth(token))
}

func WithDefaultHTTPClient() Option {
	cl := http.DefaultClient
	cl.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // provider specs
	}
	return WithHTTPClient(cl)
}
