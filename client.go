package geverse

import (
	"context"
	"io"
	"net/http"
)

type Geverse struct {
	httpClient *http.Client
	token      string
}

func NewGeverse(httpClient *http.Client, token string) *Geverse {
	return &Geverse{
		httpClient: httpClient,
		token:      token,
	}
}

func (g *Geverse) newRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := g.newBasicRequest(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+g.token)

	return req, nil
}

func (g *Geverse) newBasicRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", "geverse/0.0.1")

	return req, nil
}
