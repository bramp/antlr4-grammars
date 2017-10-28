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
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

// PrintAllTokens consumes all the tokens from the lexer and prints them to stdout.
func PrintAllTokens(lexer antlr.Lexer) {
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%d %s\n", t.GetChannel(), t.GetText())
	}
}

// TestingErrorListener is a antlr.ErrorListener which fails a given test if a
// SyntaxError is found. The other Report errors are ignored.
type TestingErrorListener struct {
	*antlr.DefaultErrorListener

	t    *testing.T
	name string
}

func NewTestingErrorListener(t *testing.T, name string) *TestingErrorListener {
	return &TestingErrorListener{t: t, name: name}
}

func (l *TestingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.t.Errorf("%s: SyntaxError %d:%d: %s", l.name, line, column, msg)
}
