package client

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/errs"
)

//Client is class object for HTTP client
type Client struct {
	ctx    context.Context
	client *http.Client
}

//New function returns Client instance.
func New(ctx context.Context, cli *http.Client) *Client {
	if ctx == nil {
		ctx = context.Background()
	}
	if cli == nil {
		cli = http.DefaultClient
	}
	return &Client{ctx: ctx, client: cli}
}

//Default function returns defauilt Client instance.
func Default() *Client { return New(context.Background(), nil) }

//Get method returns respons data from URL.
func (c *Client) Get(rawurl string) (io.ReadCloser, error) {
	if c == nil {
		c = Default()
	}
	if len(rawurl) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidRequest, errs.WithContext("url", rawurl))
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, errs.Wrap(ecode.ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", rawurl))
	}
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errs.Wrap(ecode.ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", rawurl))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(ecode.ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", rawurl))
	}
	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		resp.Body.Close()
		return nil, errs.Wrap(ecode.ErrInvalidRequest, errs.WithCause(ecode.ErrHTTPStatus), errs.WithContext("url", rawurl), errs.WithContext("status", resp.Status))
	}
	return resp.Body, nil
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
