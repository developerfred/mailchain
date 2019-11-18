// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encodingtest

import (
	"reflect"
	"testing"

	"github.com/mailchain/mailchain/internal/testutil"
)

// MustDecodeBase58 decodes a Base58 string.
// It panics for invalid input.
func TestMustDecodeBase58(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"success",
			args{
				"5CLmNK8f16nagFeF2h3iNeeChaxPiAsJu7piNYJgdPpmaRzPD",
			},
			testutil.MustDecodeBase58("5CLmNK8f16nagFeF2h3iNeeChaxPiAsJu7piNYJgdPpmaRzPD"),
			false,
		},
		{
			"err-empty",
			args{
				"",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testutil.MustDecodeBase58(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("MustDecodeBase58() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustDecodeBase58() = %v, want %v", got, tt.want)
			}

		})
	}
}
