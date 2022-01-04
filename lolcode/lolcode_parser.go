// Code generated from lolcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lolcode // lolcode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 276,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 5, 2, 64, 10, 2, 3, 3, 6, 3,
	67, 10, 3, 13, 3, 14, 3, 68, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 80, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 95, 10, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 102, 10, 7, 3, 8, 3, 8, 7, 8, 106, 10, 8, 12, 8, 14,
	8, 109, 11, 8, 3, 8, 5, 8, 112, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 125, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 138, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 150, 10, 12, 12, 12, 14, 12, 153, 11, 12, 5, 12, 155, 10, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 182, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 247, 10, 27, 12, 27,
	14, 27, 250, 11, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 258,
	10, 28, 12, 28, 14, 28, 261, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 6, 30, 270, 10, 30, 13, 30, 14, 30, 271, 3, 30, 3, 30, 3,
	30, 2, 2, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 2, 2, 2, 285, 2, 60,
	3, 2, 2, 2, 4, 66, 3, 2, 2, 2, 6, 79, 3, 2, 2, 2, 8, 81, 3, 2, 2, 2, 10,
	94, 3, 2, 2, 2, 12, 101, 3, 2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 124, 3, 2,
	2, 2, 18, 137, 3, 2, 2, 2, 20, 139, 3, 2, 2, 2, 22, 142, 3, 2, 2, 2, 24,
	159, 3, 2, 2, 2, 26, 181, 3, 2, 2, 2, 28, 183, 3, 2, 2, 2, 30, 188, 3,
	2, 2, 2, 32, 193, 3, 2, 2, 2, 34, 198, 3, 2, 2, 2, 36, 203, 3, 2, 2, 2,
	38, 208, 3, 2, 2, 2, 40, 213, 3, 2, 2, 2, 42, 218, 3, 2, 2, 2, 44, 223,
	3, 2, 2, 2, 46, 228, 3, 2, 2, 2, 48, 233, 3, 2, 2, 2, 50, 238, 3, 2, 2,
	2, 52, 242, 3, 2, 2, 2, 54, 253, 3, 2, 2, 2, 56, 264, 3, 2, 2, 2, 58, 267,
	3, 2, 2, 2, 60, 61, 7, 3, 2, 2, 61, 63, 5, 4, 3, 2, 62, 64, 7, 4, 2, 2,
	63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 3, 3, 2, 2, 2, 65, 67, 5, 6,
	4, 2, 66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69,
	3, 2, 2, 2, 69, 5, 3, 2, 2, 2, 70, 80, 5, 8, 5, 2, 71, 80, 5, 10, 6, 2,
	72, 80, 5, 12, 7, 2, 73, 80, 5, 14, 8, 2, 74, 80, 5, 16, 9, 2, 75, 80,
	5, 20, 11, 2, 76, 80, 5, 22, 12, 2, 77, 80, 5, 24, 13, 2, 78, 80, 5, 26,
	14, 2, 79, 70, 3, 2, 2, 2, 79, 71, 3, 2, 2, 2, 79, 72, 3, 2, 2, 2, 79,
	73, 3, 2, 2, 2, 79, 74, 3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 79, 76, 3, 2, 2,
	2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 7, 3, 2, 2, 2, 81, 82, 7,
	5, 2, 2, 82, 83, 7, 43, 2, 2, 83, 84, 7, 6, 2, 2, 84, 85, 5, 26, 14, 2,
	85, 86, 5, 4, 3, 2, 86, 87, 7, 7, 2, 2, 87, 88, 7, 43, 2, 2, 88, 9, 3,
	2, 2, 2, 89, 90, 7, 8, 2, 2, 90, 95, 7, 43, 2, 2, 91, 92, 7, 8, 2, 2, 92,
	93, 7, 43, 2, 2, 93, 95, 7, 9, 2, 2, 94, 89, 3, 2, 2, 2, 94, 91, 3, 2,
	2, 2, 95, 11, 3, 2, 2, 2, 96, 97, 7, 10, 2, 2, 97, 102, 7, 45, 2, 2, 98,
	99, 7, 11, 2, 2, 99, 100, 7, 45, 2, 2, 100, 102, 7, 12, 2, 2, 101, 96,
	3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 102, 13, 3, 2, 2, 2, 103, 107, 7, 13,
	2, 2, 104, 106, 5, 26, 14, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2,
	2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109,
	107, 3, 2, 2, 2, 110, 112, 7, 14, 2, 2, 111, 110, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 15, 3, 2, 2, 2, 113, 114, 7, 15, 2, 2, 114, 115, 7, 16,
	2, 2, 115, 116, 5, 4, 3, 2, 116, 117, 7, 17, 2, 2, 117, 125, 3, 2, 2, 2,
	118, 119, 7, 15, 2, 2, 119, 120, 7, 16, 2, 2, 120, 121, 5, 4, 3, 2, 121,
	122, 5, 18, 10, 2, 122, 123, 7, 17, 2, 2, 123, 125, 3, 2, 2, 2, 124, 113,
	3, 2, 2, 2, 124, 118, 3, 2, 2, 2, 125, 17, 3, 2, 2, 2, 126, 127, 7, 18,
	2, 2, 127, 128, 5, 26, 14, 2, 128, 129, 5, 4, 3, 2, 129, 130, 5, 18, 10,
	2, 130, 138, 3, 2, 2, 2, 131, 132, 7, 19, 2, 2, 132, 138, 5, 4, 3, 2, 133,
	134, 7, 18, 2, 2, 134, 135, 5, 26, 14, 2, 135, 136, 5, 4, 3, 2, 136, 138,
	3, 2, 2, 2, 137, 126, 3, 2, 2, 2, 137, 131, 3, 2, 2, 2, 137, 133, 3, 2,
	2, 2, 138, 19, 3, 2, 2, 2, 139, 140, 7, 20, 2, 2, 140, 141, 7, 43, 2, 2,
	141, 21, 3, 2, 2, 2, 142, 143, 7, 21, 2, 2, 143, 154, 7, 43, 2, 2, 144,
	145, 7, 22, 2, 2, 145, 146, 7, 43, 2, 2, 146, 151, 3, 2, 2, 2, 147, 148,
	7, 23, 2, 2, 148, 150, 7, 43, 2, 2, 149, 147, 3, 2, 2, 2, 150, 153, 3,
	2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 155, 3, 2, 2,
	2, 153, 151, 3, 2, 2, 2, 154, 144, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 157, 5, 4, 3, 2, 157, 158, 7, 24, 2, 2, 158, 23,
	3, 2, 2, 2, 159, 160, 7, 43, 2, 2, 160, 161, 7, 25, 2, 2, 161, 162, 5,
	26, 14, 2, 162, 25, 3, 2, 2, 2, 163, 182, 5, 28, 15, 2, 164, 182, 5, 32,
	17, 2, 165, 182, 5, 30, 16, 2, 166, 182, 5, 36, 19, 2, 167, 182, 5, 38,
	20, 2, 168, 182, 5, 40, 21, 2, 169, 182, 5, 42, 22, 2, 170, 182, 5, 44,
	23, 2, 171, 182, 5, 46, 24, 2, 172, 182, 5, 48, 25, 2, 173, 182, 5, 50,
	26, 2, 174, 182, 5, 34, 18, 2, 175, 182, 5, 52, 27, 2, 176, 182, 5, 54,
	28, 2, 177, 182, 5, 56, 29, 2, 178, 182, 5, 58, 30, 2, 179, 182, 7, 43,
	2, 2, 180, 182, 7, 44, 2, 2, 181, 163, 3, 2, 2, 2, 181, 164, 3, 2, 2, 2,
	181, 165, 3, 2, 2, 2, 181, 166, 3, 2, 2, 2, 181, 167, 3, 2, 2, 2, 181,
	168, 3, 2, 2, 2, 181, 169, 3, 2, 2, 2, 181, 170, 3, 2, 2, 2, 181, 171,
	3, 2, 2, 2, 181, 172, 3, 2, 2, 2, 181, 173, 3, 2, 2, 2, 181, 174, 3, 2,
	2, 2, 181, 175, 3, 2, 2, 2, 181, 176, 3, 2, 2, 2, 181, 177, 3, 2, 2, 2,
	181, 178, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182,
	27, 3, 2, 2, 2, 183, 184, 7, 26, 2, 2, 184, 185, 5, 26, 14, 2, 185, 186,
	7, 27, 2, 2, 186, 187, 5, 26, 14, 2, 187, 29, 3, 2, 2, 2, 188, 189, 7,
	28, 2, 2, 189, 190, 5, 26, 14, 2, 190, 191, 7, 27, 2, 2, 191, 192, 5, 26,
	14, 2, 192, 31, 3, 2, 2, 2, 193, 194, 7, 29, 2, 2, 194, 195, 5, 26, 14,
	2, 195, 196, 7, 27, 2, 2, 196, 197, 5, 26, 14, 2, 197, 33, 3, 2, 2, 2,
	198, 199, 7, 30, 2, 2, 199, 200, 5, 26, 14, 2, 200, 201, 7, 27, 2, 2, 201,
	202, 5, 26, 14, 2, 202, 35, 3, 2, 2, 2, 203, 204, 7, 31, 2, 2, 204, 205,
	5, 26, 14, 2, 205, 206, 7, 27, 2, 2, 206, 207, 5, 26, 14, 2, 207, 37, 3,
	2, 2, 2, 208, 209, 7, 32, 2, 2, 209, 210, 5, 26, 14, 2, 210, 211, 7, 27,
	2, 2, 211, 212, 5, 26, 14, 2, 212, 39, 3, 2, 2, 2, 213, 214, 7, 33, 2,
	2, 214, 215, 5, 26, 14, 2, 215, 216, 7, 27, 2, 2, 216, 217, 5, 26, 14,
	2, 217, 41, 3, 2, 2, 2, 218, 219, 7, 34, 2, 2, 219, 220, 5, 26, 14, 2,
	220, 221, 7, 27, 2, 2, 221, 222, 5, 26, 14, 2, 222, 43, 3, 2, 2, 2, 223,
	224, 7, 35, 2, 2, 224, 225, 5, 26, 14, 2, 225, 226, 7, 27, 2, 2, 226, 227,
	5, 26, 14, 2, 227, 45, 3, 2, 2, 2, 228, 229, 7, 36, 2, 2, 229, 230, 5,
	26, 14, 2, 230, 231, 7, 27, 2, 2, 231, 232, 5, 26, 14, 2, 232, 47, 3, 2,
	2, 2, 233, 234, 7, 37, 2, 2, 234, 235, 5, 26, 14, 2, 235, 236, 7, 27, 2,
	2, 236, 237, 5, 26, 14, 2, 237, 49, 3, 2, 2, 2, 238, 239, 7, 38, 2, 2,
	239, 240, 5, 26, 14, 2, 240, 241, 7, 39, 2, 2, 241, 51, 3, 2, 2, 2, 242,
	243, 7, 40, 2, 2, 243, 248, 5, 26, 14, 2, 244, 245, 7, 27, 2, 2, 245, 247,
	5, 26, 14, 2, 246, 244, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3,
	2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 251, 3, 2, 2, 2, 250, 248, 3, 2, 2,
	2, 251, 252, 7, 14, 2, 2, 252, 53, 3, 2, 2, 2, 253, 254, 7, 41, 2, 2, 254,
	259, 5, 26, 14, 2, 255, 256, 7, 27, 2, 2, 256, 258, 5, 26, 14, 2, 257,
	255, 3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 262, 263, 7, 14,
	2, 2, 263, 55, 3, 2, 2, 2, 264, 265, 7, 42, 2, 2, 265, 266, 5, 26, 14,
	2, 266, 57, 3, 2, 2, 2, 267, 269, 7, 43, 2, 2, 268, 270, 5, 26, 14, 2,
	269, 268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271,
	272, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 7, 14, 2, 2, 274, 59,
	3, 2, 2, 2, 17, 63, 68, 79, 94, 101, 107, 111, 124, 137, 151, 154, 181,
	248, 259, 271,
}
var literalNames = []string{
	"", "'HAI'", "'KTHXBYE'", "'IM IN YR'", "'WILE'", "'IM OUTTA YR'", "'I HAS A'",
	"'ITZ'", "'BTW'", "'OBTW'", "'TLDR'", "'VISIBLE'", "'MKAY?'", "'O RLY?'",
	"'YA RLY'", "'OIC'", "'MEBBE'", "'NO WAI'", "'GIMMEH'", "'HOW DUZ I'",
	"'YR'", "'AN YR'", "'IF U SAY SO'", "'R'", "'BOTH SAEM'", "'AN'", "'DIFFRINT'",
	"'BOTH OF'", "'EITHER OF'", "'BIGGR OF'", "'SMALLR OF'", "'SUM OF'", "'DIFF OF'",
	"'PRODUKT OF'", "'QUOSHUNT OF'", "'MOD OF'", "'MAEK'", "'A'", "'ALL OF'",
	"'ANY OF'", "'NOT'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "LABEL", "ATOM", "STRING", "WS",
}

var ruleNames = []string{
	"program", "code_block", "statement", "loop", "declaration", "comment",
	"print_block", "if_block", "else_if_block", "input_block", "func_decl",
	"assignment", "expression", "equals", "not_equals", "both", "either", "greater",
	"less", "add", "sub", "mul", "div", "mod", "cast", "all_", "any_", "not_",
	"func_",
}

type lolcodeParser struct {
	*antlr.BaseParser
}

// NewlolcodeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *lolcodeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlolcodeParser(input antlr.TokenStream) *lolcodeParser {
	this := new(lolcodeParser)
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
	this.GrammarFileName = "lolcode.g4"

	return this
}

// lolcodeParser tokens.
const (
	lolcodeParserEOF    = antlr.TokenEOF
	lolcodeParserT__0   = 1
	lolcodeParserT__1   = 2
	lolcodeParserT__2   = 3
	lolcodeParserT__3   = 4
	lolcodeParserT__4   = 5
	lolcodeParserT__5   = 6
	lolcodeParserT__6   = 7
	lolcodeParserT__7   = 8
	lolcodeParserT__8   = 9
	lolcodeParserT__9   = 10
	lolcodeParserT__10  = 11
	lolcodeParserT__11  = 12
	lolcodeParserT__12  = 13
	lolcodeParserT__13  = 14
	lolcodeParserT__14  = 15
	lolcodeParserT__15  = 16
	lolcodeParserT__16  = 17
	lolcodeParserT__17  = 18
	lolcodeParserT__18  = 19
	lolcodeParserT__19  = 20
	lolcodeParserT__20  = 21
	lolcodeParserT__21  = 22
	lolcodeParserT__22  = 23
	lolcodeParserT__23  = 24
	lolcodeParserT__24  = 25
	lolcodeParserT__25  = 26
	lolcodeParserT__26  = 27
	lolcodeParserT__27  = 28
	lolcodeParserT__28  = 29
	lolcodeParserT__29  = 30
	lolcodeParserT__30  = 31
	lolcodeParserT__31  = 32
	lolcodeParserT__32  = 33
	lolcodeParserT__33  = 34
	lolcodeParserT__34  = 35
	lolcodeParserT__35  = 36
	lolcodeParserT__36  = 37
	lolcodeParserT__37  = 38
	lolcodeParserT__38  = 39
	lolcodeParserT__39  = 40
	lolcodeParserLABEL  = 41
	lolcodeParserATOM   = 42
	lolcodeParserSTRING = 43
	lolcodeParserWS     = 44
)

// lolcodeParser rules.
const (
	lolcodeParserRULE_program       = 0
	lolcodeParserRULE_code_block    = 1
	lolcodeParserRULE_statement     = 2
	lolcodeParserRULE_loop          = 3
	lolcodeParserRULE_declaration   = 4
	lolcodeParserRULE_comment       = 5
	lolcodeParserRULE_print_block   = 6
	lolcodeParserRULE_if_block      = 7
	lolcodeParserRULE_else_if_block = 8
	lolcodeParserRULE_input_block   = 9
	lolcodeParserRULE_func_decl     = 10
	lolcodeParserRULE_assignment    = 11
	lolcodeParserRULE_expression    = 12
	lolcodeParserRULE_equals        = 13
	lolcodeParserRULE_not_equals    = 14
	lolcodeParserRULE_both          = 15
	lolcodeParserRULE_either        = 16
	lolcodeParserRULE_greater       = 17
	lolcodeParserRULE_less          = 18
	lolcodeParserRULE_add           = 19
	lolcodeParserRULE_sub           = 20
	lolcodeParserRULE_mul           = 21
	lolcodeParserRULE_div           = 22
	lolcodeParserRULE_mod           = 23
	lolcodeParserRULE_cast          = 24
	lolcodeParserRULE_all_          = 25
	lolcodeParserRULE_any_          = 26
	lolcodeParserRULE_not_          = 27
	lolcodeParserRULE_func_         = 28
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
	p.RuleIndex = lolcodeParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Code_block() ICode_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *lolcodeParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lolcodeParserRULE_program)
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
		p.SetState(58)
		p.Match(lolcodeParserT__0)
	}
	{
		p.SetState(59)
		p.Code_block()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lolcodeParserT__1 {
		{
			p.SetState(60)
			p.Match(lolcodeParserT__1)
		}

	}

	return localctx
}

// ICode_blockContext is an interface to support dynamic dispatch.
type ICode_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCode_blockContext differentiates from other interfaces.
	IsCode_blockContext()
}

type Code_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_blockContext() *Code_blockContext {
	var p = new(Code_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_code_block
	return p
}

func (*Code_blockContext) IsCode_blockContext() {}

func NewCode_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_blockContext {
	var p = new(Code_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_code_block

	return p
}

func (s *Code_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_blockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *Code_blockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Code_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterCode_block(s)
	}
}

func (s *Code_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitCode_block(s)
	}
}

func (p *lolcodeParser) Code_block() (localctx ICode_blockContext) {
	this := p
	_ = this

	localctx = NewCode_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lolcodeParserRULE_code_block)
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lolcodeParserT__2)|(1<<lolcodeParserT__5)|(1<<lolcodeParserT__7)|(1<<lolcodeParserT__8)|(1<<lolcodeParserT__10)|(1<<lolcodeParserT__12)|(1<<lolcodeParserT__17)|(1<<lolcodeParserT__18)|(1<<lolcodeParserT__23)|(1<<lolcodeParserT__25)|(1<<lolcodeParserT__26)|(1<<lolcodeParserT__27)|(1<<lolcodeParserT__28)|(1<<lolcodeParserT__29)|(1<<lolcodeParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(lolcodeParserT__31-32))|(1<<(lolcodeParserT__32-32))|(1<<(lolcodeParserT__33-32))|(1<<(lolcodeParserT__34-32))|(1<<(lolcodeParserT__35-32))|(1<<(lolcodeParserT__37-32))|(1<<(lolcodeParserT__38-32))|(1<<(lolcodeParserT__39-32))|(1<<(lolcodeParserLABEL-32))|(1<<(lolcodeParserATOM-32)))) != 0) {
		{
			p.SetState(63)
			p.Statement()
		}

		p.SetState(66)
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
	p.RuleIndex = lolcodeParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *StatementContext) Print_block() IPrint_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint_blockContext)
}

func (s *StatementContext) If_block() IIf_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_blockContext)
}

func (s *StatementContext) Input_block() IInput_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInput_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInput_blockContext)
}

func (s *StatementContext) Func_decl() IFunc_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_declContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *lolcodeParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lolcodeParserRULE_statement)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Loop()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Comment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Print_block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.If_block()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.Input_block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(74)
			p.Func_decl()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.Assignment()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.Expression()
		}

	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) AllLABEL() []antlr.TerminalNode {
	return s.GetTokens(lolcodeParserLABEL)
}

func (s *LoopContext) LABEL(i int) antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, i)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) Code_block() ICode_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *lolcodeParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lolcodeParserRULE_loop)

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
		p.Match(lolcodeParserT__2)
	}
	{
		p.SetState(80)
		p.Match(lolcodeParserLABEL)
	}
	{
		p.SetState(81)
		p.Match(lolcodeParserT__3)
	}
	{
		p.SetState(82)
		p.Expression()
	}
	{
		p.SetState(83)
		p.Code_block()
	}
	{
		p.SetState(84)
		p.Match(lolcodeParserT__4)
	}
	{
		p.SetState(85)
		p.Match(lolcodeParserLABEL)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) LABEL() antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *lolcodeParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lolcodeParserRULE_declaration)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(lolcodeParserT__5)
		}
		{
			p.SetState(88)
			p.Match(lolcodeParserLABEL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(lolcodeParserT__5)
		}
		{
			p.SetState(90)
			p.Match(lolcodeParserLABEL)
		}
		{
			p.SetState(91)
			p.Match(lolcodeParserT__6)
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
	p.RuleIndex = lolcodeParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) STRING() antlr.TerminalNode {
	return s.GetToken(lolcodeParserSTRING, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *lolcodeParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lolcodeParserRULE_comment)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lolcodeParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(lolcodeParserT__7)
		}
		{
			p.SetState(95)
			p.Match(lolcodeParserSTRING)
		}

	case lolcodeParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(lolcodeParserT__8)
		}
		{
			p.SetState(97)
			p.Match(lolcodeParserSTRING)
		}
		{
			p.SetState(98)
			p.Match(lolcodeParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrint_blockContext is an interface to support dynamic dispatch.
type IPrint_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrint_blockContext differentiates from other interfaces.
	IsPrint_blockContext()
}

type Print_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_blockContext() *Print_blockContext {
	var p = new(Print_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_print_block
	return p
}

func (*Print_blockContext) IsPrint_blockContext() {}

func NewPrint_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_blockContext {
	var p = new(Print_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_print_block

	return p
}

func (s *Print_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_blockContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Print_blockContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Print_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterPrint_block(s)
	}
}

func (s *Print_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitPrint_block(s)
	}
}

func (p *lolcodeParser) Print_block() (localctx IPrint_blockContext) {
	this := p
	_ = this

	localctx = NewPrint_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lolcodeParserRULE_print_block)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(lolcodeParserT__10)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.Expression()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lolcodeParserT__11 {
		{
			p.SetState(108)
			p.Match(lolcodeParserT__11)
		}

	}

	return localctx
}

// IIf_blockContext is an interface to support dynamic dispatch.
type IIf_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_blockContext differentiates from other interfaces.
	IsIf_blockContext()
}

type If_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_blockContext() *If_blockContext {
	var p = new(If_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_if_block
	return p
}

func (*If_blockContext) IsIf_blockContext() {}

func NewIf_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_blockContext {
	var p = new(If_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_if_block

	return p
}

func (s *If_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *If_blockContext) Code_block() ICode_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *If_blockContext) Else_if_block() IElse_if_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_if_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElse_if_blockContext)
}

func (s *If_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterIf_block(s)
	}
}

func (s *If_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitIf_block(s)
	}
}

func (p *lolcodeParser) If_block() (localctx IIf_blockContext) {
	this := p
	_ = this

	localctx = NewIf_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lolcodeParserRULE_if_block)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(lolcodeParserT__12)
		}
		{
			p.SetState(112)
			p.Match(lolcodeParserT__13)
		}
		{
			p.SetState(113)
			p.Code_block()
		}
		{
			p.SetState(114)
			p.Match(lolcodeParserT__14)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(lolcodeParserT__12)
		}
		{
			p.SetState(117)
			p.Match(lolcodeParserT__13)
		}
		{
			p.SetState(118)
			p.Code_block()
		}
		{
			p.SetState(119)
			p.Else_if_block()
		}
		{
			p.SetState(120)
			p.Match(lolcodeParserT__14)
		}

	}

	return localctx
}

// IElse_if_blockContext is an interface to support dynamic dispatch.
type IElse_if_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElse_if_blockContext differentiates from other interfaces.
	IsElse_if_blockContext()
}

type Else_if_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_if_blockContext() *Else_if_blockContext {
	var p = new(Else_if_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_else_if_block
	return p
}

func (*Else_if_blockContext) IsElse_if_blockContext() {}

func NewElse_if_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_if_blockContext {
	var p = new(Else_if_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_else_if_block

	return p
}

func (s *Else_if_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_if_blockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Else_if_blockContext) Code_block() ICode_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Else_if_blockContext) Else_if_block() IElse_if_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_if_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElse_if_blockContext)
}

func (s *Else_if_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_if_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_if_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterElse_if_block(s)
	}
}

func (s *Else_if_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitElse_if_block(s)
	}
}

func (p *lolcodeParser) Else_if_block() (localctx IElse_if_blockContext) {
	this := p
	_ = this

	localctx = NewElse_if_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lolcodeParserRULE_else_if_block)

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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(lolcodeParserT__15)
		}
		{
			p.SetState(125)
			p.Expression()
		}
		{
			p.SetState(126)
			p.Code_block()
		}
		{
			p.SetState(127)
			p.Else_if_block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(lolcodeParserT__16)
		}
		{
			p.SetState(130)
			p.Code_block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(lolcodeParserT__15)
		}
		{
			p.SetState(132)
			p.Expression()
		}
		{
			p.SetState(133)
			p.Code_block()
		}

	}

	return localctx
}

// IInput_blockContext is an interface to support dynamic dispatch.
type IInput_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInput_blockContext differentiates from other interfaces.
	IsInput_blockContext()
}

type Input_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_blockContext() *Input_blockContext {
	var p = new(Input_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_input_block
	return p
}

func (*Input_blockContext) IsInput_blockContext() {}

func NewInput_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_blockContext {
	var p = new(Input_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_input_block

	return p
}

func (s *Input_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_blockContext) LABEL() antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, 0)
}

func (s *Input_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterInput_block(s)
	}
}

func (s *Input_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitInput_block(s)
	}
}

func (p *lolcodeParser) Input_block() (localctx IInput_blockContext) {
	this := p
	_ = this

	localctx = NewInput_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lolcodeParserRULE_input_block)

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
		p.SetState(137)
		p.Match(lolcodeParserT__17)
	}
	{
		p.SetState(138)
		p.Match(lolcodeParserLABEL)
	}

	return localctx
}

// IFunc_declContext is an interface to support dynamic dispatch.
type IFunc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_declContext differentiates from other interfaces.
	IsFunc_declContext()
}

type Func_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declContext() *Func_declContext {
	var p = new(Func_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_func_decl
	return p
}

func (*Func_declContext) IsFunc_declContext() {}

func NewFunc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declContext {
	var p = new(Func_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_func_decl

	return p
}

func (s *Func_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declContext) AllLABEL() []antlr.TerminalNode {
	return s.GetTokens(lolcodeParserLABEL)
}

func (s *Func_declContext) LABEL(i int) antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, i)
}

func (s *Func_declContext) Code_block() ICode_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Func_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterFunc_decl(s)
	}
}

func (s *Func_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitFunc_decl(s)
	}
}

func (p *lolcodeParser) Func_decl() (localctx IFunc_declContext) {
	this := p
	_ = this

	localctx = NewFunc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lolcodeParserRULE_func_decl)
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
		p.Match(lolcodeParserT__18)
	}
	{
		p.SetState(141)
		p.Match(lolcodeParserLABEL)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lolcodeParserT__19 {
		{
			p.SetState(142)
			p.Match(lolcodeParserT__19)
		}
		{
			p.SetState(143)
			p.Match(lolcodeParserLABEL)
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == lolcodeParserT__20 {
			{
				p.SetState(145)
				p.Match(lolcodeParserT__20)
			}
			{
				p.SetState(146)
				p.Match(lolcodeParserLABEL)
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(154)
		p.Code_block()
	}
	{
		p.SetState(155)
		p.Match(lolcodeParserT__21)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) LABEL() antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *lolcodeParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, lolcodeParserRULE_assignment)

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
		p.SetState(157)
		p.Match(lolcodeParserLABEL)
	}
	{
		p.SetState(158)
		p.Match(lolcodeParserT__22)
	}
	{
		p.SetState(159)
		p.Expression()
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
	p.RuleIndex = lolcodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Equals() IEqualsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualsContext)
}

func (s *ExpressionContext) Both() IBothContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBothContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBothContext)
}

func (s *ExpressionContext) Not_equals() INot_equalsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_equalsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_equalsContext)
}

func (s *ExpressionContext) Greater() IGreaterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGreaterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGreaterContext)
}

func (s *ExpressionContext) Less() ILessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILessContext)
}

func (s *ExpressionContext) Add() IAddContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddContext)
}

func (s *ExpressionContext) Sub() ISubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubContext)
}

func (s *ExpressionContext) Mul() IMulContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulContext)
}

func (s *ExpressionContext) Div() IDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDivContext)
}

func (s *ExpressionContext) Mod() IModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModContext)
}

func (s *ExpressionContext) Cast() ICastContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICastContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICastContext)
}

func (s *ExpressionContext) Either() IEitherContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEitherContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEitherContext)
}

func (s *ExpressionContext) All_() IAll_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAll_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAll_Context)
}

func (s *ExpressionContext) Any_() IAny_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_Context)
}

func (s *ExpressionContext) Not_() INot_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_Context)
}

func (s *ExpressionContext) Func_() IFunc_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_Context)
}

func (s *ExpressionContext) LABEL() antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, 0)
}

func (s *ExpressionContext) ATOM() antlr.TerminalNode {
	return s.GetToken(lolcodeParserATOM, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *lolcodeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, lolcodeParserRULE_expression)

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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Equals()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Both()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Not_equals()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.Greater()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.Less()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.Add()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(167)
			p.Sub()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)
			p.Mul()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Div()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(170)
			p.Mod()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(171)
			p.Cast()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(172)
			p.Either()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(173)
			p.All_()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(174)
			p.Any_()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(175)
			p.Not_()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(176)
			p.Func_()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(177)
			p.Match(lolcodeParserLABEL)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(178)
			p.Match(lolcodeParserATOM)
		}

	}

	return localctx
}

// IEqualsContext is an interface to support dynamic dispatch.
type IEqualsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualsContext differentiates from other interfaces.
	IsEqualsContext()
}

type EqualsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualsContext() *EqualsContext {
	var p = new(EqualsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_equals
	return p
}

func (*EqualsContext) IsEqualsContext() {}

func NewEqualsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualsContext {
	var p = new(EqualsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_equals

	return p
}

func (s *EqualsContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterEquals(s)
	}
}

func (s *EqualsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitEquals(s)
	}
}

func (p *lolcodeParser) Equals() (localctx IEqualsContext) {
	this := p
	_ = this

	localctx = NewEqualsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, lolcodeParserRULE_equals)

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
		p.Match(lolcodeParserT__23)
	}
	{
		p.SetState(182)
		p.Expression()
	}
	{
		p.SetState(183)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(184)
		p.Expression()
	}

	return localctx
}

// INot_equalsContext is an interface to support dynamic dispatch.
type INot_equalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_equalsContext differentiates from other interfaces.
	IsNot_equalsContext()
}

type Not_equalsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_equalsContext() *Not_equalsContext {
	var p = new(Not_equalsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_not_equals
	return p
}

func (*Not_equalsContext) IsNot_equalsContext() {}

func NewNot_equalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_equalsContext {
	var p = new(Not_equalsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_not_equals

	return p
}

func (s *Not_equalsContext) GetParser() antlr.Parser { return s.parser }

func (s *Not_equalsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Not_equalsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Not_equalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_equalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_equalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterNot_equals(s)
	}
}

func (s *Not_equalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitNot_equals(s)
	}
}

func (p *lolcodeParser) Not_equals() (localctx INot_equalsContext) {
	this := p
	_ = this

	localctx = NewNot_equalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, lolcodeParserRULE_not_equals)

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
		p.SetState(186)
		p.Match(lolcodeParserT__25)
	}
	{
		p.SetState(187)
		p.Expression()
	}
	{
		p.SetState(188)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(189)
		p.Expression()
	}

	return localctx
}

// IBothContext is an interface to support dynamic dispatch.
type IBothContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBothContext differentiates from other interfaces.
	IsBothContext()
}

type BothContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBothContext() *BothContext {
	var p = new(BothContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_both
	return p
}

func (*BothContext) IsBothContext() {}

func NewBothContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BothContext {
	var p = new(BothContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_both

	return p
}

func (s *BothContext) GetParser() antlr.Parser { return s.parser }

func (s *BothContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BothContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BothContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BothContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BothContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterBoth(s)
	}
}

func (s *BothContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitBoth(s)
	}
}

func (p *lolcodeParser) Both() (localctx IBothContext) {
	this := p
	_ = this

	localctx = NewBothContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, lolcodeParserRULE_both)

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
		p.SetState(191)
		p.Match(lolcodeParserT__26)
	}
	{
		p.SetState(192)
		p.Expression()
	}
	{
		p.SetState(193)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(194)
		p.Expression()
	}

	return localctx
}

// IEitherContext is an interface to support dynamic dispatch.
type IEitherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEitherContext differentiates from other interfaces.
	IsEitherContext()
}

type EitherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEitherContext() *EitherContext {
	var p = new(EitherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_either
	return p
}

func (*EitherContext) IsEitherContext() {}

func NewEitherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EitherContext {
	var p = new(EitherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_either

	return p
}

func (s *EitherContext) GetParser() antlr.Parser { return s.parser }

func (s *EitherContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EitherContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EitherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EitherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EitherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterEither(s)
	}
}

func (s *EitherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitEither(s)
	}
}

func (p *lolcodeParser) Either() (localctx IEitherContext) {
	this := p
	_ = this

	localctx = NewEitherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, lolcodeParserRULE_either)

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
		p.SetState(196)
		p.Match(lolcodeParserT__27)
	}
	{
		p.SetState(197)
		p.Expression()
	}
	{
		p.SetState(198)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(199)
		p.Expression()
	}

	return localctx
}

// IGreaterContext is an interface to support dynamic dispatch.
type IGreaterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGreaterContext differentiates from other interfaces.
	IsGreaterContext()
}

type GreaterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterContext() *GreaterContext {
	var p = new(GreaterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_greater
	return p
}

func (*GreaterContext) IsGreaterContext() {}

func NewGreaterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterContext {
	var p = new(GreaterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_greater

	return p
}

func (s *GreaterContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GreaterContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GreaterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterGreater(s)
	}
}

func (s *GreaterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitGreater(s)
	}
}

func (p *lolcodeParser) Greater() (localctx IGreaterContext) {
	this := p
	_ = this

	localctx = NewGreaterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, lolcodeParserRULE_greater)

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
		p.SetState(201)
		p.Match(lolcodeParserT__28)
	}
	{
		p.SetState(202)
		p.Expression()
	}
	{
		p.SetState(203)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(204)
		p.Expression()
	}

	return localctx
}

// ILessContext is an interface to support dynamic dispatch.
type ILessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLessContext differentiates from other interfaces.
	IsLessContext()
}

type LessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessContext() *LessContext {
	var p = new(LessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_less
	return p
}

func (*LessContext) IsLessContext() {}

func NewLessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessContext {
	var p = new(LessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_less

	return p
}

func (s *LessContext) GetParser() antlr.Parser { return s.parser }

func (s *LessContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterLess(s)
	}
}

func (s *LessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitLess(s)
	}
}

func (p *lolcodeParser) Less() (localctx ILessContext) {
	this := p
	_ = this

	localctx = NewLessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, lolcodeParserRULE_less)

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
		p.SetState(206)
		p.Match(lolcodeParserT__29)
	}
	{
		p.SetState(207)
		p.Expression()
	}
	{
		p.SetState(208)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(209)
		p.Expression()
	}

	return localctx
}

// IAddContext is an interface to support dynamic dispatch.
type IAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddContext differentiates from other interfaces.
	IsAddContext()
}

type AddContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddContext() *AddContext {
	var p = new(AddContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_add
	return p
}

func (*AddContext) IsAddContext() {}

func NewAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddContext {
	var p = new(AddContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_add

	return p
}

func (s *AddContext) GetParser() antlr.Parser { return s.parser }

func (s *AddContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitAdd(s)
	}
}

func (p *lolcodeParser) Add() (localctx IAddContext) {
	this := p
	_ = this

	localctx = NewAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, lolcodeParserRULE_add)

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
		p.SetState(211)
		p.Match(lolcodeParserT__30)
	}
	{
		p.SetState(212)
		p.Expression()
	}
	{
		p.SetState(213)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(214)
		p.Expression()
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
	p.RuleIndex = lolcodeParserRULE_sub
	return p
}

func (*SubContext) IsSubContext() {}

func NewSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubContext {
	var p = new(SubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_sub

	return p
}

func (s *SubContext) GetParser() antlr.Parser { return s.parser }

func (s *SubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitSub(s)
	}
}

func (p *lolcodeParser) Sub() (localctx ISubContext) {
	this := p
	_ = this

	localctx = NewSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, lolcodeParserRULE_sub)

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
		p.SetState(216)
		p.Match(lolcodeParserT__31)
	}
	{
		p.SetState(217)
		p.Expression()
	}
	{
		p.SetState(218)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(219)
		p.Expression()
	}

	return localctx
}

// IMulContext is an interface to support dynamic dispatch.
type IMulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulContext differentiates from other interfaces.
	IsMulContext()
}

type MulContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulContext() *MulContext {
	var p = new(MulContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_mul
	return p
}

func (*MulContext) IsMulContext() {}

func NewMulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulContext {
	var p = new(MulContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_mul

	return p
}

func (s *MulContext) GetParser() antlr.Parser { return s.parser }

func (s *MulContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitMul(s)
	}
}

func (p *lolcodeParser) Mul() (localctx IMulContext) {
	this := p
	_ = this

	localctx = NewMulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, lolcodeParserRULE_mul)

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
		p.SetState(221)
		p.Match(lolcodeParserT__32)
	}
	{
		p.SetState(222)
		p.Expression()
	}
	{
		p.SetState(223)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(224)
		p.Expression()
	}

	return localctx
}

// IDivContext is an interface to support dynamic dispatch.
type IDivContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDivContext differentiates from other interfaces.
	IsDivContext()
}

type DivContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDivContext() *DivContext {
	var p = new(DivContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_div
	return p
}

func (*DivContext) IsDivContext() {}

func NewDivContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DivContext {
	var p = new(DivContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_div

	return p
}

func (s *DivContext) GetParser() antlr.Parser { return s.parser }

func (s *DivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterDiv(s)
	}
}

func (s *DivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitDiv(s)
	}
}

func (p *lolcodeParser) Div() (localctx IDivContext) {
	this := p
	_ = this

	localctx = NewDivContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, lolcodeParserRULE_div)

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
		p.SetState(226)
		p.Match(lolcodeParserT__33)
	}
	{
		p.SetState(227)
		p.Expression()
	}
	{
		p.SetState(228)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(229)
		p.Expression()
	}

	return localctx
}

// IModContext is an interface to support dynamic dispatch.
type IModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModContext differentiates from other interfaces.
	IsModContext()
}

type ModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModContext() *ModContext {
	var p = new(ModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_mod
	return p
}

func (*ModContext) IsModContext() {}

func NewModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModContext {
	var p = new(ModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_mod

	return p
}

func (s *ModContext) GetParser() antlr.Parser { return s.parser }

func (s *ModContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ModContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterMod(s)
	}
}

func (s *ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitMod(s)
	}
}

func (p *lolcodeParser) Mod() (localctx IModContext) {
	this := p
	_ = this

	localctx = NewModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, lolcodeParserRULE_mod)

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
		p.SetState(231)
		p.Match(lolcodeParserT__34)
	}
	{
		p.SetState(232)
		p.Expression()
	}
	{
		p.SetState(233)
		p.Match(lolcodeParserT__24)
	}
	{
		p.SetState(234)
		p.Expression()
	}

	return localctx
}

// ICastContext is an interface to support dynamic dispatch.
type ICastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCastContext differentiates from other interfaces.
	IsCastContext()
}

type CastContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastContext() *CastContext {
	var p = new(CastContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_cast
	return p
}

func (*CastContext) IsCastContext() {}

func NewCastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastContext {
	var p = new(CastContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_cast

	return p
}

func (s *CastContext) GetParser() antlr.Parser { return s.parser }

func (s *CastContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterCast(s)
	}
}

func (s *CastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitCast(s)
	}
}

func (p *lolcodeParser) Cast() (localctx ICastContext) {
	this := p
	_ = this

	localctx = NewCastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, lolcodeParserRULE_cast)

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
		p.SetState(236)
		p.Match(lolcodeParserT__35)
	}
	{
		p.SetState(237)
		p.Expression()
	}
	{
		p.SetState(238)
		p.Match(lolcodeParserT__36)
	}

	return localctx
}

// IAll_Context is an interface to support dynamic dispatch.
type IAll_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAll_Context differentiates from other interfaces.
	IsAll_Context()
}

type All_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAll_Context() *All_Context {
	var p = new(All_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_all_
	return p
}

func (*All_Context) IsAll_Context() {}

func NewAll_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *All_Context {
	var p = new(All_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_all_

	return p
}

func (s *All_Context) GetParser() antlr.Parser { return s.parser }

func (s *All_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *All_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *All_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *All_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *All_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterAll_(s)
	}
}

func (s *All_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitAll_(s)
	}
}

func (p *lolcodeParser) All_() (localctx IAll_Context) {
	this := p
	_ = this

	localctx = NewAll_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, lolcodeParserRULE_all_)
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
		p.SetState(240)
		p.Match(lolcodeParserT__37)
	}
	{
		p.SetState(241)
		p.Expression()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lolcodeParserT__24 {
		{
			p.SetState(242)
			p.Match(lolcodeParserT__24)
		}
		{
			p.SetState(243)
			p.Expression()
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(249)
		p.Match(lolcodeParserT__11)
	}

	return localctx
}

// IAny_Context is an interface to support dynamic dispatch.
type IAny_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_Context differentiates from other interfaces.
	IsAny_Context()
}

type Any_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_Context() *Any_Context {
	var p = new(Any_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_any_
	return p
}

func (*Any_Context) IsAny_Context() {}

func NewAny_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_Context {
	var p = new(Any_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_any_

	return p
}

func (s *Any_Context) GetParser() antlr.Parser { return s.parser }

func (s *Any_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Any_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Any_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterAny_(s)
	}
}

func (s *Any_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitAny_(s)
	}
}

func (p *lolcodeParser) Any_() (localctx IAny_Context) {
	this := p
	_ = this

	localctx = NewAny_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, lolcodeParserRULE_any_)
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
		p.SetState(251)
		p.Match(lolcodeParserT__38)
	}
	{
		p.SetState(252)
		p.Expression()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lolcodeParserT__24 {
		{
			p.SetState(253)
			p.Match(lolcodeParserT__24)
		}
		{
			p.SetState(254)
			p.Expression()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
		p.Match(lolcodeParserT__11)
	}

	return localctx
}

// INot_Context is an interface to support dynamic dispatch.
type INot_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_Context differentiates from other interfaces.
	IsNot_Context()
}

type Not_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_Context() *Not_Context {
	var p = new(Not_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_not_
	return p
}

func (*Not_Context) IsNot_Context() {}

func NewNot_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_Context {
	var p = new(Not_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_not_

	return p
}

func (s *Not_Context) GetParser() antlr.Parser { return s.parser }

func (s *Not_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Not_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterNot_(s)
	}
}

func (s *Not_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitNot_(s)
	}
}

func (p *lolcodeParser) Not_() (localctx INot_Context) {
	this := p
	_ = this

	localctx = NewNot_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, lolcodeParserRULE_not_)

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
		p.SetState(262)
		p.Match(lolcodeParserT__39)
	}
	{
		p.SetState(263)
		p.Expression()
	}

	return localctx
}

// IFunc_Context is an interface to support dynamic dispatch.
type IFunc_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_Context differentiates from other interfaces.
	IsFunc_Context()
}

type Func_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_Context() *Func_Context {
	var p = new(Func_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lolcodeParserRULE_func_
	return p
}

func (*Func_Context) IsFunc_Context() {}

func NewFunc_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_Context {
	var p = new(Func_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lolcodeParserRULE_func_

	return p
}

func (s *Func_Context) GetParser() antlr.Parser { return s.parser }

func (s *Func_Context) LABEL() antlr.TerminalNode {
	return s.GetToken(lolcodeParserLABEL, 0)
}

func (s *Func_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Func_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Func_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.EnterFunc_(s)
	}
}

func (s *Func_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lolcodeListener); ok {
		listenerT.ExitFunc_(s)
	}
}

func (p *lolcodeParser) Func_() (localctx IFunc_Context) {
	this := p
	_ = this

	localctx = NewFunc_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, lolcodeParserRULE_func_)
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
		p.SetState(265)
		p.Match(lolcodeParserLABEL)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(lolcodeParserT__23-24))|(1<<(lolcodeParserT__25-24))|(1<<(lolcodeParserT__26-24))|(1<<(lolcodeParserT__27-24))|(1<<(lolcodeParserT__28-24))|(1<<(lolcodeParserT__29-24))|(1<<(lolcodeParserT__30-24))|(1<<(lolcodeParserT__31-24))|(1<<(lolcodeParserT__32-24))|(1<<(lolcodeParserT__33-24))|(1<<(lolcodeParserT__34-24))|(1<<(lolcodeParserT__35-24))|(1<<(lolcodeParserT__37-24))|(1<<(lolcodeParserT__38-24))|(1<<(lolcodeParserT__39-24))|(1<<(lolcodeParserLABEL-24))|(1<<(lolcodeParserATOM-24)))) != 0) {
		{
			p.SetState(266)
			p.Expression()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(271)
		p.Match(lolcodeParserT__11)
	}

	return localctx
}
