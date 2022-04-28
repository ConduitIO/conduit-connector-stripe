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

package source

import (
	"context"

	sdk "github.com/conduitio/conduit-connector-sdk"

	"github.com/ConduitIO/conduit-connector-stripe/clients/http"
	"github.com/ConduitIO/conduit-connector-stripe/config"
)

// A Source represents the source connector.
type Source struct {
	sdk.UnimplementedSource
	config     *config.Config
	httpClient *http.Client
}

// NewSource initialises a new source.
func NewSource() sdk.Source {
	return &Source{}
}

// Configure parses and stores configurations, returns an error in case of invalid configuration.
func (s *Source) Configure(ctx context.Context, cfgRaw map[string]string) error {
	cfg, err := config.Parse(cfgRaw)
	if err != nil {
		return err
	}

	s.config = &cfg

	s.httpClient = http.NewClient(s.config)

	return nil
}