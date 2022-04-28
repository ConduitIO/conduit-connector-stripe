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
	"github.com/hashicorp/go-retryablehttp"

	"github.com/ConduitIO/conduit-connector-stripe/config"
)

// A Client represents retryable http client.
type Client struct {
	config     *config.Config
	httpClient *retryablehttp.Client
}

// NewClient returns a new retryable http client.
func NewClient(config *config.Config) *Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = config.HTTPClientRetryMax

	return &Client{
		httpClient: retryClient,
		config:     config,
	}
}