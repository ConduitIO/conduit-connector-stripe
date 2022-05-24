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

package iterator

import (
	"errors"

	sdk "github.com/conduitio/conduit-connector-sdk"

	"github.com/conduitio/conduit-connector-stripe/source/position"
)

const idKey = "id"

// A Combined represents a struct of combined iterator.
type Combined struct {
	snapshot *Snapshot
	cdc      *CDC
	position *position.Position
}

// New initializes a combined iterator.
func New(stripeSvc Stripe, pos *position.Position) *Combined {
	combined := &Combined{
		position: pos,
		cdc:      NewCDC(stripeSvc, pos),
	}

	if pos.IteratorType == position.SnapshotType {
		combined.snapshot = NewSnapshot(stripeSvc, pos)
	}

	return combined
}

// Next returns the next record.
func (iter *Combined) Next() (sdk.Record, error) {
	switch iter.position.IteratorType {
	case position.SnapshotType:
		return iter.snapshot.Next()
	case position.CDCType:
		return iter.cdc.Next()
	}

	// haris: nitpick: would be good to also return the unexpected iterator type
	return sdk.Record{}, errors.New("the iterator type is wrong")
}
