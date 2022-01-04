// Code generated from lisa.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lisa // lisa
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 235,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 6, 4, 72, 10, 4, 13, 4, 14, 4, 73, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 6, 8, 88, 10,
	8, 13, 8, 14, 8, 89, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 100, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 7, 15, 133, 10, 15, 12, 15, 14, 15, 136, 11, 15, 3, 15, 5,
	15, 139, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 146, 10, 17,
	12, 17, 14, 17, 149, 11, 17, 3, 18, 3, 18, 3, 18, 7, 18, 154, 10, 18, 12,
	18, 14, 18, 157, 11, 18, 3, 19, 5, 19, 160, 10, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 7, 20, 168, 10, 20, 12, 20, 14, 20, 171, 11, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 179, 10, 22, 12, 22, 14,
	22, 182, 11, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 190,
	10, 24, 12, 24, 14, 24, 193, 11, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 202, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 5, 27, 213, 10, 27, 3, 28, 5, 28, 216, 10, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 7,
	30, 228, 10, 30, 12, 30, 14, 30, 231, 11, 30, 3, 31, 3, 31, 3, 31, 2, 2,
	32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 2, 8, 3, 2, 7, 12, 3, 2,
	20, 21, 3, 2, 23, 27, 3, 2, 31, 35, 3, 2, 36, 37, 3, 2, 38, 40, 2, 230,
	2, 62, 3, 2, 2, 2, 4, 65, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 75, 3, 2, 2,
	2, 10, 79, 3, 2, 2, 2, 12, 81, 3, 2, 2, 2, 14, 87, 3, 2, 2, 2, 16, 99,
	3, 2, 2, 2, 18, 101, 3, 2, 2, 2, 20, 111, 3, 2, 2, 2, 22, 113, 3, 2, 2,
	2, 24, 116, 3, 2, 2, 2, 26, 122, 3, 2, 2, 2, 28, 138, 3, 2, 2, 2, 30, 140,
	3, 2, 2, 2, 32, 142, 3, 2, 2, 2, 34, 150, 3, 2, 2, 2, 36, 159, 3, 2, 2,
	2, 38, 163, 3, 2, 2, 2, 40, 172, 3, 2, 2, 2, 42, 174, 3, 2, 2, 2, 44, 183,
	3, 2, 2, 2, 46, 185, 3, 2, 2, 2, 48, 194, 3, 2, 2, 2, 50, 201, 3, 2, 2,
	2, 52, 212, 3, 2, 2, 2, 54, 215, 3, 2, 2, 2, 56, 219, 3, 2, 2, 2, 58, 224,
	3, 2, 2, 2, 60, 232, 3, 2, 2, 2, 62, 63, 5, 4, 3, 2, 63, 64, 5, 12, 7,
	2, 64, 3, 3, 2, 2, 2, 65, 66, 7, 3, 2, 2, 66, 67, 7, 4, 2, 2, 67, 68, 5,
	6, 4, 2, 68, 69, 7, 5, 2, 2, 69, 5, 3, 2, 2, 2, 70, 72, 5, 8, 5, 2, 71,
	70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2,
	2, 74, 7, 3, 2, 2, 2, 75, 76, 5, 10, 6, 2, 76, 77, 7, 47, 2, 2, 77, 78,
	7, 6, 2, 2, 78, 9, 3, 2, 2, 2, 79, 80, 9, 2, 2, 2, 80, 11, 3, 2, 2, 2,
	81, 82, 7, 13, 2, 2, 82, 83, 7, 4, 2, 2, 83, 84, 5, 14, 8, 2, 84, 85, 7,
	5, 2, 2, 85, 13, 3, 2, 2, 2, 86, 88, 5, 16, 9, 2, 87, 86, 3, 2, 2, 2, 88,
	89, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 15, 3, 2, 2,
	2, 91, 100, 5, 22, 12, 2, 92, 100, 5, 24, 13, 2, 93, 100, 5, 26, 14, 2,
	94, 100, 5, 18, 10, 2, 95, 96, 7, 14, 2, 2, 96, 100, 7, 6, 2, 2, 97, 98,
	7, 15, 2, 2, 98, 100, 7, 6, 2, 2, 99, 91, 3, 2, 2, 2, 99, 92, 3, 2, 2,
	2, 99, 93, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 100, 17, 3, 2, 2, 2, 101, 102, 7, 16, 2, 2, 102, 103, 7, 17,
	2, 2, 103, 104, 5, 20, 11, 2, 104, 105, 7, 18, 2, 2, 105, 106, 7, 7, 2,
	2, 106, 107, 7, 18, 2, 2, 107, 108, 7, 7, 2, 2, 108, 109, 7, 19, 2, 2,
	109, 110, 5, 16, 9, 2, 110, 19, 3, 2, 2, 2, 111, 112, 9, 3, 2, 2, 112,
	21, 3, 2, 2, 2, 113, 114, 5, 28, 15, 2, 114, 115, 7, 6, 2, 2, 115, 23,
	3, 2, 2, 2, 116, 117, 7, 22, 2, 2, 117, 118, 7, 17, 2, 2, 118, 119, 5,
	28, 15, 2, 119, 120, 7, 19, 2, 2, 120, 121, 5, 16, 9, 2, 121, 25, 3, 2,
	2, 2, 122, 123, 7, 45, 2, 2, 123, 124, 7, 17, 2, 2, 124, 125, 5, 28, 15,
	2, 125, 126, 7, 19, 2, 2, 126, 127, 5, 16, 9, 2, 127, 27, 3, 2, 2, 2, 128,
	134, 5, 60, 31, 2, 129, 130, 5, 30, 16, 2, 130, 131, 5, 28, 15, 2, 131,
	133, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132,
	3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 139, 3, 2, 2, 2, 136, 134, 3, 2,
	2, 2, 137, 139, 5, 32, 17, 2, 138, 128, 3, 2, 2, 2, 138, 137, 3, 2, 2,
	2, 139, 29, 3, 2, 2, 2, 140, 141, 9, 4, 2, 2, 141, 31, 3, 2, 2, 2, 142,
	147, 5, 34, 18, 2, 143, 144, 7, 28, 2, 2, 144, 146, 5, 34, 18, 2, 145,
	143, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148,
	3, 2, 2, 2, 148, 33, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 155, 5, 36,
	19, 2, 151, 152, 7, 29, 2, 2, 152, 154, 5, 36, 19, 2, 153, 151, 3, 2, 2,
	2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156,
	35, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 160, 7, 30, 2, 2, 159, 158,
	3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 5, 38,
	20, 2, 162, 37, 3, 2, 2, 2, 163, 169, 5, 42, 22, 2, 164, 165, 5, 40, 21,
	2, 165, 166, 5, 42, 22, 2, 166, 168, 3, 2, 2, 2, 167, 164, 3, 2, 2, 2,
	168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	39, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173, 9, 5, 2, 2, 173, 41, 3,
	2, 2, 2, 174, 180, 5, 46, 24, 2, 175, 176, 5, 44, 23, 2, 176, 177, 5, 46,
	24, 2, 177, 179, 3, 2, 2, 2, 178, 175, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2,
	180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 43, 3, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 183, 184, 9, 6, 2, 2, 184, 45, 3, 2, 2, 2, 185, 191, 5, 50,
	26, 2, 186, 187, 5, 48, 25, 2, 187, 188, 5, 50, 26, 2, 188, 190, 3, 2,
	2, 2, 189, 186, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2,
	191, 192, 3, 2, 2, 2, 192, 47, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 195,
	9, 7, 2, 2, 195, 49, 3, 2, 2, 2, 196, 197, 7, 17, 2, 2, 197, 198, 5, 32,
	17, 2, 198, 199, 7, 19, 2, 2, 199, 202, 3, 2, 2, 2, 200, 202, 5, 52, 27,
	2, 201, 196, 3, 2, 2, 2, 201, 200, 3, 2, 2, 2, 202, 51, 3, 2, 2, 2, 203,
	213, 5, 54, 28, 2, 204, 213, 7, 43, 2, 2, 205, 213, 7, 44, 2, 2, 206, 213,
	7, 41, 2, 2, 207, 213, 7, 42, 2, 2, 208, 213, 5, 60, 31, 2, 209, 213, 7,
	48, 2, 2, 210, 213, 5, 56, 29, 2, 211, 213, 5, 10, 6, 2, 212, 203, 3, 2,
	2, 2, 212, 204, 3, 2, 2, 2, 212, 205, 3, 2, 2, 2, 212, 206, 3, 2, 2, 2,
	212, 207, 3, 2, 2, 2, 212, 208, 3, 2, 2, 2, 212, 209, 3, 2, 2, 2, 212,
	210, 3, 2, 2, 2, 212, 211, 3, 2, 2, 2, 213, 53, 3, 2, 2, 2, 214, 216, 9,
	6, 2, 2, 215, 214, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2,
	2, 217, 218, 7, 46, 2, 2, 218, 55, 3, 2, 2, 2, 219, 220, 7, 47, 2, 2, 220,
	221, 7, 17, 2, 2, 221, 222, 5, 58, 30, 2, 222, 223, 7, 19, 2, 2, 223, 57,
	3, 2, 2, 2, 224, 229, 5, 32, 17, 2, 225, 226, 7, 18, 2, 2, 226, 228, 5,
	32, 17, 2, 227, 225, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 229, 230, 3, 2, 2, 2, 230, 59, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2,
	232, 233, 7, 47, 2, 2, 233, 61, 3, 2, 2, 2, 17, 73, 89, 99, 134, 138, 147,
	155, 159, 169, 180, 191, 201, 212, 215, 229,
}
var literalNames = []string{
	"", "'declare'", "'{'", "'}'", "';'", "'int'", "'dfa'", "'nfa'", "'regex'",
	"'bool'", "'string'", "'program'", "'break'", "'continue'", "'generate'",
	"'('", "','", "')'", "'random'", "'enumerate'", "'if'", "'='", "'-+'",
	"'*='", "'/='", "'+='", "'||'", "'&&'", "'!'", "'<='", "'>='", "'=='",
	"'!='", "'>'", "'-'", "'+'", "'*'", "'/'", "'%'", "'next'", "'hasnext'",
	"'TRUE'", "'FALSE'", "'WHILE'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "TRUE", "FALSE", "WHILE", "INTEGER", "ID", "STRINGLITERAL",
	"COMMENT", "WS",
}

var ruleNames = []string{
	"program", "declaration_block", "declaration_statements", "declaration_statement",
	"type_", "program_block", "statements", "statement", "generating_statement",
	"generator_type", "expression_statement", "if_statement", "while_statement",
	"expression", "exprop", "simple_expression", "or_expression", "unary_relationexpression",
	"relation_expression", "relop", "add_expression", "addop", "term", "multop",
	"factor", "constant", "integer", "function_", "parameter_list", "variable",
}

type lisaParser struct {
	*antlr.BaseParser
}

// NewlisaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *lisaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlisaParser(input antlr.TokenStream) *lisaParser {
	this := new(lisaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "lisa.g4"

	return this
}

// lisaParser tokens.
const (
	lisaParserEOF           = antlr.TokenEOF
	lisaParserT__0          = 1
	lisaParserT__1          = 2
	lisaParserT__2          = 3
	lisaParserT__3          = 4
	lisaParserT__4          = 5
	lisaParserT__5          = 6
	lisaParserT__6          = 7
	lisaParserT__7          = 8
	lisaParserT__8          = 9
	lisaParserT__9          = 10
	lisaParserT__10         = 11
	lisaParserT__11         = 12
	lisaParserT__12         = 13
	lisaParserT__13         = 14
	lisaParserT__14         = 15
	lisaParserT__15         = 16
	lisaParserT__16         = 17
	lisaParserT__17         = 18
	lisaParserT__18         = 19
	lisaParserT__19         = 20
	lisaParserT__20         = 21
	lisaParserT__21         = 22
	lisaParserT__22         = 23
	lisaParserT__23         = 24
	lisaParserT__24         = 25
	lisaParserT__25         = 26
	lisaParserT__26         = 27
	lisaParserT__27         = 28
	lisaParserT__28         = 29
	lisaParserT__29         = 30
	lisaParserT__30         = 31
	lisaParserT__31         = 32
	lisaParserT__32         = 33
	lisaParserT__33         = 34
	lisaParserT__34         = 35
	lisaParserT__35         = 36
	lisaParserT__36         = 37
	lisaParserT__37         = 38
	lisaParserT__38         = 39
	lisaParserT__39         = 40
	lisaParserTRUE          = 41
	lisaParserFALSE         = 42
	lisaParserWHILE         = 43
	lisaParserINTEGER       = 44
	lisaParserID            = 45
	lisaParserSTRINGLITERAL = 46
	lisaParserCOMMENT       = 47
	lisaParserWS            = 48
)

// lisaParser rules.
const (
	lisaParserRULE_program                  = 0
	lisaParserRULE_declaration_block        = 1
	lisaParserRULE_declaration_statements   = 2
	lisaParserRULE_declaration_statement    = 3
	lisaParserRULE_type_                    = 4
	lisaParserRULE_program_block            = 5
	lisaParserRULE_statements               = 6
	lisaParserRULE_statement                = 7
	lisaParserRULE_generating_statement     = 8
	lisaParserRULE_generator_type           = 9
	lisaParserRULE_expression_statement     = 10
	lisaParserRULE_if_statement             = 11
	lisaParserRULE_while_statement          = 12
	lisaParserRULE_expression               = 13
	lisaParserRULE_exprop                   = 14
	lisaParserRULE_simple_expression        = 15
	lisaParserRULE_or_expression            = 16
	lisaParserRULE_unary_relationexpression = 17
	lisaParserRULE_relation_expression      = 18
	lisaParserRULE_relop                    = 19
	lisaParserRULE_add_expression           = 20
	lisaParserRULE_addop                    = 21
	lisaParserRULE_term                     = 22
	lisaParserRULE_multop                   = 23
	lisaParserRULE_factor                   = 24
	lisaParserRULE_constant                 = 25
	lisaParserRULE_integer                  = 26
	lisaParserRULE_function_                = 27
	lisaParserRULE_parameter_list           = 28
	lisaParserRULE_variable                 = 29
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Declaration_block() IDeclaration_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaration_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaration_blockContext)
}

func (s *ProgramContext) Program_block() IProgram_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram_blockContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *lisaParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lisaParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Declaration_block()
	}
	{
		p.SetState(61)
		p.Program_block()
	}

	return localctx
}

// IDeclaration_blockContext is an interface to support dynamic dispatch.
type IDeclaration_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaration_blockContext differentiates from other interfaces.
	IsDeclaration_blockContext()
}

type Declaration_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaration_blockContext() *Declaration_blockContext {
	var p = new(Declaration_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_declaration_block
	return p
}

func (*Declaration_blockContext) IsDeclaration_blockContext() {}

func NewDeclaration_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaration_blockContext {
	var p = new(Declaration_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_declaration_block

	return p
}

func (s *Declaration_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaration_blockContext) Declaration_statements() IDeclaration_statementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaration_statementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaration_statementsContext)
}

func (s *Declaration_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaration_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaration_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterDeclaration_block(s)
	}
}

func (s *Declaration_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitDeclaration_block(s)
	}
}

func (p *lisaParser) Declaration_block() (localctx IDeclaration_blockContext) {
	this := p
	_ = this

	localctx = NewDeclaration_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lisaParserRULE_declaration_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(lisaParserT__0)
	}
	{
		p.SetState(64)
		p.Match(lisaParserT__1)
	}
	{
		p.SetState(65)
		p.Declaration_statements()
	}
	{
		p.SetState(66)
		p.Match(lisaParserT__2)
	}

	return localctx
}

// IDeclaration_statementsContext is an interface to support dynamic dispatch.
type IDeclaration_statementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaration_statementsContext differentiates from other interfaces.
	IsDeclaration_statementsContext()
}

type Declaration_statementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaration_statementsContext() *Declaration_statementsContext {
	var p = new(Declaration_statementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_declaration_statements
	return p
}

func (*Declaration_statementsContext) IsDeclaration_statementsContext() {}

func NewDeclaration_statementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaration_statementsContext {
	var p = new(Declaration_statementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_declaration_statements

	return p
}

func (s *Declaration_statementsContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaration_statementsContext) AllDeclaration_statement() []IDeclaration_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclaration_statementContext)(nil)).Elem())
	var tst = make([]IDeclaration_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclaration_statementContext)
		}
	}

	return tst
}

func (s *Declaration_statementsContext) Declaration_statement(i int) IDeclaration_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaration_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclaration_statementContext)
}

func (s *Declaration_statementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaration_statementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaration_statementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterDeclaration_statements(s)
	}
}

func (s *Declaration_statementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitDeclaration_statements(s)
	}
}

func (p *lisaParser) Declaration_statements() (localctx IDeclaration_statementsContext) {
	this := p
	_ = this

	localctx = NewDeclaration_statementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lisaParserRULE_declaration_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lisaParserT__4)|(1<<lisaParserT__5)|(1<<lisaParserT__6)|(1<<lisaParserT__7)|(1<<lisaParserT__8)|(1<<lisaParserT__9))) != 0) {
		{
			p.SetState(68)
			p.Declaration_statement()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclaration_statementContext is an interface to support dynamic dispatch.
type IDeclaration_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaration_statementContext differentiates from other interfaces.
	IsDeclaration_statementContext()
}

type Declaration_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaration_statementContext() *Declaration_statementContext {
	var p = new(Declaration_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_declaration_statement
	return p
}

func (*Declaration_statementContext) IsDeclaration_statementContext() {}

func NewDeclaration_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaration_statementContext {
	var p = new(Declaration_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_declaration_statement

	return p
}

func (s *Declaration_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaration_statementContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Declaration_statementContext) ID() antlr.TerminalNode {
	return s.GetToken(lisaParserID, 0)
}

func (s *Declaration_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaration_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaration_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterDeclaration_statement(s)
	}
}

func (s *Declaration_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitDeclaration_statement(s)
	}
}

func (p *lisaParser) Declaration_statement() (localctx IDeclaration_statementContext) {
	this := p
	_ = this

	localctx = NewDeclaration_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lisaParserRULE_declaration_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Type_()
	}
	{
		p.SetState(74)
		p.Match(lisaParserID)
	}
	{
		p.SetState(75)
		p.Match(lisaParserT__3)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }
func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *lisaParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lisaParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lisaParserT__4)|(1<<lisaParserT__5)|(1<<lisaParserT__6)|(1<<lisaParserT__7)|(1<<lisaParserT__8)|(1<<lisaParserT__9))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProgram_blockContext is an interface to support dynamic dispatch.
type IProgram_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgram_blockContext differentiates from other interfaces.
	IsProgram_blockContext()
}

type Program_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgram_blockContext() *Program_blockContext {
	var p = new(Program_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_program_block
	return p
}

func (*Program_blockContext) IsProgram_blockContext() {}

func NewProgram_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program_blockContext {
	var p = new(Program_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_program_block

	return p
}

func (s *Program_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Program_blockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *Program_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Program_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Program_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterProgram_block(s)
	}
}

func (s *Program_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitProgram_block(s)
	}
}

func (p *lisaParser) Program_block() (localctx IProgram_blockContext) {
	this := p
	_ = this

	localctx = NewProgram_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lisaParserRULE_program_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(lisaParserT__10)
	}
	{
		p.SetState(80)
		p.Match(lisaParserT__1)
	}
	{
		p.SetState(81)
		p.Statements()
	}
	{
		p.SetState(82)
		p.Match(lisaParserT__2)
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *lisaParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lisaParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lisaParserT__4)|(1<<lisaParserT__5)|(1<<lisaParserT__6)|(1<<lisaParserT__7)|(1<<lisaParserT__8)|(1<<lisaParserT__9)|(1<<lisaParserT__11)|(1<<lisaParserT__12)|(1<<lisaParserT__13)|(1<<lisaParserT__14)|(1<<lisaParserT__19)|(1<<lisaParserT__27))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(lisaParserT__33-34))|(1<<(lisaParserT__34-34))|(1<<(lisaParserT__38-34))|(1<<(lisaParserT__39-34))|(1<<(lisaParserTRUE-34))|(1<<(lisaParserFALSE-34))|(1<<(lisaParserWHILE-34))|(1<<(lisaParserINTEGER-34))|(1<<(lisaParserID-34))|(1<<(lisaParserSTRINGLITERAL-34)))) != 0) {
		{
			p.SetState(84)
			p.Statement()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression_statement() IExpression_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) While_statement() IWhile_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *StatementContext) Generating_statement() IGenerating_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenerating_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenerating_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *lisaParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lisaParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lisaParserT__4, lisaParserT__5, lisaParserT__6, lisaParserT__7, lisaParserT__8, lisaParserT__9, lisaParserT__14, lisaParserT__27, lisaParserT__33, lisaParserT__34, lisaParserT__38, lisaParserT__39, lisaParserTRUE, lisaParserFALSE, lisaParserINTEGER, lisaParserID, lisaParserSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Expression_statement()
		}

	case lisaParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.If_statement()
		}

	case lisaParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.While_statement()
		}

	case lisaParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Generating_statement()
		}

	case lisaParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(lisaParserT__11)
		}
		{
			p.SetState(94)
			p.Match(lisaParserT__3)
		}

	case lisaParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(lisaParserT__12)
		}
		{
			p.SetState(96)
			p.Match(lisaParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGenerating_statementContext is an interface to support dynamic dispatch.
type IGenerating_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenerating_statementContext differentiates from other interfaces.
	IsGenerating_statementContext()
}

type Generating_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenerating_statementContext() *Generating_statementContext {
	var p = new(Generating_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_generating_statement
	return p
}

func (*Generating_statementContext) IsGenerating_statementContext() {}

func NewGenerating_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Generating_statementContext {
	var p = new(Generating_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_generating_statement

	return p
}

func (s *Generating_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Generating_statementContext) Generator_type() IGenerator_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenerator_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenerator_typeContext)
}

func (s *Generating_statementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Generating_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Generating_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Generating_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterGenerating_statement(s)
	}
}

func (s *Generating_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitGenerating_statement(s)
	}
}

func (p *lisaParser) Generating_statement() (localctx IGenerating_statementContext) {
	this := p
	_ = this

	localctx = NewGenerating_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lisaParserRULE_generating_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(lisaParserT__13)
	}
	{
		p.SetState(100)
		p.Match(lisaParserT__14)
	}
	{
		p.SetState(101)
		p.Generator_type()
	}
	{
		p.SetState(102)
		p.Match(lisaParserT__15)
	}
	{
		p.SetState(103)
		p.Match(lisaParserT__4)
	}
	{
		p.SetState(104)
		p.Match(lisaParserT__15)
	}
	{
		p.SetState(105)
		p.Match(lisaParserT__4)
	}
	{
		p.SetState(106)
		p.Match(lisaParserT__16)
	}
	{
		p.SetState(107)
		p.Statement()
	}

	return localctx
}

// IGenerator_typeContext is an interface to support dynamic dispatch.
type IGenerator_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenerator_typeContext differentiates from other interfaces.
	IsGenerator_typeContext()
}

type Generator_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenerator_typeContext() *Generator_typeContext {
	var p = new(Generator_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_generator_type
	return p
}

func (*Generator_typeContext) IsGenerator_typeContext() {}

func NewGenerator_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Generator_typeContext {
	var p = new(Generator_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_generator_type

	return p
}

func (s *Generator_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Generator_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Generator_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Generator_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterGenerator_type(s)
	}
}

func (s *Generator_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitGenerator_type(s)
	}
}

func (p *lisaParser) Generator_type() (localctx IGenerator_typeContext) {
	this := p
	_ = this

	localctx = NewGenerator_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lisaParserRULE_generator_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lisaParserT__17 || _la == lisaParserT__18) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpression_statementContext is an interface to support dynamic dispatch.
type IExpression_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_statementContext differentiates from other interfaces.
	IsExpression_statementContext()
}

type Expression_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_statementContext() *Expression_statementContext {
	var p = new(Expression_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_expression_statement
	return p
}

func (*Expression_statementContext) IsExpression_statementContext() {}

func NewExpression_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_statementContext {
	var p = new(Expression_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_expression_statement

	return p
}

func (s *Expression_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterExpression_statement(s)
	}
}

func (s *Expression_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitExpression_statement(s)
	}
}

func (p *lisaParser) Expression_statement() (localctx IExpression_statementContext) {
	this := p
	_ = this

	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lisaParserRULE_expression_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Expression()
	}
	{
		p.SetState(112)
		p.Match(lisaParserT__3)
	}

	return localctx
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_if_statement
	return p
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_statementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *lisaParser) If_statement() (localctx IIf_statementContext) {
	this := p
	_ = this

	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, lisaParserRULE_if_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(lisaParserT__19)
	}
	{
		p.SetState(115)
		p.Match(lisaParserT__14)
	}
	{
		p.SetState(116)
		p.Expression()
	}
	{
		p.SetState(117)
		p.Match(lisaParserT__16)
	}
	{
		p.SetState(118)
		p.Statement()
	}

	return localctx
}

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_while_statement
	return p
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(lisaParserWHILE, 0)
}

func (s *While_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_statementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterWhile_statement(s)
	}
}

func (s *While_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitWhile_statement(s)
	}
}

func (p *lisaParser) While_statement() (localctx IWhile_statementContext) {
	this := p
	_ = this

	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, lisaParserRULE_while_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(lisaParserWHILE)
	}
	{
		p.SetState(121)
		p.Match(lisaParserT__14)
	}
	{
		p.SetState(122)
		p.Expression()
	}
	{
		p.SetState(123)
		p.Match(lisaParserT__16)
	}
	{
		p.SetState(124)
		p.Statement()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionContext) AllExprop() []IExpropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpropContext)(nil)).Elem())
	var tst = make([]IExpropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpropContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Exprop(i int) IExpropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpropContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Simple_expression() ISimple_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *lisaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, lisaParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Variable()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(127)
					p.Exprop()
				}
				{
					p.SetState(128)
					p.Expression()
				}

			}
			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Simple_expression()
		}

	}

	return localctx
}

// IExpropContext is an interface to support dynamic dispatch.
type IExpropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpropContext differentiates from other interfaces.
	IsExpropContext()
}

type ExpropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpropContext() *ExpropContext {
	var p = new(ExpropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_exprop
	return p
}

func (*ExpropContext) IsExpropContext() {}

func NewExpropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpropContext {
	var p = new(ExpropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_exprop

	return p
}

func (s *ExpropContext) GetParser() antlr.Parser { return s.parser }
func (s *ExpropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterExprop(s)
	}
}

func (s *ExpropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitExprop(s)
	}
}

func (p *lisaParser) Exprop() (localctx IExpropContext) {
	this := p
	_ = this

	localctx = NewExpropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, lisaParserRULE_exprop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lisaParserT__20)|(1<<lisaParserT__21)|(1<<lisaParserT__22)|(1<<lisaParserT__23)|(1<<lisaParserT__24))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISimple_expressionContext is an interface to support dynamic dispatch.
type ISimple_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_expressionContext differentiates from other interfaces.
	IsSimple_expressionContext()
}

type Simple_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_expressionContext() *Simple_expressionContext {
	var p = new(Simple_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_simple_expression
	return p
}

func (*Simple_expressionContext) IsSimple_expressionContext() {}

func NewSimple_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_expressionContext {
	var p = new(Simple_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_simple_expression

	return p
}

func (s *Simple_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_expressionContext) AllOr_expression() []IOr_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem())
	var tst = make([]IOr_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOr_expressionContext)
		}
	}

	return tst
}

func (s *Simple_expressionContext) Or_expression(i int) IOr_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *Simple_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterSimple_expression(s)
	}
}

func (s *Simple_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitSimple_expression(s)
	}
}

func (p *lisaParser) Simple_expression() (localctx ISimple_expressionContext) {
	this := p
	_ = this

	localctx = NewSimple_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, lisaParserRULE_simple_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Or_expression()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lisaParserT__25 {
		{
			p.SetState(141)
			p.Match(lisaParserT__25)
		}
		{
			p.SetState(142)
			p.Or_expression()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOr_expressionContext is an interface to support dynamic dispatch.
type IOr_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_expressionContext differentiates from other interfaces.
	IsOr_expressionContext()
}

type Or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_expressionContext() *Or_expressionContext {
	var p = new(Or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_or_expression
	return p
}

func (*Or_expressionContext) IsOr_expressionContext() {}

func NewOr_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_expressionContext {
	var p = new(Or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_or_expression

	return p
}

func (s *Or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_expressionContext) AllUnary_relationexpression() []IUnary_relationexpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnary_relationexpressionContext)(nil)).Elem())
	var tst = make([]IUnary_relationexpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnary_relationexpressionContext)
		}
	}

	return tst
}

func (s *Or_expressionContext) Unary_relationexpression(i int) IUnary_relationexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_relationexpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnary_relationexpressionContext)
}

func (s *Or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterOr_expression(s)
	}
}

func (s *Or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitOr_expression(s)
	}
}

func (p *lisaParser) Or_expression() (localctx IOr_expressionContext) {
	this := p
	_ = this

	localctx = NewOr_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, lisaParserRULE_or_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Unary_relationexpression()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lisaParserT__26 {
		{
			p.SetState(149)
			p.Match(lisaParserT__26)
		}
		{
			p.SetState(150)
			p.Unary_relationexpression()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnary_relationexpressionContext is an interface to support dynamic dispatch.
type IUnary_relationexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_relationexpressionContext differentiates from other interfaces.
	IsUnary_relationexpressionContext()
}

type Unary_relationexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_relationexpressionContext() *Unary_relationexpressionContext {
	var p = new(Unary_relationexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_unary_relationexpression
	return p
}

func (*Unary_relationexpressionContext) IsUnary_relationexpressionContext() {}

func NewUnary_relationexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_relationexpressionContext {
	var p = new(Unary_relationexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_unary_relationexpression

	return p
}

func (s *Unary_relationexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_relationexpressionContext) Relation_expression() IRelation_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelation_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelation_expressionContext)
}

func (s *Unary_relationexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_relationexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_relationexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterUnary_relationexpression(s)
	}
}

func (s *Unary_relationexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitUnary_relationexpression(s)
	}
}

func (p *lisaParser) Unary_relationexpression() (localctx IUnary_relationexpressionContext) {
	this := p
	_ = this

	localctx = NewUnary_relationexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, lisaParserRULE_unary_relationexpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lisaParserT__27 {
		{
			p.SetState(156)
			p.Match(lisaParserT__27)
		}

	}
	{
		p.SetState(159)
		p.Relation_expression()
	}

	return localctx
}

// IRelation_expressionContext is an interface to support dynamic dispatch.
type IRelation_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelation_expressionContext differentiates from other interfaces.
	IsRelation_expressionContext()
}

type Relation_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelation_expressionContext() *Relation_expressionContext {
	var p = new(Relation_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_relation_expression
	return p
}

func (*Relation_expressionContext) IsRelation_expressionContext() {}

func NewRelation_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relation_expressionContext {
	var p = new(Relation_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_relation_expression

	return p
}

func (s *Relation_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Relation_expressionContext) AllAdd_expression() []IAdd_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdd_expressionContext)(nil)).Elem())
	var tst = make([]IAdd_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdd_expressionContext)
		}
	}

	return tst
}

func (s *Relation_expressionContext) Add_expression(i int) IAdd_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdd_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdd_expressionContext)
}

func (s *Relation_expressionContext) AllRelop() []IRelopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelopContext)(nil)).Elem())
	var tst = make([]IRelopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelopContext)
		}
	}

	return tst
}

func (s *Relation_expressionContext) Relop(i int) IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *Relation_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relation_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relation_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterRelation_expression(s)
	}
}

func (s *Relation_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitRelation_expression(s)
	}
}

func (p *lisaParser) Relation_expression() (localctx IRelation_expressionContext) {
	this := p
	_ = this

	localctx = NewRelation_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, lisaParserRULE_relation_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Add_expression()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(lisaParserT__28-29))|(1<<(lisaParserT__29-29))|(1<<(lisaParserT__30-29))|(1<<(lisaParserT__31-29))|(1<<(lisaParserT__32-29)))) != 0 {
		{
			p.SetState(162)
			p.Relop()
		}
		{
			p.SetState(163)
			p.Add_expression()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *lisaParser) Relop() (localctx IRelopContext) {
	this := p
	_ = this

	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, lisaParserRULE_relop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(lisaParserT__28-29))|(1<<(lisaParserT__29-29))|(1<<(lisaParserT__30-29))|(1<<(lisaParserT__31-29))|(1<<(lisaParserT__32-29)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdd_expressionContext is an interface to support dynamic dispatch.
type IAdd_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdd_expressionContext differentiates from other interfaces.
	IsAdd_expressionContext()
}

type Add_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_expressionContext() *Add_expressionContext {
	var p = new(Add_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_add_expression
	return p
}

func (*Add_expressionContext) IsAdd_expressionContext() {}

func NewAdd_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_expressionContext {
	var p = new(Add_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_add_expression

	return p
}

func (s *Add_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_expressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Add_expressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Add_expressionContext) AllAddop() []IAddopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddopContext)(nil)).Elem())
	var tst = make([]IAddopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddopContext)
		}
	}

	return tst
}

func (s *Add_expressionContext) Addop(i int) IAddopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *Add_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Add_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterAdd_expression(s)
	}
}

func (s *Add_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitAdd_expression(s)
	}
}

func (p *lisaParser) Add_expression() (localctx IAdd_expressionContext) {
	this := p
	_ = this

	localctx = NewAdd_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, lisaParserRULE_add_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Term()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lisaParserT__33 || _la == lisaParserT__34 {
		{
			p.SetState(173)
			p.Addop()
		}
		{
			p.SetState(174)
			p.Term()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAddopContext is an interface to support dynamic dispatch.
type IAddopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddopContext differentiates from other interfaces.
	IsAddopContext()
}

type AddopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddopContext() *AddopContext {
	var p = new(AddopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_addop
	return p
}

func (*AddopContext) IsAddopContext() {}

func NewAddopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddopContext {
	var p = new(AddopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_addop

	return p
}

func (s *AddopContext) GetParser() antlr.Parser { return s.parser }
func (s *AddopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterAddop(s)
	}
}

func (s *AddopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitAddop(s)
	}
}

func (p *lisaParser) Addop() (localctx IAddopContext) {
	this := p
	_ = this

	localctx = NewAddopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, lisaParserRULE_addop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lisaParserT__33 || _la == lisaParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMultop() []IMultopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultopContext)(nil)).Elem())
	var tst = make([]IMultopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultopContext)
		}
	}

	return tst
}

func (s *TermContext) Multop(i int) IMultopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultopContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *lisaParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, lisaParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Factor()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(lisaParserT__35-36))|(1<<(lisaParserT__36-36))|(1<<(lisaParserT__37-36)))) != 0 {
		{
			p.SetState(184)
			p.Multop()
		}
		{
			p.SetState(185)
			p.Factor()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultopContext is an interface to support dynamic dispatch.
type IMultopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultopContext differentiates from other interfaces.
	IsMultopContext()
}

type MultopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultopContext() *MultopContext {
	var p = new(MultopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_multop
	return p
}

func (*MultopContext) IsMultopContext() {}

func NewMultopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultopContext {
	var p = new(MultopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_multop

	return p
}

func (s *MultopContext) GetParser() antlr.Parser { return s.parser }
func (s *MultopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterMultop(s)
	}
}

func (s *MultopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitMultop(s)
	}
}

func (p *lisaParser) Multop() (localctx IMultopContext) {
	this := p
	_ = this

	localctx = NewMultopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, lisaParserRULE_multop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(lisaParserT__35-36))|(1<<(lisaParserT__36-36))|(1<<(lisaParserT__37-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Simple_expression() ISimple_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *FactorContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *lisaParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, lisaParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lisaParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Match(lisaParserT__14)
		}
		{
			p.SetState(195)
			p.Simple_expression()
		}
		{
			p.SetState(196)
			p.Match(lisaParserT__16)
		}

	case lisaParserT__4, lisaParserT__5, lisaParserT__6, lisaParserT__7, lisaParserT__8, lisaParserT__9, lisaParserT__33, lisaParserT__34, lisaParserT__38, lisaParserT__39, lisaParserTRUE, lisaParserFALSE, lisaParserINTEGER, lisaParserID, lisaParserSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ConstantContext) TRUE() antlr.TerminalNode {
	return s.GetToken(lisaParserTRUE, 0)
}

func (s *ConstantContext) FALSE() antlr.TerminalNode {
	return s.GetToken(lisaParserFALSE, 0)
}

func (s *ConstantContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ConstantContext) STRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(lisaParserSTRINGLITERAL, 0)
}

func (s *ConstantContext) Function_() IFunction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *ConstantContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *lisaParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, lisaParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(lisaParserTRUE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(lisaParserFALSE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(lisaParserT__38)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)
			p.Match(lisaParserT__39)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(206)
			p.Variable()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(207)
			p.Match(lisaParserSTRINGLITERAL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(208)
			p.Function_()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(209)
			p.Type_()
		}

	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(lisaParserINTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *lisaParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, lisaParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lisaParserT__33 || _la == lisaParserT__34 {
		{
			p.SetState(212)
			_la = p.GetTokenStream().LA(1)

			if !(_la == lisaParserT__33 || _la == lisaParserT__34) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(215)
		p.Match(lisaParserINTEGER)
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) ID() antlr.TerminalNode {
	return s.GetToken(lisaParserID, 0)
}

func (s *Function_Context) Parameter_list() IParameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_listContext)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *lisaParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, lisaParserRULE_function_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(lisaParserID)
	}
	{
		p.SetState(218)
		p.Match(lisaParserT__14)
	}
	{
		p.SetState(219)
		p.Parameter_list()
	}
	{
		p.SetState(220)
		p.Match(lisaParserT__16)
	}

	return localctx
}

// IParameter_listContext is an interface to support dynamic dispatch.
type IParameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_listContext differentiates from other interfaces.
	IsParameter_listContext()
}

type Parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_listContext() *Parameter_listContext {
	var p = new(Parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_parameter_list
	return p
}

func (*Parameter_listContext) IsParameter_listContext() {}

func NewParameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_listContext {
	var p = new(Parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_parameter_list

	return p
}

func (s *Parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_listContext) AllSimple_expression() []ISimple_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem())
	var tst = make([]ISimple_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimple_expressionContext)
		}
	}

	return tst
}

func (s *Parameter_listContext) Simple_expression(i int) ISimple_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimple_expressionContext)
}

func (s *Parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterParameter_list(s)
	}
}

func (s *Parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitParameter_list(s)
	}
}

func (p *lisaParser) Parameter_list() (localctx IParameter_listContext) {
	this := p
	_ = this

	localctx = NewParameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, lisaParserRULE_parameter_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Simple_expression()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lisaParserT__15 {
		{
			p.SetState(223)
			p.Match(lisaParserT__15)
		}
		{
			p.SetState(224)
			p.Simple_expression()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lisaParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lisaParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(lisaParserID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lisaListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *lisaParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, lisaParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(lisaParserID)
	}

	return localctx
}
