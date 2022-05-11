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

package position

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

var errParseCreatedAt = errors.New("the fourth part of position must be an int64")

const (
	SnapshotType = "s"
	CDCType      = "c"
)

// A Position represents a Stripe position.
type Position struct {
	IteratorType string
	Cursor       string
	CreatedAt    int64
}

// ParseSDKPosition parses SDK position and returns Position.
func ParseSDKPosition(p sdk.Position) (Position, error) {
	if p == nil {
		return Position{
			IteratorType: SnapshotType,
			CreatedAt:    time.Now().Unix(),
		}, nil
	}

	parts := strings.Split(string(p), ".")

	if len(parts) != reflect.TypeOf(Position{}).NumField() {
		return Position{}, fmt.Errorf("the number of position elements must be equal to %d, now it is %d",
			reflect.TypeOf(Position{}).NumField(), len(parts))
	}

	started, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return Position{}, errParseCreatedAt
	}

	return Position{
		IteratorType: parts[0],
		Cursor:       parts[1],
		CreatedAt:    started,
	}, nil
}

// FormatSDKPosition formats and returns sdk.Position.
func (p Position) FormatSDKPosition() sdk.Position {
	return sdk.Position(fmt.Sprintf("%s.%s.%d", p.IteratorType, p.Cursor, p.CreatedAt))
}
