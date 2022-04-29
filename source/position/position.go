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

	sdk "github.com/conduitio/conduit-connector-sdk"
)

// A Position represents a Stripe position.
type Position struct {
	HasMore       bool
	StartingAfter string
}

// ParseSDKPosition parses SDK position and returns Position.
func ParseSDKPosition(p sdk.Position) (Position, error) {
	if p == nil {
		return Position{}, nil
	}

	parts := strings.Split(string(p), ".")

	if len(parts) != reflect.TypeOf(Position{}).NumField() {
		return Position{}, fmt.Errorf("the number of position elements must be equal to %d, now it is %d",
			reflect.TypeOf(Position{}).NumField(), len(parts))
	}

	hasMore, err := strconv.ParseBool(parts[0])
	if err != nil {
		return Position{}, errors.New("the first part of position must be a bool")
	}

	return Position{
		HasMore:       hasMore,
		StartingAfter: parts[1],
	}, nil
}

// FormatSDKPosition formats and returns sdk.Position.
func (p Position) FormatSDKPosition() sdk.Position {
	return sdk.Position(fmt.Sprintf("%t.%s", p.HasMore, p.StartingAfter))
}
