// Code generated from b.g4 by ANTLR 4.9.3. DO NOT EDIT.

package b // b
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 282,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 7, 2, 62, 10, 2, 12, 2, 14, 2, 65, 11, 2,
	3, 3, 3, 3, 5, 3, 69, 10, 3, 3, 3, 3, 3, 3, 3, 7, 3, 74, 10, 3, 12, 3,
	14, 3, 77, 11, 3, 7, 3, 79, 10, 3, 12, 3, 14, 3, 82, 11, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 91, 10, 3, 12, 3, 14, 3, 94, 11, 3,
	5, 3, 96, 10, 3, 3, 3, 3, 3, 3, 3, 5, 3, 101, 10, 3, 3, 4, 3, 4, 5, 4,
	105, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 122, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 7, 8, 131, 10, 8, 12, 8, 14, 8, 134, 11, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 143, 10, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 168,
	10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 179, 10, 15, 12, 15, 14, 15, 182, 11, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 5, 16, 189, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 194, 10, 16,
	7, 16, 196, 10, 16, 12, 16, 14, 16, 199, 11, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 5, 17, 207, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 241, 10, 21, 3, 22, 3,
	22, 3, 22, 5, 22, 246, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23,
	253, 10, 23, 12, 23, 14, 23, 256, 11, 23, 3, 24, 3, 24, 5, 24, 260, 10,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 276, 10, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 2, 2, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 2, 6, 3, 2,
	22, 23, 3, 2, 24, 25, 5, 2, 20, 20, 24, 24, 26, 38, 3, 2, 42, 44, 2, 293,
	2, 63, 3, 2, 2, 2, 4, 100, 3, 2, 2, 2, 6, 104, 3, 2, 2, 2, 8, 121, 3, 2,
	2, 2, 10, 123, 3, 2, 2, 2, 12, 125, 3, 2, 2, 2, 14, 128, 3, 2, 2, 2, 16,
	137, 3, 2, 2, 2, 18, 146, 3, 2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 154, 3,
	2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 169, 3, 2, 2, 2, 28, 174, 3, 2, 2, 2,
	30, 185, 3, 2, 2, 2, 32, 206, 3, 2, 2, 2, 34, 208, 3, 2, 2, 2, 36, 214,
	3, 2, 2, 2, 38, 218, 3, 2, 2, 2, 40, 240, 3, 2, 2, 2, 42, 242, 3, 2, 2,
	2, 44, 249, 3, 2, 2, 2, 46, 257, 3, 2, 2, 2, 48, 261, 3, 2, 2, 2, 50, 263,
	3, 2, 2, 2, 52, 265, 3, 2, 2, 2, 54, 275, 3, 2, 2, 2, 56, 277, 3, 2, 2,
	2, 58, 279, 3, 2, 2, 2, 60, 62, 5, 4, 3, 2, 61, 60, 3, 2, 2, 2, 62, 65,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 3, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 66, 68, 5, 58, 30, 2, 67, 69, 5, 56, 29, 2, 68, 67,
	3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 80, 3, 2, 2, 2, 70, 75, 5, 6, 4, 2,
	71, 72, 7, 3, 2, 2, 72, 74, 5, 6, 4, 2, 73, 71, 3, 2, 2, 2, 74, 77, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 78, 70, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2,
	2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84,
	7, 4, 2, 2, 84, 101, 3, 2, 2, 2, 85, 86, 5, 58, 30, 2, 86, 95, 7, 5, 2,
	2, 87, 92, 5, 58, 30, 2, 88, 89, 7, 3, 2, 2, 89, 91, 5, 58, 30, 2, 90,
	88, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 87, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 7, 6, 2, 2, 98, 99, 5, 8, 5, 2,
	99, 101, 3, 2, 2, 2, 100, 66, 3, 2, 2, 2, 100, 85, 3, 2, 2, 2, 101, 5,
	3, 2, 2, 2, 102, 105, 5, 56, 29, 2, 103, 105, 5, 58, 30, 2, 104, 102, 3,
	2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 7, 3, 2, 2, 2, 106, 122, 5, 28, 15,
	2, 107, 122, 5, 30, 16, 2, 108, 109, 5, 58, 30, 2, 109, 110, 7, 7, 2, 2,
	110, 111, 5, 8, 5, 2, 111, 122, 3, 2, 2, 2, 112, 122, 5, 26, 14, 2, 113,
	122, 5, 14, 8, 2, 114, 122, 5, 24, 13, 2, 115, 122, 5, 22, 12, 2, 116,
	122, 5, 20, 11, 2, 117, 122, 5, 18, 10, 2, 118, 122, 5, 16, 9, 2, 119,
	122, 5, 12, 7, 2, 120, 122, 5, 10, 6, 2, 121, 106, 3, 2, 2, 2, 121, 107,
	3, 2, 2, 2, 121, 108, 3, 2, 2, 2, 121, 112, 3, 2, 2, 2, 121, 113, 3, 2,
	2, 2, 121, 114, 3, 2, 2, 2, 121, 115, 3, 2, 2, 2, 121, 116, 3, 2, 2, 2,
	121, 117, 3, 2, 2, 2, 121, 118, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121,
	120, 3, 2, 2, 2, 122, 9, 3, 2, 2, 2, 123, 124, 7, 4, 2, 2, 124, 11, 3,
	2, 2, 2, 125, 126, 5, 32, 17, 2, 126, 127, 7, 4, 2, 2, 127, 13, 3, 2, 2,
	2, 128, 132, 7, 8, 2, 2, 129, 131, 5, 8, 5, 2, 130, 129, 3, 2, 2, 2, 131,
	134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135,
	3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 136, 7, 9, 2, 2, 136, 15, 3, 2,
	2, 2, 137, 142, 7, 10, 2, 2, 138, 139, 7, 5, 2, 2, 139, 140, 5, 32, 17,
	2, 140, 141, 7, 6, 2, 2, 141, 143, 3, 2, 2, 2, 142, 138, 3, 2, 2, 2, 142,
	143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 4, 2, 2, 145, 17, 3,
	2, 2, 2, 146, 147, 7, 11, 2, 2, 147, 148, 5, 32, 17, 2, 148, 149, 7, 4,
	2, 2, 149, 19, 3, 2, 2, 2, 150, 151, 7, 12, 2, 2, 151, 152, 5, 32, 17,
	2, 152, 153, 5, 8, 5, 2, 153, 21, 3, 2, 2, 2, 154, 155, 7, 13, 2, 2, 155,
	156, 7, 5, 2, 2, 156, 157, 5, 32, 17, 2, 157, 158, 7, 6, 2, 2, 158, 159,
	5, 8, 5, 2, 159, 23, 3, 2, 2, 2, 160, 161, 7, 14, 2, 2, 161, 162, 7, 5,
	2, 2, 162, 163, 5, 32, 17, 2, 163, 164, 7, 6, 2, 2, 164, 167, 5, 8, 5,
	2, 165, 166, 7, 15, 2, 2, 166, 168, 5, 8, 5, 2, 167, 165, 3, 2, 2, 2, 167,
	168, 3, 2, 2, 2, 168, 25, 3, 2, 2, 2, 169, 170, 7, 16, 2, 2, 170, 171,
	5, 56, 29, 2, 171, 172, 7, 7, 2, 2, 172, 173, 5, 8, 5, 2, 173, 27, 3, 2,
	2, 2, 174, 175, 7, 17, 2, 2, 175, 180, 5, 58, 30, 2, 176, 177, 7, 3, 2,
	2, 177, 179, 5, 58, 30, 2, 178, 176, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2,
	180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182,
	180, 3, 2, 2, 2, 183, 184, 7, 4, 2, 2, 184, 29, 3, 2, 2, 2, 185, 186, 7,
	18, 2, 2, 186, 188, 5, 58, 30, 2, 187, 189, 5, 56, 29, 2, 188, 187, 3,
	2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 197, 3, 2, 2, 2, 190, 191, 7, 3, 2,
	2, 191, 193, 5, 58, 30, 2, 192, 194, 5, 56, 29, 2, 193, 192, 3, 2, 2, 2,
	193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 190, 3, 2, 2, 2, 196,
	199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 200,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 201, 7, 4, 2, 2, 201, 31, 3, 2,
	2, 2, 202, 207, 5, 40, 21, 2, 203, 207, 5, 36, 19, 2, 204, 207, 5, 34,
	18, 2, 205, 207, 5, 38, 20, 2, 206, 202, 3, 2, 2, 2, 206, 203, 3, 2, 2,
	2, 206, 204, 3, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 33, 3, 2, 2, 2, 208,
	209, 5, 40, 21, 2, 209, 210, 7, 19, 2, 2, 210, 211, 5, 32, 17, 2, 211,
	212, 7, 7, 2, 2, 212, 213, 5, 32, 17, 2, 213, 35, 3, 2, 2, 2, 214, 215,
	5, 40, 21, 2, 215, 216, 5, 52, 27, 2, 216, 217, 5, 32, 17, 2, 217, 37,
	3, 2, 2, 2, 218, 219, 5, 58, 30, 2, 219, 220, 5, 46, 24, 2, 220, 221, 5,
	32, 17, 2, 221, 39, 3, 2, 2, 2, 222, 223, 7, 5, 2, 2, 223, 224, 5, 32,
	17, 2, 224, 225, 7, 6, 2, 2, 225, 241, 3, 2, 2, 2, 226, 241, 5, 58, 30,
	2, 227, 241, 5, 56, 29, 2, 228, 229, 5, 48, 25, 2, 229, 230, 5, 58, 30,
	2, 230, 241, 3, 2, 2, 2, 231, 232, 5, 58, 30, 2, 232, 233, 5, 48, 25, 2,
	233, 241, 3, 2, 2, 2, 234, 235, 5, 50, 26, 2, 235, 236, 5, 32, 17, 2, 236,
	241, 3, 2, 2, 2, 237, 238, 7, 20, 2, 2, 238, 241, 5, 58, 30, 2, 239, 241,
	5, 42, 22, 2, 240, 222, 3, 2, 2, 2, 240, 226, 3, 2, 2, 2, 240, 227, 3,
	2, 2, 2, 240, 228, 3, 2, 2, 2, 240, 231, 3, 2, 2, 2, 240, 234, 3, 2, 2,
	2, 240, 237, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 41, 3, 2, 2, 2, 242,
	243, 5, 58, 30, 2, 243, 245, 7, 5, 2, 2, 244, 246, 5, 44, 23, 2, 245, 244,
	3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 7, 6,
	2, 2, 248, 43, 3, 2, 2, 2, 249, 254, 5, 32, 17, 2, 250, 251, 7, 3, 2, 2,
	251, 253, 5, 32, 17, 2, 252, 250, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254,
	252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 45, 3, 2, 2, 2, 256, 254, 3,
	2, 2, 2, 257, 259, 7, 21, 2, 2, 258, 260, 5, 52, 27, 2, 259, 258, 3, 2,
	2, 2, 259, 260, 3, 2, 2, 2, 260, 47, 3, 2, 2, 2, 261, 262, 9, 2, 2, 2,
	262, 49, 3, 2, 2, 2, 263, 264, 9, 3, 2, 2, 264, 51, 3, 2, 2, 2, 265, 266,
	9, 4, 2, 2, 266, 53, 3, 2, 2, 2, 267, 276, 5, 58, 30, 2, 268, 269, 7, 37,
	2, 2, 269, 276, 5, 32, 17, 2, 270, 271, 5, 32, 17, 2, 271, 272, 7, 39,
	2, 2, 272, 273, 5, 32, 17, 2, 273, 274, 7, 40, 2, 2, 274, 276, 3, 2, 2,
	2, 275, 267, 3, 2, 2, 2, 275, 268, 3, 2, 2, 2, 275, 270, 3, 2, 2, 2, 276,
	55, 3, 2, 2, 2, 277, 278, 9, 5, 2, 2, 278, 57, 3, 2, 2, 2, 279, 280, 7,
	41, 2, 2, 280, 59, 3, 2, 2, 2, 24, 63, 68, 75, 80, 92, 95, 100, 104, 121,
	132, 142, 167, 180, 188, 193, 197, 206, 240, 245, 254, 259, 275,
}
var literalNames = []string{
	"", "','", "';'", "'('", "')'", "':'", "'{'", "'}'", "'return'", "'goto'",
	"'switch'", "'while'", "'if'", "'else'", "'case'", "'extrn'", "'auto'",
	"'?'", "'&'", "'='", "'++'", "'--'", "'-'", "'!'", "'|'", "'=='", "'!='",
	"'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'%'", "'*'", "'/'",
	"'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "NAME", "INT", "STRING1", "STRING2", "BLOCKCOMMENT", "WS",
}

var ruleNames = []string{
	"program", "definition", "ival", "statement", "nullstmt", "expressionstmt",
	"blockstmt", "returnstmt", "gotostmt", "switchstmt", "whilestmt", "ifstmt",
	"casestmt", "externsmt", "autosmt", "rvalue", "ternary", "comparison",
	"assignment", "expression", "functioninvocation", "functionparameters",
	"assign", "incdec", "unary", "binary", "lvalue", "constant", "name",
}

type bParser struct {
	*antlr.BaseParser
}

// NewbParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *bParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbParser(input antlr.TokenStream) *bParser {
	this := new(bParser)
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
	this.GrammarFileName = "b.g4"

	return this
}

// bParser tokens.
const (
	bParserEOF          = antlr.TokenEOF
	bParserT__0         = 1
	bParserT__1         = 2
	bParserT__2         = 3
	bParserT__3         = 4
	bParserT__4         = 5
	bParserT__5         = 6
	bParserT__6         = 7
	bParserT__7         = 8
	bParserT__8         = 9
	bParserT__9         = 10
	bParserT__10        = 11
	bParserT__11        = 12
	bParserT__12        = 13
	bParserT__13        = 14
	bParserT__14        = 15
	bParserT__15        = 16
	bParserT__16        = 17
	bParserT__17        = 18
	bParserT__18        = 19
	bParserT__19        = 20
	bParserT__20        = 21
	bParserT__21        = 22
	bParserT__22        = 23
	bParserT__23        = 24
	bParserT__24        = 25
	bParserT__25        = 26
	bParserT__26        = 27
	bParserT__27        = 28
	bParserT__28        = 29
	bParserT__29        = 30
	bParserT__30        = 31
	bParserT__31        = 32
	bParserT__32        = 33
	bParserT__33        = 34
	bParserT__34        = 35
	bParserT__35        = 36
	bParserT__36        = 37
	bParserT__37        = 38
	bParserNAME         = 39
	bParserINT          = 40
	bParserSTRING1      = 41
	bParserSTRING2      = 42
	bParserBLOCKCOMMENT = 43
	bParserWS           = 44
)

// bParser rules.
const (
	bParserRULE_program            = 0
	bParserRULE_definition         = 1
	bParserRULE_ival               = 2
	bParserRULE_statement          = 3
	bParserRULE_nullstmt           = 4
	bParserRULE_expressionstmt     = 5
	bParserRULE_blockstmt          = 6
	bParserRULE_returnstmt         = 7
	bParserRULE_gotostmt           = 8
	bParserRULE_switchstmt         = 9
	bParserRULE_whilestmt          = 10
	bParserRULE_ifstmt             = 11
	bParserRULE_casestmt           = 12
	bParserRULE_externsmt          = 13
	bParserRULE_autosmt            = 14
	bParserRULE_rvalue             = 15
	bParserRULE_ternary            = 16
	bParserRULE_comparison         = 17
	bParserRULE_assignment         = 18
	bParserRULE_expression         = 19
	bParserRULE_functioninvocation = 20
	bParserRULE_functionparameters = 21
	bParserRULE_assign             = 22
	bParserRULE_incdec             = 23
	bParserRULE_unary              = 24
	bParserRULE_binary             = 25
	bParserRULE_lvalue             = 26
	bParserRULE_constant           = 27
	bParserRULE_name               = 28
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
	p.RuleIndex = bParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *bParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bParserRULE_program)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bParserNAME {
		{
			p.SetState(58)
			p.Definition()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = bParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_definition

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

func (s *DefinitionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DefinitionContext) AllIval() []IIvalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIvalContext)(nil)).Elem())
	var tst = make([]IIvalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIvalContext)
		}
	}

	return tst
}

func (s *DefinitionContext) Ival(i int) IIvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIvalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIvalContext)
}

func (s *DefinitionContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *bParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bParserRULE_definition)
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Name()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(65)
				p.Constant()
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(bParserNAME-39))|(1<<(bParserINT-39))|(1<<(bParserSTRING1-39))|(1<<(bParserSTRING2-39)))) != 0 {
			{
				p.SetState(68)
				p.Ival()
			}
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == bParserT__0 {
				{
					p.SetState(69)
					p.Match(bParserT__0)
				}
				{
					p.SetState(70)
					p.Ival()
				}

				p.SetState(75)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)
			p.Match(bParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Name()
		}
		{
			p.SetState(84)
			p.Match(bParserT__2)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == bParserNAME {
			{
				p.SetState(85)
				p.Name()
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == bParserT__0 {
				{
					p.SetState(86)
					p.Match(bParserT__0)
				}
				{
					p.SetState(87)
					p.Name()
				}

				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(95)
			p.Match(bParserT__3)
		}
		{
			p.SetState(96)
			p.Statement()
		}

	}

	return localctx
}

// IIvalContext is an interface to support dynamic dispatch.
type IIvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIvalContext differentiates from other interfaces.
	IsIvalContext()
}

type IvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIvalContext() *IvalContext {
	var p = new(IvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_ival
	return p
}

func (*IvalContext) IsIvalContext() {}

func NewIvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IvalContext {
	var p = new(IvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ival

	return p
}

func (s *IvalContext) GetParser() antlr.Parser { return s.parser }

func (s *IvalContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *IvalContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *IvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIval(s)
	}
}

func (s *IvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIval(s)
	}
}

func (p *bParser) Ival() (localctx IIvalContext) {
	this := p
	_ = this

	localctx = NewIvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bParserRULE_ival)

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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bParserINT, bParserSTRING1, bParserSTRING2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Constant()
		}

	case bParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = bParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Externsmt() IExternsmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternsmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternsmtContext)
}

func (s *StatementContext) Autosmt() IAutosmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAutosmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAutosmtContext)
}

func (s *StatementContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) Casestmt() ICasestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasestmtContext)
}

func (s *StatementContext) Blockstmt() IBlockstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockstmtContext)
}

func (s *StatementContext) Ifstmt() IIfstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StatementContext) Whilestmt() IWhilestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhilestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StatementContext) Switchstmt() ISwitchstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *StatementContext) Gotostmt() IGotostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotostmtContext)
}

func (s *StatementContext) Returnstmt() IReturnstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnstmtContext)
}

func (s *StatementContext) Expressionstmt() IExpressionstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionstmtContext)
}

func (s *StatementContext) Nullstmt() INullstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullstmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *bParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bParserRULE_statement)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Externsmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Autosmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Name()
		}
		{
			p.SetState(107)
			p.Match(bParserT__4)
		}
		{
			p.SetState(108)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Casestmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.Blockstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(112)
			p.Ifstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(113)
			p.Whilestmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(114)
			p.Switchstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(115)
			p.Gotostmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(116)
			p.Returnstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(117)
			p.Expressionstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(118)
			p.Nullstmt()
		}

	}

	return localctx
}

// INullstmtContext is an interface to support dynamic dispatch.
type INullstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullstmtContext differentiates from other interfaces.
	IsNullstmtContext()
}

type NullstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullstmtContext() *NullstmtContext {
	var p = new(NullstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_nullstmt
	return p
}

func (*NullstmtContext) IsNullstmtContext() {}

func NewNullstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullstmtContext {
	var p = new(NullstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_nullstmt

	return p
}

func (s *NullstmtContext) GetParser() antlr.Parser { return s.parser }
func (s *NullstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterNullstmt(s)
	}
}

func (s *NullstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitNullstmt(s)
	}
}

func (p *bParser) Nullstmt() (localctx INullstmtContext) {
	this := p
	_ = this

	localctx = NewNullstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bParserRULE_nullstmt)

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
		p.Match(bParserT__1)
	}

	return localctx
}

// IExpressionstmtContext is an interface to support dynamic dispatch.
type IExpressionstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionstmtContext differentiates from other interfaces.
	IsExpressionstmtContext()
}

type ExpressionstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionstmtContext() *ExpressionstmtContext {
	var p = new(ExpressionstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_expressionstmt
	return p
}

func (*ExpressionstmtContext) IsExpressionstmtContext() {}

func NewExpressionstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionstmtContext {
	var p = new(ExpressionstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expressionstmt

	return p
}

func (s *ExpressionstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionstmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ExpressionstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpressionstmt(s)
	}
}

func (s *ExpressionstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpressionstmt(s)
	}
}

func (p *bParser) Expressionstmt() (localctx IExpressionstmtContext) {
	this := p
	_ = this

	localctx = NewExpressionstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bParserRULE_expressionstmt)

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
		p.Rvalue()
	}
	{
		p.SetState(124)
		p.Match(bParserT__1)
	}

	return localctx
}

// IBlockstmtContext is an interface to support dynamic dispatch.
type IBlockstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockstmtContext differentiates from other interfaces.
	IsBlockstmtContext()
}

type BlockstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockstmtContext() *BlockstmtContext {
	var p = new(BlockstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_blockstmt
	return p
}

func (*BlockstmtContext) IsBlockstmtContext() {}

func NewBlockstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockstmtContext {
	var p = new(BlockstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_blockstmt

	return p
}

func (s *BlockstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterBlockstmt(s)
	}
}

func (s *BlockstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitBlockstmt(s)
	}
}

func (p *bParser) Blockstmt() (localctx IBlockstmtContext) {
	this := p
	_ = this

	localctx = NewBlockstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bParserRULE_blockstmt)
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
		p.SetState(126)
		p.Match(bParserT__5)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<bParserT__1)|(1<<bParserT__2)|(1<<bParserT__5)|(1<<bParserT__7)|(1<<bParserT__8)|(1<<bParserT__9)|(1<<bParserT__10)|(1<<bParserT__11)|(1<<bParserT__13)|(1<<bParserT__14)|(1<<bParserT__15)|(1<<bParserT__17)|(1<<bParserT__19)|(1<<bParserT__20)|(1<<bParserT__21)|(1<<bParserT__22))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(bParserNAME-39))|(1<<(bParserINT-39))|(1<<(bParserSTRING1-39))|(1<<(bParserSTRING2-39)))) != 0) {
		{
			p.SetState(127)
			p.Statement()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(bParserT__6)
	}

	return localctx
}

// IReturnstmtContext is an interface to support dynamic dispatch.
type IReturnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnstmtContext differentiates from other interfaces.
	IsReturnstmtContext()
}

type ReturnstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstmtContext() *ReturnstmtContext {
	var p = new(ReturnstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_returnstmt
	return p
}

func (*ReturnstmtContext) IsReturnstmtContext() {}

func NewReturnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstmtContext {
	var p = new(ReturnstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_returnstmt

	return p
}

func (s *ReturnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ReturnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterReturnstmt(s)
	}
}

func (s *ReturnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitReturnstmt(s)
	}
}

func (p *bParser) Returnstmt() (localctx IReturnstmtContext) {
	this := p
	_ = this

	localctx = NewReturnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bParserRULE_returnstmt)
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
		p.SetState(135)
		p.Match(bParserT__7)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == bParserT__2 {
		{
			p.SetState(136)
			p.Match(bParserT__2)
		}
		{
			p.SetState(137)
			p.Rvalue()
		}
		{
			p.SetState(138)
			p.Match(bParserT__3)
		}

	}
	{
		p.SetState(142)
		p.Match(bParserT__1)
	}

	return localctx
}

// IGotostmtContext is an interface to support dynamic dispatch.
type IGotostmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotostmtContext differentiates from other interfaces.
	IsGotostmtContext()
}

type GotostmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotostmtContext() *GotostmtContext {
	var p = new(GotostmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_gotostmt
	return p
}

func (*GotostmtContext) IsGotostmtContext() {}

func NewGotostmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotostmtContext {
	var p = new(GotostmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_gotostmt

	return p
}

func (s *GotostmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GotostmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *GotostmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotostmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotostmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterGotostmt(s)
	}
}

func (s *GotostmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitGotostmt(s)
	}
}

func (p *bParser) Gotostmt() (localctx IGotostmtContext) {
	this := p
	_ = this

	localctx = NewGotostmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bParserRULE_gotostmt)

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
		p.SetState(144)
		p.Match(bParserT__8)
	}
	{
		p.SetState(145)
		p.Rvalue()
	}
	{
		p.SetState(146)
		p.Match(bParserT__1)
	}

	return localctx
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_switchstmt
	return p
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *SwitchstmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *bParser) Switchstmt() (localctx ISwitchstmtContext) {
	this := p
	_ = this

	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bParserRULE_switchstmt)

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
		p.Match(bParserT__9)
	}
	{
		p.SetState(149)
		p.Rvalue()
	}
	{
		p.SetState(150)
		p.Statement()
	}

	return localctx
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_whilestmt
	return p
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *WhilestmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *bParser) Whilestmt() (localctx IWhilestmtContext) {
	this := p
	_ = this

	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bParserRULE_whilestmt)

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
		p.SetState(152)
		p.Match(bParserT__10)
	}
	{
		p.SetState(153)
		p.Match(bParserT__2)
	}
	{
		p.SetState(154)
		p.Rvalue()
	}
	{
		p.SetState(155)
		p.Match(bParserT__3)
	}
	{
		p.SetState(156)
		p.Statement()
	}

	return localctx
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_ifstmt
	return p
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *IfstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *bParser) Ifstmt() (localctx IIfstmtContext) {
	this := p
	_ = this

	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bParserRULE_ifstmt)

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
		p.SetState(158)
		p.Match(bParserT__11)
	}
	{
		p.SetState(159)
		p.Match(bParserT__2)
	}
	{
		p.SetState(160)
		p.Rvalue()
	}
	{
		p.SetState(161)
		p.Match(bParserT__3)
	}
	{
		p.SetState(162)
		p.Statement()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(163)
			p.Match(bParserT__12)
		}
		{
			p.SetState(164)
			p.Statement()
		}

	}

	return localctx
}

// ICasestmtContext is an interface to support dynamic dispatch.
type ICasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCasestmtContext differentiates from other interfaces.
	IsCasestmtContext()
}

type CasestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasestmtContext() *CasestmtContext {
	var p = new(CasestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_casestmt
	return p
}

func (*CasestmtContext) IsCasestmtContext() {}

func NewCasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasestmtContext {
	var p = new(CasestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_casestmt

	return p
}

func (s *CasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CasestmtContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *CasestmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterCasestmt(s)
	}
}

func (s *CasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitCasestmt(s)
	}
}

func (p *bParser) Casestmt() (localctx ICasestmtContext) {
	this := p
	_ = this

	localctx = NewCasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bParserRULE_casestmt)

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
		p.SetState(167)
		p.Match(bParserT__13)
	}
	{
		p.SetState(168)
		p.Constant()
	}
	{
		p.SetState(169)
		p.Match(bParserT__4)
	}
	{
		p.SetState(170)
		p.Statement()
	}

	return localctx
}

// IExternsmtContext is an interface to support dynamic dispatch.
type IExternsmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternsmtContext differentiates from other interfaces.
	IsExternsmtContext()
}

type ExternsmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternsmtContext() *ExternsmtContext {
	var p = new(ExternsmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_externsmt
	return p
}

func (*ExternsmtContext) IsExternsmtContext() {}

func NewExternsmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternsmtContext {
	var p = new(ExternsmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_externsmt

	return p
}

func (s *ExternsmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternsmtContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ExternsmtContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExternsmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternsmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternsmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExternsmt(s)
	}
}

func (s *ExternsmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExternsmt(s)
	}
}

func (p *bParser) Externsmt() (localctx IExternsmtContext) {
	this := p
	_ = this

	localctx = NewExternsmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, bParserRULE_externsmt)
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
		p.Match(bParserT__14)
	}
	{
		p.SetState(173)
		p.Name()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__0 {
		{
			p.SetState(174)
			p.Match(bParserT__0)
		}
		{
			p.SetState(175)
			p.Name()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.Match(bParserT__1)
	}

	return localctx
}

// IAutosmtContext is an interface to support dynamic dispatch.
type IAutosmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAutosmtContext differentiates from other interfaces.
	IsAutosmtContext()
}

type AutosmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAutosmtContext() *AutosmtContext {
	var p = new(AutosmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_autosmt
	return p
}

func (*AutosmtContext) IsAutosmtContext() {}

func NewAutosmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AutosmtContext {
	var p = new(AutosmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_autosmt

	return p
}

func (s *AutosmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AutosmtContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *AutosmtContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AutosmtContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *AutosmtContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AutosmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AutosmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AutosmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAutosmt(s)
	}
}

func (s *AutosmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAutosmt(s)
	}
}

func (p *bParser) Autosmt() (localctx IAutosmtContext) {
	this := p
	_ = this

	localctx = NewAutosmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, bParserRULE_autosmt)
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
		p.Match(bParserT__15)
	}
	{
		p.SetState(184)
		p.Name()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(bParserINT-40))|(1<<(bParserSTRING1-40))|(1<<(bParserSTRING2-40)))) != 0 {
		{
			p.SetState(185)
			p.Constant()
		}

	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__0 {
		{
			p.SetState(188)
			p.Match(bParserT__0)
		}
		{
			p.SetState(189)
			p.Name()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(bParserINT-40))|(1<<(bParserSTRING1-40))|(1<<(bParserSTRING2-40)))) != 0 {
			{
				p.SetState(190)
				p.Constant()
			}

		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(198)
		p.Match(bParserT__1)
	}

	return localctx
}

// IRvalueContext is an interface to support dynamic dispatch.
type IRvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRvalueContext differentiates from other interfaces.
	IsRvalueContext()
}

type RvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRvalueContext() *RvalueContext {
	var p = new(RvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_rvalue
	return p
}

func (*RvalueContext) IsRvalueContext() {}

func NewRvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RvalueContext {
	var p = new(RvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_rvalue

	return p
}

func (s *RvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *RvalueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RvalueContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *RvalueContext) Ternary() ITernaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITernaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITernaryContext)
}

func (s *RvalueContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *RvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterRvalue(s)
	}
}

func (s *RvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitRvalue(s)
	}
}

func (p *bParser) Rvalue() (localctx IRvalueContext) {
	this := p
	_ = this

	localctx = NewRvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, bParserRULE_rvalue)

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Comparison()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Ternary()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(203)
			p.Assignment()
		}

	}

	return localctx
}

// ITernaryContext is an interface to support dynamic dispatch.
type ITernaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTernaryContext differentiates from other interfaces.
	IsTernaryContext()
}

type TernaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTernaryContext() *TernaryContext {
	var p = new(TernaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_ternary
	return p
}

func (*TernaryContext) IsTernaryContext() {}

func NewTernaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TernaryContext {
	var p = new(TernaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ternary

	return p
}

func (s *TernaryContext) GetParser() antlr.Parser { return s.parser }

func (s *TernaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryContext) AllRvalue() []IRvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRvalueContext)(nil)).Elem())
	var tst = make([]IRvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRvalueContext)
		}
	}

	return tst
}

func (s *TernaryContext) Rvalue(i int) IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *TernaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TernaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterTernary(s)
	}
}

func (s *TernaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitTernary(s)
	}
}

func (p *bParser) Ternary() (localctx ITernaryContext) {
	this := p
	_ = this

	localctx = NewTernaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, bParserRULE_ternary)

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
		p.Expression()
	}
	{
		p.SetState(207)
		p.Match(bParserT__16)
	}
	{
		p.SetState(208)
		p.Rvalue()
	}
	{
		p.SetState(209)
		p.Match(bParserT__4)
	}
	{
		p.SetState(210)
		p.Rvalue()
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonContext) Binary() IBinaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *ComparisonContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *bParser) Comparison() (localctx IComparisonContext) {
	this := p
	_ = this

	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, bParserRULE_comparison)

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
		p.SetState(212)
		p.Expression()
	}
	{
		p.SetState(213)
		p.Binary()
	}
	{
		p.SetState(214)
		p.Rvalue()
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
	p.RuleIndex = bParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AssignmentContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignmentContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *bParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, bParserRULE_assignment)

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
		p.Name()
	}
	{
		p.SetState(217)
		p.Assign()
	}
	{
		p.SetState(218)
		p.Rvalue()
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
	p.RuleIndex = bParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ExpressionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExpressionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionContext) Incdec() IIncdecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncdecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncdecContext)
}

func (s *ExpressionContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *ExpressionContext) Functioninvocation() IFunctioninvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctioninvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctioninvocationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *bParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, bParserRULE_expression)

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

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(bParserT__2)
		}
		{
			p.SetState(221)
			p.Rvalue()
		}
		{
			p.SetState(222)
			p.Match(bParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Incdec()
		}
		{
			p.SetState(227)
			p.Name()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(229)
			p.Name()
		}
		{
			p.SetState(230)
			p.Incdec()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(232)
			p.Unary()
		}
		{
			p.SetState(233)
			p.Rvalue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(235)
			p.Match(bParserT__17)
		}
		{
			p.SetState(236)
			p.Name()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(237)
			p.Functioninvocation()
		}

	}

	return localctx
}

// IFunctioninvocationContext is an interface to support dynamic dispatch.
type IFunctioninvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioninvocationContext differentiates from other interfaces.
	IsFunctioninvocationContext()
}

type FunctioninvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioninvocationContext() *FunctioninvocationContext {
	var p = new(FunctioninvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_functioninvocation
	return p
}

func (*FunctioninvocationContext) IsFunctioninvocationContext() {}

func NewFunctioninvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioninvocationContext {
	var p = new(FunctioninvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_functioninvocation

	return p
}

func (s *FunctioninvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioninvocationContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FunctioninvocationContext) Functionparameters() IFunctionparametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionparametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionparametersContext)
}

func (s *FunctioninvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioninvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioninvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterFunctioninvocation(s)
	}
}

func (s *FunctioninvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitFunctioninvocation(s)
	}
}

func (p *bParser) Functioninvocation() (localctx IFunctioninvocationContext) {
	this := p
	_ = this

	localctx = NewFunctioninvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, bParserRULE_functioninvocation)
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
		p.Name()
	}
	{
		p.SetState(241)
		p.Match(bParserT__2)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<bParserT__2)|(1<<bParserT__17)|(1<<bParserT__19)|(1<<bParserT__20)|(1<<bParserT__21)|(1<<bParserT__22))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(bParserNAME-39))|(1<<(bParserINT-39))|(1<<(bParserSTRING1-39))|(1<<(bParserSTRING2-39)))) != 0) {
		{
			p.SetState(242)
			p.Functionparameters()
		}

	}
	{
		p.SetState(245)
		p.Match(bParserT__3)
	}

	return localctx
}

// IFunctionparametersContext is an interface to support dynamic dispatch.
type IFunctionparametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionparametersContext differentiates from other interfaces.
	IsFunctionparametersContext()
}

type FunctionparametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionparametersContext() *FunctionparametersContext {
	var p = new(FunctionparametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_functionparameters
	return p
}

func (*FunctionparametersContext) IsFunctionparametersContext() {}

func NewFunctionparametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionparametersContext {
	var p = new(FunctionparametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_functionparameters

	return p
}

func (s *FunctionparametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionparametersContext) AllRvalue() []IRvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRvalueContext)(nil)).Elem())
	var tst = make([]IRvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRvalueContext)
		}
	}

	return tst
}

func (s *FunctionparametersContext) Rvalue(i int) IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *FunctionparametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionparametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionparametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterFunctionparameters(s)
	}
}

func (s *FunctionparametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitFunctionparameters(s)
	}
}

func (p *bParser) Functionparameters() (localctx IFunctionparametersContext) {
	this := p
	_ = this

	localctx = NewFunctionparametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, bParserRULE_functionparameters)
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
		p.SetState(247)
		p.Rvalue()
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__0 {
		{
			p.SetState(248)
			p.Match(bParserT__0)
		}
		{
			p.SetState(249)
			p.Rvalue()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Binary() IBinaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *bParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, bParserRULE_assign)

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
		p.SetState(255)
		p.Match(bParserT__18)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(256)
			p.Binary()
		}

	}

	return localctx
}

// IIncdecContext is an interface to support dynamic dispatch.
type IIncdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncdecContext differentiates from other interfaces.
	IsIncdecContext()
}

type IncdecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncdecContext() *IncdecContext {
	var p = new(IncdecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_incdec
	return p
}

func (*IncdecContext) IsIncdecContext() {}

func NewIncdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncdecContext {
	var p = new(IncdecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_incdec

	return p
}

func (s *IncdecContext) GetParser() antlr.Parser { return s.parser }
func (s *IncdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIncdec(s)
	}
}

func (s *IncdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIncdec(s)
	}
}

func (p *bParser) Incdec() (localctx IIncdecContext) {
	this := p
	_ = this

	localctx = NewIncdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, bParserRULE_incdec)
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
		p.SetState(259)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bParserT__19 || _la == bParserT__20) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_unary
	return p
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (p *bParser) Unary() (localctx IUnaryContext) {
	this := p
	_ = this

	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, bParserRULE_unary)
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
		p.SetState(261)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bParserT__21 || _la == bParserT__22) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }
func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *bParser) Binary() (localctx IBinaryContext) {
	this := p
	_ = this

	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, bParserRULE_binary)
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
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(bParserT__17-18))|(1<<(bParserT__21-18))|(1<<(bParserT__23-18))|(1<<(bParserT__24-18))|(1<<(bParserT__25-18))|(1<<(bParserT__26-18))|(1<<(bParserT__27-18))|(1<<(bParserT__28-18))|(1<<(bParserT__29-18))|(1<<(bParserT__30-18))|(1<<(bParserT__31-18))|(1<<(bParserT__32-18))|(1<<(bParserT__33-18))|(1<<(bParserT__34-18))|(1<<(bParserT__35-18)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LvalueContext) AllRvalue() []IRvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRvalueContext)(nil)).Elem())
	var tst = make([]IRvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRvalueContext)
		}
	}

	return tst
}

func (s *LvalueContext) Rvalue(i int) IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *bParser) Lvalue() (localctx ILvalueContext) {
	this := p
	_ = this

	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, bParserRULE_lvalue)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(bParserT__34)
		}
		{
			p.SetState(267)
			p.Rvalue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)
			p.Rvalue()
		}
		{
			p.SetState(269)
			p.Match(bParserT__36)
		}
		{
			p.SetState(270)
			p.Rvalue()
		}
		{
			p.SetState(271)
			p.Match(bParserT__37)
		}

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
	p.RuleIndex = bParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(bParserINT, 0)
}

func (s *ConstantContext) STRING1() antlr.TerminalNode {
	return s.GetToken(bParserSTRING1, 0)
}

func (s *ConstantContext) STRING2() antlr.TerminalNode {
	return s.GetToken(bParserSTRING2, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *bParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, bParserRULE_constant)
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
		p.SetState(275)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(bParserINT-40))|(1<<(bParserSTRING1-40))|(1<<(bParserSTRING2-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = bParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(bParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *bParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, bParserRULE_name)

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
		p.Match(bParserNAME)
	}

	return localctx
}
