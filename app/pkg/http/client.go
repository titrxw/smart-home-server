package http

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type Option interface {
	apply(client *resty.Client)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*resty.Client)

func (f optionFunc) apply(client *resty.Client) {
	f(client)
}

func WithOptions(client *resty.Client, opts ...Option) *resty.Client {
	for _, opt := range opts {
		opt.apply(client)
	}

	return client
}

func PostWithWWWForm(ctx context.Context, url string, body map[string]string, headers map[string][]string, options ...Option) (*resty.Response, error) {
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Type"] = []string{
		"application/x-www-form-urlencoded",
	}

	client := resty.New()
	WithOptions(client, options...)

	return client.R().
		SetContext(ctx).
		SetHeaderMultiValues(headers).
		SetFormData(body).
		Post(url)
}

func PostWithJson(ctx context.Context, url string, body any, headers map[string][]string, options ...Option) (*resty.Response, error) {
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Type"] = []string{
		"application/json",
	}

	client := resty.New()
	WithOptions(client, options...)

	return client.R().
		SetContext(ctx).
		SetHeaderMultiValues(headers).
		SetBody(body).
		Post(url)
}

func Get(ctx context.Context, url string, params map[string]string, headers map[string][]string, options ...Option) (*resty.Response, error) {
	client := resty.New()
	WithOptions(client, options...)

	return client.R().
		SetContext(ctx).
		SetHeaderMultiValues(headers).
		SetQueryParams(params).
		Get(url)
}
