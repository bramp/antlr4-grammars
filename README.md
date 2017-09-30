# antlr4test-go
by Andrew Brampton

Precompiled Go versions of most of the grammars on github.com/antlr/grammars-v4.

Example
```go

import (
	"bramp.net/antlr4/java"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type exampleListener struct {
	*java.BaseJavaListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := java.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := java.NewJavaParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.CompilationUnit()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
}
```

## Supported

| Status | Language     | Notes                                                                       |
| --- | --------------- | --------------------------------------------------------------------------- |
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
| ✅  | cookie          |                                                                             |
| ✅  | cool            |                                                                             |
| ✅  | creole          |                                                                             |
| ✅  | csv             |                                                                             |
| ✅  | datetime        |                                                                             |
| ✅  | dcm             |                                                                             |
| ✅  | dice            |                                                                             |
| ✅  | emailaddress    |                                                                             |
| ✅  | fasta           |                                                                             |
| ✅  | fol             |                                                                             |
| ✅  | gml             |                                                                             |
| ✅  | gtin            |                                                                             |
| ✅  | idl             |                                                                             |
| ✅  | iri             |                                                                             |
| ✅  | istc            |                                                                             |
| ✅  | java            |                                                                             |
| ✅  | json            |                                                                             |
| ✅  | lambda          |                                                                             |
| ✅  | lcc             |                                                                             |
| ✅  | mdx             |                                                                             |
| ✅  | memcached_protocol |                                                                             |
| ✅  | metric          |                                                                             |
| ✅  | modelica        |                                                                             |
| ✅  | molecule        |                                                                             |
| ✅  | morsecode       |                                                                             |
| ✅  | mps             |                                                                             |
| ✅  | mumath          |                                                                             |
| ✅  | mumps           |                                                                             |
| ✅  | p               |                                                                             |
| ✅  | pcre            |                                                                             |
| ✅  | peoplecode      |                                                                             |
| ✅  | postalcode      |                                                                             |
| ✅  | powerbuilder    |                                                                             |
| ✅  | propcalc        |                                                                             |
| ✅  | r               |                                                                             |
| ✅  | rcs             |                                                                             |
| ✅  | robotwars       |                                                                             |
| ✅  | romannumerals   |                                                                             |
| ✅  | rpn             |                                                                             |
| ✅  | ruby            |                                                                             |
| ✅  | sexpression     |                                                                             |
| ✅  | snobol          |                                                                             |
| ✅  | suokif          |                                                                             |
| ✅  | telephone       |                                                                             |
| ✅  | tiny            |                                                                             |
| ✅  | tinyc           |                                                                             |
| ✅  | tnt             |                                                                             |
| ✅  | useragent       |                                                                             |
| ✅  | vhdl            |                                                                             |
| ✅  | wavefront       |                                                                             |
| ✅  | xpath           |                                                                             |
| ❌  | antlr3          | antlr: error(134): ANTLRv3.g4:32:102: symbol action conflicts with generated code in target language or runtime |
| ❌  | antlr4          | antlr: error(134): ANTLRv4Parser.g4:62:5: symbol action conflicts with generated code in target language or runtime |
| ❌  | apex            | antlr: error(134): apex.g4:544:30: symbol type conflicts with generated code in target language or runtime |
| ❌  | asm6502         | antlr: error(134): asm6502.g4:71:30: symbol string conflicts with generated code in target language or runtime |
| ❌  | asn             | antlr: error(134): ASN.g4:166:2: symbol type conflicts with generated code in target language or runtime |
| ❌  | basic           | antlr: error(134): jvmBasic.g4:481:8: symbol var conflicts with generated code in target language or runtime |
| ❌  | calculator      | antlr: error(134): calculator.g4:54:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | clojure         | antlr: error(134): Clojure.g4:99:12: symbol map conflicts with generated code in target language or runtime |
| ❌  | cpp             | build: cpp/cpp14_parser.go:28859:6: syntax error: unexpected NewEmptyBaseclauseContext, expecting ( |
| ❌  | csharp          | antlr: error(134): CSharpParser.g4:330:8: symbol type conflicts with generated code in target language or runtime |
| ❌  | css3            | antlr: error(134): css3.g4:196:6: symbol var conflicts with generated code in target language or runtime |
| ❌  | ecmascript      | build: ecmascript/ecmascript_lexer.go:761:5: invalid character U+0040 '@'   |
| ❌  | erlang          | antlr: error(134): Erlang.g4:156:36: symbol type conflicts with generated code in target language or runtime |
| ❌  | fortran77       | antlr: error(134): fortran77.g4:89:6: symbol type conflicts with generated code in target language or runtime |
| ❌  | fusion-tables   | build: fusion-tables/fusiontablessql_base_listener.go:3:15: syntax error: unexpected -, expecting ; |
| ❌  | gff3            | antlr: error(134): gff3.g4:45:28: symbol type conflicts with generated code in target language or runtime |
| ❌  | golang          | antlr: error(134): Golang.g4:151:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | graphql         | antlr: error(8): Graphql.g4:33:8: grammar name GraphQL and file name Graphql.g4 differ |
| ❌  | icalendar       | antlr: error(8): iCalendar.g4:35:8: grammar name ICalendar and file name iCalendar.g4 differ |
| ❌  | informix        | antlr: error(134): informix.g4:601:266: symbol string conflicts with generated code in target language or runtime |
| ❌  | java8           | antlr: error(134): Java8.g4:73:0: symbol type conflicts with generated code in target language or runtime |
| ❌  | java8-pt        | build: java8-pt/java_lexer.go:3:14: syntax error: unexpected -, expecting ; |
| ❌  | javadoc         | build: javadoc/javadoc_lexer.go:167:10: undefined: _input                   |
| ❌  | jpa             | antlr: error(8): jpa.g4:1:8: grammar name JPA and file name jpa.g4 differ   |
| ❌  | kotlin          | antlr: error(134): KotlinParser.g4:56:21: symbol type conflicts with generated code in target language or runtime |
| ❌  | lolcode         | antlr: error(134): lolcode.g4:111:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | lua             | antlr: error(134): Lua.g4:90:15: symbol var conflicts with generated code in target language or runtime |
| ❌  | masm            | build: masm/masm_lexer.go:12:4: syntax error: non-declaration statement outside function body |
| ❌  | muparser        |  test: FAIL  bramp.net/antlr4test-go/muparser [build failed]                 |
| ❌  | mysql           | build:       previous declaration at mysql/mysql_parser.go:12215:6               |
| ❌  | pascal          | antlr: error(134): pascal.g4:115:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | pddl            | antlr: error(134): Pddl.g4:92:25: symbol type conflicts with generated code in target language or runtime |
| ❌  | pdn             | antlr: error(134): pdn.g4:44:14: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdp7            | antlr: error(134): pdp7.g4:87:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | php             | antlr: error(134): PHPParser.g4:478:6: symbol string conflicts with generated code in target language or runtime |
| ❌  | plsql           | antlr: error(134): PlSqlParser.g4:779:41: symbol range conflicts with generated code in target language or runtime |
| ❌  | prolog          | antlr: error(134): prolog.g4:72:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | protobuf3       | antlr: error(134): Protobuf3.g4:138:19: symbol range conflicts with generated code in target language or runtime |
| ❌  | python2         | build: python2/python2_lexer.go:639:14: too many errors                     |
| ❌  | python3         | build: python3/python3_lexer.go:828:13: too many errors                     |
| ❌  | python3-js      | build: python3-js/python3_base_listener.go:3:16: syntax error: unexpected -, expecting ; |
| ❌  | python3-py      | build: python3-py/python3_base_listener.go:3:16: syntax error: unexpected -, expecting ; |
| ❌  | python3-ts      | build: python3-ts/python3_base_listener.go:3:16: syntax error: unexpected -, expecting ; |
| ❌  | python3alt      | build: python3alt/altpython3_lexer.go:812:13: too many errors               |
| ❌  | quakemap        | antlr: error(134): quakemap.g4:44:12: symbol string conflicts with generated code in target language or runtime |
| ❌  | scala           | antlr: error(134): Scala.g4:433:17: symbol type conflicts with generated code in target language or runtime |
| ❌  | smalltalk       | antlr: error(134): Smalltalk.g4:50:57: symbol string conflicts with generated code in target language or runtime |
| ❌  | smtlibv2        | antlr: error(134): SMTLIBv2.g4:1089:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | sparql          | antlr: error(134): Sparql.g4:57:45: symbol var conflicts with generated code in target language or runtime |
| ❌  | sqlite          | antlr: error(134): SQLite.g4:34:21: symbol error conflicts with generated code in target language or runtime |
| ❌  | tsql            | build:       previous declaration at tsql/tsql_parser.go:123881:8                |
| ❌  | ucb-logo        | build: ucb-logo/ucblogo_base_listener.go:3:12: syntax error: unexpected -, expecting ; |
| ❌  | url             | antlr: error(134): url.g4:95:18: symbol string conflicts with generated code in target language or runtime |
| ❌  | vb6             | antlr: error(134): VisualBasic6.g4:555:35: symbol type conflicts with generated code in target language or runtime |
| ❌  | vba             | antlr: error(134): vba.g4:569:44: symbol type conflicts with generated code in target language or runtime |
| ❌  | verilog         | antlr: error(134): Verilog2001.g4:1944:24: symbol range conflicts with generated code in target language or runtime |
| ❌  | webidl          | antlr: error(134): WebIDL.g4:584:3: symbol type conflicts with generated code in target language or runtime |

## To develop

```bash
make <name of grammar>
```

## Update
To pull down the latest grammars and compile them:

```bash
git submodule init
git submodule update

# TODO Instructions to download antlr
# Update the Makefile
go run makemake.go

# Now build all the grammars
make clean
make all -k -j8 2> /dev/null

# Update the table at the top with new successes or failures
```
