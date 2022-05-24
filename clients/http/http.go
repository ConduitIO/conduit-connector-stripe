// Copyright © 2022 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"fmt"
	"io/ioutil"

	"github.com/conduitio/conduit-connector-stripe/config"
	"github.com/hashicorp/go-retryablehttp"
)

const methodGet = "GET"

// A Client represents retryable http client.
type Client struct {
	HTTPClient *retryablehttp.Client
}

// NewClient returns a new retryable http client.
func NewClient(cfg *config.Config) Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = cfg.HTTPClientRetryMax

	return Client{
		HTTPClient: retryClient,
	}
}

// Get makes a GET http-request to the URL with headers.
func (cli Client) Get(url string, header ...map[string]string) ([]byte, error) {
	req, err := retryablehttp.NewRequest(methodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	for i := range header {
		for k, v := range header[i] {
			req.Header.Add(k, v)
		}
	}

	r, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read all response body: %w", err)
	}

	return data, r.Body.Close()
}
