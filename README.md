# antlr4 Go parsers [![Build Status](https://img.shields.io/travis/bramp/antlr4-grammars.svg)](https://travis-ci.org/bramp/antlr4-grammars) [![Coverage](https://img.shields.io/coveralls/bramp/antlr4-grammars.svg)](https://coveralls.io/github/bramp/antlr4-grammars) [![Report card](https://goreportcard.com/badge/bramp.net/antlr4)](https://goreportcard.com/report/bramp.net/antlr4) [![GoDoc](https://godoc.org/bramp.net/antlr4?status.svg)](https://godoc.org/bramp.net/antlr4)
Compiled by Andrew Brampton ([bramp.net](https://bramp.net))

Precompiled Go parsers of many of the grammars on [github.com/antlr/grammars-v4](http://github.com/antlr/grammars-v4).

The Antlr's Go Target is still a work in progress. As such, many of the grammars fail to compile, or only pass simple tests. To report issues with the grammar [go here](https://github.com/antlr/grammars-v4), to report issues with Antlr's Go Target [go here](https://github.com/antlr/antlr4).

## Quick Start

Just import one of the parser listed in the table below:

```go
import (
	"bramp.net/antlr4/<name of grammar>"
)
```

Then the Antlr Lexer, Parser, and Listeners are available. More information on each Grammar's API is found on the [godocs](https://godoc.org/bramp.net/antlr4) with examples.

## Full Example (using the json parser)

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
| ✅  | asn             |                                                                             |
| ✅  | atl             |                                                                             |
| ✅  | b               |                                                                             |
| ✅  | bnf             |                                                                             |
| ✅  | brainfuck       |                                                                             |
| ✅  | c               |                                                                             |
| ✅  | clf             |                                                                             |
| ✅  | clif            |                                                                             |
| ✅  | clu             |                                                                             |
| ✅  | cmake           |                                                                             |
| ✅  | cobol85         |                                                                             |
| ✅  | cobol85preprocessor |                                                                             |
| ✅  | cookie          |                                                                             |
| ✅  | cool            |                                                                             |
| ✅  | corundum        |                                                                             |
| ✅  | creole          |                                                                             |
| ✅  | csv             |                                                                             |
| ✅  | dart2           |                                                                             |
| ✅  | databank        |                                                                             |
| ✅  | datetime        |                                                                             |
| ✅  | dcm_2_0_grammar |                                                                             |
| ✅  | dgs             |                                                                             |
| ✅  | dot             |                                                                             |
| ✅  | ecmascript      |                                                                             |
| ✅  | emailaddress    |                                                                             |
| ✅  | fasta           |                                                                             |
| ✅  | fen             |                                                                             |
| ✅  | fol             |                                                                             |
| ✅  | fusiontablessql |                                                                             |
| ✅  | gml             |                                                                             |
| ✅  | graphemes       |                                                                             |
| ✅  | gtin            |                                                                             |
| ✅  | guido           |                                                                             |
| ✅  | http            |                                                                             |
| ✅  | icon            |                                                                             |
| ✅  | idl             |                                                                             |
| ✅  | iri             |                                                                             |
| ✅  | istc            |                                                                             |
| ✅  | jpa             |                                                                             |
| ✅  | json            |                                                                             |
| ✅  | lambda          |                                                                             |
| ✅  | lcc             |                                                                             |
| ✅  | less            |                                                                             |
| ✅  | lexunicode      |                                                                             |
| ✅  | matlab          |                                                                             |
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
| ✅  | pl0             |                                                                             |
| ✅  | postalcode      |                                                                             |
| ✅  | powerbuilder    |                                                                             |
| ✅  | prolog          |                                                                             |
| ✅  | propcalc        |                                                                             |
| ✅  | properties      |                                                                             |
| ✅  | prov_n          |                                                                             |
| ✅  | r               |                                                                             |
| ✅  | rcs             |                                                                             |
| ✅  | redcode         |                                                                             |
| ✅  | regex           |                                                                             |
| ✅  | restructuredtext |                                                                             |
| ✅  | robotwar        |                                                                             |
| ✅  | romannumerals   |                                                                             |
| ✅  | rpn             |                                                                             |
| ✅  | scss            |                                                                             |
| ✅  | sexpression     |                                                                             |
| ✅  | sharc           |                                                                             |
| ✅  | smiles          |                                                                             |
| ✅  | snobol          |                                                                             |
| ✅  | solidity        |                                                                             |
| ✅  | stacktrace      |                                                                             |
| ✅  | suokif          |                                                                             |
| ✅  | telephone       |                                                                             |
| ✅  | tiny            |                                                                             |
| ✅  | tinybasic       |                                                                             |
| ✅  | tinyc           |                                                                             |
| ✅  | tnsnames        |                                                                             |
| ✅  | tnt             |                                                                             |
| ✅  | tsv             |                                                                             |
| ✅  | unicodeclasses  |                                                                             |
| ✅  | upnp            |                                                                             |
| ✅  | useragent       |                                                                             |
| ✅  | wavefrontobj    |                                                                             |
| ✅  | wkt             |                                                                             |
| ✅  | xml             |                                                                             |
| ❌  | algol60         | antlr: error(134): algol60.g4:155:5: symbol type conflicts with generated code in target language or runtime |
| ❌  | altpython3      | build: altpython3/altpython3_lexer.go:812:13: too many errors               |
| ❌  | antlrv3         | antlr: error(134): ANTLRv3.g4:31:102: symbol action conflicts with generated code in target language or runtime |
| ❌  | antlrv4         | build: antlrv4/antlrv4_lexer.go:506:2: undefined: LexerAdaptor              |
| ❌  | antlrv4lexerpythontarget | build: antlrv4lexerpythontarget/antlrv4lexerpythontarget_lexer.go:12:1: syntax error: non-declaration statement outside function body |
| ❌  | apex            | antlr: error(134): apex.g4:544:30: symbol type conflicts with generated code in target language or runtime |
| ❌  | asm6502         | antlr: error(134): asm6502.g4:71:30: symbol string conflicts with generated code in target language or runtime |
| ❌  | asm8080         | antlr: error(134): asm8080.g4:91:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | asm8086         | antlr: error(134): asm8086.g4:167:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | asmmasm         | antlr: error(134): asmMASM.g4:68:6: symbol type conflicts with generated code in target language or runtime |
| ❌  | asmz80          | antlr: error(134): asmZ80.g4:91:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | asn_3gpp        | build: asn_3gpp/asn_3gpp_lexer.go:680:10: undefined: getCharPositionInLine  |
| ❌  | aspectj         | antlr: error(56): /Users/bramp/go/src/bramp.net/antlr4/grammars-v4/aspectj/AspectJLexer.g4:382:21: reference to undefined rule: Identifier |
| ❌  | calculator      | antlr: error(134): calculator.g4:54:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | capnproto       | antlr: error(134): CapnProto.g4:85:18: symbol type conflicts with generated code in target language or runtime |
| ❌  | classify        | test: FAIL	bramp.net/antlr4/classify	0.058s                                 |
| ❌  | clojure         | antlr: error(134): Clojure.g4:99:12: symbol map conflicts with generated code in target language or runtime |
| ❌  | cpp14           | build: cpp14/cpp14_parser.go:28332:2: syntax error: unexpected var at end of statement |
| ❌  | cql             | antlr: error(134): CqlParser.g4:339:44: symbol type conflicts with generated code in target language or runtime |
| ❌  | csharp          | build: csharp/csharp_lexer.go:11:12: expected 'STRING', found '.'           |
| ❌  | css3            | antlr: error(134): css3.g4:132:65: symbol String conflicts with generated code in target language or runtime |
| ❌  | cto             | test: FAIL	bramp.net/antlr4/cto	0.066s                                      |
| ❌  | dicenotation    | antlr: warning(125): DiceNotation.g4:56:23: implicit definition of token DSEPARATOR in parser |
| ❌  | edif300         | antlr: error(134): EDIF300.g4:3082:24: symbol string conflicts with generated code in target language or runtime |
| ❌  | erlang          | antlr: error(134): Erlang.g4:156:36: symbol type conflicts with generated code in target language or runtime |
| ❌  | flatbuffers     | antlr: error(134): FlatBuffers.g4:16:33: symbol type conflicts with generated code in target language or runtime |
| ❌  | fortran77       | antlr: error(134): Fortran77Parser.g4:81:6: symbol type conflicts with generated code in target language or runtime |
| ❌  | gff3            | antlr: error(134): gff3.g4:45:28: symbol type conflicts with generated code in target language or runtime |
| ❌  | golang          | antlr: error(134): Golang.g4:210:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | graphql         | antlr: error(134): GraphQL.g4:140:9: symbol type conflicts with generated code in target language or runtime |
| ❌  | html            | test: FAIL	bramp.net/antlr4/html	30.072s                                    |
| ❌  | hypertalk       | antlr: error(134): HyperTalk.g4:281:20: symbol range conflicts with generated code in target language or runtime |
| ❌  | icalendar       | antlr: error(134): ICalendar.g4:284:5: symbol action conflicts with generated code in target language or runtime |
| ❌  | inf             | antlr: error(134): inf.g4:57:17: symbol string conflicts with generated code in target language or runtime |
| ❌  | informix        | antlr: error(134): informix.g4:601:266: symbol string conflicts with generated code in target language or runtime |
| ❌  | java            | test: FAIL	bramp.net/antlr4/java	30.124s                                    |
| ❌  | java8           | build: java8/java8_lexer.go:768:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | java9           | build: java9/java9_lexer.go:825:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | javadoc         | build: javadoc/javadoc_lexer.go:167:10: undefined: _input                   |
| ❌  | javascript      | build: javascript/javascript_lexer.go:631:2: undefined: JavaScriptBaseLexer |
| ❌  | jvmbasic        | antlr: error(134): jvmBasic.g4:481:8: symbol var conflicts with generated code in target language or runtime |
| ❌  | kotlin          | antlr: error(134): KotlinParser.g4:325:75: symbol type conflicts with generated code in target language or runtime |
| ❌  | krl             | antlr: error(134): krl.g4:80:68: symbol type conflicts with generated code in target language or runtime |
| ❌  | logo            | antlr: error(134): logo.g4:128:38: symbol func conflicts with generated code in target language or runtime |
| ❌  | lolcode         | antlr: error(134): lolcode.g4:111:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | lpc             | antlr: error(134): LPC.g4:671:20: symbol String conflicts with generated code in target language or runtime |
| ❌  | lua             | antlr: error(134): Lua.g4:90:15: symbol var conflicts with generated code in target language or runtime |
| ❌  | m2pim4          | antlr: error(134): m2pim4.g4:356:22: symbol type conflicts with generated code in target language or runtime |
| ❌  | masm            | antlr: error(134): MASM.g4:130:31: symbol String conflicts with generated code in target language or runtime |
| ❌  | moo             | antlr: error(134): moo.g4:205:5: symbol real conflicts with generated code in target language or runtime |
| ❌  | muddb           | antlr: error(134): muddb.g4:104:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | mysql           | test: FAIL	bramp.net/antlr4/mysql	30.375s                                   |
| ❌  | pascal          | antlr: error(134): pascal.g4:120:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | pddl            | antlr: error(134): Pddl.g4:92:25: symbol type conflicts with generated code in target language or runtime |
| ❌  | pdn             | antlr: error(134): pdn.g4:44:14: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdp7            | antlr: error(134): pdp7.g4:87:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | pgn             | build: pgn/pgn_lexer.go:191:10: undefined: getCharPositionInLine            |
| ❌  | php             | build: php/php_lexer.go:1591:4: too many errors                             |
| ❌  | pike            | antlr: error(119): pike.g4::: The following sets of rules are mutually left-recursive [expression6, call, index, arrow] |
| ❌  | plsql           | build: plsql/plsql_lexer.go:15793:2: undefined: PlSqlBaseLexer              |
| ❌  | ply             | antlr: error(134): ply.g4:65:33: symbol string conflicts with generated code in target language or runtime |
| ❌  | protobuf3       | antlr: error(134): Protobuf3.g4:138:19: symbol range conflicts with generated code in target language or runtime |
| ❌  | python2         | build: python2/python2_lexer.go:635:15: too many errors                     |
| ❌  | python3         | build: python3/python3_lexer.go:828:13: too many errors                     |
| ❌  | quakemap        | antlr: error(134): quakemap.g4:44:12: symbol string conflicts with generated code in target language or runtime |
| ❌  | rego            | antlr: error(134): RegoLexer.g4:40:0: symbol String conflicts with generated code in target language or runtime |
| ❌  | rexx            | build: rexx/rexx_lexer.go:877:2: syntax error: unexpected default, expecting } |
| ❌  | scala           | antlr: error(134): Scala.g4:434:17: symbol type conflicts with generated code in target language or runtime |
| ❌  | sgf             | antlr: error(134): sgf.g4:25:20: symbol go conflicts with generated code in target language or runtime |
| ❌  | smalltalk       | antlr: error(134): Smalltalk.g4:50:57: symbol string conflicts with generated code in target language or runtime |
| ❌  | smtlibv2        | antlr: error(134): SMTLIBv2.g4:1089:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | sparql          | antlr: error(134): Sparql.g4:57:45: symbol var conflicts with generated code in target language or runtime |
| ❌  | sqlite          | antlr: error(134): SQLite.g4:34:21: symbol error conflicts with generated code in target language or runtime |
| ❌  | st              | build: st/st_lexer.go:777:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | stg             | build: stg/stg_lexer.go:714:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | swift2          | antlr: error(134): Swift2.g4:783:34: symbol type conflicts with generated code in target language or runtime |
| ❌  | swift3          | antlr: error(134): Swift3.g4:664:16: symbol type conflicts with generated code in target language or runtime |
| ❌  | swiftfin        | antlr: error(134): SwiftFinParser.g4:61:16: symbol map conflicts with generated code in target language or runtime |
| ❌  | sysveriloghdl   | maketest: exit status 1                                                     |
| ❌  | thrift          | antlr: error(134): Thrift.g4:28:49: symbol struct conflicts with generated code in target language or runtime |
| ❌  | tinymud         | antlr: error(134): tinymud.g4:40:16: symbol action conflicts with generated code in target language or runtime |
| ❌  | tjs             | antlr: error(134): TJSLexer.g4:132:0: symbol String conflicts with generated code in target language or runtime |
| ❌  | toml            | antlr: error(134): toml.g4:25:44: symbol bool conflicts with generated code in target language or runtime |
| ❌  | tsql            | test: FAIL	bramp.net/antlr4/tsql	30.370s                                    |
| ❌  | turtle          | antlr: error(134): TURTLE.g4:112:5: symbol String conflicts with generated code in target language or runtime |
| ❌  | ucblogo         | build: ucblogo/ucblogo_parser.go:14:14: expected 'STRING', found '.'        |
| ❌  | url             | antlr: error(134): url.g4:95:18: symbol string conflicts with generated code in target language or runtime |
| ❌  | v               | antlr: error(134): V.g4:751:31: symbol type conflicts with generated code in target language or runtime |
| ❌  | vba             | antlr: error(134): vba.g4:569:44: symbol type conflicts with generated code in target language or runtime |
| ❌  | verilog2001     | antlr: error(134): Verilog2001.g4:1488:5: symbol String conflicts with generated code in target language or runtime |
| ❌  | vhdl            | test: FAIL	bramp.net/antlr4/vhdl	30.197s                                    |
| ❌  | visualbasic6    | antlr: error(134): VisualBasic6.g4:554:35: symbol type conflicts with generated code in target language or runtime |
| ❌  | wat             | antlr: error(134): WatLexer.g4:39:9: symbol String conflicts with generated code in target language or runtime |
| ❌  | webidl          | antlr: error(134): WebIDL.g4:584:3: symbol type conflicts with generated code in target language or runtime |
| ❌  | xdr             | test: FAIL	bramp.net/antlr4/xdr [build failed]                              |
| ❌  | xpath           | build: 	previous declaration at xpath/xpath_parser.go:3504:6                |
| ❌  | z               | build: z/z_lexer.go:1240:2: syntax error: non-declaration statement outside function body |

The failures are broken down like so:

* **antlr** - ANTLR failed to generate Go code from the grammar.
* **build** - The generated Go code failed to build.
* **test**  - The generated Go code failed the unit tests for that language.

If you wish to help fix the situation then please submit fixes back to the [ANTLR Go target](https://github.com/antlr/antlr4/blob/master/tool/src/org/antlr/v4/codegen/target/GoTarget.java), or the [Gramamers Github Repo](https://github.com/antlr/grammars-v4).

# Build

The grammars above are ready to use, but if you wish to change them, or build them yourself for any reason just run:

```bash
make <name of grammar>
```

## To update all the grammars

```bash
# Ensure antlr is installed
mvn dependency:get -Dartifact=org.antlr:antlr4:4.7.2:jar:complete
go get -u github.com/antlr/antlr4/runtime/Go/antlr

# Ensure some other deps are available
go get -u github.com/iancoleman/strcase

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
