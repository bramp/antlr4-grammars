# antlr4 Go parsers [![Build Status](https://img.shields.io/travis/bramp/antlr4-grammars.svg)](https://travis-ci.org/bramp/antlr4-grammars) [![Coverage](https://img.shields.io/coveralls/bramp/antlr4-grammars.svg)](https://coveralls.io/github/bramp/antlr4-grammars) [![Report card](https://goreportcard.com/badge/bramp.net/antlr4)](https://goreportcard.com/report/bramp.net/antlr4) [![GoDoc](https://godoc.org/bramp.net/antlr4?status.svg)](https://godoc.org/bramp.net/antlr4)
Compiled by Andrew Brampton ([bramp.net](https://bramp.net))

Precompiled Go parsers of many of the grammars on [github.com/antlr/grammars-v4](http://github.com/antlr/grammars-v4).

The Antlr's Go Target is still a work in progress. As such, many of the grammars fail to compile, or only pass simple tests. To report issues with the grammar [go here](https://github.com/antlr/grammars-v4), to report issues with Antlr's Go Target [go here](https://github.com/antlr/antlr4).

## Example

```go
import (
	"fmt"

	"bramp.net/antlr4/json"                    // The parser
	"github.com/antlr/antlr4/runtime/Go/antlr" // The antlr library
)

// exampleListener is an event-driven callback for the parser.
type exampleListener struct {
	*json.BaseJSONListener // https://godoc.org/bramp.net/antlr4/json#BaseJSONListener
}

func (l *exampleListener) EnterObj(ctx *json.ObjContext) {
	fmt.Printf("Object: %s\n", ctx.GetText())

}

func (l *exampleListener) EnterPair(ctx *json.PairContext) {
	fmt.Printf("Pair: %s\n", ctx.GetText())

}

func (l *exampleListener) EnterArray(ctx *json.ArrayContext) {
	fmt.Printf("Array: %s\n", ctx.GetText())

}

// Example shows how to use the JSON lexer and parser.
func Example() {
	// Setup the input
	is := antlr.NewInputStream(`{"example": "json", "with": ["an", "array"]}`)

	// Create the Lexer
	lexer := json.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := json.NewJSONParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)

	// Output:
	// Object: {"example":"json","with":["an","array"]}
	// Pair: "example":"json"
	// Pair: "with":["an","array"]
	// Array: ["an","array"]
}
```

## Supported Languages

| Status | Language     | Notes                                                                       |
| :-: | --------------- | --------------------------------------------------------------------------- |
| ✅  | abnf            |                                                                             |
| ✅  | agc             |                                                                             |
| ✅  | arithmetic      |                                                                             |
| ✅  | atl             |                                                                             |
| ✅  | bnf             |                                                                             |
| ✅  | brainfuck       |                                                                             |
| ✅  | c               |                                                                             |
| ✅  | clf             |                                                                             |
| ✅  | clif            |                                                                             |
| ✅  | cobol85         |                                                                             |
| ✅  | cobol85preprocessor |                                                                             |
| ✅  | cookie          |                                                                             |
| ✅  | cool            |                                                                             |
| ✅  | corundum        |                                                                             |
| ✅  | creole          |                                                                             |
| ✅  | csv             |                                                                             |
| ✅  | datetime        |                                                                             |
| ✅  | dcm_2_0_grammar |                                                                             |
| ✅  | dgs             |                                                                             |
| ✅  | dot             |                                                                             |
| ✅  | ecmascript      |                                                                             |
| ✅  | emailaddress    |                                                                             |
| ✅  | fasta           |                                                                             |
| ✅  | fol             |                                                                             |
| ✅  | fusiontablessql |                                                                             |
| ✅  | gml             |                                                                             |
| ✅  | graphemes       |                                                                             |
| ✅  | gtin            |                                                                             |
| ✅  | idl             |                                                                             |
| ✅  | iri             |                                                                             |
| ✅  | istc            |                                                                             |
| ✅  | jpa             |                                                                             |
| ✅  | json            |                                                                             |
| ✅  | lambda          |                                                                             |
| ✅  | lcc             |                                                                             |
| ✅  | less            |                                                                             |
| ✅  | lexunicode      |                                                                             |
| ✅  | mdx             |                                                                             |
| ✅  | memcached_protocol |                                                                             |
| ✅  | metric          |                                                                             |
| ✅  | modelica        |                                                                             |
| ✅  | molecule        |                                                                             |
| ✅  | morsecode       |                                                                             |
| ✅  | mps             |                                                                             |
| ✅  | mu              |                                                                             |
| ✅  | mumath          |                                                                             |
| ✅  | mumps           |                                                                             |
| ✅  | objectivec      |                                                                             |
| ✅  | oncrpcv2        |                                                                             |
| ✅  | p               |                                                                             |
| ✅  | pcre            |                                                                             |
| ✅  | peoplecode      |                                                                             |
| ✅  | postalcode      |                                                                             |
| ✅  | powerbuilder    |                                                                             |
| ✅  | propcalc        |                                                                             |
| ✅  | properties      |                                                                             |
| ✅  | r               |                                                                             |
| ✅  | rcs             |                                                                             |
| ✅  | redcode         |                                                                             |
| ✅  | robotwar        |                                                                             |
| ✅  | romannumerals   |                                                                             |
| ✅  | rpn             |                                                                             |
| ✅  | scss            |                                                                             |
| ✅  | sexpression     |                                                                             |
| ✅  | sharc           |                                                                             |
| ✅  | snobol          |                                                                             |
| ✅  | stacktrace      |                                                                             |
| ✅  | suokif          |                                                                             |
| ✅  | telephone       |                                                                             |
| ✅  | tiny            |                                                                             |
| ✅  | tinyc           |                                                                             |
| ✅  | tnsnames        |                                                                             |
| ✅  | tnt             |                                                                             |
| ✅  | unicodeclasses  |                                                                             |
| ✅  | useragent       |                                                                             |
| ✅  | wavefrontobj    |                                                                             |
| ✅  | xml             |                                                                             |
| ✅  | xpath           |                                                                             |
| ❌  | altpython3      | build: altpython3/altpython3_lexer.go:812:13: too many errors               |
| ❌  | antlrv3         | antlr: error(134): ANTLRv3.g4:32:102: symbol action conflicts with generated code in target language or runtime |
| ❌  | antlrv4         | build: antlrv4/antlrv4_lexer.go:506:2: undefined: LexerAdaptor              |
| ❌  | apex            | antlr: error(134): apex.g4:544:30: symbol type conflicts with generated code in target language or runtime |
| ❌  | asm6502         | antlr: error(134): asm6502.g4:71:30: symbol string conflicts with generated code in target language or runtime |
| ❌  | asn             | antlr: error(134): ASN.g4:166:2: symbol type conflicts with generated code in target language or runtime |
| ❌  | aspectj         | antlr: error(56): AspectJLexer.g4:382:21: reference to undefined rule: Identifier |
| ❌  | calculator      | antlr: error(134): calculator.g4:54:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | classify        | test: FAIL   bramp.net/antlr4/classify       0.021s                                 |
| ❌  | clojure         | antlr: error(134): Clojure.g4:99:12: symbol map conflicts with generated code in target language or runtime |
| ❌  | cpp14           | build: cpp14/cpp14_parser.go:27309:6: syntax error: unexpected NewEmptyBaseclauseContext, expecting ( |
| ❌  | csharp          | build: csharp/csharp_lexer.go:11:12: expected 'STRING', found '.'           |
| ❌  | css3            | antlr: error(134): css3.g4:196:6: symbol var conflicts with generated code in target language or runtime |
| ❌  | dicenotation    | antlr: warning(125): DiceNotation.g4:50:7: implicit definition of token DSEPARATOR in parser |
| ❌  | erlang          | antlr: error(134): Erlang.g4:156:36: symbol type conflicts with generated code in target language or runtime |
| ❌  | fortran77       | antlr: error(134): fortran77.g4:89:6: symbol type conflicts with generated code in target language or runtime |
| ❌  | gff3            | antlr: error(134): gff3.g4:45:28: symbol type conflicts with generated code in target language or runtime |
| ❌  | golang          | antlr: error(134): Golang.g4:151:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | graphql         | antlr: error(134): GraphQL.g4:140:9: symbol type conflicts with generated code in target language or runtime |
| ❌  | html            | test: FAIL   bramp.net/antlr4/html   30.028s                                    |
| ❌  | icalendar       | antlr: error(134): ICalendar.g4:305:3: symbol action conflicts with generated code in target language or runtime |
| ❌  | inf             | antlr: error(134): inf.g4:57:17: symbol string conflicts with generated code in target language or runtime |
| ❌  | informix        | antlr: error(134): informix.g4:601:266: symbol string conflicts with generated code in target language or runtime |
| ❌  | java            | test: FAIL   bramp.net/antlr4/java   30.071s                                    |
| ❌  | java8           | antlr: error(134): Java8.g4:73:0: symbol type conflicts with generated code in target language or runtime |
| ❌  | java9           | antlr: error(134): Java9.g4:84:0: symbol type conflicts with generated code in target language or runtime |
| ❌  | javadoc         | build: javadoc/javadoc_lexer.go:167:10: undefined: _input                   |
| ❌  | javascript      | build: javascript/javascript_lexer.go:635:2: undefined: JavaScriptBaseLexer |
| ❌  | jvmbasic        | antlr: error(134): jvmBasic.g4:481:8: symbol var conflicts with generated code in target language or runtime |
| ❌  | kotlin          | antlr: error(134): KotlinParser.g4:261:53: symbol type conflicts with generated code in target language or runtime |
| ❌  | krl             | antlr: error(134): krl.g4:80:68: symbol type conflicts with generated code in target language or runtime |
| ❌  | logo            | antlr: error(134): logo.g4:128:38: symbol func conflicts with generated code in target language or runtime |
| ❌  | lolcode         | antlr: error(134): lolcode.g4:111:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | lua             | antlr: error(134): Lua.g4:90:15: symbol var conflicts with generated code in target language or runtime |
| ❌  | m2pim4          | antlr: error(134): m2pim4.g4:356:22: symbol type conflicts with generated code in target language or runtime |
| ❌  | masm            | build: masm/masm_lexer.go:12:4: syntax error: non-declaration statement outside function body |
| ❌  | mysql           | build:       previous declaration at mysql/mysql_parser.go:11965:6               |
| ❌  | pascal          | antlr: error(134): pascal.g4:115:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | pddl            | antlr: error(134): Pddl.g4:92:25: symbol type conflicts with generated code in target language or runtime |
| ❌  | pdn             | antlr: error(134): pdn.g4:44:14: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdp7            | antlr: error(134): pdp7.g4:87:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | pgn             | build: pgn/pgn_lexer.go:191:10: undefined: getCharPositionInLine            |
| ❌  | php             | build: php/php_lexer.go:1589:4: too many errors                             |
| ❌  | plsql           | antlr: error(134): PlSqlParser.g4:844:41: symbol range conflicts with generated code in target language or runtime |
| ❌  | prolog          | antlr: error(134): prolog.g4:72:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | protobuf3       | antlr: error(134): Protobuf3.g4:138:19: symbol range conflicts with generated code in target language or runtime |
| ❌  | python2         | build: python2/python2_lexer.go:639:14: too many errors                     |
| ❌  | python3         | build: python3/python3_lexer.go:828:13: too many errors                     |
| ❌  | quakemap        | antlr: error(134): quakemap.g4:44:12: symbol string conflicts with generated code in target language or runtime |
| ❌  | scala           | antlr: error(134): Scala.g4:433:17: symbol type conflicts with generated code in target language or runtime |
| ❌  | smalltalk       | antlr: error(134): Smalltalk.g4:50:57: symbol string conflicts with generated code in target language or runtime |
| ❌  | smtlibv2        | antlr: error(134): SMTLIBv2.g4:1089:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | sparql          | antlr: error(134): Sparql.g4:57:45: symbol var conflicts with generated code in target language or runtime |
| ❌  | sqlite          | antlr: error(134): SQLite.g4:34:21: symbol error conflicts with generated code in target language or runtime |
| ❌  | st              | build: st/st_lexer.go:779:2: syntax error: unexpected default, expecting :  |
| ❌  | stg             | build: stg/stg_lexer.go:716:2: syntax error: unexpected default, expecting : |
| ❌  | swift2          | antlr: error(134): Swift2.g4:783:34: symbol type conflicts with generated code in target language or runtime |
| ❌  | swift3          | antlr: error(134): Swift3.g4:664:16: symbol type conflicts with generated code in target language or runtime |
| ❌  | swiftfin        | antlr: error(134): SwiftFinParser.g4:61:16: symbol map conflicts with generated code in target language or runtime |
| ❌  | tsql            | build:       previous declaration at tsql/tsql_parser.go:119031:8                |
| ❌  | turtle          | build:               want String([]string, antlr.RuleContext) string                    |
| ❌  | ucblogo         | build: ucblogo/ucblogo_parser.go:14:14: expected 'STRING', found '.'        |
| ❌  | url             | antlr: error(134): url.g4:95:18: symbol string conflicts with generated code in target language or runtime |
| ❌  | vba             | antlr: error(134): vba.g4:569:44: symbol type conflicts with generated code in target language or runtime |
| ❌  | verilog2001     | antlr: error(134): Verilog2001.g4:1944:24: symbol range conflicts with generated code in target language or runtime |
| ❌  | vhdl            | test: FAIL   bramp.net/antlr4/vhdl   30.069s                                    |
| ❌  | visualbasic6    | antlr: error(134): VisualBasic6.g4:555:35: symbol type conflicts with generated code in target language or runtime |
| ❌  | webidl          | antlr: error(134): WebIDL.g4:584:3: symbol type conflicts with generated code in target language or runtime |
| ❌  | xdr             | test: FAIL   bramp.net/antlr4/xdr [build failed]                              |
| ❌  | z               | build: z/z_lexer.go:1231:2: syntax error: non-declaration statement outside function body |

The failures are broken down like so:

* **antlr** - ANTLR failed to generate Go code from the grammar.
* **build** - The generated Go code failed to build.
* **test**  - The generated Go code failed the unit tests for that language.

If you wish to help fix the situation then please submit fixes back to the [ANTLR Go target](https://github.com/antlr/antlr4/blob/master/tool/src/org/antlr/v4/codegen/target/GoTarget.java), or the [Gramamrs Github Repo](https://github.com/antlr/grammars-v4).

# Build

To generate, build and test a single grammar, just run:

```bash
make <name of grammar>
```

## To update all the grammars

```bash
# Fetch the latest set of grammars
git submodule init
git submodule update

# Run make and all the magic will happen
make

# Update the table in the README.md with the output
```

## Licence (Apache 2)

Each grammar has its [own licence](https://github.com/antlr/grammars-v4#license), but the compiled code is licenced under Apache 2.

*This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.*

```
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
