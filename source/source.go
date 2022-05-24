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

	"github.com/conduitio/conduit-connector-stripe/config"
	"github.com/conduitio/conduit-connector-stripe/source/iterator"
	"github.com/conduitio/conduit-connector-stripe/source/position"
	"github.com/conduitio/conduit-connector-stripe/stripe"
)

// A Source represents the source connector.
type Source struct {
	sdk.UnimplementedSource
	cfg      *config.Config
	iterator iterator.Iterator
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

	s.cfg = &cfg

	return nil
}

// Open parses sdk.Position and initializes a Snapshot iterator.
func (s *Source) Open(ctx context.Context, rp sdk.Position) error {
	pos, err := position.ParseSDKPosition(rp)
	if err != nil {
		return err
	}

	s.iterator = iterator.New(stripe.New(s.cfg), &pos)

	return nil
}

// Read returns the next sdk.Record.
func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	r, err := s.iterator.Next()
	if err != nil {
		return sdk.Record{}, err
	}

	return r, nil
}

// Teardown does nothing.
func (s *Source) Teardown(ctx context.Context) error {
	// maha: are all the connections closed after this call?
	return nil
}

// Ack does nothing.
func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	sdk.Logger(ctx).Debug().Str("position", string(position)).Msg("got ack")
	return nil
}
