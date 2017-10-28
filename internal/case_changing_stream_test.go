// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kylelemons/godebug/pretty"
)

func TestCaseChangingStream(t *testing.T) {
	tests := []struct {
		input string
		upper bool
		want  []int
	}{
		{"abcd", true, []int{'A', 'B', 'C', 'D', antlr.TokenEOF}},
		{"ABCD", true, []int{'A', 'B', 'C', 'D', antlr.TokenEOF}},
		{"abcd", false, []int{'a', 'b', 'c', 'd', antlr.TokenEOF}},
		{"ABCD", false, []int{'a', 'b', 'c', 'd', antlr.TokenEOF}},
		{"", false, []int{antlr.TokenEOF}},
	}

	for _, test := range tests {
		var got []int
		is := NewCaseChangingStream(antlr.NewInputStream(test.input), test.upper)
		for i := 1; i <= is.Size()+1; i++ {
			got = append(got, is.LA(i))
		}

		if diff := pretty.Compare(test.want, got); diff != "" {
			t.Errorf("NewCaseChangingStream(%q, %v) diff: (-got +want)\n%s", test.input, test.upper, diff)
		}
	}
}
