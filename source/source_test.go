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
	"reflect"
	"testing"

	"go.uber.org/multierr"

	"github.com/conduitio/conduit-connector-stripe/config"
)

func TestSource_Configure(t *testing.T) {
	underTestConfig := config.Config{}
	underTestSource := Source{}

	tests := []struct {
		name        string
		in          map[string]string
		want        Source
		wantErr     bool
		expectedErr string
	}{
		{
			name: "valid config",
			in: map[string]string{
				config.SecretKey:    "sk_51JB",
				config.ResourceName: "subscription",
			},
			want: Source{
				cfg: &config.Config{
					SecretKey:          "sk_51JB",
					ResourceName:       "subscription",
					HTTPClientRetryMax: config.RetryMaxDefault,
					Limit:              config.LimitDefault,
				},
			},
		},
		{
			name: "no secret key and resource name",
			in: map[string]string{
				config.SecretKey:    "",
				config.ResourceName: "",
			},
			wantErr: true,
			expectedErr: multierr.Combine(underTestConfig.RequiredConfigErr(config.SecretKey),
				underTestConfig.RequiredConfigErr(config.ResourceName)).Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := underTestSource.Configure(context.Background(), tt.in)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("parse error = \"%s\", wantErr %t", err.Error(), tt.wantErr)

					return
				}

				if err.Error() != tt.expectedErr {
					t.Errorf("expected error \"%s\", got \"%s\"", tt.expectedErr, err.Error())

					return
				}

				return
			}

			if !reflect.DeepEqual(underTestSource.cfg, tt.want.cfg) {
				t.Errorf("parse = %v, want %v", underTestSource.cfg, tt.want.cfg)

				return
			}
		})
	}
}

// maha: we need more tests for real scenarios.. check S3 connector source tests for ideas
// like, what happens when source is empty, what happens when it stops and continue, what happens when it stops
// exactly after snapshot and before cdc, and more..
// https://github.com/ConduitIO/conduit-connector-s3/blob/main/source/source_integration_test.go
