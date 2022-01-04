// Code generated from BSL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bsl // BSL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 246,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 28, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4,
	35, 10, 4, 13, 4, 14, 4, 36, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 56, 10,
	4, 13, 4, 14, 4, 57, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 70, 10, 4, 12, 4, 14, 4, 73, 11, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 78, 10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 83, 10, 5, 13, 5, 14, 5, 84, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 96, 10, 5, 13, 5,
	14, 5, 97, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	109, 10, 5, 12, 5, 14, 5, 112, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6,
	5, 131, 10, 5, 13, 5, 14, 5, 132, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6,
	5, 141, 10, 5, 13, 5, 14, 5, 142, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 153, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 6, 6, 178, 10, 6, 13, 6, 14, 6, 179, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 194, 10,
	6, 3, 6, 3, 6, 5, 6, 198, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 216, 10,
	7, 13, 7, 14, 7, 217, 3, 7, 5, 7, 221, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 235, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8,
	10, 12, 14, 16, 2, 3, 3, 2, 23, 24, 2, 275, 2, 19, 3, 2, 2, 2, 4, 27, 3,
	2, 2, 2, 6, 77, 3, 2, 2, 2, 8, 152, 3, 2, 2, 2, 10, 197, 3, 2, 2, 2, 12,
	234, 3, 2, 2, 2, 14, 236, 3, 2, 2, 2, 16, 243, 3, 2, 2, 2, 18, 20, 5, 4,
	3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22,
	3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 28, 5, 6, 4, 2, 24, 28, 5, 8, 5, 2,
	25, 28, 5, 10, 6, 2, 26, 28, 5, 12, 7, 2, 27, 23, 3, 2, 2, 2, 27, 24, 3,
	2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29,
	30, 7, 3, 2, 2, 30, 31, 7, 4, 2, 2, 31, 32, 7, 3, 2, 2, 32, 34, 5, 16,
	9, 2, 33, 35, 7, 24, 2, 2, 34, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36,
	34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 7, 5, 2,
	2, 39, 40, 5, 8, 5, 2, 40, 41, 7, 5, 2, 2, 41, 78, 3, 2, 2, 2, 42, 43,
	7, 3, 2, 2, 43, 44, 7, 4, 2, 2, 44, 45, 5, 16, 9, 2, 45, 46, 5, 8, 5, 2,
	46, 47, 7, 5, 2, 2, 47, 78, 3, 2, 2, 2, 48, 49, 7, 3, 2, 2, 49, 50, 7,
	4, 2, 2, 50, 51, 5, 16, 9, 2, 51, 52, 7, 3, 2, 2, 52, 53, 7, 6, 2, 2, 53,
	55, 7, 3, 2, 2, 54, 56, 7, 24, 2, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2,
	2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60,
	7, 5, 2, 2, 60, 61, 5, 8, 5, 2, 61, 62, 7, 5, 2, 2, 62, 63, 7, 5, 2, 2,
	63, 78, 3, 2, 2, 2, 64, 65, 7, 3, 2, 2, 65, 66, 7, 7, 2, 2, 66, 67, 5,
	16, 9, 2, 67, 71, 7, 3, 2, 2, 68, 70, 5, 16, 9, 2, 69, 68, 3, 2, 2, 2,
	70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3,
	2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 7, 5, 2, 2, 75, 76, 7, 5, 2, 2, 76,
	78, 3, 2, 2, 2, 77, 29, 3, 2, 2, 2, 77, 42, 3, 2, 2, 2, 77, 48, 3, 2, 2,
	2, 77, 64, 3, 2, 2, 2, 78, 7, 3, 2, 2, 2, 79, 80, 7, 3, 2, 2, 80, 82, 5,
	16, 9, 2, 81, 83, 5, 8, 5, 2, 82, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84,
	82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 7, 5, 2,
	2, 87, 153, 3, 2, 2, 2, 88, 89, 7, 3, 2, 2, 89, 95, 7, 8, 2, 2, 90, 91,
	7, 9, 2, 2, 91, 92, 5, 8, 5, 2, 92, 93, 5, 8, 5, 2, 93, 94, 7, 10, 2, 2,
	94, 96, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 95, 3,
	2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 7, 5, 2, 2, 100,
	153, 3, 2, 2, 2, 101, 102, 7, 3, 2, 2, 102, 110, 7, 8, 2, 2, 103, 104,
	7, 9, 2, 2, 104, 105, 5, 8, 5, 2, 105, 106, 5, 8, 5, 2, 106, 107, 7, 10,
	2, 2, 107, 109, 3, 2, 2, 2, 108, 103, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2,
	110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 113, 114, 7, 9, 2, 2, 114, 115, 7, 11, 2, 2, 115, 116,
	5, 8, 5, 2, 116, 117, 7, 10, 2, 2, 117, 118, 7, 5, 2, 2, 118, 153, 3, 2,
	2, 2, 119, 120, 7, 3, 2, 2, 120, 121, 7, 12, 2, 2, 121, 122, 5, 8, 5, 2,
	122, 123, 5, 8, 5, 2, 123, 124, 5, 8, 5, 2, 124, 125, 7, 5, 2, 2, 125,
	153, 3, 2, 2, 2, 126, 127, 7, 3, 2, 2, 127, 128, 7, 13, 2, 2, 128, 130,
	5, 8, 5, 2, 129, 131, 5, 8, 5, 2, 130, 129, 3, 2, 2, 2, 131, 132, 3, 2,
	2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2,
	134, 135, 7, 5, 2, 2, 135, 153, 3, 2, 2, 2, 136, 137, 7, 3, 2, 2, 137,
	138, 7, 14, 2, 2, 138, 140, 5, 8, 5, 2, 139, 141, 5, 8, 5, 2, 140, 139,
	3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2,
	2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 5, 2, 2, 145, 153, 3, 2, 2, 2,
	146, 153, 7, 15, 2, 2, 147, 153, 5, 16, 9, 2, 148, 153, 7, 25, 2, 2, 149,
	153, 7, 27, 2, 2, 150, 153, 7, 28, 2, 2, 151, 153, 7, 29, 2, 2, 152, 79,
	3, 2, 2, 2, 152, 88, 3, 2, 2, 2, 152, 101, 3, 2, 2, 2, 152, 119, 3, 2,
	2, 2, 152, 126, 3, 2, 2, 2, 152, 136, 3, 2, 2, 2, 152, 146, 3, 2, 2, 2,
	152, 147, 3, 2, 2, 2, 152, 148, 3, 2, 2, 2, 152, 149, 3, 2, 2, 2, 152,
	150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 9, 3, 2, 2, 2, 154, 155, 7,
	3, 2, 2, 155, 156, 7, 16, 2, 2, 156, 157, 5, 8, 5, 2, 157, 158, 5, 8, 5,
	2, 158, 159, 7, 5, 2, 2, 159, 198, 3, 2, 2, 2, 160, 161, 7, 3, 2, 2, 161,
	162, 7, 17, 2, 2, 162, 163, 5, 8, 5, 2, 163, 164, 5, 8, 5, 2, 164, 165,
	7, 5, 2, 2, 165, 198, 3, 2, 2, 2, 166, 167, 7, 3, 2, 2, 167, 168, 7, 18,
	2, 2, 168, 169, 5, 8, 5, 2, 169, 170, 5, 8, 5, 2, 170, 171, 5, 8, 5, 2,
	171, 172, 7, 5, 2, 2, 172, 198, 3, 2, 2, 2, 173, 174, 7, 3, 2, 2, 174,
	175, 7, 19, 2, 2, 175, 177, 5, 8, 5, 2, 176, 178, 5, 8, 5, 2, 177, 176,
	3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2,
	2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 5, 2, 2, 182, 198, 3, 2, 2, 2,
	183, 184, 7, 3, 2, 2, 184, 185, 7, 20, 2, 2, 185, 186, 5, 8, 5, 2, 186,
	187, 5, 16, 9, 2, 187, 188, 7, 5, 2, 2, 188, 198, 3, 2, 2, 2, 189, 190,
	7, 3, 2, 2, 190, 191, 7, 21, 2, 2, 191, 193, 5, 8, 5, 2, 192, 194, 5, 8,
	5, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2,
	195, 196, 7, 5, 2, 2, 196, 198, 3, 2, 2, 2, 197, 154, 3, 2, 2, 2, 197,
	160, 3, 2, 2, 2, 197, 166, 3, 2, 2, 2, 197, 173, 3, 2, 2, 2, 197, 183,
	3, 2, 2, 2, 197, 189, 3, 2, 2, 2, 198, 11, 3, 2, 2, 2, 199, 200, 7, 3,
	2, 2, 200, 201, 7, 22, 2, 2, 201, 202, 7, 28, 2, 2, 202, 235, 7, 5, 2,
	2, 203, 204, 7, 3, 2, 2, 204, 205, 7, 22, 2, 2, 205, 206, 5, 16, 9, 2,
	206, 207, 7, 5, 2, 2, 207, 235, 3, 2, 2, 2, 208, 209, 7, 3, 2, 2, 209,
	210, 7, 22, 2, 2, 210, 211, 7, 3, 2, 2, 211, 212, 5, 16, 9, 2, 212, 220,
	7, 28, 2, 2, 213, 215, 7, 3, 2, 2, 214, 216, 7, 28, 2, 2, 215, 214, 3,
	2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2,
	2, 218, 219, 3, 2, 2, 2, 219, 221, 7, 5, 2, 2, 220, 213, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 7, 5, 2, 2, 223, 224,
	7, 5, 2, 2, 224, 235, 3, 2, 2, 2, 225, 226, 7, 3, 2, 2, 226, 227, 7, 22,
	2, 2, 227, 228, 7, 3, 2, 2, 228, 229, 5, 16, 9, 2, 229, 230, 7, 28, 2,
	2, 230, 231, 5, 14, 8, 2, 231, 232, 7, 5, 2, 2, 232, 233, 7, 5, 2, 2, 233,
	235, 3, 2, 2, 2, 234, 199, 3, 2, 2, 2, 234, 203, 3, 2, 2, 2, 234, 208,
	3, 2, 2, 2, 234, 225, 3, 2, 2, 2, 235, 13, 3, 2, 2, 2, 236, 237, 7, 3,
	2, 2, 237, 238, 7, 28, 2, 2, 238, 239, 7, 28, 2, 2, 239, 240, 7, 25, 2,
	2, 240, 241, 7, 25, 2, 2, 241, 242, 7, 5, 2, 2, 242, 15, 3, 2, 2, 2, 243,
	244, 9, 2, 2, 2, 244, 17, 3, 2, 2, 2, 20, 21, 27, 36, 57, 71, 77, 84, 97,
	110, 132, 142, 152, 179, 193, 197, 217, 220, 234,
}
var literalNames = []string{
	"", "'('", "'define'", "')'", "'lambda'", "'define-struct'", "'cond'",
	"'['", "']'", "'else '", "'if'", "'and'", "'or'", "'\u2019()'", "'check-expect'",
	"'check-random'", "'check-within'", "'check-member-of'", "'check-satisfied'",
	"'check-error'", "'require'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "SYMBOL", "NAME", "NUMBER", "INT", "BOOLEAN", "STRING", "CHARACTER",
	"LANG", "COMMENT", "WS",
}

var ruleNames = []string{
	"program", "defOrExpr", "definition", "expr", "testCase", "libraryRequire",
	"pkg", "name",
}

type BSLParser struct {
	*antlr.BaseParser
}

// NewBSLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BSLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBSLParser(input antlr.TokenStream) *BSLParser {
	this := new(BSLParser)
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
	this.GrammarFileName = "BSL.g4"

	return this
}

// BSLParser tokens.
const (
	BSLParserEOF       = antlr.TokenEOF
	BSLParserT__0      = 1
	BSLParserT__1      = 2
	BSLParserT__2      = 3
	BSLParserT__3      = 4
	BSLParserT__4      = 5
	BSLParserT__5      = 6
	BSLParserT__6      = 7
	BSLParserT__7      = 8
	BSLParserT__8      = 9
	BSLParserT__9      = 10
	BSLParserT__10     = 11
	BSLParserT__11     = 12
	BSLParserT__12     = 13
	BSLParserT__13     = 14
	BSLParserT__14     = 15
	BSLParserT__15     = 16
	BSLParserT__16     = 17
	BSLParserT__17     = 18
	BSLParserT__18     = 19
	BSLParserT__19     = 20
	BSLParserSYMBOL    = 21
	BSLParserNAME      = 22
	BSLParserNUMBER    = 23
	BSLParserINT       = 24
	BSLParserBOOLEAN   = 25
	BSLParserSTRING    = 26
	BSLParserCHARACTER = 27
	BSLParserLANG      = 28
	BSLParserCOMMENT   = 29
	BSLParserWS        = 30
)

// BSLParser rules.
const (
	BSLParserRULE_program        = 0
	BSLParserRULE_defOrExpr      = 1
	BSLParserRULE_definition     = 2
	BSLParserRULE_expr           = 3
	BSLParserRULE_testCase       = 4
	BSLParserRULE_libraryRequire = 5
	BSLParserRULE_pkg            = 6
	BSLParserRULE_name           = 7
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
	p.RuleIndex = BSLParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDefOrExpr() []IDefOrExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefOrExprContext)(nil)).Elem())
	var tst = make([]IDefOrExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefOrExprContext)
		}
	}

	return tst
}

func (s *ProgramContext) DefOrExpr(i int) IDefOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefOrExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefOrExprContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BSLParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BSLParserRULE_program)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0) {
		{
			p.SetState(16)
			p.DefOrExpr()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefOrExprContext is an interface to support dynamic dispatch.
type IDefOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefOrExprContext differentiates from other interfaces.
	IsDefOrExprContext()
}

type DefOrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefOrExprContext() *DefOrExprContext {
	var p = new(DefOrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_defOrExpr
	return p
}

func (*DefOrExprContext) IsDefOrExprContext() {}

func NewDefOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefOrExprContext {
	var p = new(DefOrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_defOrExpr

	return p
}

func (s *DefOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DefOrExprContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefOrExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefOrExprContext) TestCase() ITestCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestCaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestCaseContext)
}

func (s *DefOrExprContext) LibraryRequire() ILibraryRequireContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryRequireContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryRequireContext)
}

func (s *DefOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterDefOrExpr(s)
	}
}

func (s *DefOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitDefOrExpr(s)
	}
}

func (p *BSLParser) DefOrExpr() (localctx IDefOrExprContext) {
	this := p
	_ = this

	localctx = NewDefOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BSLParserRULE_defOrExpr)

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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(23)
			p.TestCase()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(24)
			p.LibraryRequire()
		}

	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *DefinitionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DefinitionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefinitionContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(BSLParserNAME)
}

func (s *DefinitionContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(BSLParserNAME, i)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *BSLParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BSLParserRULE_definition)
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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(28)
			p.Match(BSLParserT__1)
		}
		{
			p.SetState(29)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(30)
			p.Name()
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BSLParserNAME {
			{
				p.SetState(31)
				p.Match(BSLParserNAME)
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(36)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(37)
			p.Expr()
		}
		{
			p.SetState(38)
			p.Match(BSLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(41)
			p.Match(BSLParserT__1)
		}
		{
			p.SetState(42)
			p.Name()
		}
		{
			p.SetState(43)
			p.Expr()
		}
		{
			p.SetState(44)
			p.Match(BSLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(47)
			p.Match(BSLParserT__1)
		}
		{
			p.SetState(48)
			p.Name()
		}
		{
			p.SetState(49)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(50)
			p.Match(BSLParserT__3)
		}
		{
			p.SetState(51)
			p.Match(BSLParserT__0)
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BSLParserNAME {
			{
				p.SetState(52)
				p.Match(BSLParserNAME)
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(57)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(58)
			p.Expr()
		}
		{
			p.SetState(59)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(60)
			p.Match(BSLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(63)
			p.Match(BSLParserT__4)
		}
		{
			p.SetState(64)
			p.Name()
		}
		{
			p.SetState(65)
			p.Match(BSLParserT__0)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BSLParserSYMBOL || _la == BSLParserNAME {
			{
				p.SetState(66)
				p.Name()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(72)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(73)
			p.Match(BSLParserT__2)
		}

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
	p.RuleIndex = BSLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

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

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BSLParserNUMBER, 0)
}

func (s *ExprContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(BSLParserBOOLEAN, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(BSLParserSTRING, 0)
}

func (s *ExprContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(BSLParserCHARACTER, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *BSLParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BSLParserRULE_expr)
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

	var _alt int

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(78)
			p.Name()
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0) {
			{
				p.SetState(79)
				p.Expr()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(84)
			p.Match(BSLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(87)
			p.Match(BSLParserT__5)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BSLParserT__6 {
			{
				p.SetState(88)
				p.Match(BSLParserT__6)
			}
			{
				p.SetState(89)
				p.Expr()
			}
			{
				p.SetState(90)
				p.Expr()
			}
			{
				p.SetState(91)
				p.Match(BSLParserT__7)
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(BSLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(100)
			p.Match(BSLParserT__5)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(101)
					p.Match(BSLParserT__6)
				}
				{
					p.SetState(102)
					p.Expr()
				}
				{
					p.SetState(103)
					p.Expr()
				}
				{
					p.SetState(104)
					p.Match(BSLParserT__7)
				}

			}
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}
		{
			p.SetState(111)
			p.Match(BSLParserT__6)
		}
		{
			p.SetState(112)
			p.Match(BSLParserT__8)
		}
		{
			p.SetState(113)
			p.Expr()
		}
		{
			p.SetState(114)
			p.Match(BSLParserT__7)
		}
		{
			p.SetState(115)
			p.Match(BSLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(118)
			p.Match(BSLParserT__9)
		}
		{
			p.SetState(119)
			p.Expr()
		}
		{
			p.SetState(120)
			p.Expr()
		}
		{
			p.SetState(121)
			p.Expr()
		}
		{
			p.SetState(122)
			p.Match(BSLParserT__2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(125)
			p.Match(BSLParserT__10)
		}
		{
			p.SetState(126)
			p.Expr()
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0) {
			{
				p.SetState(127)
				p.Expr()
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(132)
			p.Match(BSLParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(134)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(135)
			p.Match(BSLParserT__11)
		}
		{
			p.SetState(136)
			p.Expr()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0) {
			{
				p.SetState(137)
				p.Expr()
			}

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(142)
			p.Match(BSLParserT__2)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.Match(BSLParserT__12)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(145)
			p.Name()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)
			p.Match(BSLParserNUMBER)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(147)
			p.Match(BSLParserBOOLEAN)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(148)
			p.Match(BSLParserSTRING)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(149)
			p.Match(BSLParserCHARACTER)
		}

	}

	return localctx
}

// ITestCaseContext is an interface to support dynamic dispatch.
type ITestCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestCaseContext differentiates from other interfaces.
	IsTestCaseContext()
}

type TestCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestCaseContext() *TestCaseContext {
	var p = new(TestCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_testCase
	return p
}

func (*TestCaseContext) IsTestCaseContext() {}

func NewTestCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestCaseContext {
	var p = new(TestCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_testCase

	return p
}

func (s *TestCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *TestCaseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TestCaseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TestCaseContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TestCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterTestCase(s)
	}
}

func (s *TestCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitTestCase(s)
	}
}

func (p *BSLParser) TestCase() (localctx ITestCaseContext) {
	this := p
	_ = this

	localctx = NewTestCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BSLParserRULE_testCase)
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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(153)
			p.Match(BSLParserT__13)
		}
		{
			p.SetState(154)
			p.Expr()
		}
		{
			p.SetState(155)
			p.Expr()
		}
		{
			p.SetState(156)
			p.Match(BSLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(159)
			p.Match(BSLParserT__14)
		}
		{
			p.SetState(160)
			p.Expr()
		}
		{
			p.SetState(161)
			p.Expr()
		}
		{
			p.SetState(162)
			p.Match(BSLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(165)
			p.Match(BSLParserT__15)
		}
		{
			p.SetState(166)
			p.Expr()
		}
		{
			p.SetState(167)
			p.Expr()
		}
		{
			p.SetState(168)
			p.Expr()
		}
		{
			p.SetState(169)
			p.Match(BSLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(172)
			p.Match(BSLParserT__16)
		}
		{
			p.SetState(173)
			p.Expr()
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0) {
			{
				p.SetState(174)
				p.Expr()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(179)
			p.Match(BSLParserT__2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(181)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(182)
			p.Match(BSLParserT__17)
		}
		{
			p.SetState(183)
			p.Expr()
		}
		{
			p.SetState(184)
			p.Name()
		}
		{
			p.SetState(185)
			p.Match(BSLParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(187)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(188)
			p.Match(BSLParserT__18)
		}
		{
			p.SetState(189)
			p.Expr()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BSLParserT__0)|(1<<BSLParserT__12)|(1<<BSLParserSYMBOL)|(1<<BSLParserNAME)|(1<<BSLParserNUMBER)|(1<<BSLParserBOOLEAN)|(1<<BSLParserSTRING)|(1<<BSLParserCHARACTER))) != 0 {
			{
				p.SetState(190)
				p.Expr()
			}

		}
		{
			p.SetState(193)
			p.Match(BSLParserT__2)
		}

	}

	return localctx
}

// ILibraryRequireContext is an interface to support dynamic dispatch.
type ILibraryRequireContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryRequireContext differentiates from other interfaces.
	IsLibraryRequireContext()
}

type LibraryRequireContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryRequireContext() *LibraryRequireContext {
	var p = new(LibraryRequireContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_libraryRequire
	return p
}

func (*LibraryRequireContext) IsLibraryRequireContext() {}

func NewLibraryRequireContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryRequireContext {
	var p = new(LibraryRequireContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_libraryRequire

	return p
}

func (s *LibraryRequireContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryRequireContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BSLParserSTRING)
}

func (s *LibraryRequireContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BSLParserSTRING, i)
}

func (s *LibraryRequireContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LibraryRequireContext) Pkg() IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

func (s *LibraryRequireContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryRequireContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryRequireContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterLibraryRequire(s)
	}
}

func (s *LibraryRequireContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitLibraryRequire(s)
	}
}

func (p *BSLParser) LibraryRequire() (localctx ILibraryRequireContext) {
	this := p
	_ = this

	localctx = NewLibraryRequireContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BSLParserRULE_libraryRequire)
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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(198)
			p.Match(BSLParserT__19)
		}
		{
			p.SetState(199)
			p.Match(BSLParserSTRING)
		}
		{
			p.SetState(200)
			p.Match(BSLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(202)
			p.Match(BSLParserT__19)
		}
		{
			p.SetState(203)
			p.Name()
		}
		{
			p.SetState(204)
			p.Match(BSLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(207)
			p.Match(BSLParserT__19)
		}
		{
			p.SetState(208)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(209)
			p.Name()
		}
		{
			p.SetState(210)
			p.Match(BSLParserSTRING)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BSLParserT__0 {
			{
				p.SetState(211)
				p.Match(BSLParserT__0)
			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == BSLParserSTRING {
				{
					p.SetState(212)
					p.Match(BSLParserSTRING)
				}

				p.SetState(215)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(217)
				p.Match(BSLParserT__2)
			}

		}
		{
			p.SetState(220)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(221)
			p.Match(BSLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(223)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(224)
			p.Match(BSLParserT__19)
		}
		{
			p.SetState(225)
			p.Match(BSLParserT__0)
		}
		{
			p.SetState(226)
			p.Name()
		}
		{
			p.SetState(227)
			p.Match(BSLParserSTRING)
		}
		{
			p.SetState(228)
			p.Pkg()
		}
		{
			p.SetState(229)
			p.Match(BSLParserT__2)
		}
		{
			p.SetState(230)
			p.Match(BSLParserT__2)
		}

	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BSLParserSTRING)
}

func (s *PkgContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BSLParserSTRING, i)
}

func (s *PkgContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(BSLParserNUMBER)
}

func (s *PkgContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(BSLParserNUMBER, i)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *BSLParser) Pkg() (localctx IPkgContext) {
	this := p
	_ = this

	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BSLParserRULE_pkg)

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
		p.SetState(234)
		p.Match(BSLParserT__0)
	}
	{
		p.SetState(235)
		p.Match(BSLParserSTRING)
	}
	{
		p.SetState(236)
		p.Match(BSLParserSTRING)
	}
	{
		p.SetState(237)
		p.Match(BSLParserNUMBER)
	}
	{
		p.SetState(238)
		p.Match(BSLParserNUMBER)
	}
	{
		p.SetState(239)
		p.Match(BSLParserT__2)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BSLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BSLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(BSLParserSYMBOL, 0)
}

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(BSLParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BSLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *BSLParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BSLParserRULE_name)
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
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BSLParserSYMBOL || _la == BSLParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
