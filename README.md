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
| ✅  | abb             |                                                                             |
| ✅  | abnf            |                                                                             |
| ✅  | agc             |                                                                             |
| ✅  | alef            |                                                                             |
| ✅  | alloy           |                                                                             |
| ✅  | alpaca          |                                                                             |
| ✅  | angelscript     |                                                                             |
| ✅  | apt             |                                                                             |
| ✅  | argus           |                                                                             |
| ✅  | arithmetic      |                                                                             |
| ✅  | asl             |                                                                             |
| ✅  | asm6502         |                                                                             |
| ✅  | asm8080         |                                                                             |
| ✅  | asm8086         |                                                                             |
| ✅  | asmz80          |                                                                             |
| ✅  | asn             |                                                                             |
| ✅  | atl             |                                                                             |
| ✅  | b               |                                                                             |
| ✅  | bcl             |                                                                             |
| ✅  | bcpl            |                                                                             |
| ✅  | bdf             |                                                                             |
| ✅  | beep            |                                                                             |
| ✅  | bibcode         |                                                                             |
| ✅  | bnf             |                                                                             |
| ✅  | brainflak       |                                                                             |
| ✅  | brainfuck       |                                                                             |
| ✅  | bsl             |                                                                             |
| ✅  | c               |                                                                             |
| ✅  | calculator      |                                                                             |
| ✅  | callable_       |                                                                             |
| ✅  | cayenne         |                                                                             |
| ✅  | classify        |                                                                             |
| ✅  | clf             |                                                                             |
| ✅  | clojure         |                                                                             |
| ✅  | clu             |                                                                             |
| ✅  | cmake           |                                                                             |
| ✅  | cobol85         |                                                                             |
| ✅  | cobol85preprocessor |                                                                             |
| ✅  | cookie          |                                                                             |
| ✅  | cool            |                                                                             |
| ✅  | corundum        |                                                                             |
| ✅  | cql             |                                                                             |
| ✅  | creole          |                                                                             |
| ✅  | css3            |                                                                             |
| ✅  | csv             |                                                                             |
| ✅  | ctl             |                                                                             |
| ✅  | dart2           |                                                                             |
| ✅  | databank        |                                                                             |
| ✅  | datetime        |                                                                             |
| ✅  | dcm_2_0_grammar |                                                                             |
| ✅  | dgol            |                                                                             |
| ✅  | dgs             |                                                                             |
| ✅  | dif             |                                                                             |
| ✅  | doiurl          |                                                                             |
| ✅  | domain          |                                                                             |
| ✅  | dot             |                                                                             |
| ✅  | ecmascript      |                                                                             |
| ✅  | emailaddress    |                                                                             |
| ✅  | erlang          |                                                                             |
| ✅  | fasta           |                                                                             |
| ✅  | fen             |                                                                             |
| ✅  | filter          |                                                                             |
| ✅  | flowmatic       |                                                                             |
| ✅  | focal           |                                                                             |
| ✅  | fol             |                                                                             |
| ✅  | fusiontablessql |                                                                             |
| ✅  | gedcom          |                                                                             |
| ✅  | gff3            |                                                                             |
| ✅  | gml             |                                                                             |
| ✅  | graphemes       |                                                                             |
| ✅  | graphql         |                                                                             |
| ✅  | gtin            |                                                                             |
| ✅  | guido           |                                                                             |
| ✅  | http            |                                                                             |
| ✅  | icon            |                                                                             |
| ✅  | idl             |                                                                             |
| ✅  | infosapient     |                                                                             |
| ✅  | iri             |                                                                             |
| ✅  | isl             |                                                                             |
| ✅  | iso8601         |                                                                             |
| ✅  | istc            |                                                                             |
| ✅  | itn             |                                                                             |
| ✅  | janus           |                                                                             |
| ✅  | java8           |                                                                             |
| ✅  | joss            |                                                                             |
| ✅  | jpa             |                                                                             |
| ✅  | json            |                                                                             |
| ✅  | json5           |                                                                             |
| ✅  | karel           |                                                                             |
| ✅  | krl             |                                                                             |
| ✅  | lambda          |                                                                             |
| ✅  | lark            |                                                                             |
| ✅  | lcc             |                                                                             |
| ✅  | less            |                                                                             |
| ✅  | lexunicode      |                                                                             |
| ✅  | limbo           |                                                                             |
| ✅  | lisa            |                                                                             |
| ✅  | lolcode         |                                                                             |
| ✅  | loop            |                                                                             |
| ✅  | lrc             |                                                                             |
| ✅  | ltl             |                                                                             |
| ✅  | matlab          |                                                                             |
| ✅  | mdx             |                                                                             |
| ✅  | memcached_protocol |                                                                             |
| ✅  | metamath        |                                                                             |
| ✅  | metric          |                                                                             |
| ✅  | microc          |                                                                             |
| ✅  | modelica        |                                                                             |
| ✅  | molecule        |                                                                             |
| ✅  | morsecode       |                                                                             |
| ✅  | mps             |                                                                             |
| ✅  | mu              |                                                                             |
| ✅  | mumath          |                                                                             |
| ✅  | mumps           |                                                                             |
| ✅  | nanofuck        |                                                                             |
| ✅  | newick          |                                                                             |
| ✅  | objectivec      |                                                                             |
| ✅  | oncrpcv2        |                                                                             |
| ✅  | orwell          |                                                                             |
| ✅  | p               |                                                                             |
| ✅  | pcre            |                                                                             |
| ✅  | pddl            |                                                                             |
| ✅  | peoplecode      |                                                                             |
| ✅  | pii             |                                                                             |
| ✅  | pike            |                                                                             |
| ✅  | pl0             |                                                                             |
| ✅  | plucid          |                                                                             |
| ✅  | pmmn            |                                                                             |
| ✅  | postalcode      |                                                                             |
| ✅  | powerbuilderdw  |                                                                             |
| ✅  | powerquery      |                                                                             |
| ✅  | prolog          |                                                                             |
| ✅  | promql          |                                                                             |
| ✅  | propcalc        |                                                                             |
| ✅  | properties      |                                                                             |
| ✅  | protobuf3       |                                                                             |
| ✅  | prov_n          |                                                                             |
| ✅  | qif             |                                                                             |
| ✅  | r               |                                                                             |
| ✅  | rcs             |                                                                             |
| ✅  | redcode         |                                                                             |
| ✅  | refal           |                                                                             |
| ✅  | regex           |                                                                             |
| ✅  | restructuredtext |                                                                             |
| ✅  | robotwar        |                                                                             |
| ✅  | romannumerals   |                                                                             |
| ✅  | rpn             |                                                                             |
| ✅  | scala           |                                                                             |
| ✅  | scotty          |                                                                             |
| ✅  | scss            |                                                                             |
| ✅  | sexpression     |                                                                             |
| ✅  | sgf             |                                                                             |
| ✅  | sharc           |                                                                             |
| ✅  | sici            |                                                                             |
| ✅  | sickbay         |                                                                             |
| ✅  | smiles          |                                                                             |
| ✅  | snobol          |                                                                             |
| ✅  | snowball        |                                                                             |
| ✅  | solidity        |                                                                             |
| ✅  | sqlite          |                                                                             |
| ✅  | stacktrace      |                                                                             |
| ✅  | stellaris       |                                                                             |
| ✅  | stl             |                                                                             |
| ✅  | suokif          |                                                                             |
| ✅  | systemverilog   |                                                                             |
| ✅  | teal            |                                                                             |
| ✅  | telephone       |                                                                             |
| ✅  | tiny            |                                                                             |
| ✅  | tinybasic       |                                                                             |
| ✅  | tinyc           |                                                                             |
| ✅  | tl              |                                                                             |
| ✅  | tnsnames        |                                                                             |
| ✅  | tnt             |                                                                             |
| ✅  | tsv             |                                                                             |
| ✅  | unicodeclasses  |                                                                             |
| ✅  | upnp            |                                                                             |
| ✅  | useragent       |                                                                             |
| ✅  | vba             |                                                                             |
| ✅  | verilog         |                                                                             |
| ✅  | vhdl            |                                                                             |
| ✅  | vmf             |                                                                             |
| ✅  | wat             |                                                                             |
| ✅  | wavefrontobj    |                                                                             |
| ✅  | webidl          |                                                                             |
| ✅  | wkt             |                                                                             |
| ✅  | wln             |                                                                             |
| ✅  | xml             |                                                                             |
| ❌  | acme            | test: FAIL                                                                  |
| ❌  | algol60         | antlr: error(134): algol60.g4:144:19: symbol String conflicts with generated code in target language or runtime |
| ❌  | altpython3      | build: altpython3/altpython3_lexer.go:816:13: too many errors               |
| ❌  | antlrv2         | build: antlrv2/antlrv2_lexer.go:591:2: undefined: LexerAdaptor              |
| ❌  | antlrv3         | build: antlrv3/antlrv3_lexer.go:545:2: undefined: LexerAdaptor              |
| ❌  | antlrv4         | build: antlrv4/antlrv4_lexer.go:410:2: undefined: LexerAdaptor              |
| ❌  | apex            | build: apex/apex_lexer.go:930:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | asmmasm         | antlr: error(134): asmMASM.g4:182:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | asn_3gpp        | build: asn_3gpp/asn_3gpp_lexer.go:682:10: undefined: getCharPositionInLine  |
| ❌  | aspectj         | antlr: error(56): /Users/bramp/personal/antlr4-grammars/grammars-v4/aspectj/AspectJLexer.g4:382:21: reference to undefined rule: Identifier |
| ❌  | capnproto       | test: FAIL                                                                  |
| ❌  | clif            | test: FAIL                                                                  |
| ❌  | cpp14           | build: 	/Users/bramp/personal/antlr4-grammars/cpp14/cpp14_parser.go:26954:49: previous declaration |
| ❌  | csharp          | build: csharp/csharp_lexer.go:1194:2: undefined: CSharpLexerBase            |
| ❌  | cto             | test: FAIL                                                                  |
| ❌  | dicenotation    | test: FAIL                                                                  |
| ❌  | edif300         | antlr: error(134): EDIF300.g4:3082:24: symbol string conflicts with generated code in target language or runtime |
| ❌  | edn             | test: FAIL                                                                  |
| ❌  | file.           | build: malformed import path "bramp.net/antlr4/file.": trailing dot in path element |
| ❌  | flatbuffers     | test: FAIL                                                                  |
| ❌  | fortran77       | build: fortran77/fortran77_lexer.go:1070:10: undefined: getCharPositionInLine |
| ❌  | gdscript        | build: gdscript/gdscript_lexer.go:543:2: undefined: GDScriptLexerBase       |
| ❌  | go              | build: go/go_lexer.go:3:9: expected 'IDENT', found 'go'                     |
| ❌  | guitartab       | antlr: error(134): guitartab.g4:35:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | haskell         | build: haskell/haskell_lexer.go:871:2: undefined: HaskellBaseLexer          |
| ❌  | hive            | build: 	/Users/bramp/personal/antlr4-grammars/hive/hive_parser.go:49454:6: previous declaration |
| ❌  | html            | test: FAIL                                                                  |
| ❌  | hypertalk       | antlr: error(134): HyperTalk.g4:281:20: symbol range conflicts with generated code in target language or runtime |
| ❌  | icalendar       | antlr: error(134): ICalendar.g4:284:5: symbol action conflicts with generated code in target language or runtime |
| ❌  | inf             | antlr: error(134): inf.g4:57:17: symbol string conflicts with generated code in target language or runtime |
| ❌  | informix        | antlr: error(134): informix.g4:601:266: symbol string conflicts with generated code in target language or runtime |
| ❌  | java            | test: FAIL                                                                  |
| ❌  | java9           | build: java9/java9_lexer.go:635:2: undefined: Java9LexerBase                |
| ❌  | javadoc         | build: javadoc/javadoc_lexer.go:169:10: undefined: _input                   |
| ❌  | javascript      | build: javascript/javascript_lexer.go:894:2: undefined: JavaScriptLexerBase |
| ❌  | jvmbasic        | test: FAIL                                                                  |
| ❌  | kotlin          | antlr: error(134): KotlinParser.g4:325:75: symbol type conflicts with generated code in target language or runtime |
| ❌  | logo            | antlr: error(134): logo.g4:60:5: symbol make conflicts with generated code in target language or runtime |
| ❌  | lpc             | antlr: error(134): LPC.g4:671:20: symbol String conflicts with generated code in target language or runtime |
| ❌  | lua             | antlr: error(134): Lua.g4:112:6: symbol string conflicts with generated code in target language or runtime |
| ❌  | m2pim4          | antlr: error(134): m2pim4.g4:243:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | masm            | antlr: error(134): MASM.g4:127:31: symbol String conflicts with generated code in target language or runtime |
| ❌  | mckeemanform    | antlr: error(134): McKeemanForm.g4:43:6: symbol String conflicts with generated code in target language or runtime |
| ❌  | moo             | antlr: error(134): moo.g4:205:5: symbol real conflicts with generated code in target language or runtime |
| ❌  | muddb           | antlr: error(134): muddb.g4:104:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | mysql           | maketest: exit status 1                                                     |
| ❌  | oberon          | antlr: error(134): oberon.g4:65:5: symbol real conflicts with generated code in target language or runtime |
| ❌  | pascal          | antlr: error(134): pascal.g4:367:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdn             | antlr: error(134): pdn.g4:44:14: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdp7            | antlr: error(134): pdp7.g4:87:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | pgn             | build: pgn/pgn_lexer.go:193:10: undefined: getCharPositionInLine            |
| ❌  | php             | build: php/php_lexer.go:1159:2: undefined: PhpLexerBase                     |
| ❌  | plsql           | build: plsql/plsql_lexer.go:15810:2: undefined: PlSqlLexerBase              |
| ❌  | ply             | antlr: error(134): ply.g4:65:35: symbol string conflicts with generated code in target language or runtime |
| ❌  | postgresql      | build: postgresql/postgresql_lexer.go:3695:73: syntax error: unexpected InputStream, expecting comma or ) |
| ❌  | powerbuilder    | test: FAIL                                                                  |
| ❌  | python          | build: python/python_lexer.go:548:2: undefined: PythonLexerBase             |
| ❌  | python2         | build: python2/python2_lexer.go:636:8: too many errors                      |
| ❌  | python3         | build: python3/python3_lexer.go:742:2: undefined: Python3LexerBase          |
| ❌  | quakemap        | antlr: error(134): quakemap.g4:44:12: symbol string conflicts with generated code in target language or runtime |
| ❌  | rego            | antlr: error(134): RegoLexer.g4:40:0: symbol String conflicts with generated code in target language or runtime |
| ❌  | rexx            | build: rexx/rexx_lexer.go:880:2: syntax error: unexpected default, expecting } |
| ❌  | rust            | build: rust/rust_lexer.go:635:2: undefined: RustLexerBase                   |
| ❌  | sieve           | antlr: error(134): sieve.g4:70:21: symbol string conflicts with generated code in target language or runtime |
| ❌  | smalltalk       | antlr: error(134): Smalltalk.g4:50:57: symbol string conflicts with generated code in target language or runtime |
| ❌  | smtlibv2        | antlr: error(134): SMTLIBv2.g4:1089:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | sparql          | antlr: error(134): Sparql.g4:300:6: symbol string conflicts with generated code in target language or runtime |
| ❌  | st              | build: st/st_lexer.go:790:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | stg             | build: stg/stg_lexer.go:715:70: syntax error: unexpected _input, expecting comma or ) |
| ❌  | swift2          | build: swift2/swift2_parser.go:35927:31: too many errors                    |
| ❌  | swift3          | build: swift3/swift3_parser.go:40532:7: too many errors                     |
| ❌  | swift5          | build: swift5/swift5_lexer.go:12:13: expected 'STRING', found '.'           |
| ❌  | swiftfin        | antlr: error(134): SwiftFinParser.g4:61:16: symbol map conflicts with generated code in target language or runtime |
| ❌  | tcp             | test: FAIL                                                                  |
| ❌  | terraform       | antlr: error(134): terraform.g4:133:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | thrift          | test: FAIL                                                                  |
| ❌  | tinymud         | antlr: error(134): tinymud.g4:40:16: symbol action conflicts with generated code in target language or runtime |
| ❌  | tinyos          | build: tinyos/tinyos_parser.go:3633:3: undefined: System                    |
| ❌  | tjs             | antlr: error(134): TJSLexer.g4:132:0: symbol String conflicts with generated code in target language or runtime |
| ❌  | toml            | antlr: error(134): toml.g4:44:8: symbol string conflicts with generated code in target language or runtime |
| ❌  | trac            | antlr: error(134): trac.g4:62:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | tsql            | test: FAIL                                                                  |
| ❌  | ttm             | antlr: error(134): ttm.g4:62:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | turing          | antlr: error(134): turing.g4:193:31: symbol string conflicts with generated code in target language or runtime |
| ❌  | turtle          | antlr: error(134): TURTLE.g4:112:5: symbol String conflicts with generated code in target language or runtime |
| ❌  | typescript      | antlr: error(134): TypeScriptLexer.g4:150:0: symbol String conflicts with generated code in target language or runtime |
| ❌  | ucblogo         | build: ucblogo/ucblogo_parser.go:14:14: expected 'STRING', found '.'        |
| ❌  | url             | antlr: error(134): url.g4:96:18: symbol string conflicts with generated code in target language or runtime |
| ❌  | v               | build: v/v_parser.go:419:23: syntax error: unexpected ==, expecting name    |
| ❌  | visualbasic6    | test: FAIL                                                                  |
| ❌  | vtl             | build: vtl/vtl_lexer.go:700:14: this._input undefined (type *VTLLexer has no field or method _input) |
| ❌  | xdr             | test: FAIL                                                                  |
| ❌  | xpath           | build: 	/Users/bramp/personal/antlr4-grammars/xpath/xpath_parser.go:3574:6: previous declaration |
| ❌  | xpath20         | build: xpath20/xpath20_parser.go:11259:2: syntax error: unexpected default, expecting : |
| ❌  | xpath31         | build: xpath31/xpath31_parser.go:17223:2: syntax error: unexpected default, expecting : |
| ❌  | z               | build: z/z_lexer.go:1266:2: syntax error: non-declaration statement outside function body |

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
# Fetch the latest set of grammars
git submodule init
git submodule update
git submodule foreach git pull origin master

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
