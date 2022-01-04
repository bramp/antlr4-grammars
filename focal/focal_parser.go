// Code generated from focal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package focal // focal
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 247,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 6, 2, 52, 10, 2, 13, 2, 14, 2, 53, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 76, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 82, 10, 7, 12, 7, 14, 7, 85, 11, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 7, 8, 92, 10, 8, 12, 8, 14, 8, 95, 11, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 101, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 111, 10, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 124, 10, 12, 3, 13, 3, 13, 5,
	13, 128, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 141, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7,
	15, 147, 10, 15, 12, 15, 14, 15, 150, 11, 15, 3, 15, 3, 15, 5, 15, 154,
	10, 15, 3, 16, 3, 16, 6, 16, 158, 10, 16, 13, 16, 14, 16, 159, 3, 16, 6,
	16, 163, 10, 16, 13, 16, 14, 16, 164, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	5, 16, 172, 10, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 179, 10,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 186, 10, 20, 12, 20, 14,
	20, 189, 11, 20, 3, 21, 3, 21, 3, 21, 7, 21, 194, 10, 21, 12, 21, 14, 21,
	197, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 223, 10, 22, 3, 23, 3, 23, 3,
	23, 5, 23, 228, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 5, 24, 240, 10, 24, 3, 25, 5, 25, 243, 10, 25, 3,
	25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 8, 3, 2, 4, 5, 3, 2,
	7, 8, 3, 2, 10, 11, 3, 2, 15, 16, 3, 2, 17, 18, 3, 2, 20, 21, 2, 264, 2,
	51, 3, 2, 2, 2, 4, 55, 3, 2, 2, 2, 6, 58, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2,
	10, 75, 3, 2, 2, 2, 12, 77, 3, 2, 2, 2, 14, 86, 3, 2, 2, 2, 16, 96, 3,
	2, 2, 2, 18, 102, 3, 2, 2, 2, 20, 115, 3, 2, 2, 2, 22, 117, 3, 2, 2, 2,
	24, 125, 3, 2, 2, 2, 26, 129, 3, 2, 2, 2, 28, 142, 3, 2, 2, 2, 30, 171,
	3, 2, 2, 2, 32, 173, 3, 2, 2, 2, 34, 175, 3, 2, 2, 2, 36, 180, 3, 2, 2,
	2, 38, 182, 3, 2, 2, 2, 40, 190, 3, 2, 2, 2, 42, 222, 3, 2, 2, 2, 44, 224,
	3, 2, 2, 2, 46, 239, 3, 2, 2, 2, 48, 242, 3, 2, 2, 2, 50, 52, 5, 4, 3,
	2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54,
	3, 2, 2, 2, 54, 3, 3, 2, 2, 2, 55, 56, 5, 8, 5, 2, 56, 57, 5, 10, 6, 2,
	57, 5, 3, 2, 2, 2, 58, 59, 7, 37, 2, 2, 59, 7, 3, 2, 2, 2, 60, 61, 5, 6,
	4, 2, 61, 62, 7, 3, 2, 2, 62, 63, 7, 37, 2, 2, 63, 9, 3, 2, 2, 2, 64, 76,
	5, 12, 7, 2, 65, 76, 5, 16, 9, 2, 66, 76, 5, 18, 10, 2, 67, 76, 5, 22,
	12, 2, 68, 76, 5, 20, 11, 2, 69, 76, 5, 24, 13, 2, 70, 76, 5, 26, 14, 2,
	71, 76, 5, 28, 15, 2, 72, 76, 5, 32, 17, 2, 73, 76, 5, 34, 18, 2, 74, 76,
	5, 36, 19, 2, 75, 64, 3, 2, 2, 2, 75, 65, 3, 2, 2, 2, 75, 66, 3, 2, 2,
	2, 75, 67, 3, 2, 2, 2, 75, 68, 3, 2, 2, 2, 75, 69, 3, 2, 2, 2, 75, 70,
	3, 2, 2, 2, 75, 71, 3, 2, 2, 2, 75, 72, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 74, 3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77, 78, 9, 2, 2, 2, 78, 83, 5,
	14, 8, 2, 79, 80, 7, 6, 2, 2, 80, 82, 5, 14, 8, 2, 81, 79, 3, 2, 2, 2,
	82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 13, 3,
	2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7, 41, 2, 2, 87, 88, 7, 6, 2, 2, 88,
	93, 7, 36, 2, 2, 89, 90, 7, 6, 2, 2, 90, 92, 7, 36, 2, 2, 91, 89, 3, 2,
	2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 15,
	3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 100, 9, 3, 2, 2, 97, 101, 7, 9, 2,
	2, 98, 101, 5, 6, 4, 2, 99, 101, 5, 8, 5, 2, 100, 97, 3, 2, 2, 2, 100,
	98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 17, 3, 2, 2, 2, 102, 103, 9,
	4, 2, 2, 103, 104, 7, 36, 2, 2, 104, 105, 7, 12, 2, 2, 105, 106, 5, 38,
	20, 2, 106, 107, 7, 6, 2, 2, 107, 110, 5, 38, 20, 2, 108, 109, 7, 6, 2,
	2, 109, 111, 5, 38, 20, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 113, 7, 13, 2, 2, 113, 114, 5, 10, 6, 2, 114,
	19, 3, 2, 2, 2, 115, 116, 7, 14, 2, 2, 116, 21, 3, 2, 2, 2, 117, 118, 9,
	5, 2, 2, 118, 119, 7, 36, 2, 2, 119, 120, 7, 12, 2, 2, 120, 123, 5, 38,
	20, 2, 121, 122, 7, 13, 2, 2, 122, 124, 5, 10, 6, 2, 123, 121, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 23, 3, 2, 2, 2, 125, 127, 9, 6, 2, 2, 126,
	128, 5, 8, 5, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 25, 3,
	2, 2, 2, 129, 130, 7, 19, 2, 2, 130, 131, 5, 38, 20, 2, 131, 132, 5, 8,
	5, 2, 132, 133, 7, 6, 2, 2, 133, 134, 5, 8, 5, 2, 134, 135, 3, 2, 2, 2,
	135, 136, 7, 6, 2, 2, 136, 137, 5, 8, 5, 2, 137, 140, 3, 2, 2, 2, 138,
	139, 7, 13, 2, 2, 139, 141, 5, 10, 6, 2, 140, 138, 3, 2, 2, 2, 140, 141,
	3, 2, 2, 2, 141, 27, 3, 2, 2, 2, 142, 143, 9, 7, 2, 2, 143, 148, 5, 30,
	16, 2, 144, 145, 7, 6, 2, 2, 145, 147, 5, 30, 16, 2, 146, 144, 3, 2, 2,
	2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	153, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152, 7, 13, 2, 2, 152, 154,
	5, 10, 6, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 29, 3, 2,
	2, 2, 155, 172, 5, 38, 20, 2, 156, 158, 7, 22, 2, 2, 157, 156, 3, 2, 2,
	2, 158, 159, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	172, 3, 2, 2, 2, 161, 163, 7, 23, 2, 2, 162, 161, 3, 2, 2, 2, 163, 164,
	3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 172, 3, 2,
	2, 2, 166, 172, 7, 41, 2, 2, 167, 168, 7, 24, 2, 2, 168, 169, 7, 37, 2,
	2, 169, 170, 7, 3, 2, 2, 170, 172, 7, 37, 2, 2, 171, 155, 3, 2, 2, 2, 171,
	157, 3, 2, 2, 2, 171, 162, 3, 2, 2, 2, 171, 166, 3, 2, 2, 2, 171, 167,
	3, 2, 2, 2, 172, 31, 3, 2, 2, 2, 173, 174, 7, 25, 2, 2, 174, 33, 3, 2,
	2, 2, 175, 178, 7, 26, 2, 2, 176, 179, 5, 6, 4, 2, 177, 179, 5, 8, 5, 2,
	178, 176, 3, 2, 2, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	35, 3, 2, 2, 2, 180, 181, 7, 42, 2, 2, 181, 37, 3, 2, 2, 2, 182, 187, 5,
	40, 21, 2, 183, 184, 7, 34, 2, 2, 184, 186, 5, 40, 21, 2, 185, 183, 3,
	2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2,
	2, 188, 39, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 195, 5, 42, 22, 2, 191,
	192, 7, 35, 2, 2, 192, 194, 5, 42, 22, 2, 193, 191, 3, 2, 2, 2, 194, 197,
	3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 41, 3, 2,
	2, 2, 197, 195, 3, 2, 2, 2, 198, 199, 7, 27, 2, 2, 199, 200, 5, 38, 20,
	2, 200, 201, 7, 28, 2, 2, 201, 223, 3, 2, 2, 2, 202, 203, 7, 29, 2, 2,
	203, 204, 5, 38, 20, 2, 204, 205, 7, 30, 2, 2, 205, 223, 3, 2, 2, 2, 206,
	207, 7, 31, 2, 2, 207, 208, 5, 38, 20, 2, 208, 209, 7, 32, 2, 2, 209, 223,
	3, 2, 2, 2, 210, 223, 5, 44, 23, 2, 211, 223, 7, 36, 2, 2, 212, 213, 7,
	36, 2, 2, 213, 214, 7, 27, 2, 2, 214, 215, 5, 38, 20, 2, 215, 216, 7, 28,
	2, 2, 216, 223, 3, 2, 2, 2, 217, 218, 7, 40, 2, 2, 218, 219, 7, 27, 2,
	2, 219, 220, 5, 38, 20, 2, 220, 221, 7, 28, 2, 2, 221, 223, 3, 2, 2, 2,
	222, 198, 3, 2, 2, 2, 222, 202, 3, 2, 2, 2, 222, 206, 3, 2, 2, 2, 222,
	210, 3, 2, 2, 2, 222, 211, 3, 2, 2, 2, 222, 212, 3, 2, 2, 2, 222, 217,
	3, 2, 2, 2, 223, 43, 3, 2, 2, 2, 224, 227, 5, 46, 24, 2, 225, 226, 7, 33,
	2, 2, 226, 228, 5, 48, 25, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2,
	2, 228, 45, 3, 2, 2, 2, 229, 240, 5, 48, 25, 2, 230, 231, 5, 48, 25, 2,
	231, 232, 7, 3, 2, 2, 232, 240, 3, 2, 2, 2, 233, 234, 7, 3, 2, 2, 234,
	240, 5, 48, 25, 2, 235, 236, 5, 48, 25, 2, 236, 237, 7, 3, 2, 2, 237, 238,
	5, 48, 25, 2, 238, 240, 3, 2, 2, 2, 239, 229, 3, 2, 2, 2, 239, 230, 3,
	2, 2, 2, 239, 233, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2, 240, 47, 3, 2, 2,
	2, 241, 243, 7, 34, 2, 2, 242, 241, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243,
	244, 3, 2, 2, 2, 244, 245, 7, 37, 2, 2, 245, 49, 3, 2, 2, 2, 23, 53, 75,
	83, 93, 100, 110, 123, 127, 140, 148, 153, 159, 164, 171, 178, 187, 195,
	222, 227, 239, 242,
}
var literalNames = []string{
	"", "'.'", "'ASK'", "'A'", "','", "'DO'", "'D'", "'all'", "'FOR'", "'F'",
	"'='", "';'", "'QUIT'", "'SET'", "'S'", "'GOTO'", "'G'", "'IF'", "'TYPE'",
	"'T'", "'!'", "'#'", "'%'", "'RETURN'", "'WRITE'", "'('", "')'", "'['",
	"']'", "'<'", "'>'", "'e'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "PLUSMIN", "MULOP",
	"VARIABLE", "INTEGER", "ALPHA", "DIGIT", "BUILTIN", "STRING_LITERAL", "COMMENT",
	"WS",
}

var ruleNames = []string{
	"prog", "statement", "grpnum", "linenum", "command", "ask", "askpair",
	"do_", "for_", "quit", "set_", "goto_", "if_", "type_", "typeexpression",
	"return_", "write_", "comment", "expression", "primary", "term", "number",
	"mantissa", "signed_",
}

type focalParser struct {
	*antlr.BaseParser
}

// NewfocalParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *focalParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfocalParser(input antlr.TokenStream) *focalParser {
	this := new(focalParser)
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
	this.GrammarFileName = "focal.g4"

	return this
}

// focalParser tokens.
const (
	focalParserEOF            = antlr.TokenEOF
	focalParserT__0           = 1
	focalParserT__1           = 2
	focalParserT__2           = 3
	focalParserT__3           = 4
	focalParserT__4           = 5
	focalParserT__5           = 6
	focalParserT__6           = 7
	focalParserT__7           = 8
	focalParserT__8           = 9
	focalParserT__9           = 10
	focalParserT__10          = 11
	focalParserT__11          = 12
	focalParserT__12          = 13
	focalParserT__13          = 14
	focalParserT__14          = 15
	focalParserT__15          = 16
	focalParserT__16          = 17
	focalParserT__17          = 18
	focalParserT__18          = 19
	focalParserT__19          = 20
	focalParserT__20          = 21
	focalParserT__21          = 22
	focalParserT__22          = 23
	focalParserT__23          = 24
	focalParserT__24          = 25
	focalParserT__25          = 26
	focalParserT__26          = 27
	focalParserT__27          = 28
	focalParserT__28          = 29
	focalParserT__29          = 30
	focalParserT__30          = 31
	focalParserPLUSMIN        = 32
	focalParserMULOP          = 33
	focalParserVARIABLE       = 34
	focalParserINTEGER        = 35
	focalParserALPHA          = 36
	focalParserDIGIT          = 37
	focalParserBUILTIN        = 38
	focalParserSTRING_LITERAL = 39
	focalParserCOMMENT        = 40
	focalParserWS             = 41
)

// focalParser rules.
const (
	focalParserRULE_prog           = 0
	focalParserRULE_statement      = 1
	focalParserRULE_grpnum         = 2
	focalParserRULE_linenum        = 3
	focalParserRULE_command        = 4
	focalParserRULE_ask            = 5
	focalParserRULE_askpair        = 6
	focalParserRULE_do_            = 7
	focalParserRULE_for_           = 8
	focalParserRULE_quit           = 9
	focalParserRULE_set_           = 10
	focalParserRULE_goto_          = 11
	focalParserRULE_if_            = 12
	focalParserRULE_type_          = 13
	focalParserRULE_typeexpression = 14
	focalParserRULE_return_        = 15
	focalParserRULE_write_         = 16
	focalParserRULE_comment        = 17
	focalParserRULE_expression     = 18
	focalParserRULE_primary        = 19
	focalParserRULE_term           = 20
	focalParserRULE_number         = 21
	focalParserRULE_mantissa       = 22
	focalParserRULE_signed_        = 23
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
	p.RuleIndex = focalParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *focalParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, focalParserRULE_prog)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == focalParserINTEGER {
		{
			p.SetState(48)
			p.Statement()
		}

		p.SetState(51)
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
	p.RuleIndex = focalParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Linenum() ILinenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinenumContext)
}

func (s *StatementContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *focalParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, focalParserRULE_statement)

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
		p.SetState(53)
		p.Linenum()
	}
	{
		p.SetState(54)
		p.Command()
	}

	return localctx
}

// IGrpnumContext is an interface to support dynamic dispatch.
type IGrpnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGrpnumContext differentiates from other interfaces.
	IsGrpnumContext()
}

type GrpnumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrpnumContext() *GrpnumContext {
	var p = new(GrpnumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_grpnum
	return p
}

func (*GrpnumContext) IsGrpnumContext() {}

func NewGrpnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrpnumContext {
	var p = new(GrpnumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_grpnum

	return p
}

func (s *GrpnumContext) GetParser() antlr.Parser { return s.parser }

func (s *GrpnumContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(focalParserINTEGER, 0)
}

func (s *GrpnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrpnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrpnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterGrpnum(s)
	}
}

func (s *GrpnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitGrpnum(s)
	}
}

func (p *focalParser) Grpnum() (localctx IGrpnumContext) {
	this := p
	_ = this

	localctx = NewGrpnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, focalParserRULE_grpnum)

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
		p.SetState(56)
		p.Match(focalParserINTEGER)
	}

	return localctx
}

// ILinenumContext is an interface to support dynamic dispatch.
type ILinenumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinenumContext differentiates from other interfaces.
	IsLinenumContext()
}

type LinenumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinenumContext() *LinenumContext {
	var p = new(LinenumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_linenum
	return p
}

func (*LinenumContext) IsLinenumContext() {}

func NewLinenumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinenumContext {
	var p = new(LinenumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_linenum

	return p
}

func (s *LinenumContext) GetParser() antlr.Parser { return s.parser }

func (s *LinenumContext) Grpnum() IGrpnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGrpnumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGrpnumContext)
}

func (s *LinenumContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(focalParserINTEGER, 0)
}

func (s *LinenumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinenumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinenumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterLinenum(s)
	}
}

func (s *LinenumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitLinenum(s)
	}
}

func (p *focalParser) Linenum() (localctx ILinenumContext) {
	this := p
	_ = this

	localctx = NewLinenumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, focalParserRULE_linenum)

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
		p.SetState(58)
		p.Grpnum()
	}
	{
		p.SetState(59)
		p.Match(focalParserT__0)
	}
	{
		p.SetState(60)
		p.Match(focalParserINTEGER)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Ask() IAskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAskContext)
}

func (s *CommandContext) Do_() IDo_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDo_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDo_Context)
}

func (s *CommandContext) For_() IFor_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_Context)
}

func (s *CommandContext) Set_() ISet_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_Context)
}

func (s *CommandContext) Quit() IQuitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuitContext)
}

func (s *CommandContext) Goto_() IGoto_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoto_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoto_Context)
}

func (s *CommandContext) If_() IIf_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_Context)
}

func (s *CommandContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *CommandContext) Return_() IReturn_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_Context)
}

func (s *CommandContext) Write_() IWrite_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWrite_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWrite_Context)
}

func (s *CommandContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *focalParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, focalParserRULE_command)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case focalParserT__1, focalParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Ask()
		}

	case focalParserT__4, focalParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Do_()
		}

	case focalParserT__7, focalParserT__8:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.For_()
		}

	case focalParserT__12, focalParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Set_()
		}

	case focalParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.Quit()
		}

	case focalParserT__14, focalParserT__15:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.Goto_()
		}

	case focalParserT__16:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(68)
			p.If_()
		}

	case focalParserT__17, focalParserT__18:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(69)
			p.Type_()
		}

	case focalParserT__22:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(70)
			p.Return_()
		}

	case focalParserT__23:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(71)
			p.Write_()
		}

	case focalParserCOMMENT:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(72)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAskContext is an interface to support dynamic dispatch.
type IAskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAskContext differentiates from other interfaces.
	IsAskContext()
}

type AskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAskContext() *AskContext {
	var p = new(AskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_ask
	return p
}

func (*AskContext) IsAskContext() {}

func NewAskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AskContext {
	var p = new(AskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_ask

	return p
}

func (s *AskContext) GetParser() antlr.Parser { return s.parser }

func (s *AskContext) AllAskpair() []IAskpairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAskpairContext)(nil)).Elem())
	var tst = make([]IAskpairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAskpairContext)
		}
	}

	return tst
}

func (s *AskContext) Askpair(i int) IAskpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAskpairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAskpairContext)
}

func (s *AskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterAsk(s)
	}
}

func (s *AskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitAsk(s)
	}
}

func (p *focalParser) Ask() (localctx IAskContext) {
	this := p
	_ = this

	localctx = NewAskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, focalParserRULE_ask)
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
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__1 || _la == focalParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(76)
		p.Askpair()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == focalParserT__3 {
		{
			p.SetState(77)
			p.Match(focalParserT__3)
		}
		{
			p.SetState(78)
			p.Askpair()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAskpairContext is an interface to support dynamic dispatch.
type IAskpairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAskpairContext differentiates from other interfaces.
	IsAskpairContext()
}

type AskpairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAskpairContext() *AskpairContext {
	var p = new(AskpairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_askpair
	return p
}

func (*AskpairContext) IsAskpairContext() {}

func NewAskpairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AskpairContext {
	var p = new(AskpairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_askpair

	return p
}

func (s *AskpairContext) GetParser() antlr.Parser { return s.parser }

func (s *AskpairContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(focalParserSTRING_LITERAL, 0)
}

func (s *AskpairContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(focalParserVARIABLE)
}

func (s *AskpairContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(focalParserVARIABLE, i)
}

func (s *AskpairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AskpairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AskpairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterAskpair(s)
	}
}

func (s *AskpairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitAskpair(s)
	}
}

func (p *focalParser) Askpair() (localctx IAskpairContext) {
	this := p
	_ = this

	localctx = NewAskpairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, focalParserRULE_askpair)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(focalParserSTRING_LITERAL)
	}
	{
		p.SetState(85)
		p.Match(focalParserT__3)
	}
	{
		p.SetState(86)
		p.Match(focalParserVARIABLE)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(87)
				p.Match(focalParserT__3)
			}
			{
				p.SetState(88)
				p.Match(focalParserVARIABLE)
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IDo_Context is an interface to support dynamic dispatch.
type IDo_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDo_Context differentiates from other interfaces.
	IsDo_Context()
}

type Do_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDo_Context() *Do_Context {
	var p = new(Do_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_do_
	return p
}

func (*Do_Context) IsDo_Context() {}

func NewDo_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Do_Context {
	var p = new(Do_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_do_

	return p
}

func (s *Do_Context) GetParser() antlr.Parser { return s.parser }

func (s *Do_Context) Grpnum() IGrpnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGrpnumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGrpnumContext)
}

func (s *Do_Context) Linenum() ILinenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinenumContext)
}

func (s *Do_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Do_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Do_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterDo_(s)
	}
}

func (s *Do_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitDo_(s)
	}
}

func (p *focalParser) Do_() (localctx IDo_Context) {
	this := p
	_ = this

	localctx = NewDo_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, focalParserRULE_do_)
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
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__4 || _la == focalParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(95)
			p.Match(focalParserT__6)
		}

	case 2:
		{
			p.SetState(96)
			p.Grpnum()
		}

	case 3:
		{
			p.SetState(97)
			p.Linenum()
		}

	}

	return localctx
}

// IFor_Context is an interface to support dynamic dispatch.
type IFor_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_Context differentiates from other interfaces.
	IsFor_Context()
}

type For_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_Context() *For_Context {
	var p = new(For_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_for_
	return p
}

func (*For_Context) IsFor_Context() {}

func NewFor_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_Context {
	var p = new(For_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_for_

	return p
}

func (s *For_Context) GetParser() antlr.Parser { return s.parser }

func (s *For_Context) VARIABLE() antlr.TerminalNode {
	return s.GetToken(focalParserVARIABLE, 0)
}

func (s *For_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *For_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_Context) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *For_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterFor_(s)
	}
}

func (s *For_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitFor_(s)
	}
}

func (p *focalParser) For_() (localctx IFor_Context) {
	this := p
	_ = this

	localctx = NewFor_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, focalParserRULE_for_)
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
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__7 || _la == focalParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(101)
		p.Match(focalParserVARIABLE)
	}
	{
		p.SetState(102)
		p.Match(focalParserT__9)
	}
	{
		p.SetState(103)
		p.Expression()
	}
	{
		p.SetState(104)
		p.Match(focalParserT__3)
	}
	{
		p.SetState(105)
		p.Expression()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserT__3 {
		{
			p.SetState(106)
			p.Match(focalParserT__3)
		}
		{
			p.SetState(107)
			p.Expression()
		}

	}
	{
		p.SetState(110)
		p.Match(focalParserT__10)
	}
	{
		p.SetState(111)
		p.Command()
	}

	return localctx
}

// IQuitContext is an interface to support dynamic dispatch.
type IQuitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuitContext differentiates from other interfaces.
	IsQuitContext()
}

type QuitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuitContext() *QuitContext {
	var p = new(QuitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_quit
	return p
}

func (*QuitContext) IsQuitContext() {}

func NewQuitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuitContext {
	var p = new(QuitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_quit

	return p
}

func (s *QuitContext) GetParser() antlr.Parser { return s.parser }
func (s *QuitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterQuit(s)
	}
}

func (s *QuitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitQuit(s)
	}
}

func (p *focalParser) Quit() (localctx IQuitContext) {
	this := p
	_ = this

	localctx = NewQuitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, focalParserRULE_quit)

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
		p.SetState(113)
		p.Match(focalParserT__11)
	}

	return localctx
}

// ISet_Context is an interface to support dynamic dispatch.
type ISet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_Context differentiates from other interfaces.
	IsSet_Context()
}

type Set_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_Context() *Set_Context {
	var p = new(Set_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_set_
	return p
}

func (*Set_Context) IsSet_Context() {}

func NewSet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_Context {
	var p = new(Set_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_set_

	return p
}

func (s *Set_Context) GetParser() antlr.Parser { return s.parser }

func (s *Set_Context) VARIABLE() antlr.TerminalNode {
	return s.GetToken(focalParserVARIABLE, 0)
}

func (s *Set_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Set_Context) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *Set_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterSet_(s)
	}
}

func (s *Set_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitSet_(s)
	}
}

func (p *focalParser) Set_() (localctx ISet_Context) {
	this := p
	_ = this

	localctx = NewSet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, focalParserRULE_set_)
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
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__12 || _la == focalParserT__13) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(116)
		p.Match(focalParserVARIABLE)
	}
	{
		p.SetState(117)
		p.Match(focalParserT__9)
	}
	{
		p.SetState(118)
		p.Expression()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserT__10 {
		{
			p.SetState(119)
			p.Match(focalParserT__10)
		}
		{
			p.SetState(120)
			p.Command()
		}

	}

	return localctx
}

// IGoto_Context is an interface to support dynamic dispatch.
type IGoto_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGoto_Context differentiates from other interfaces.
	IsGoto_Context()
}

type Goto_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoto_Context() *Goto_Context {
	var p = new(Goto_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_goto_
	return p
}

func (*Goto_Context) IsGoto_Context() {}

func NewGoto_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Goto_Context {
	var p = new(Goto_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_goto_

	return p
}

func (s *Goto_Context) GetParser() antlr.Parser { return s.parser }

func (s *Goto_Context) Linenum() ILinenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinenumContext)
}

func (s *Goto_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Goto_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Goto_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterGoto_(s)
	}
}

func (s *Goto_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitGoto_(s)
	}
}

func (p *focalParser) Goto_() (localctx IGoto_Context) {
	this := p
	_ = this

	localctx = NewGoto_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, focalParserRULE_goto_)
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
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__14 || _la == focalParserT__15) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.Linenum()
		}

	}

	return localctx
}

// IIf_Context is an interface to support dynamic dispatch.
type IIf_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_Context differentiates from other interfaces.
	IsIf_Context()
}

type If_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_Context() *If_Context {
	var p = new(If_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_if_
	return p
}

func (*If_Context) IsIf_Context() {}

func NewIf_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_Context {
	var p = new(If_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_if_

	return p
}

func (s *If_Context) GetParser() antlr.Parser { return s.parser }

func (s *If_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_Context) AllLinenum() []ILinenumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILinenumContext)(nil)).Elem())
	var tst = make([]ILinenumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILinenumContext)
		}
	}

	return tst
}

func (s *If_Context) Linenum(i int) ILinenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinenumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILinenumContext)
}

func (s *If_Context) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *If_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterIf_(s)
	}
}

func (s *If_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitIf_(s)
	}
}

func (p *focalParser) If_() (localctx IIf_Context) {
	this := p
	_ = this

	localctx = NewIf_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, focalParserRULE_if_)
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
		p.SetState(127)
		p.Match(focalParserT__16)
	}

	{
		p.SetState(128)
		p.Expression()
	}
	{
		p.SetState(129)
		p.Linenum()
	}

	{
		p.SetState(130)
		p.Match(focalParserT__3)
	}
	{
		p.SetState(131)
		p.Linenum()
	}

	{
		p.SetState(133)
		p.Match(focalParserT__3)
	}
	{
		p.SetState(134)
		p.Linenum()
	}

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserT__10 {
		{
			p.SetState(136)
			p.Match(focalParserT__10)
		}
		{
			p.SetState(137)
			p.Command()
		}

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
	p.RuleIndex = focalParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) AllTypeexpression() []ITypeexpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeexpressionContext)(nil)).Elem())
	var tst = make([]ITypeexpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeexpressionContext)
		}
	}

	return tst
}

func (s *Type_Context) Typeexpression(i int) ITypeexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeexpressionContext)
}

func (s *Type_Context) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *focalParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, focalParserRULE_type_)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == focalParserT__17 || _la == focalParserT__18) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(141)
		p.Typeexpression()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == focalParserT__3 {
		{
			p.SetState(142)
			p.Match(focalParserT__3)
		}
		{
			p.SetState(143)
			p.Typeexpression()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserT__10 {
		{
			p.SetState(149)
			p.Match(focalParserT__10)
		}
		{
			p.SetState(150)
			p.Command()
		}

	}

	return localctx
}

// ITypeexpressionContext is an interface to support dynamic dispatch.
type ITypeexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeexpressionContext differentiates from other interfaces.
	IsTypeexpressionContext()
}

type TypeexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeexpressionContext() *TypeexpressionContext {
	var p = new(TypeexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_typeexpression
	return p
}

func (*TypeexpressionContext) IsTypeexpressionContext() {}

func NewTypeexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeexpressionContext {
	var p = new(TypeexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_typeexpression

	return p
}

func (s *TypeexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeexpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeexpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(focalParserSTRING_LITERAL, 0)
}

func (s *TypeexpressionContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(focalParserINTEGER)
}

func (s *TypeexpressionContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(focalParserINTEGER, i)
}

func (s *TypeexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterTypeexpression(s)
	}
}

func (s *TypeexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitTypeexpression(s)
	}
}

func (p *focalParser) Typeexpression() (localctx ITypeexpressionContext) {
	this := p
	_ = this

	localctx = NewTypeexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, focalParserRULE_typeexpression)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case focalParserT__0, focalParserT__24, focalParserT__26, focalParserT__28, focalParserPLUSMIN, focalParserVARIABLE, focalParserINTEGER, focalParserBUILTIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Expression()
		}

	case focalParserT__19:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == focalParserT__19 {
			{
				p.SetState(154)
				p.Match(focalParserT__19)
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case focalParserT__20:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == focalParserT__20 {
			{
				p.SetState(159)
				p.Match(focalParserT__20)
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case focalParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.Match(focalParserSTRING_LITERAL)
		}

	case focalParserT__21:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.Match(focalParserT__21)
		}
		{
			p.SetState(166)
			p.Match(focalParserINTEGER)
		}
		{
			p.SetState(167)
			p.Match(focalParserT__0)
		}
		{
			p.SetState(168)
			p.Match(focalParserINTEGER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReturn_Context is an interface to support dynamic dispatch.
type IReturn_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_Context differentiates from other interfaces.
	IsReturn_Context()
}

type Return_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_Context() *Return_Context {
	var p = new(Return_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_return_
	return p
}

func (*Return_Context) IsReturn_Context() {}

func NewReturn_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_Context {
	var p = new(Return_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_return_

	return p
}

func (s *Return_Context) GetParser() antlr.Parser { return s.parser }
func (s *Return_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterReturn_(s)
	}
}

func (s *Return_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitReturn_(s)
	}
}

func (p *focalParser) Return_() (localctx IReturn_Context) {
	this := p
	_ = this

	localctx = NewReturn_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, focalParserRULE_return_)

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
		p.SetState(171)
		p.Match(focalParserT__22)
	}

	return localctx
}

// IWrite_Context is an interface to support dynamic dispatch.
type IWrite_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWrite_Context differentiates from other interfaces.
	IsWrite_Context()
}

type Write_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWrite_Context() *Write_Context {
	var p = new(Write_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_write_
	return p
}

func (*Write_Context) IsWrite_Context() {}

func NewWrite_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Write_Context {
	var p = new(Write_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_write_

	return p
}

func (s *Write_Context) GetParser() antlr.Parser { return s.parser }

func (s *Write_Context) Grpnum() IGrpnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGrpnumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGrpnumContext)
}

func (s *Write_Context) Linenum() ILinenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinenumContext)
}

func (s *Write_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Write_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Write_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterWrite_(s)
	}
}

func (s *Write_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitWrite_(s)
	}
}

func (p *focalParser) Write_() (localctx IWrite_Context) {
	this := p
	_ = this

	localctx = NewWrite_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, focalParserRULE_write_)

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
		p.SetState(173)
		p.Match(focalParserT__23)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(174)
			p.Grpnum()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(175)
			p.Linenum()
		}

	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(focalParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *focalParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, focalParserRULE_comment)

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
		p.SetState(178)
		p.Match(focalParserCOMMENT)
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
	p.RuleIndex = focalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllPrimary() []IPrimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryContext)(nil)).Elem())
	var tst = make([]IPrimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Primary(i int) IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) AllPLUSMIN() []antlr.TerminalNode {
	return s.GetTokens(focalParserPLUSMIN)
}

func (s *ExpressionContext) PLUSMIN(i int) antlr.TerminalNode {
	return s.GetToken(focalParserPLUSMIN, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *focalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, focalParserRULE_expression)
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
		p.Primary()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == focalParserPLUSMIN {
		{
			p.SetState(181)
			p.Match(focalParserPLUSMIN)
		}
		{
			p.SetState(182)
			p.Primary()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *PrimaryContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PrimaryContext) AllMULOP() []antlr.TerminalNode {
	return s.GetTokens(focalParserMULOP)
}

func (s *PrimaryContext) MULOP(i int) antlr.TerminalNode {
	return s.GetToken(focalParserMULOP, i)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *focalParser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, focalParserRULE_primary)
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
		p.SetState(188)
		p.Term()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == focalParserMULOP {
		{
			p.SetState(189)
			p.Match(focalParserMULOP)
		}
		{
			p.SetState(190)
			p.Term()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = focalParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *TermContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(focalParserVARIABLE, 0)
}

func (s *TermContext) BUILTIN() antlr.TerminalNode {
	return s.GetToken(focalParserBUILTIN, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *focalParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, focalParserRULE_term)

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

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(focalParserT__24)
		}
		{
			p.SetState(197)
			p.Expression()
		}
		{
			p.SetState(198)
			p.Match(focalParserT__25)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Match(focalParserT__26)
		}
		{
			p.SetState(201)
			p.Expression()
		}
		{
			p.SetState(202)
			p.Match(focalParserT__27)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(204)
			p.Match(focalParserT__28)
		}
		{
			p.SetState(205)
			p.Expression()
		}
		{
			p.SetState(206)
			p.Match(focalParserT__29)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Number()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(209)
			p.Match(focalParserVARIABLE)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(210)
			p.Match(focalParserVARIABLE)
		}
		{
			p.SetState(211)
			p.Match(focalParserT__24)
		}
		{
			p.SetState(212)
			p.Expression()
		}
		{
			p.SetState(213)
			p.Match(focalParserT__25)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(215)
			p.Match(focalParserBUILTIN)
		}
		{
			p.SetState(216)
			p.Match(focalParserT__24)
		}
		{
			p.SetState(217)
			p.Expression()
		}
		{
			p.SetState(218)
			p.Match(focalParserT__25)
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Mantissa() IMantissaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMantissaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMantissaContext)
}

func (s *NumberContext) Signed_() ISigned_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigned_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISigned_Context)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *focalParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, focalParserRULE_number)
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
		p.Mantissa()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserT__30 {
		{
			p.SetState(223)
			p.Match(focalParserT__30)
		}
		{
			p.SetState(224)
			p.Signed_()
		}

	}

	return localctx
}

// IMantissaContext is an interface to support dynamic dispatch.
type IMantissaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMantissaContext differentiates from other interfaces.
	IsMantissaContext()
}

type MantissaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMantissaContext() *MantissaContext {
	var p = new(MantissaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_mantissa
	return p
}

func (*MantissaContext) IsMantissaContext() {}

func NewMantissaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MantissaContext {
	var p = new(MantissaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_mantissa

	return p
}

func (s *MantissaContext) GetParser() antlr.Parser { return s.parser }

func (s *MantissaContext) AllSigned_() []ISigned_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISigned_Context)(nil)).Elem())
	var tst = make([]ISigned_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISigned_Context)
		}
	}

	return tst
}

func (s *MantissaContext) Signed_(i int) ISigned_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigned_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISigned_Context)
}

func (s *MantissaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MantissaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MantissaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterMantissa(s)
	}
}

func (s *MantissaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitMantissa(s)
	}
}

func (p *focalParser) Mantissa() (localctx IMantissaContext) {
	this := p
	_ = this

	localctx = NewMantissaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, focalParserRULE_mantissa)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Signed_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Signed_()
		}
		{
			p.SetState(229)
			p.Match(focalParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(focalParserT__0)
		}
		{
			p.SetState(232)
			p.Signed_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.Signed_()
		}
		{
			p.SetState(234)
			p.Match(focalParserT__0)
		}
		{
			p.SetState(235)
			p.Signed_()
		}

	}

	return localctx
}

// ISigned_Context is an interface to support dynamic dispatch.
type ISigned_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSigned_Context differentiates from other interfaces.
	IsSigned_Context()
}

type Signed_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySigned_Context() *Signed_Context {
	var p = new(Signed_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = focalParserRULE_signed_
	return p
}

func (*Signed_Context) IsSigned_Context() {}

func NewSigned_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Signed_Context {
	var p = new(Signed_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = focalParserRULE_signed_

	return p
}

func (s *Signed_Context) GetParser() antlr.Parser { return s.parser }

func (s *Signed_Context) INTEGER() antlr.TerminalNode {
	return s.GetToken(focalParserINTEGER, 0)
}

func (s *Signed_Context) PLUSMIN() antlr.TerminalNode {
	return s.GetToken(focalParserPLUSMIN, 0)
}

func (s *Signed_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Signed_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Signed_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.EnterSigned_(s)
	}
}

func (s *Signed_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(focalListener); ok {
		listenerT.ExitSigned_(s)
	}
}

func (p *focalParser) Signed_() (localctx ISigned_Context) {
	this := p
	_ = this

	localctx = NewSigned_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, focalParserRULE_signed_)
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
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == focalParserPLUSMIN {
		{
			p.SetState(239)
			p.Match(focalParserPLUSMIN)
		}

	}
	{
		p.SetState(242)
		p.Match(focalParserINTEGER)
	}

	return localctx
}
