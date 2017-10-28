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
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CaseChangingStream wraps an existing CharStream, but upper cases, or lower
// cases the input before it is tokenized.
//
// It would be useful if this was part of the core antlr library, but it was
// rejected in https://github.com/antlr/antlr4/pull/2046 as users can trivial
// strings.ToUpper(...) the input, or write their own version of this stream.
type CaseChangingStream struct {
	antlr.CharStream

	upper bool
}

func NewCaseChangingStream(in antlr.CharStream, upper bool) *CaseChangingStream {
	return &CaseChangingStream{
		in, upper,
	}
}

func (is *CaseChangingStream) LA(offset int) int {
	in := is.CharStream.LA(offset)
	if in < 0 {
		// Such as antlr.TokenEOF which is -1
		return in
	}
	if is.upper {
		return int(unicode.ToUpper(rune(in)))
	}
	return int(unicode.ToLower(rune(in)))
}
