// Generated from morsecode.g4 by ANTLR 4.7.

package morsecode // morsecode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 292, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18,
	4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 4,
	24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 29,
	9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 9,
	34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 39,
	3, 2, 3, 2, 3, 2, 6, 2, 82, 10, 2, 13, 2, 14, 2, 83, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 122, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 2, 2, 40, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	2, 2, 2, 289, 2, 78, 3, 2, 2, 2, 4, 121, 3, 2, 2, 2, 6, 123, 3, 2, 2, 2,
	8, 126, 3, 2, 2, 2, 10, 131, 3, 2, 2, 2, 12, 136, 3, 2, 2, 2, 14, 140,
	3, 2, 2, 2, 16, 142, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20, 151, 3, 2, 2,
	2, 22, 156, 3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 164, 3, 2, 2, 2, 28, 168,
	3, 2, 2, 2, 30, 173, 3, 2, 2, 2, 32, 176, 3, 2, 2, 2, 34, 179, 3, 2, 2,
	2, 36, 183, 3, 2, 2, 2, 38, 188, 3, 2, 2, 2, 40, 193, 3, 2, 2, 2, 42, 197,
	3, 2, 2, 2, 44, 201, 3, 2, 2, 2, 46, 203, 3, 2, 2, 2, 48, 207, 3, 2, 2,
	2, 50, 212, 3, 2, 2, 2, 52, 216, 3, 2, 2, 2, 54, 221, 3, 2, 2, 2, 56, 226,
	3, 2, 2, 2, 58, 231, 3, 2, 2, 2, 60, 237, 3, 2, 2, 2, 62, 243, 3, 2, 2,
	2, 64, 249, 3, 2, 2, 2, 66, 255, 3, 2, 2, 2, 68, 261, 3, 2, 2, 2, 70, 267,
	3, 2, 2, 2, 72, 273, 3, 2, 2, 2, 74, 279, 3, 2, 2, 2, 76, 285, 3, 2, 2,
	2, 78, 81, 5, 4, 3, 2, 79, 80, 7, 5, 2, 2, 80, 82, 5, 4, 3, 2, 81, 79,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2,
	84, 3, 3, 2, 2, 2, 85, 122, 5, 6, 4, 2, 86, 122, 5, 8, 5, 2, 87, 122, 5,
	10, 6, 2, 88, 122, 5, 12, 7, 2, 89, 122, 5, 14, 8, 2, 90, 122, 5, 16, 9,
	2, 91, 122, 5, 18, 10, 2, 92, 122, 5, 20, 11, 2, 93, 122, 5, 22, 12, 2,
	94, 122, 5, 24, 13, 2, 95, 122, 5, 26, 14, 2, 96, 122, 5, 28, 15, 2, 97,
	122, 5, 30, 16, 2, 98, 122, 5, 32, 17, 2, 99, 122, 5, 34, 18, 2, 100, 122,
	5, 36, 19, 2, 101, 122, 5, 38, 20, 2, 102, 122, 5, 40, 21, 2, 103, 122,
	5, 42, 22, 2, 104, 122, 5, 44, 23, 2, 105, 122, 5, 46, 24, 2, 106, 122,
	5, 48, 25, 2, 107, 122, 5, 50, 26, 2, 108, 122, 5, 52, 27, 2, 109, 122,
	5, 54, 28, 2, 110, 122, 5, 56, 29, 2, 111, 122, 5, 58, 30, 2, 112, 122,
	5, 60, 31, 2, 113, 122, 5, 62, 32, 2, 114, 122, 5, 64, 33, 2, 115, 122,
	5, 66, 34, 2, 116, 122, 5, 68, 35, 2, 117, 122, 5, 70, 36, 2, 118, 122,
	5, 72, 37, 2, 119, 122, 5, 74, 38, 2, 120, 122, 5, 76, 39, 2, 121, 85,
	3, 2, 2, 2, 121, 86, 3, 2, 2, 2, 121, 87, 3, 2, 2, 2, 121, 88, 3, 2, 2,
	2, 121, 89, 3, 2, 2, 2, 121, 90, 3, 2, 2, 2, 121, 91, 3, 2, 2, 2, 121,
	92, 3, 2, 2, 2, 121, 93, 3, 2, 2, 2, 121, 94, 3, 2, 2, 2, 121, 95, 3, 2,
	2, 2, 121, 96, 3, 2, 2, 2, 121, 97, 3, 2, 2, 2, 121, 98, 3, 2, 2, 2, 121,
	99, 3, 2, 2, 2, 121, 100, 3, 2, 2, 2, 121, 101, 3, 2, 2, 2, 121, 102, 3,
	2, 2, 2, 121, 103, 3, 2, 2, 2, 121, 104, 3, 2, 2, 2, 121, 105, 3, 2, 2,
	2, 121, 106, 3, 2, 2, 2, 121, 107, 3, 2, 2, 2, 121, 108, 3, 2, 2, 2, 121,
	109, 3, 2, 2, 2, 121, 110, 3, 2, 2, 2, 121, 111, 3, 2, 2, 2, 121, 112,
	3, 2, 2, 2, 121, 113, 3, 2, 2, 2, 121, 114, 3, 2, 2, 2, 121, 115, 3, 2,
	2, 2, 121, 116, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 121, 118, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 5, 3, 2, 2, 2, 123, 124,
	7, 3, 2, 2, 124, 125, 7, 4, 2, 2, 125, 7, 3, 2, 2, 2, 126, 127, 7, 4, 2,
	2, 127, 128, 7, 3, 2, 2, 128, 129, 7, 3, 2, 2, 129, 130, 7, 3, 2, 2, 130,
	9, 3, 2, 2, 2, 131, 132, 7, 4, 2, 2, 132, 133, 7, 3, 2, 2, 133, 134, 7,
	4, 2, 2, 134, 135, 7, 3, 2, 2, 135, 11, 3, 2, 2, 2, 136, 137, 7, 4, 2,
	2, 137, 138, 7, 3, 2, 2, 138, 139, 7, 3, 2, 2, 139, 13, 3, 2, 2, 2, 140,
	141, 7, 3, 2, 2, 141, 15, 3, 2, 2, 2, 142, 143, 7, 3, 2, 2, 143, 144, 7,
	3, 2, 2, 144, 145, 7, 4, 2, 2, 145, 146, 7, 3, 2, 2, 146, 17, 3, 2, 2,
	2, 147, 148, 7, 4, 2, 2, 148, 149, 7, 4, 2, 2, 149, 150, 7, 3, 2, 2, 150,
	19, 3, 2, 2, 2, 151, 152, 7, 3, 2, 2, 152, 153, 7, 3, 2, 2, 153, 154, 7,
	3, 2, 2, 154, 155, 7, 3, 2, 2, 155, 21, 3, 2, 2, 2, 156, 157, 7, 3, 2,
	2, 157, 158, 7, 3, 2, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7, 3, 2, 2, 160,
	161, 7, 4, 2, 2, 161, 162, 7, 4, 2, 2, 162, 163, 7, 4, 2, 2, 163, 25, 3,
	2, 2, 2, 164, 165, 7, 4, 2, 2, 165, 166, 7, 3, 2, 2, 166, 167, 7, 4, 2,
	2, 167, 27, 3, 2, 2, 2, 168, 169, 7, 3, 2, 2, 169, 170, 7, 4, 2, 2, 170,
	171, 7, 3, 2, 2, 171, 172, 7, 3, 2, 2, 172, 29, 3, 2, 2, 2, 173, 174, 7,
	4, 2, 2, 174, 175, 7, 4, 2, 2, 175, 31, 3, 2, 2, 2, 176, 177, 7, 4, 2,
	2, 177, 178, 7, 3, 2, 2, 178, 33, 3, 2, 2, 2, 179, 180, 7, 4, 2, 2, 180,
	181, 7, 4, 2, 2, 181, 182, 7, 4, 2, 2, 182, 35, 3, 2, 2, 2, 183, 184, 7,
	3, 2, 2, 184, 185, 7, 4, 2, 2, 185, 186, 7, 4, 2, 2, 186, 187, 7, 3, 2,
	2, 187, 37, 3, 2, 2, 2, 188, 189, 7, 4, 2, 2, 189, 190, 7, 4, 2, 2, 190,
	191, 7, 3, 2, 2, 191, 192, 7, 4, 2, 2, 192, 39, 3, 2, 2, 2, 193, 194, 7,
	3, 2, 2, 194, 195, 7, 4, 2, 2, 195, 196, 7, 3, 2, 2, 196, 41, 3, 2, 2,
	2, 197, 198, 7, 3, 2, 2, 198, 199, 7, 3, 2, 2, 199, 200, 7, 3, 2, 2, 200,
	43, 3, 2, 2, 2, 201, 202, 7, 4, 2, 2, 202, 45, 3, 2, 2, 2, 203, 204, 7,
	3, 2, 2, 204, 205, 7, 3, 2, 2, 205, 206, 7, 4, 2, 2, 206, 47, 3, 2, 2,
	2, 207, 208, 7, 3, 2, 2, 208, 209, 7, 3, 2, 2, 209, 210, 7, 3, 2, 2, 210,
	211, 7, 4, 2, 2, 211, 49, 3, 2, 2, 2, 212, 213, 7, 3, 2, 2, 213, 214, 7,
	4, 2, 2, 214, 215, 7, 4, 2, 2, 215, 51, 3, 2, 2, 2, 216, 217, 7, 4, 2,
	2, 217, 218, 7, 3, 2, 2, 218, 219, 7, 3, 2, 2, 219, 220, 7, 4, 2, 2, 220,
	53, 3, 2, 2, 2, 221, 222, 7, 4, 2, 2, 222, 223, 7, 3, 2, 2, 223, 224, 7,
	4, 2, 2, 224, 225, 7, 4, 2, 2, 225, 55, 3, 2, 2, 2, 226, 227, 7, 4, 2,
	2, 227, 228, 7, 4, 2, 2, 228, 229, 7, 3, 2, 2, 229, 230, 7, 3, 2, 2, 230,
	57, 3, 2, 2, 2, 231, 232, 7, 3, 2, 2, 232, 233, 7, 4, 2, 2, 233, 234, 7,
	4, 2, 2, 234, 235, 7, 4, 2, 2, 235, 236, 7, 4, 2, 2, 236, 59, 3, 2, 2,
	2, 237, 238, 7, 3, 2, 2, 238, 239, 7, 3, 2, 2, 239, 240, 7, 4, 2, 2, 240,
	241, 7, 4, 2, 2, 241, 242, 7, 4, 2, 2, 242, 61, 3, 2, 2, 2, 243, 244, 7,
	3, 2, 2, 244, 245, 7, 3, 2, 2, 245, 246, 7, 3, 2, 2, 246, 247, 7, 4, 2,
	2, 247, 248, 7, 4, 2, 2, 248, 63, 3, 2, 2, 2, 249, 250, 7, 3, 2, 2, 250,
	251, 7, 3, 2, 2, 251, 252, 7, 3, 2, 2, 252, 253, 7, 3, 2, 2, 253, 254,
	7, 4, 2, 2, 254, 65, 3, 2, 2, 2, 255, 256, 7, 3, 2, 2, 256, 257, 7, 3,
	2, 2, 257, 258, 7, 3, 2, 2, 258, 259, 7, 3, 2, 2, 259, 260, 7, 3, 2, 2,
	260, 67, 3, 2, 2, 2, 261, 262, 7, 4, 2, 2, 262, 263, 7, 3, 2, 2, 263, 264,
	7, 3, 2, 2, 264, 265, 7, 3, 2, 2, 265, 266, 7, 3, 2, 2, 266, 69, 3, 2,
	2, 2, 267, 268, 7, 4, 2, 2, 268, 269, 7, 4, 2, 2, 269, 270, 7, 3, 2, 2,
	270, 271, 7, 3, 2, 2, 271, 272, 7, 3, 2, 2, 272, 71, 3, 2, 2, 2, 273, 274,
	7, 4, 2, 2, 274, 275, 7, 4, 2, 2, 275, 276, 7, 4, 2, 2, 276, 277, 7, 3,
	2, 2, 277, 278, 7, 3, 2, 2, 278, 73, 3, 2, 2, 2, 279, 280, 7, 4, 2, 2,
	280, 281, 7, 4, 2, 2, 281, 282, 7, 4, 2, 2, 282, 283, 7, 4, 2, 2, 283,
	284, 7, 3, 2, 2, 284, 75, 3, 2, 2, 2, 285, 286, 7, 4, 2, 2, 286, 287, 7,
	4, 2, 2, 287, 288, 7, 4, 2, 2, 288, 289, 7, 4, 2, 2, 289, 290, 7, 4, 2,
	2, 290, 77, 3, 2, 2, 2, 4, 83, 121,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'-'", "' '",
}
var symbolicNames = []string{
	"", "DOT", "DASH", "SPACE", "WS",
}

var ruleNames = []string{
	"morsecode", "letter", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y",
	"z", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"zero",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type morsecodeParser struct {
	*antlr.BaseParser
}

func NewmorsecodeParser(input antlr.TokenStream) *morsecodeParser {
	this := new(morsecodeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "morsecode.g4"

	return this
}

// morsecodeParser tokens.
const (
	morsecodeParserEOF   = antlr.TokenEOF
	morsecodeParserDOT   = 1
	morsecodeParserDASH  = 2
	morsecodeParserSPACE = 3
	morsecodeParserWS    = 4
)

// morsecodeParser rules.
const (
	morsecodeParserRULE_morsecode = 0
	morsecodeParserRULE_letter    = 1
	morsecodeParserRULE_a         = 2
	morsecodeParserRULE_b         = 3
	morsecodeParserRULE_c         = 4
	morsecodeParserRULE_d         = 5
	morsecodeParserRULE_e         = 6
	morsecodeParserRULE_f         = 7
	morsecodeParserRULE_g         = 8
	morsecodeParserRULE_h         = 9
	morsecodeParserRULE_i         = 10
	morsecodeParserRULE_j         = 11
	morsecodeParserRULE_k         = 12
	morsecodeParserRULE_l         = 13
	morsecodeParserRULE_m         = 14
	morsecodeParserRULE_n         = 15
	morsecodeParserRULE_o         = 16
	morsecodeParserRULE_p         = 17
	morsecodeParserRULE_q         = 18
	morsecodeParserRULE_r         = 19
	morsecodeParserRULE_s         = 20
	morsecodeParserRULE_t         = 21
	morsecodeParserRULE_u         = 22
	morsecodeParserRULE_v         = 23
	morsecodeParserRULE_w         = 24
	morsecodeParserRULE_x         = 25
	morsecodeParserRULE_y         = 26
	morsecodeParserRULE_z         = 27
	morsecodeParserRULE_one       = 28
	morsecodeParserRULE_two       = 29
	morsecodeParserRULE_three     = 30
	morsecodeParserRULE_four      = 31
	morsecodeParserRULE_five      = 32
	morsecodeParserRULE_six       = 33
	morsecodeParserRULE_seven     = 34
	morsecodeParserRULE_eight     = 35
	morsecodeParserRULE_nine      = 36
	morsecodeParserRULE_zero      = 37
)

// IMorsecodeContext is an interface to support dynamic dispatch.
type IMorsecodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMorsecodeContext differentiates from other interfaces.
	IsMorsecodeContext()
}

type MorsecodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMorsecodeContext() *MorsecodeContext {
	var p = new(MorsecodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_morsecode
	return p
}

func (*MorsecodeContext) IsMorsecodeContext() {}

func NewMorsecodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MorsecodeContext {
	var p = new(MorsecodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_morsecode

	return p
}

func (s *MorsecodeContext) GetParser() antlr.Parser { return s.parser }

func (s *MorsecodeContext) AllLetter() []ILetterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILetterContext)(nil)).Elem())
	var tst = make([]ILetterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILetterContext)
		}
	}

	return tst
}

func (s *MorsecodeContext) Letter(i int) ILetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILetterContext)
}

func (s *MorsecodeContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserSPACE)
}

func (s *MorsecodeContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserSPACE, i)
}

func (s *MorsecodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MorsecodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MorsecodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterMorsecode(s)
	}
}

func (s *MorsecodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitMorsecode(s)
	}
}

func (p *morsecodeParser) Morsecode() (localctx IMorsecodeContext) {
	localctx = NewMorsecodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, morsecodeParserRULE_morsecode)
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
		p.SetState(76)
		p.Letter()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == morsecodeParserSPACE {
		{
			p.SetState(77)
			p.Match(morsecodeParserSPACE)
		}
		{
			p.SetState(78)
			p.Letter()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILetterContext is an interface to support dynamic dispatch.
type ILetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetterContext differentiates from other interfaces.
	IsLetterContext()
}

type LetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetterContext() *LetterContext {
	var p = new(LetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_letter
	return p
}

func (*LetterContext) IsLetterContext() {}

func NewLetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetterContext {
	var p = new(LetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_letter

	return p
}

func (s *LetterContext) GetParser() antlr.Parser { return s.parser }

func (s *LetterContext) A() IAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *LetterContext) B() IBContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBContext)
}

func (s *LetterContext) C() ICContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICContext)
}

func (s *LetterContext) D() IDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDContext)
}

func (s *LetterContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *LetterContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *LetterContext) G() IGContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGContext)
}

func (s *LetterContext) H() IHContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHContext)
}

func (s *LetterContext) I() IIContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIContext)
}

func (s *LetterContext) J() IJContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJContext)
}

func (s *LetterContext) K() IKContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKContext)
}

func (s *LetterContext) L() ILContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILContext)
}

func (s *LetterContext) M() IMContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMContext)
}

func (s *LetterContext) N() INContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INContext)
}

func (s *LetterContext) O() IOContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOContext)
}

func (s *LetterContext) P() IPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPContext)
}

func (s *LetterContext) Q() IQContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQContext)
}

func (s *LetterContext) R() IRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRContext)
}

func (s *LetterContext) S() ISContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISContext)
}

func (s *LetterContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *LetterContext) U() IUContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUContext)
}

func (s *LetterContext) V() IVContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVContext)
}

func (s *LetterContext) W() IWContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWContext)
}

func (s *LetterContext) X() IXContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXContext)
}

func (s *LetterContext) Y() IYContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYContext)
}

func (s *LetterContext) Z() IZContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZContext)
}

func (s *LetterContext) One() IOneContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneContext)
}

func (s *LetterContext) Two() ITwoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITwoContext)
}

func (s *LetterContext) Three() IThreeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThreeContext)
}

func (s *LetterContext) Four() IFourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFourContext)
}

func (s *LetterContext) Five() IFiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFiveContext)
}

func (s *LetterContext) Six() ISixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISixContext)
}

func (s *LetterContext) Seven() ISevenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISevenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISevenContext)
}

func (s *LetterContext) Eight() IEightContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEightContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEightContext)
}

func (s *LetterContext) Nine() INineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INineContext)
}

func (s *LetterContext) Zero() IZeroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZeroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZeroContext)
}

func (s *LetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterLetter(s)
	}
}

func (s *LetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitLetter(s)
	}
}

func (p *morsecodeParser) Letter() (localctx ILetterContext) {
	localctx = NewLetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, morsecodeParserRULE_letter)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.A()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.B()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.C()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.D()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(87)
			p.E()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.F()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)
			p.G()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(90)
			p.H()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(91)
			p.I()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(92)
			p.J()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(93)
			p.K()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(94)
			p.L()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(95)
			p.M()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(96)
			p.N()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(97)
			p.O()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(98)
			p.P()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(99)
			p.Q()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(100)
			p.R()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(101)
			p.S()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(102)
			p.T()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(103)
			p.U()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(104)
			p.V()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(105)
			p.W()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(106)
			p.X()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(107)
			p.Y()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(108)
			p.Z()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(109)
			p.One()
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(110)
			p.Two()
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(111)
			p.Three()
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(112)
			p.Four()
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(113)
			p.Five()
		}

	case 32:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(114)
			p.Six()
		}

	case 33:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(115)
			p.Seven()
		}

	case 34:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(116)
			p.Eight()
		}

	case 35:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(117)
			p.Nine()
		}

	case 36:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(118)
			p.Zero()
		}

	}

	return localctx
}

// IAContext is an interface to support dynamic dispatch.
type IAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAContext differentiates from other interfaces.
	IsAContext()
}

type AContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAContext() *AContext {
	var p = new(AContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_a
	return p
}

func (*AContext) IsAContext() {}

func NewAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AContext {
	var p = new(AContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_a

	return p
}

func (s *AContext) GetParser() antlr.Parser { return s.parser }

func (s *AContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *AContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *AContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterA(s)
	}
}

func (s *AContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitA(s)
	}
}

func (p *morsecodeParser) A() (localctx IAContext) {
	localctx = NewAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, morsecodeParserRULE_a)

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
		p.SetState(121)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(122)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IBContext is an interface to support dynamic dispatch.
type IBContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBContext differentiates from other interfaces.
	IsBContext()
}

type BContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBContext() *BContext {
	var p = new(BContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_b
	return p
}

func (*BContext) IsBContext() {}

func NewBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BContext {
	var p = new(BContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_b

	return p
}

func (s *BContext) GetParser() antlr.Parser { return s.parser }

func (s *BContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *BContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *BContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *BContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterB(s)
	}
}

func (s *BContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitB(s)
	}
}

func (p *morsecodeParser) B() (localctx IBContext) {
	localctx = NewBContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, morsecodeParserRULE_b)

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
		p.SetState(124)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(125)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(126)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(127)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// ICContext is an interface to support dynamic dispatch.
type ICContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCContext differentiates from other interfaces.
	IsCContext()
}

type CContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCContext() *CContext {
	var p = new(CContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_c
	return p
}

func (*CContext) IsCContext() {}

func NewCContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CContext {
	var p = new(CContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_c

	return p
}

func (s *CContext) GetParser() antlr.Parser { return s.parser }

func (s *CContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *CContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *CContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *CContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *CContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterC(s)
	}
}

func (s *CContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitC(s)
	}
}

func (p *morsecodeParser) C() (localctx ICContext) {
	localctx = NewCContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, morsecodeParserRULE_c)

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
		p.SetState(129)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(130)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(131)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(132)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IDContext is an interface to support dynamic dispatch.
type IDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDContext differentiates from other interfaces.
	IsDContext()
}

type DContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDContext() *DContext {
	var p = new(DContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_d
	return p
}

func (*DContext) IsDContext() {}

func NewDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DContext {
	var p = new(DContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_d

	return p
}

func (s *DContext) GetParser() antlr.Parser { return s.parser }

func (s *DContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *DContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *DContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *DContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterD(s)
	}
}

func (s *DContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitD(s)
	}
}

func (p *morsecodeParser) D() (localctx IDContext) {
	localctx = NewDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, morsecodeParserRULE_d)

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
		p.SetState(134)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(135)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(136)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterE(s)
	}
}

func (s *EContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitE(s)
	}
}

func (p *morsecodeParser) E() (localctx IEContext) {
	localctx = NewEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, morsecodeParserRULE_e)

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
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IFContext is an interface to support dynamic dispatch.
type IFContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFContext differentiates from other interfaces.
	IsFContext()
}

type FContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_f
	return p
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *FContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *FContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterF(s)
	}
}

func (s *FContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitF(s)
	}
}

func (p *morsecodeParser) F() (localctx IFContext) {
	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, morsecodeParserRULE_f)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(141)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(142)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(143)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IGContext is an interface to support dynamic dispatch.
type IGContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGContext differentiates from other interfaces.
	IsGContext()
}

type GContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGContext() *GContext {
	var p = new(GContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_g
	return p
}

func (*GContext) IsGContext() {}

func NewGContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GContext {
	var p = new(GContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_g

	return p
}

func (s *GContext) GetParser() antlr.Parser { return s.parser }

func (s *GContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *GContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *GContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *GContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterG(s)
	}
}

func (s *GContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitG(s)
	}
}

func (p *morsecodeParser) G() (localctx IGContext) {
	localctx = NewGContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, morsecodeParserRULE_g)

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
		p.SetState(145)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(146)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(147)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IHContext is an interface to support dynamic dispatch.
type IHContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHContext differentiates from other interfaces.
	IsHContext()
}

type HContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHContext() *HContext {
	var p = new(HContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_h
	return p
}

func (*HContext) IsHContext() {}

func NewHContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HContext {
	var p = new(HContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_h

	return p
}

func (s *HContext) GetParser() antlr.Parser { return s.parser }

func (s *HContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *HContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *HContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterH(s)
	}
}

func (s *HContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitH(s)
	}
}

func (p *morsecodeParser) H() (localctx IHContext) {
	localctx = NewHContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, morsecodeParserRULE_h)

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
		p.SetState(149)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(150)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(151)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(152)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IIContext is an interface to support dynamic dispatch.
type IIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIContext differentiates from other interfaces.
	IsIContext()
}

type IContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIContext() *IContext {
	var p = new(IContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_i
	return p
}

func (*IContext) IsIContext() {}

func NewIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IContext {
	var p = new(IContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_i

	return p
}

func (s *IContext) GetParser() antlr.Parser { return s.parser }

func (s *IContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *IContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *IContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterI(s)
	}
}

func (s *IContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitI(s)
	}
}

func (p *morsecodeParser) I() (localctx IIContext) {
	localctx = NewIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, morsecodeParserRULE_i)

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
		p.SetState(154)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(155)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IJContext is an interface to support dynamic dispatch.
type IJContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJContext differentiates from other interfaces.
	IsJContext()
}

type JContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJContext() *JContext {
	var p = new(JContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_j
	return p
}

func (*JContext) IsJContext() {}

func NewJContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JContext {
	var p = new(JContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_j

	return p
}

func (s *JContext) GetParser() antlr.Parser { return s.parser }

func (s *JContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *JContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *JContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *JContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterJ(s)
	}
}

func (s *JContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitJ(s)
	}
}

func (p *morsecodeParser) J() (localctx IJContext) {
	localctx = NewJContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, morsecodeParserRULE_j)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(158)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(159)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(160)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IKContext is an interface to support dynamic dispatch.
type IKContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKContext differentiates from other interfaces.
	IsKContext()
}

type KContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKContext() *KContext {
	var p = new(KContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_k
	return p
}

func (*KContext) IsKContext() {}

func NewKContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KContext {
	var p = new(KContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_k

	return p
}

func (s *KContext) GetParser() antlr.Parser { return s.parser }

func (s *KContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *KContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *KContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *KContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterK(s)
	}
}

func (s *KContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitK(s)
	}
}

func (p *morsecodeParser) K() (localctx IKContext) {
	localctx = NewKContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, morsecodeParserRULE_k)

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
		p.SetState(162)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(163)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(164)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// ILContext is an interface to support dynamic dispatch.
type ILContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLContext differentiates from other interfaces.
	IsLContext()
}

type LContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLContext() *LContext {
	var p = new(LContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_l
	return p
}

func (*LContext) IsLContext() {}

func NewLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LContext {
	var p = new(LContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_l

	return p
}

func (s *LContext) GetParser() antlr.Parser { return s.parser }

func (s *LContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *LContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *LContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterL(s)
	}
}

func (s *LContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitL(s)
	}
}

func (p *morsecodeParser) L() (localctx ILContext) {
	localctx = NewLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, morsecodeParserRULE_l)

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
		p.SetState(166)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(167)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(168)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(169)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IMContext is an interface to support dynamic dispatch.
type IMContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMContext differentiates from other interfaces.
	IsMContext()
}

type MContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMContext() *MContext {
	var p = new(MContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_m
	return p
}

func (*MContext) IsMContext() {}

func NewMContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MContext {
	var p = new(MContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_m

	return p
}

func (s *MContext) GetParser() antlr.Parser { return s.parser }

func (s *MContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *MContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *MContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterM(s)
	}
}

func (s *MContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitM(s)
	}
}

func (p *morsecodeParser) M() (localctx IMContext) {
	localctx = NewMContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, morsecodeParserRULE_m)

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
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(172)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// INContext is an interface to support dynamic dispatch.
type INContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNContext differentiates from other interfaces.
	IsNContext()
}

type NContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNContext() *NContext {
	var p = new(NContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_n
	return p
}

func (*NContext) IsNContext() {}

func NewNContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NContext {
	var p = new(NContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_n

	return p
}

func (s *NContext) GetParser() antlr.Parser { return s.parser }

func (s *NContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *NContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *NContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterN(s)
	}
}

func (s *NContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitN(s)
	}
}

func (p *morsecodeParser) N() (localctx INContext) {
	localctx = NewNContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, morsecodeParserRULE_n)

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
		p.SetState(174)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(175)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IOContext is an interface to support dynamic dispatch.
type IOContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOContext differentiates from other interfaces.
	IsOContext()
}

type OContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOContext() *OContext {
	var p = new(OContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_o
	return p
}

func (*OContext) IsOContext() {}

func NewOContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OContext {
	var p = new(OContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_o

	return p
}

func (s *OContext) GetParser() antlr.Parser { return s.parser }

func (s *OContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *OContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *OContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterO(s)
	}
}

func (s *OContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitO(s)
	}
}

func (p *morsecodeParser) O() (localctx IOContext) {
	localctx = NewOContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, morsecodeParserRULE_o)

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
		p.SetState(177)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(178)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(179)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IPContext is an interface to support dynamic dispatch.
type IPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPContext differentiates from other interfaces.
	IsPContext()
}

type PContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPContext() *PContext {
	var p = new(PContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_p
	return p
}

func (*PContext) IsPContext() {}

func NewPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PContext {
	var p = new(PContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_p

	return p
}

func (s *PContext) GetParser() antlr.Parser { return s.parser }

func (s *PContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *PContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *PContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *PContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterP(s)
	}
}

func (s *PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitP(s)
	}
}

func (p *morsecodeParser) P() (localctx IPContext) {
	localctx = NewPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, morsecodeParserRULE_p)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(182)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(183)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(184)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IQContext is an interface to support dynamic dispatch.
type IQContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQContext differentiates from other interfaces.
	IsQContext()
}

type QContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQContext() *QContext {
	var p = new(QContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_q
	return p
}

func (*QContext) IsQContext() {}

func NewQContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QContext {
	var p = new(QContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_q

	return p
}

func (s *QContext) GetParser() antlr.Parser { return s.parser }

func (s *QContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *QContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *QContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *QContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterQ(s)
	}
}

func (s *QContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitQ(s)
	}
}

func (p *morsecodeParser) Q() (localctx IQContext) {
	localctx = NewQContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, morsecodeParserRULE_q)

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
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(187)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(188)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(189)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IRContext is an interface to support dynamic dispatch.
type IRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRContext differentiates from other interfaces.
	IsRContext()
}

type RContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRContext() *RContext {
	var p = new(RContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_r
	return p
}

func (*RContext) IsRContext() {}

func NewRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RContext {
	var p = new(RContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_r

	return p
}

func (s *RContext) GetParser() antlr.Parser { return s.parser }

func (s *RContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *RContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *RContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterR(s)
	}
}

func (s *RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitR(s)
	}
}

func (p *morsecodeParser) R() (localctx IRContext) {
	localctx = NewRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, morsecodeParserRULE_r)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(192)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(193)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *SContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *morsecodeParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, morsecodeParserRULE_s)

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
		p.SetState(195)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(196)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(197)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterT(s)
	}
}

func (s *TContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitT(s)
	}
}

func (p *morsecodeParser) T() (localctx ITContext) {
	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, morsecodeParserRULE_t)

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
		p.SetState(199)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IUContext is an interface to support dynamic dispatch.
type IUContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUContext differentiates from other interfaces.
	IsUContext()
}

type UContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUContext() *UContext {
	var p = new(UContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_u
	return p
}

func (*UContext) IsUContext() {}

func NewUContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UContext {
	var p = new(UContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_u

	return p
}

func (s *UContext) GetParser() antlr.Parser { return s.parser }

func (s *UContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *UContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *UContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *UContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterU(s)
	}
}

func (s *UContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitU(s)
	}
}

func (p *morsecodeParser) U() (localctx IUContext) {
	localctx = NewUContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, morsecodeParserRULE_u)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(202)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(203)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IVContext is an interface to support dynamic dispatch.
type IVContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVContext differentiates from other interfaces.
	IsVContext()
}

type VContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVContext() *VContext {
	var p = new(VContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_v
	return p
}

func (*VContext) IsVContext() {}

func NewVContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VContext {
	var p = new(VContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_v

	return p
}

func (s *VContext) GetParser() antlr.Parser { return s.parser }

func (s *VContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *VContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *VContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *VContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterV(s)
	}
}

func (s *VContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitV(s)
	}
}

func (p *morsecodeParser) V() (localctx IVContext) {
	localctx = NewVContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, morsecodeParserRULE_v)

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
		p.SetState(205)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(206)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(207)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(208)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IWContext is an interface to support dynamic dispatch.
type IWContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWContext differentiates from other interfaces.
	IsWContext()
}

type WContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWContext() *WContext {
	var p = new(WContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_w
	return p
}

func (*WContext) IsWContext() {}

func NewWContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WContext {
	var p = new(WContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_w

	return p
}

func (s *WContext) GetParser() antlr.Parser { return s.parser }

func (s *WContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *WContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *WContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *WContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterW(s)
	}
}

func (s *WContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitW(s)
	}
}

func (p *morsecodeParser) W() (localctx IWContext) {
	localctx = NewWContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, morsecodeParserRULE_w)

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
		p.SetState(210)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(211)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(212)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IXContext is an interface to support dynamic dispatch.
type IXContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXContext differentiates from other interfaces.
	IsXContext()
}

type XContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXContext() *XContext {
	var p = new(XContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_x
	return p
}

func (*XContext) IsXContext() {}

func NewXContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XContext {
	var p = new(XContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_x

	return p
}

func (s *XContext) GetParser() antlr.Parser { return s.parser }

func (s *XContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *XContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *XContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *XContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *XContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterX(s)
	}
}

func (s *XContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitX(s)
	}
}

func (p *morsecodeParser) X() (localctx IXContext) {
	localctx = NewXContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, morsecodeParserRULE_x)

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
		p.SetState(214)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(215)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(216)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(217)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IYContext is an interface to support dynamic dispatch.
type IYContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYContext differentiates from other interfaces.
	IsYContext()
}

type YContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYContext() *YContext {
	var p = new(YContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_y
	return p
}

func (*YContext) IsYContext() {}

func NewYContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YContext {
	var p = new(YContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_y

	return p
}

func (s *YContext) GetParser() antlr.Parser { return s.parser }

func (s *YContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *YContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *YContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *YContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterY(s)
	}
}

func (s *YContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitY(s)
	}
}

func (p *morsecodeParser) Y() (localctx IYContext) {
	localctx = NewYContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, morsecodeParserRULE_y)

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
		p.SetState(219)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(220)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(221)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(222)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IZContext is an interface to support dynamic dispatch.
type IZContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZContext differentiates from other interfaces.
	IsZContext()
}

type ZContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZContext() *ZContext {
	var p = new(ZContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_z
	return p
}

func (*ZContext) IsZContext() {}

func NewZContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZContext {
	var p = new(ZContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_z

	return p
}

func (s *ZContext) GetParser() antlr.Parser { return s.parser }

func (s *ZContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *ZContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *ZContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *ZContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *ZContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterZ(s)
	}
}

func (s *ZContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitZ(s)
	}
}

func (p *morsecodeParser) Z() (localctx IZContext) {
	localctx = NewZContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, morsecodeParserRULE_z)

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
		p.SetState(224)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(225)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(226)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(227)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IOneContext is an interface to support dynamic dispatch.
type IOneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneContext differentiates from other interfaces.
	IsOneContext()
}

type OneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneContext() *OneContext {
	var p = new(OneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_one
	return p
}

func (*OneContext) IsOneContext() {}

func NewOneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneContext {
	var p = new(OneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_one

	return p
}

func (s *OneContext) GetParser() antlr.Parser { return s.parser }

func (s *OneContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *OneContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *OneContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *OneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterOne(s)
	}
}

func (s *OneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitOne(s)
	}
}

func (p *morsecodeParser) One() (localctx IOneContext) {
	localctx = NewOneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, morsecodeParserRULE_one)

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
		p.SetState(229)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(230)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(231)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(232)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(233)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// ITwoContext is an interface to support dynamic dispatch.
type ITwoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTwoContext differentiates from other interfaces.
	IsTwoContext()
}

type TwoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTwoContext() *TwoContext {
	var p = new(TwoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_two
	return p
}

func (*TwoContext) IsTwoContext() {}

func NewTwoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TwoContext {
	var p = new(TwoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_two

	return p
}

func (s *TwoContext) GetParser() antlr.Parser { return s.parser }

func (s *TwoContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *TwoContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *TwoContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *TwoContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *TwoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TwoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TwoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterTwo(s)
	}
}

func (s *TwoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitTwo(s)
	}
}

func (p *morsecodeParser) Two() (localctx ITwoContext) {
	localctx = NewTwoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, morsecodeParserRULE_two)

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
		p.SetState(235)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(236)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(237)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(238)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(239)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IThreeContext is an interface to support dynamic dispatch.
type IThreeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThreeContext differentiates from other interfaces.
	IsThreeContext()
}

type ThreeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThreeContext() *ThreeContext {
	var p = new(ThreeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_three
	return p
}

func (*ThreeContext) IsThreeContext() {}

func NewThreeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThreeContext {
	var p = new(ThreeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_three

	return p
}

func (s *ThreeContext) GetParser() antlr.Parser { return s.parser }

func (s *ThreeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *ThreeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *ThreeContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *ThreeContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *ThreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThreeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterThree(s)
	}
}

func (s *ThreeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitThree(s)
	}
}

func (p *morsecodeParser) Three() (localctx IThreeContext) {
	localctx = NewThreeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, morsecodeParserRULE_three)

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
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(242)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(243)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(244)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(245)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IFourContext is an interface to support dynamic dispatch.
type IFourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFourContext differentiates from other interfaces.
	IsFourContext()
}

type FourContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFourContext() *FourContext {
	var p = new(FourContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_four
	return p
}

func (*FourContext) IsFourContext() {}

func NewFourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FourContext {
	var p = new(FourContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_four

	return p
}

func (s *FourContext) GetParser() antlr.Parser { return s.parser }

func (s *FourContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *FourContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *FourContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *FourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterFour(s)
	}
}

func (s *FourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitFour(s)
	}
}

func (p *morsecodeParser) Four() (localctx IFourContext) {
	localctx = NewFourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, morsecodeParserRULE_four)

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
		p.SetState(247)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(248)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(249)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(250)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(251)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}

// IFiveContext is an interface to support dynamic dispatch.
type IFiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiveContext differentiates from other interfaces.
	IsFiveContext()
}

type FiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiveContext() *FiveContext {
	var p = new(FiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_five
	return p
}

func (*FiveContext) IsFiveContext() {}

func NewFiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiveContext {
	var p = new(FiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_five

	return p
}

func (s *FiveContext) GetParser() antlr.Parser { return s.parser }

func (s *FiveContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *FiveContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *FiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterFive(s)
	}
}

func (s *FiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitFive(s)
	}
}

func (p *morsecodeParser) Five() (localctx IFiveContext) {
	localctx = NewFiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, morsecodeParserRULE_five)

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
		p.SetState(253)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(254)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(255)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(256)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(257)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// ISixContext is an interface to support dynamic dispatch.
type ISixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSixContext differentiates from other interfaces.
	IsSixContext()
}

type SixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySixContext() *SixContext {
	var p = new(SixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_six
	return p
}

func (*SixContext) IsSixContext() {}

func NewSixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SixContext {
	var p = new(SixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_six

	return p
}

func (s *SixContext) GetParser() antlr.Parser { return s.parser }

func (s *SixContext) DASH() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, 0)
}

func (s *SixContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *SixContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *SixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterSix(s)
	}
}

func (s *SixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitSix(s)
	}
}

func (p *morsecodeParser) Six() (localctx ISixContext) {
	localctx = NewSixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, morsecodeParserRULE_six)

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
		p.SetState(259)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(260)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(261)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(262)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(263)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// ISevenContext is an interface to support dynamic dispatch.
type ISevenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSevenContext differentiates from other interfaces.
	IsSevenContext()
}

type SevenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySevenContext() *SevenContext {
	var p = new(SevenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_seven
	return p
}

func (*SevenContext) IsSevenContext() {}

func NewSevenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SevenContext {
	var p = new(SevenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_seven

	return p
}

func (s *SevenContext) GetParser() antlr.Parser { return s.parser }

func (s *SevenContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *SevenContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *SevenContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *SevenContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *SevenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SevenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SevenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterSeven(s)
	}
}

func (s *SevenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitSeven(s)
	}
}

func (p *morsecodeParser) Seven() (localctx ISevenContext) {
	localctx = NewSevenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, morsecodeParserRULE_seven)

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
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(266)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(267)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(268)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(269)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IEightContext is an interface to support dynamic dispatch.
type IEightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEightContext differentiates from other interfaces.
	IsEightContext()
}

type EightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEightContext() *EightContext {
	var p = new(EightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_eight
	return p
}

func (*EightContext) IsEightContext() {}

func NewEightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EightContext {
	var p = new(EightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_eight

	return p
}

func (s *EightContext) GetParser() antlr.Parser { return s.parser }

func (s *EightContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *EightContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *EightContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDOT)
}

func (s *EightContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, i)
}

func (s *EightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterEight(s)
	}
}

func (s *EightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitEight(s)
	}
}

func (p *morsecodeParser) Eight() (localctx IEightContext) {
	localctx = NewEightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, morsecodeParserRULE_eight)

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
		p.SetState(271)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(272)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(273)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(274)
		p.Match(morsecodeParserDOT)
	}
	{
		p.SetState(275)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// INineContext is an interface to support dynamic dispatch.
type INineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNineContext differentiates from other interfaces.
	IsNineContext()
}

type NineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNineContext() *NineContext {
	var p = new(NineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_nine
	return p
}

func (*NineContext) IsNineContext() {}

func NewNineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NineContext {
	var p = new(NineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_nine

	return p
}

func (s *NineContext) GetParser() antlr.Parser { return s.parser }

func (s *NineContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *NineContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *NineContext) DOT() antlr.TerminalNode {
	return s.GetToken(morsecodeParserDOT, 0)
}

func (s *NineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterNine(s)
	}
}

func (s *NineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitNine(s)
	}
}

func (p *morsecodeParser) Nine() (localctx INineContext) {
	localctx = NewNineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, morsecodeParserRULE_nine)

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
		p.SetState(277)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(278)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(279)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(280)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(281)
		p.Match(morsecodeParserDOT)
	}

	return localctx
}

// IZeroContext is an interface to support dynamic dispatch.
type IZeroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZeroContext differentiates from other interfaces.
	IsZeroContext()
}

type ZeroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZeroContext() *ZeroContext {
	var p = new(ZeroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = morsecodeParserRULE_zero
	return p
}

func (*ZeroContext) IsZeroContext() {}

func NewZeroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZeroContext {
	var p = new(ZeroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = morsecodeParserRULE_zero

	return p
}

func (s *ZeroContext) GetParser() antlr.Parser { return s.parser }

func (s *ZeroContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(morsecodeParserDASH)
}

func (s *ZeroContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(morsecodeParserDASH, i)
}

func (s *ZeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZeroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.EnterZero(s)
	}
}

func (s *ZeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(morsecodeListener); ok {
		listenerT.ExitZero(s)
	}
}

func (p *morsecodeParser) Zero() (localctx IZeroContext) {
	localctx = NewZeroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, morsecodeParserRULE_zero)

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
		p.SetState(283)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(284)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(285)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(286)
		p.Match(morsecodeParserDASH)
	}
	{
		p.SetState(287)
		p.Match(morsecodeParserDASH)
	}

	return localctx
}
