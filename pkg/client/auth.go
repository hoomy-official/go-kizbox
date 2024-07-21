package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vanyda-official/go-shared/pkg/net/do"
)

type AuthAdapter do.PreRequestHandlerFunc

func TokenAuth(token string) AuthAdapter {
	return func(_ context.Context, r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}
