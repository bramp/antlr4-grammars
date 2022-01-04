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

// Package cql_test contains tests for the Cql grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by make.go
//
package cql_test

import (
	"bramp.net/antlr4/cql"
	"bramp.net/antlr4/internal"

	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"strings"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/cql3/examples/alterKeyspace.cql",
	"grammars-v4/cql3/examples/alterMaterializedView.cql",
	"grammars-v4/cql3/examples/alterRole.cql",
	"grammars-v4/cql3/examples/alterTable.cql",
	"grammars-v4/cql3/examples/alterType.cql",
	"grammars-v4/cql3/examples/alterUser.cql",
	"grammars-v4/cql3/examples/applyBatch.cql",
	"grammars-v4/cql3/examples/createAggregate.cql",
	"grammars-v4/cql3/examples/createFunction.cql",
	"grammars-v4/cql3/examples/createIndex.cql",
	"grammars-v4/cql3/examples/createKeyspace.cql",
	"grammars-v4/cql3/examples/createMaterializedView.cql",
	"grammars-v4/cql3/examples/createRole.cql",
	"grammars-v4/cql3/examples/createTable.cql",
	"grammars-v4/cql3/examples/createTrigger.cql",
	"grammars-v4/cql3/examples/createType.cql",
	"grammars-v4/cql3/examples/createUser.cql",
	"grammars-v4/cql3/examples/delete.cql",
	"grammars-v4/cql3/examples/dropAggregate.cql",
	"grammars-v4/cql3/examples/dropFunction.cql",
	"grammars-v4/cql3/examples/dropIndex.cql",
	"grammars-v4/cql3/examples/dropKeyspace.cql",
	"grammars-v4/cql3/examples/dropMaterializedView.cql",
	"grammars-v4/cql3/examples/dropRole.cql",
	"grammars-v4/cql3/examples/dropTable.cql",
	"grammars-v4/cql3/examples/dropTrigger.cql",
	"grammars-v4/cql3/examples/dropType.cql",
	"grammars-v4/cql3/examples/dropUser.cql",
	"grammars-v4/cql3/examples/grant.cql",
	"grammars-v4/cql3/examples/insert.cql",
	"grammars-v4/cql3/examples/listPermissions.cql",
	"grammars-v4/cql3/examples/listRoles.cql",
	"grammars-v4/cql3/examples/revoke.cql",
	"grammars-v4/cql3/examples/select.cql",
	"grammars-v4/cql3/examples/truncate.cql",
	"grammars-v4/cql3/examples/update.cql",
	"grammars-v4/cql3/examples/use.cql",
}

type exampleListener struct {
	*cql.BaseCqlParserListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
func Example() {
	// Setup the input (which this parser expects to be uppercased).
	is := antlr.NewInputStream(strings.ToUpper("...some text to parse..."))

	// Create the Lexer
	lexer := cql.NewCqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := cql.NewCqlParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Root()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("..", filename))
	if err != nil {
		return nil, err
	}

	input = internal.NewCaseChangingStream(input, true)
	return input, nil
}

func TestCqlLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := cql.NewCqlLexer(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			tok := lexer.NextToken()
			if tok.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i >= MAX_TOKENS {
			t.Errorf("NewCqlLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestCqlParser(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := cql.NewCqlLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := cql.NewCqlParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(internal.NewTestingErrorListener(t, file))

		// Finally test
		p.Root()

		// TODO(bramp): If there is a "file.tree", then compare the output
		// TODO(bramp): If there is a "file.errors", then check the error
	}
}
