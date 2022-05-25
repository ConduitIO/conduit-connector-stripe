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
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/conduitio/conduit-connector-sdk"

	"github.com/conduitio/conduit-connector-stripe/models"
	"github.com/conduitio/conduit-connector-stripe/source/position"
)

// A Snapshot represents a struct of snapshot iterator.
// haris: nitpick: it might be good to rename it to SnapshotIterator.
// Snapshot (at least to me) sounds like it refers to the snapshot itself,
// not an iterator over all of the snapshot data.
type Snapshot struct {
	stripeSvc Stripe
	position  *position.Position
	response  *models.ResourceResponse
	index     int
}

// NewSnapshot initializes snapshot iterator.
func NewSnapshot(stripeSvc Stripe, pos *position.Position) *Snapshot {
	return &Snapshot{
		stripeSvc: stripeSvc,
		position:  pos,
	}
}

// Next returns the next record.
func (iter *Snapshot) Next() (sdk.Record, error) {
	if iter.response == nil || len(iter.response.Data) == iter.index {
		// haris: maybe rename populateWithResource to refreshData
		// that would make reading this code here easier
		if err := iter.populateWithResource(); err != nil {
			return sdk.Record{}, fmt.Errorf("populate with the resource: %w", err)
		}

		if len(iter.response.Data) == 0 {
			iter.position.IteratorType = position.CDCType
			iter.position.Cursor = ""

			return sdk.Record{}, sdk.ErrBackoffRetry
		}
	}

	// haris: why do we marshall it into JSON? what's the data format of iter.response.Data[iter.index]?
	payload, err := json.Marshal(iter.response.Data[iter.index])
	if err != nil {
		return sdk.Record{}, fmt.Errorf("marshal payload: %w", err)
	}

	iter.position.Cursor = iter.response.Data[iter.index][idKey].(string)

	created, ok := iter.response.Data[iter.index]["created"].(float64)
	if !ok {
		created = float64(time.Now().Unix())
	}

	output := sdk.Record{
		Position: iter.position.FormatSDKPosition(),
		Metadata: map[string]string{
			models.ActionKey: models.InsertAction,
		},
		CreatedAt: time.Unix(int64(created), 0),
		Key: sdk.StructuredData{
			idKey: iter.response.Data[iter.index][idKey].(string),
		},
		Payload: sdk.RawData(payload),
	}

	iter.index++

	return output, nil
}

func (iter *Snapshot) populateWithResource() error {
	resp, err := iter.stripeSvc.GetResource(iter.position.Cursor)
	if err != nil {
		return fmt.Errorf("get list of resource objects: %w", err)
	}

	iter.response = &resp
	iter.index = 0

	return nil
}
