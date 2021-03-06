// Copyright 2017 CoreOS, Inc.
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

package dbtester

import (
	"reflect"
	"testing"
)

func Test_assignRequest(t *testing.T) {
	ranges := []int64{1, 10, 50, 100, 300, 500, 700, 1000}
	total := int64(1000000)
	rs := assignRequest(ranges, total)
	expected := []int64{120000, 120000, 120000, 120000, 120000, 120000, 120000, 160000}
	if !reflect.DeepEqual(rs, expected) {
		t.Fatalf("expected %+v, got %+v", expected, rs)
	}

	cur := int64(0)
	for _, v := range expected {
		cur += v
	}
	if cur != total {
		t.Fatalf("sum must be %d, got %d", total, cur)
	}
}
