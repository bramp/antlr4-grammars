// Generated from R.g4 by ANTLR 4.7.

package r // R
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 211,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 21, 10, 2, 12, 2, 14, 2, 24,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 38, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 96, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 150, 10, 3, 12, 3, 14,
	3, 153, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 158, 10, 4, 7, 4, 160, 10, 4, 12,
	4, 14, 4, 163, 11, 4, 3, 4, 5, 4, 166, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 171,
	10, 5, 12, 5, 14, 5, 174, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 181,
	10, 6, 3, 7, 3, 7, 3, 7, 7, 7, 186, 10, 7, 12, 7, 14, 7, 189, 11, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 209, 10, 8, 3, 8, 2, 3, 4, 9, 2,
	4, 6, 8, 10, 12, 14, 2, 11, 4, 2, 3, 3, 64, 64, 3, 2, 12, 13, 3, 2, 7,
	8, 3, 2, 9, 10, 3, 2, 15, 16, 3, 2, 17, 22, 3, 2, 24, 25, 3, 2, 26, 27,
	3, 2, 29, 34, 2, 261, 2, 22, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2, 6, 165, 3,
	2, 2, 2, 8, 167, 3, 2, 2, 2, 10, 180, 3, 2, 2, 2, 12, 182, 3, 2, 2, 2,
	14, 208, 3, 2, 2, 2, 16, 17, 5, 4, 3, 2, 17, 18, 9, 2, 2, 2, 18, 21, 3,
	2, 2, 2, 19, 21, 7, 64, 2, 2, 20, 16, 3, 2, 2, 2, 20, 19, 3, 2, 2, 2, 21,
	24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 25, 3, 2, 2,
	2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 8,
	3, 1, 2, 28, 29, 9, 3, 2, 2, 29, 96, 5, 4, 3, 38, 30, 31, 7, 23, 2, 2,
	31, 96, 5, 4, 3, 32, 32, 33, 7, 28, 2, 2, 33, 96, 5, 4, 3, 29, 34, 35,
	7, 35, 2, 2, 35, 37, 7, 36, 2, 2, 36, 38, 5, 8, 5, 2, 37, 36, 3, 2, 2,
	2, 37, 38, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 7, 37, 2, 2, 40, 96,
	5, 4, 3, 26, 41, 42, 7, 38, 2, 2, 42, 43, 5, 6, 4, 2, 43, 44, 7, 39, 2,
	2, 44, 96, 3, 2, 2, 2, 45, 46, 7, 40, 2, 2, 46, 47, 7, 36, 2, 2, 47, 48,
	5, 4, 3, 2, 48, 49, 7, 37, 2, 2, 49, 50, 5, 4, 3, 23, 50, 96, 3, 2, 2,
	2, 51, 52, 7, 40, 2, 2, 52, 53, 7, 36, 2, 2, 53, 54, 5, 4, 3, 2, 54, 55,
	7, 37, 2, 2, 55, 56, 5, 4, 3, 2, 56, 57, 7, 41, 2, 2, 57, 58, 5, 4, 3,
	22, 58, 96, 3, 2, 2, 2, 59, 60, 7, 42, 2, 2, 60, 61, 7, 36, 2, 2, 61, 62,
	7, 62, 2, 2, 62, 63, 7, 43, 2, 2, 63, 64, 5, 4, 3, 2, 64, 65, 7, 37, 2,
	2, 65, 66, 5, 4, 3, 21, 66, 96, 3, 2, 2, 2, 67, 68, 7, 44, 2, 2, 68, 69,
	7, 36, 2, 2, 69, 70, 5, 4, 3, 2, 70, 71, 7, 37, 2, 2, 71, 72, 5, 4, 3,
	20, 72, 96, 3, 2, 2, 2, 73, 74, 7, 45, 2, 2, 74, 96, 5, 4, 3, 19, 75, 76,
	7, 46, 2, 2, 76, 96, 5, 4, 3, 18, 77, 96, 7, 47, 2, 2, 78, 96, 7, 48, 2,
	2, 79, 80, 7, 36, 2, 2, 80, 81, 5, 4, 3, 2, 81, 82, 7, 37, 2, 2, 82, 96,
	3, 2, 2, 2, 83, 96, 7, 62, 2, 2, 84, 96, 7, 61, 2, 2, 85, 96, 7, 57, 2,
	2, 86, 96, 7, 58, 2, 2, 87, 96, 7, 59, 2, 2, 88, 96, 7, 60, 2, 2, 89, 96,
	7, 49, 2, 2, 90, 96, 7, 50, 2, 2, 91, 96, 7, 51, 2, 2, 92, 96, 7, 52, 2,
	2, 93, 96, 7, 53, 2, 2, 94, 96, 7, 54, 2, 2, 95, 27, 3, 2, 2, 2, 95, 30,
	3, 2, 2, 2, 95, 32, 3, 2, 2, 2, 95, 34, 3, 2, 2, 2, 95, 41, 3, 2, 2, 2,
	95, 45, 3, 2, 2, 2, 95, 51, 3, 2, 2, 2, 95, 59, 3, 2, 2, 2, 95, 67, 3,
	2, 2, 2, 95, 73, 3, 2, 2, 2, 95, 75, 3, 2, 2, 2, 95, 77, 3, 2, 2, 2, 95,
	78, 3, 2, 2, 2, 95, 79, 3, 2, 2, 2, 95, 83, 3, 2, 2, 2, 95, 84, 3, 2, 2,
	2, 95, 85, 3, 2, 2, 2, 95, 86, 3, 2, 2, 2, 95, 87, 3, 2, 2, 2, 95, 88,
	3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2,
	95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 151, 3,
	2, 2, 2, 97, 98, 12, 41, 2, 2, 98, 99, 9, 4, 2, 2, 99, 150, 5, 4, 3, 42,
	100, 101, 12, 40, 2, 2, 101, 102, 9, 5, 2, 2, 102, 150, 5, 4, 3, 41, 103,
	104, 12, 39, 2, 2, 104, 105, 7, 11, 2, 2, 105, 150, 5, 4, 3, 39, 106, 107,
	12, 37, 2, 2, 107, 108, 7, 14, 2, 2, 108, 150, 5, 4, 3, 38, 109, 110, 12,
	36, 2, 2, 110, 111, 7, 63, 2, 2, 111, 150, 5, 4, 3, 37, 112, 113, 12, 35,
	2, 2, 113, 114, 9, 6, 2, 2, 114, 150, 5, 4, 3, 36, 115, 116, 12, 34, 2,
	2, 116, 117, 9, 3, 2, 2, 117, 150, 5, 4, 3, 35, 118, 119, 12, 33, 2, 2,
	119, 120, 9, 7, 2, 2, 120, 150, 5, 4, 3, 34, 121, 122, 12, 31, 2, 2, 122,
	123, 9, 8, 2, 2, 123, 150, 5, 4, 3, 32, 124, 125, 12, 30, 2, 2, 125, 126,
	9, 9, 2, 2, 126, 150, 5, 4, 3, 31, 127, 128, 12, 28, 2, 2, 128, 129, 7,
	28, 2, 2, 129, 150, 5, 4, 3, 29, 130, 131, 12, 27, 2, 2, 131, 132, 9, 10,
	2, 2, 132, 150, 5, 4, 3, 28, 133, 134, 12, 43, 2, 2, 134, 135, 7, 4, 2,
	2, 135, 136, 5, 12, 7, 2, 136, 137, 7, 5, 2, 2, 137, 138, 7, 5, 2, 2, 138,
	150, 3, 2, 2, 2, 139, 140, 12, 42, 2, 2, 140, 141, 7, 6, 2, 2, 141, 142,
	5, 12, 7, 2, 142, 143, 7, 5, 2, 2, 143, 150, 3, 2, 2, 2, 144, 145, 12,
	25, 2, 2, 145, 146, 7, 36, 2, 2, 146, 147, 5, 12, 7, 2, 147, 148, 7, 37,
	2, 2, 148, 150, 3, 2, 2, 2, 149, 97, 3, 2, 2, 2, 149, 100, 3, 2, 2, 2,
	149, 103, 3, 2, 2, 2, 149, 106, 3, 2, 2, 2, 149, 109, 3, 2, 2, 2, 149,
	112, 3, 2, 2, 2, 149, 115, 3, 2, 2, 2, 149, 118, 3, 2, 2, 2, 149, 121,
	3, 2, 2, 2, 149, 124, 3, 2, 2, 2, 149, 127, 3, 2, 2, 2, 149, 130, 3, 2,
	2, 2, 149, 133, 3, 2, 2, 2, 149, 139, 3, 2, 2, 2, 149, 144, 3, 2, 2, 2,
	150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152,
	5, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 161, 5, 4, 3, 2, 155, 157, 9,
	2, 2, 2, 156, 158, 5, 4, 3, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 160, 3, 2, 2, 2, 159, 155, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161,
	159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 166, 3, 2, 2, 2, 163, 161,
	3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 154, 3, 2, 2, 2, 165, 164, 3, 2,
	2, 2, 166, 7, 3, 2, 2, 2, 167, 172, 5, 10, 6, 2, 168, 169, 7, 55, 2, 2,
	169, 171, 5, 10, 6, 2, 170, 168, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172,
	170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 9, 3, 2, 2, 2, 174, 172, 3,
	2, 2, 2, 175, 181, 7, 62, 2, 2, 176, 177, 7, 62, 2, 2, 177, 178, 7, 31,
	2, 2, 178, 181, 5, 4, 3, 2, 179, 181, 7, 56, 2, 2, 180, 175, 3, 2, 2, 2,
	180, 176, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 11, 3, 2, 2, 2, 182, 187,
	5, 14, 8, 2, 183, 184, 7, 55, 2, 2, 184, 186, 5, 14, 8, 2, 185, 183, 3,
	2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2,
	2, 188, 13, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 209, 5, 4, 3, 2, 191,
	192, 7, 62, 2, 2, 192, 209, 7, 31, 2, 2, 193, 194, 7, 62, 2, 2, 194, 195,
	7, 31, 2, 2, 195, 209, 5, 4, 3, 2, 196, 197, 7, 61, 2, 2, 197, 209, 7,
	31, 2, 2, 198, 199, 7, 61, 2, 2, 199, 200, 7, 31, 2, 2, 200, 209, 5, 4,
	3, 2, 201, 202, 7, 49, 2, 2, 202, 209, 7, 31, 2, 2, 203, 204, 7, 49, 2,
	2, 204, 205, 7, 31, 2, 2, 205, 209, 5, 4, 3, 2, 206, 209, 7, 56, 2, 2,
	207, 209, 3, 2, 2, 2, 208, 190, 3, 2, 2, 2, 208, 191, 3, 2, 2, 2, 208,
	193, 3, 2, 2, 2, 208, 196, 3, 2, 2, 2, 208, 198, 3, 2, 2, 2, 208, 201,
	3, 2, 2, 2, 208, 203, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 207, 3, 2,
	2, 2, 209, 15, 3, 2, 2, 2, 15, 20, 22, 37, 95, 149, 151, 157, 161, 165,
	172, 180, 187, 208,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'[['", "']'", "'['", "'::'", "':::'", "'$'", "'@'", "'^'",
	"'-'", "'+'", "':'", "'*'", "'/'", "'>'", "'>='", "'<'", "'<='", "'=='",
	"'!='", "'!'", "'&'", "'&&'", "'|'", "'||'", "'~'", "'<-'", "'<<-'", "'='",
	"'->'", "'->>'", "':='", "'function'", "'('", "')'", "'{'", "'}'", "'if'",
	"'else'", "'for'", "'in'", "'while'", "'repeat'", "'?'", "'next'", "'break'",
	"'NULL'", "'NA'", "'Inf'", "'NaN'", "'TRUE'", "'FALSE'", "','", "'...'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "HEX", "INT", "FLOAT", "COMPLEX", "STRING", "ID", "USER_OP", "NL",
	"WS",
}

var ruleNames = []string{
	"prog", "expr", "exprlist", "formlist", "form", "sublist", "sub",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RParser struct {
	*antlr.BaseParser
}

func NewRParser(input antlr.TokenStream) *RParser {
	this := new(RParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "R.g4"

	return this
}

// RParser tokens.
const (
	RParserEOF     = antlr.TokenEOF
	RParserT__0    = 1
	RParserT__1    = 2
	RParserT__2    = 3
	RParserT__3    = 4
	RParserT__4    = 5
	RParserT__5    = 6
	RParserT__6    = 7
	RParserT__7    = 8
	RParserT__8    = 9
	RParserT__9    = 10
	RParserT__10   = 11
	RParserT__11   = 12
	RParserT__12   = 13
	RParserT__13   = 14
	RParserT__14   = 15
	RParserT__15   = 16
	RParserT__16   = 17
	RParserT__17   = 18
	RParserT__18   = 19
	RParserT__19   = 20
	RParserT__20   = 21
	RParserT__21   = 22
	RParserT__22   = 23
	RParserT__23   = 24
	RParserT__24   = 25
	RParserT__25   = 26
	RParserT__26   = 27
	RParserT__27   = 28
	RParserT__28   = 29
	RParserT__29   = 30
	RParserT__30   = 31
	RParserT__31   = 32
	RParserT__32   = 33
	RParserT__33   = 34
	RParserT__34   = 35
	RParserT__35   = 36
	RParserT__36   = 37
	RParserT__37   = 38
	RParserT__38   = 39
	RParserT__39   = 40
	RParserT__40   = 41
	RParserT__41   = 42
	RParserT__42   = 43
	RParserT__43   = 44
	RParserT__44   = 45
	RParserT__45   = 46
	RParserT__46   = 47
	RParserT__47   = 48
	RParserT__48   = 49
	RParserT__49   = 50
	RParserT__50   = 51
	RParserT__51   = 52
	RParserT__52   = 53
	RParserT__53   = 54
	RParserHEX     = 55
	RParserINT     = 56
	RParserFLOAT   = 57
	RParserCOMPLEX = 58
	RParserSTRING  = 59
	RParserID      = 60
	RParserUSER_OP = 61
	RParserNL      = 62
	RParserWS      = 63
)

// RParser rules.
const (
	RParserRULE_prog     = 0
	RParserRULE_expr     = 1
	RParserRULE_exprlist = 2
	RParserRULE_formlist = 3
	RParserRULE_form     = 4
	RParserRULE_sublist  = 5
	RParserRULE_sub      = 6
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(RParserEOF, 0)
}

func (s *ProgContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ProgContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RParserNL)
}

func (s *ProgContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RParserNL, i)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *RParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RParserRULE_prog)
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RParserT__9)|(1<<RParserT__10)|(1<<RParserT__20)|(1<<RParserT__25))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RParserT__32-33))|(1<<(RParserT__33-33))|(1<<(RParserT__35-33))|(1<<(RParserT__37-33))|(1<<(RParserT__39-33))|(1<<(RParserT__41-33))|(1<<(RParserT__42-33))|(1<<(RParserT__43-33))|(1<<(RParserT__44-33))|(1<<(RParserT__45-33))|(1<<(RParserT__46-33))|(1<<(RParserT__47-33))|(1<<(RParserT__48-33))|(1<<(RParserT__49-33))|(1<<(RParserT__50-33))|(1<<(RParserT__51-33))|(1<<(RParserHEX-33))|(1<<(RParserINT-33))|(1<<(RParserFLOAT-33))|(1<<(RParserCOMPLEX-33))|(1<<(RParserSTRING-33))|(1<<(RParserID-33))|(1<<(RParserNL-33)))) != 0) {
		p.SetState(18)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RParserT__9, RParserT__10, RParserT__20, RParserT__25, RParserT__32, RParserT__33, RParserT__35, RParserT__37, RParserT__39, RParserT__41, RParserT__42, RParserT__43, RParserT__44, RParserT__45, RParserT__46, RParserT__47, RParserT__48, RParserT__49, RParserT__50, RParserT__51, RParserHEX, RParserINT, RParserFLOAT, RParserCOMPLEX, RParserSTRING, RParserID:
			{
				p.SetState(14)
				p.expr(0)
			}
			p.SetState(15)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RParserT__0 || _la == RParserNL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		case RParserNL:
			{
				p.SetState(17)
				p.Match(RParserNL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(RParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Formlist() IFormlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormlistContext)
}

func (s *ExprContext) Exprlist() IExprlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprlistContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(RParserID, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(RParserSTRING, 0)
}

func (s *ExprContext) HEX() antlr.TerminalNode {
	return s.GetToken(RParserHEX, 0)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(RParserINT, 0)
}

func (s *ExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RParserFLOAT, 0)
}

func (s *ExprContext) COMPLEX() antlr.TerminalNode {
	return s.GetToken(RParserCOMPLEX, 0)
}

func (s *ExprContext) USER_OP() antlr.TerminalNode {
	return s.GetToken(RParserUSER_OP, 0)
}

func (s *ExprContext) Sublist() ISublistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISublistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISublistContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *RParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, RParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RParserT__9 || _la == RParserT__10) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(27)
			p.expr(36)
		}

	case 2:
		{
			p.SetState(28)
			p.Match(RParserT__20)
		}
		{
			p.SetState(29)
			p.expr(30)
		}

	case 3:
		{
			p.SetState(30)
			p.Match(RParserT__25)
		}
		{
			p.SetState(31)
			p.expr(27)
		}

	case 4:
		{
			p.SetState(32)
			p.Match(RParserT__32)
		}
		{
			p.SetState(33)
			p.Match(RParserT__33)
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RParserT__53 || _la == RParserID {
			{
				p.SetState(34)
				p.Formlist()
			}

		}
		{
			p.SetState(37)
			p.Match(RParserT__34)
		}
		{
			p.SetState(38)
			p.expr(24)
		}

	case 5:
		{
			p.SetState(39)
			p.Match(RParserT__35)
		}
		{
			p.SetState(40)
			p.Exprlist()
		}
		{
			p.SetState(41)
			p.Match(RParserT__36)
		}

	case 6:
		{
			p.SetState(43)
			p.Match(RParserT__37)
		}
		{
			p.SetState(44)
			p.Match(RParserT__33)
		}
		{
			p.SetState(45)
			p.expr(0)
		}
		{
			p.SetState(46)
			p.Match(RParserT__34)
		}
		{
			p.SetState(47)
			p.expr(21)
		}

	case 7:
		{
			p.SetState(49)
			p.Match(RParserT__37)
		}
		{
			p.SetState(50)
			p.Match(RParserT__33)
		}
		{
			p.SetState(51)
			p.expr(0)
		}
		{
			p.SetState(52)
			p.Match(RParserT__34)
		}
		{
			p.SetState(53)
			p.expr(0)
		}
		{
			p.SetState(54)
			p.Match(RParserT__38)
		}
		{
			p.SetState(55)
			p.expr(20)
		}

	case 8:
		{
			p.SetState(57)
			p.Match(RParserT__39)
		}
		{
			p.SetState(58)
			p.Match(RParserT__33)
		}
		{
			p.SetState(59)
			p.Match(RParserID)
		}
		{
			p.SetState(60)
			p.Match(RParserT__40)
		}
		{
			p.SetState(61)
			p.expr(0)
		}
		{
			p.SetState(62)
			p.Match(RParserT__34)
		}
		{
			p.SetState(63)
			p.expr(19)
		}

	case 9:
		{
			p.SetState(65)
			p.Match(RParserT__41)
		}
		{
			p.SetState(66)
			p.Match(RParserT__33)
		}
		{
			p.SetState(67)
			p.expr(0)
		}
		{
			p.SetState(68)
			p.Match(RParserT__34)
		}
		{
			p.SetState(69)
			p.expr(18)
		}

	case 10:
		{
			p.SetState(71)
			p.Match(RParserT__42)
		}
		{
			p.SetState(72)
			p.expr(17)
		}

	case 11:
		{
			p.SetState(73)
			p.Match(RParserT__43)
		}
		{
			p.SetState(74)
			p.expr(16)
		}

	case 12:
		{
			p.SetState(75)
			p.Match(RParserT__44)
		}

	case 13:
		{
			p.SetState(76)
			p.Match(RParserT__45)
		}

	case 14:
		{
			p.SetState(77)
			p.Match(RParserT__33)
		}
		{
			p.SetState(78)
			p.expr(0)
		}
		{
			p.SetState(79)
			p.Match(RParserT__34)
		}

	case 15:
		{
			p.SetState(81)
			p.Match(RParserID)
		}

	case 16:
		{
			p.SetState(82)
			p.Match(RParserSTRING)
		}

	case 17:
		{
			p.SetState(83)
			p.Match(RParserHEX)
		}

	case 18:
		{
			p.SetState(84)
			p.Match(RParserINT)
		}

	case 19:
		{
			p.SetState(85)
			p.Match(RParserFLOAT)
		}

	case 20:
		{
			p.SetState(86)
			p.Match(RParserCOMPLEX)
		}

	case 21:
		{
			p.SetState(87)
			p.Match(RParserT__46)
		}

	case 22:
		{
			p.SetState(88)
			p.Match(RParserT__47)
		}

	case 23:
		{
			p.SetState(89)
			p.Match(RParserT__48)
		}

	case 24:
		{
			p.SetState(90)
			p.Match(RParserT__49)
		}

	case 25:
		{
			p.SetState(91)
			p.Match(RParserT__50)
		}

	case 26:
		{
			p.SetState(92)
			p.Match(RParserT__51)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 39)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 39)", ""))
				}
				p.SetState(96)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__4 || _la == RParserT__5) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(97)
					p.expr(40)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 38)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 38)", ""))
				}
				p.SetState(99)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__6 || _la == RParserT__7) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(100)
					p.expr(39)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 37)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 37)", ""))
				}
				{
					p.SetState(102)
					p.Match(RParserT__8)
				}
				{
					p.SetState(103)
					p.expr(37)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 35)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 35)", ""))
				}
				{
					p.SetState(105)
					p.Match(RParserT__11)
				}
				{
					p.SetState(106)
					p.expr(36)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
				}
				{
					p.SetState(108)
					p.Match(RParserUSER_OP)
				}
				{
					p.SetState(109)
					p.expr(35)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
				}
				p.SetState(111)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__12 || _la == RParserT__13) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(112)
					p.expr(34)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
				}
				p.SetState(114)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__9 || _la == RParserT__10) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(115)
					p.expr(33)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
				}
				p.SetState(117)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RParserT__14)|(1<<RParserT__15)|(1<<RParserT__16)|(1<<RParserT__17)|(1<<RParserT__18)|(1<<RParserT__19))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(118)
					p.expr(32)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
				}
				p.SetState(120)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__21 || _la == RParserT__22) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(121)
					p.expr(30)
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
				}
				p.SetState(123)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RParserT__23 || _la == RParserT__24) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(124)
					p.expr(29)
				}

			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(126)
					p.Match(RParserT__25)
				}
				{
					p.SetState(127)
					p.expr(27)
				}

			case 12:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				p.SetState(129)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(RParserT__26-27))|(1<<(RParserT__27-27))|(1<<(RParserT__28-27))|(1<<(RParserT__29-27))|(1<<(RParserT__30-27))|(1<<(RParserT__31-27)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(130)
					p.expr(26)
				}

			case 13:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 41)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 41)", ""))
				}
				{
					p.SetState(132)
					p.Match(RParserT__1)
				}
				{
					p.SetState(133)
					p.Sublist()
				}
				{
					p.SetState(134)
					p.Match(RParserT__2)
				}
				{
					p.SetState(135)
					p.Match(RParserT__2)
				}

			case 14:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 40)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 40)", ""))
				}
				{
					p.SetState(138)
					p.Match(RParserT__3)
				}
				{
					p.SetState(139)
					p.Sublist()
				}
				{
					p.SetState(140)
					p.Match(RParserT__2)
				}

			case 15:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RParserRULE_expr)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(143)
					p.Match(RParserT__33)
				}
				{
					p.SetState(144)
					p.Sublist()
				}
				{
					p.SetState(145)
					p.Match(RParserT__34)
				}

			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IExprlistContext is an interface to support dynamic dispatch.
type IExprlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprlistContext differentiates from other interfaces.
	IsExprlistContext()
}

type ExprlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprlistContext() *ExprlistContext {
	var p = new(ExprlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_exprlist
	return p
}

func (*ExprlistContext) IsExprlistContext() {}

func NewExprlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprlistContext {
	var p = new(ExprlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_exprlist

	return p
}

func (s *ExprlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprlistContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprlistContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprlistContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RParserNL)
}

func (s *ExprlistContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RParserNL, i)
}

func (s *ExprlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterExprlist(s)
	}
}

func (s *ExprlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitExprlist(s)
	}
}

func (p *RParser) Exprlist() (localctx IExprlistContext) {
	localctx = NewExprlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RParserRULE_exprlist)
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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RParserT__9, RParserT__10, RParserT__20, RParserT__25, RParserT__32, RParserT__33, RParserT__35, RParserT__37, RParserT__39, RParserT__41, RParserT__42, RParserT__43, RParserT__44, RParserT__45, RParserT__46, RParserT__47, RParserT__48, RParserT__49, RParserT__50, RParserT__51, RParserHEX, RParserINT, RParserFLOAT, RParserCOMPLEX, RParserSTRING, RParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.expr(0)
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RParserT__0 || _la == RParserNL {
			p.SetState(153)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RParserT__0 || _la == RParserNL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RParserT__9)|(1<<RParserT__10)|(1<<RParserT__20)|(1<<RParserT__25))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RParserT__32-33))|(1<<(RParserT__33-33))|(1<<(RParserT__35-33))|(1<<(RParserT__37-33))|(1<<(RParserT__39-33))|(1<<(RParserT__41-33))|(1<<(RParserT__42-33))|(1<<(RParserT__43-33))|(1<<(RParserT__44-33))|(1<<(RParserT__45-33))|(1<<(RParserT__46-33))|(1<<(RParserT__47-33))|(1<<(RParserT__48-33))|(1<<(RParserT__49-33))|(1<<(RParserT__50-33))|(1<<(RParserT__51-33))|(1<<(RParserHEX-33))|(1<<(RParserINT-33))|(1<<(RParserFLOAT-33))|(1<<(RParserCOMPLEX-33))|(1<<(RParserSTRING-33))|(1<<(RParserID-33)))) != 0) {
				{
					p.SetState(154)
					p.expr(0)
				}

			}

			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case RParserT__36:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormlistContext is an interface to support dynamic dispatch.
type IFormlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormlistContext differentiates from other interfaces.
	IsFormlistContext()
}

type FormlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormlistContext() *FormlistContext {
	var p = new(FormlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_formlist
	return p
}

func (*FormlistContext) IsFormlistContext() {}

func NewFormlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormlistContext {
	var p = new(FormlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_formlist

	return p
}

func (s *FormlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FormlistContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *FormlistContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *FormlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterFormlist(s)
	}
}

func (s *FormlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitFormlist(s)
	}
}

func (p *RParser) Formlist() (localctx IFormlistContext) {
	localctx = NewFormlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RParserRULE_formlist)
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
		p.SetState(165)
		p.Form()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RParserT__52 {
		{
			p.SetState(166)
			p.Match(RParserT__52)
		}
		{
			p.SetState(167)
			p.Form()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormContext is an interface to support dynamic dispatch.
type IFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormContext differentiates from other interfaces.
	IsFormContext()
}

type FormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormContext() *FormContext {
	var p = new(FormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_form
	return p
}

func (*FormContext) IsFormContext() {}

func NewFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormContext {
	var p = new(FormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_form

	return p
}

func (s *FormContext) GetParser() antlr.Parser { return s.parser }

func (s *FormContext) ID() antlr.TerminalNode {
	return s.GetToken(RParserID, 0)
}

func (s *FormContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterForm(s)
	}
}

func (s *FormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitForm(s)
	}
}

func (p *RParser) Form() (localctx IFormContext) {
	localctx = NewFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RParserRULE_form)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(RParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(RParserID)
		}
		{
			p.SetState(175)
			p.Match(RParserT__28)
		}
		{
			p.SetState(176)
			p.expr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Match(RParserT__53)
		}

	}

	return localctx
}

// ISublistContext is an interface to support dynamic dispatch.
type ISublistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSublistContext differentiates from other interfaces.
	IsSublistContext()
}

type SublistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySublistContext() *SublistContext {
	var p = new(SublistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_sublist
	return p
}

func (*SublistContext) IsSublistContext() {}

func NewSublistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SublistContext {
	var p = new(SublistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_sublist

	return p
}

func (s *SublistContext) GetParser() antlr.Parser { return s.parser }

func (s *SublistContext) AllSub() []ISubContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubContext)(nil)).Elem())
	var tst = make([]ISubContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubContext)
		}
	}

	return tst
}

func (s *SublistContext) Sub(i int) ISubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubContext)
}

func (s *SublistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SublistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SublistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterSublist(s)
	}
}

func (s *SublistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitSublist(s)
	}
}

func (p *RParser) Sublist() (localctx ISublistContext) {
	localctx = NewSublistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RParserRULE_sublist)
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
		p.SetState(180)
		p.Sub()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RParserT__52 {
		{
			p.SetState(181)
			p.Match(RParserT__52)
		}
		{
			p.SetState(182)
			p.Sub()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISubContext is an interface to support dynamic dispatch.
type ISubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubContext differentiates from other interfaces.
	IsSubContext()
}

type SubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubContext() *SubContext {
	var p = new(SubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RParserRULE_sub
	return p
}

func (*SubContext) IsSubContext() {}

func NewSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubContext {
	var p = new(SubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RParserRULE_sub

	return p
}

func (s *SubContext) GetParser() antlr.Parser { return s.parser }

func (s *SubContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubContext) ID() antlr.TerminalNode {
	return s.GetToken(RParserID, 0)
}

func (s *SubContext) STRING() antlr.TerminalNode {
	return s.GetToken(RParserSTRING, 0)
}

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RListener); ok {
		listenerT.ExitSub(s)
	}
}

func (p *RParser) Sub() (localctx ISubContext) {
	localctx = NewSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RParserRULE_sub)

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

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(RParserID)
		}
		{
			p.SetState(190)
			p.Match(RParserT__28)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Match(RParserID)
		}
		{
			p.SetState(192)
			p.Match(RParserT__28)
		}
		{
			p.SetState(193)
			p.expr(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(RParserSTRING)
		}
		{
			p.SetState(195)
			p.Match(RParserT__28)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.Match(RParserSTRING)
		}
		{
			p.SetState(197)
			p.Match(RParserT__28)
		}
		{
			p.SetState(198)
			p.expr(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(199)
			p.Match(RParserT__46)
		}
		{
			p.SetState(200)
			p.Match(RParserT__28)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(201)
			p.Match(RParserT__46)
		}
		{
			p.SetState(202)
			p.Match(RParserT__28)
		}
		{
			p.SetState(203)
			p.expr(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(204)
			p.Match(RParserT__53)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)

	}

	return localctx
}

func (p *RParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 39)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 38)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 37)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 35)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 33)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 41)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 40)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 23)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
