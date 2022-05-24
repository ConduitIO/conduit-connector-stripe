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

package config

import (
	"reflect"
	"testing"

	"go.uber.org/multierr"
)

func TestParse(t *testing.T) {
	underTestConfig := Config{}

	tests := []struct {
		name        string
		in          map[string]string
		want        Config
		wantErr     bool
		expectedErr string
	}{
		{
			name: "valid config",
			in: map[string]string{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: "5",
				Limit:              "10",
			},
			want: Config{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: 5,
				Limit:              10,
			},
		},
		{
			name: "HTTPClientRetryMax and Limit by default",
			in: map[string]string{
				SecretKey:    "sk_51JB",
				ResourceName: "subscription",
			},
			want: Config{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: RetryMaxDefault,
				Limit:              LimitDefault,
			},
		},
		{
			// maha: this test and the one after are the same
			name: "no secret key",
			in: map[string]string{
				ResourceName: "subscription",
			},
			wantErr:     true,
			expectedErr: underTestConfig.RequiredConfigErr(SecretKey).Error(),
		},
		{
			name: "empty secret key",
			in: map[string]string{
				SecretKey:    "",
				ResourceName: "subscription",
			},
			wantErr:     true,
			expectedErr: underTestConfig.RequiredConfigErr(SecretKey).Error(),
		},
		{
			name: "no resource name",
			in: map[string]string{
				SecretKey: "sk_51JB",
			},
			wantErr:     true,
			expectedErr: underTestConfig.RequiredConfigErr(ResourceName).Error(),
		},
		{
			name: "empty resource name",
			in: map[string]string{
				SecretKey:    "sk_51JB",
				ResourceName: "",
			},
			wantErr:     true,
			expectedErr: underTestConfig.RequiredConfigErr(ResourceName).Error(),
		},
		{
			// maha: what i understood is that when you say "no" means the fields are not there.. so let's keep it the same
			// and add another test for when these two are empty
			name:    "no secret key and resource name",
			in:      map[string]string{},
			wantErr: true,
			expectedErr: multierr.Combine(underTestConfig.RequiredConfigErr(SecretKey),
				underTestConfig.RequiredConfigErr(ResourceName)).Error(),
		},
		{
			name: "HTTPClientRetryMax is greater than the value of lte tag",
			in: map[string]string{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: "12",
			},
			wantErr:     true,
			expectedErr: underTestConfig.OutOfRangeConfigErr(HTTPClientRetryMax).Error(),
		},
		{
			name: "HTTPClientRetryMax is more than the value of lte tag",
			in: map[string]string{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: "0",
			},
			wantErr:     true,
			expectedErr: underTestConfig.OutOfRangeConfigErr(HTTPClientRetryMax).Error(),
		},
		{
			name: "invalid HTTPClientRetryMax",
			in: map[string]string{
				SecretKey:          "sk_51JB",
				ResourceName:       "subscription",
				HTTPClientRetryMax: "test",
			},
			wantErr:     true,
			expectedErr: underTestConfig.IntegerTypeConfigErr(HTTPClientRetryMax).Error(),
		},
		{
			name: "wrong resource name",
			in: map[string]string{
				SecretKey:    "sk_51JB",
				ResourceName: "test",
			},
			wantErr:     true,
			expectedErr: underTestConfig.WrongResourceNameConfigErr(ResourceName).Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.in)
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

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse = %v, want %v", got, tt.want)
			}
		})
	}
}
