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

// A CDC represents a struct of cdc iterator.
type CDC struct {
	stripeSvc Stripe
	position  *position.Position
	// haris: would be good to clarify
	eventData []models.EventData
}

// NewCDC initializes cdc iterator.
// haris: given that pos is passed as a pointer, that means we are updating it
// are we sure that won't create any negative side effects?
func NewCDC(stripeSvc Stripe, pos *position.Position) *CDC {
	return &CDC{
		stripeSvc: stripeSvc,
		position:  pos,
	}
}

// Next returns the next record.
func (iter *CDC) Next() (sdk.Record, error) {
	if iter.eventData == nil || iter.position.Index == 0 {
		if err := iter.getData(); err != nil {
			return sdk.Record{}, fmt.Errorf("get event data: %w", err)
		}

		if len(iter.eventData) == 0 {
			return sdk.Record{}, sdk.ErrBackoffRetry
		}
	}

	index := iter.position.Index

	iter.position.Index++

	// update the cursor before formatting the last cached data record
	if len(iter.eventData) == iter.position.Index {
		iter.position.Index = 0
		iter.position.Cursor = iter.eventData[len(iter.eventData)-1].ID
	}

	payload, err := json.Marshal(iter.eventData[index].Data.Object)
	if err != nil {
		return sdk.Record{}, fmt.Errorf("marshal payload: %w", err)
	}

	output := sdk.Record{
		Position: iter.position.FormatSDKPosition(),
		Metadata: map[string]string{
			models.ActionKey: models.EventsAction[iter.eventData[index].Type],
		},
		CreatedAt: time.Unix(iter.eventData[index].Created, 0),
		Key: sdk.StructuredData{
			idKey: iter.eventData[index].Data.Object[idKey].(string),
		},
		Payload: sdk.RawData(payload),
	}

	return output, nil
}

func (iter *CDC) getData() error {
	if iter.position.Cursor == "" {
		return iter.getDataWithStartingAfter()
	}

	return iter.getDataWithEndingBefore()
}

func (iter *CDC) getDataWithStartingAfter() error {
	var (
		eventData []models.EventData

		startingAfter string
	)

	for {
		resp, err := iter.stripeSvc.GetEvent(iter.position.CreatedAt, startingAfter, "")
		if err != nil {
			return fmt.Errorf("get list of event objects: %w", err)
		}

		if len(resp.Data) > 0 {
			startingAfter = resp.Data[len(resp.Data)-1].ID

			eventData = append(eventData, resp.Data...)
		}

		if !resp.HasMore {
			break
		}
	}

	if len(eventData) > 0 {
		// do the reverse after all requests, because Stripe receives data ordered by DESC always
		for i, j := 0, len(eventData)-1; i < j; i, j = i+1, j-1 {
			eventData[i], eventData[j] = eventData[j], eventData[i]
		}
	}

	iter.eventData = eventData

	return nil
}

func (iter *CDC) getDataWithEndingBefore() error {
	var (
		eventData []models.EventData

		endingBefore = iter.position.Cursor
	)

	for {
		resp, err := iter.stripeSvc.GetEvent(iter.position.CreatedAt, "", endingBefore)
		if err != nil {
			return fmt.Errorf("get list of event objects: %w", err)
		}

		if len(resp.Data) > 0 {
			endingBefore = resp.Data[0].ID

			// do the reverse after each request, because we get the data from the last page of Stripe events
			for i, j := 0, len(resp.Data)-1; i < j; i, j = i+1, j-1 {
				resp.Data[i], resp.Data[j] = resp.Data[j], resp.Data[i]
			}

			eventData = append(eventData, resp.Data...)
		}

		if !resp.HasMore {
			break
		}
	}

	iter.eventData = eventData

	return nil
}
