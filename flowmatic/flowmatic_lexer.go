// Code generated from flowmatic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flowmatic

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 36, 281,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 30, 6, 30, 243, 10, 30, 13, 30, 14, 30, 244, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 6, 33, 266, 10, 33, 13,
	33, 14, 33, 267, 3, 34, 6, 34, 271, 10, 34, 13, 34, 14, 34, 272, 3, 35,
	6, 35, 276, 10, 35, 13, 35, 14, 35, 277, 3, 35, 3, 35, 2, 2, 36, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 3, 2, 5, 5, 2, 47, 47, 67, 92,
	99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 284, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 5, 73,
	3, 2, 2, 2, 7, 75, 3, 2, 2, 2, 9, 77, 3, 2, 2, 2, 11, 79, 3, 2, 2, 2, 13,
	83, 3, 2, 2, 2, 15, 90, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19, 101, 3, 2,
	2, 2, 21, 104, 3, 2, 2, 2, 23, 114, 3, 2, 2, 2, 25, 125, 3, 2, 2, 2, 27,
	134, 3, 2, 2, 2, 29, 137, 3, 2, 2, 2, 31, 147, 3, 2, 2, 2, 33, 154, 3,
	2, 2, 2, 35, 159, 3, 2, 2, 2, 37, 167, 3, 2, 2, 2, 39, 172, 3, 2, 2, 2,
	41, 177, 3, 2, 2, 2, 43, 181, 3, 2, 2, 2, 45, 186, 3, 2, 2, 2, 47, 194,
	3, 2, 2, 2, 49, 204, 3, 2, 2, 2, 51, 210, 3, 2, 2, 2, 53, 220, 3, 2, 2,
	2, 55, 226, 3, 2, 2, 2, 57, 234, 3, 2, 2, 2, 59, 242, 3, 2, 2, 2, 61, 246,
	3, 2, 2, 2, 63, 252, 3, 2, 2, 2, 65, 265, 3, 2, 2, 2, 67, 270, 3, 2, 2,
	2, 69, 275, 3, 2, 2, 2, 71, 72, 7, 61, 2, 2, 72, 4, 3, 2, 2, 2, 73, 74,
	7, 48, 2, 2, 74, 6, 3, 2, 2, 2, 75, 76, 7, 42, 2, 2, 76, 8, 3, 2, 2, 2,
	77, 78, 7, 43, 2, 2, 78, 10, 3, 2, 2, 2, 79, 80, 7, 74, 2, 2, 80, 81, 7,
	85, 2, 2, 81, 82, 7, 82, 2, 2, 82, 12, 3, 2, 2, 2, 83, 84, 7, 81, 2, 2,
	84, 85, 7, 87, 2, 2, 85, 86, 7, 86, 2, 2, 86, 87, 7, 82, 2, 2, 87, 88,
	7, 87, 2, 2, 88, 89, 7, 86, 2, 2, 89, 14, 3, 2, 2, 2, 90, 91, 7, 75, 2,
	2, 91, 92, 7, 80, 2, 2, 92, 93, 7, 82, 2, 2, 93, 94, 7, 87, 2, 2, 94, 95,
	7, 86, 2, 2, 95, 16, 3, 2, 2, 2, 96, 97, 7, 76, 2, 2, 97, 98, 7, 87, 2,
	2, 98, 99, 7, 79, 2, 2, 99, 100, 7, 82, 2, 2, 100, 18, 3, 2, 2, 2, 101,
	102, 7, 86, 2, 2, 102, 103, 7, 81, 2, 2, 103, 20, 3, 2, 2, 2, 104, 105,
	7, 84, 2, 2, 105, 106, 7, 71, 2, 2, 106, 107, 7, 67, 2, 2, 107, 108, 7,
	70, 2, 2, 108, 109, 7, 47, 2, 2, 109, 110, 7, 75, 2, 2, 110, 111, 7, 86,
	2, 2, 111, 112, 7, 71, 2, 2, 112, 113, 7, 79, 2, 2, 113, 22, 3, 2, 2, 2,
	114, 115, 7, 89, 2, 2, 115, 116, 7, 84, 2, 2, 116, 117, 7, 75, 2, 2, 117,
	118, 7, 86, 2, 2, 118, 119, 7, 71, 2, 2, 119, 120, 7, 47, 2, 2, 120, 121,
	7, 75, 2, 2, 121, 122, 7, 86, 2, 2, 122, 123, 7, 71, 2, 2, 123, 124, 7,
	79, 2, 2, 124, 24, 3, 2, 2, 2, 125, 126, 7, 86, 2, 2, 126, 127, 7, 84,
	2, 2, 127, 128, 7, 67, 2, 2, 128, 129, 7, 80, 2, 2, 129, 130, 7, 85, 2,
	2, 130, 131, 7, 72, 2, 2, 131, 132, 7, 71, 2, 2, 132, 133, 7, 84, 2, 2,
	133, 26, 3, 2, 2, 2, 134, 135, 7, 75, 2, 2, 135, 136, 7, 72, 2, 2, 136,
	28, 3, 2, 2, 2, 137, 138, 7, 81, 2, 2, 138, 139, 7, 86, 2, 2, 139, 140,
	7, 74, 2, 2, 140, 141, 7, 71, 2, 2, 141, 142, 7, 84, 2, 2, 142, 143, 7,
	89, 2, 2, 143, 144, 7, 75, 2, 2, 144, 145, 7, 85, 2, 2, 145, 146, 7, 71,
	2, 2, 146, 30, 3, 2, 2, 2, 147, 148, 7, 84, 2, 2, 148, 149, 7, 71, 2, 2,
	149, 150, 7, 89, 2, 2, 150, 151, 7, 75, 2, 2, 151, 152, 7, 80, 2, 2, 152,
	153, 7, 70, 2, 2, 153, 32, 3, 2, 2, 2, 154, 155, 7, 85, 2, 2, 155, 156,
	7, 86, 2, 2, 156, 157, 7, 81, 2, 2, 157, 158, 7, 82, 2, 2, 158, 34, 3,
	2, 2, 2, 159, 160, 7, 69, 2, 2, 160, 161, 7, 81, 2, 2, 161, 162, 7, 79,
	2, 2, 162, 163, 7, 82, 2, 2, 163, 164, 7, 67, 2, 2, 164, 165, 7, 84, 2,
	2, 165, 166, 7, 71, 2, 2, 166, 36, 3, 2, 2, 2, 167, 168, 7, 89, 2, 2, 168,
	169, 7, 75, 2, 2, 169, 170, 7, 86, 2, 2, 170, 171, 7, 74, 2, 2, 171, 38,
	3, 2, 2, 2, 172, 173, 7, 79, 2, 2, 173, 174, 7, 81, 2, 2, 174, 175, 7,
	88, 2, 2, 175, 176, 7, 71, 2, 2, 176, 40, 3, 2, 2, 2, 177, 178, 7, 85,
	2, 2, 178, 179, 7, 71, 2, 2, 179, 180, 7, 86, 2, 2, 180, 42, 3, 2, 2, 2,
	181, 182, 7, 86, 2, 2, 182, 183, 7, 71, 2, 2, 183, 184, 7, 85, 2, 2, 184,
	185, 7, 86, 2, 2, 185, 44, 3, 2, 2, 2, 186, 187, 7, 67, 2, 2, 187, 188,
	7, 73, 2, 2, 188, 189, 7, 67, 2, 2, 189, 190, 7, 75, 2, 2, 190, 191, 7,
	80, 2, 2, 191, 192, 7, 85, 2, 2, 192, 193, 7, 86, 2, 2, 193, 46, 3, 2,
	2, 2, 194, 195, 7, 69, 2, 2, 195, 196, 7, 78, 2, 2, 196, 197, 7, 81, 2,
	2, 197, 198, 7, 85, 2, 2, 198, 199, 7, 71, 2, 2, 199, 200, 7, 47, 2, 2,
	200, 201, 7, 81, 2, 2, 201, 202, 7, 87, 2, 2, 202, 203, 7, 86, 2, 2, 203,
	48, 3, 2, 2, 2, 204, 205, 7, 72, 2, 2, 205, 206, 7, 75, 2, 2, 206, 207,
	7, 78, 2, 2, 207, 208, 7, 71, 2, 2, 208, 209, 7, 85, 2, 2, 209, 50, 3,
	2, 2, 2, 210, 211, 7, 81, 2, 2, 211, 212, 7, 82, 2, 2, 212, 213, 7, 71,
	2, 2, 213, 214, 7, 84, 2, 2, 214, 215, 7, 67, 2, 2, 215, 216, 7, 86, 2,
	2, 216, 217, 7, 75, 2, 2, 217, 218, 7, 81, 2, 2, 218, 219, 7, 80, 2, 2,
	219, 52, 3, 2, 2, 2, 220, 221, 7, 71, 2, 2, 221, 222, 7, 83, 2, 2, 222,
	223, 7, 87, 2, 2, 223, 224, 7, 67, 2, 2, 224, 225, 7, 78, 2, 2, 225, 54,
	3, 2, 2, 2, 226, 227, 7, 73, 2, 2, 227, 228, 7, 84, 2, 2, 228, 229, 7,
	71, 2, 2, 229, 230, 7, 67, 2, 2, 230, 231, 7, 86, 2, 2, 231, 232, 7, 71,
	2, 2, 232, 233, 7, 84, 2, 2, 233, 56, 3, 2, 2, 2, 234, 235, 7, 42, 2, 2,
	235, 236, 7, 71, 2, 2, 236, 237, 7, 80, 2, 2, 237, 238, 7, 70, 2, 2, 238,
	239, 3, 2, 2, 2, 239, 240, 7, 43, 2, 2, 240, 58, 3, 2, 2, 2, 241, 243,
	7, 92, 2, 2, 242, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 242, 3, 2,
	2, 2, 244, 245, 3, 2, 2, 2, 245, 60, 3, 2, 2, 2, 246, 247, 7, 73, 2, 2,
	247, 248, 7, 81, 2, 2, 248, 249, 7, 34, 2, 2, 249, 250, 7, 86, 2, 2, 250,
	251, 7, 81, 2, 2, 251, 62, 3, 2, 2, 2, 252, 253, 7, 71, 2, 2, 253, 254,
	7, 80, 2, 2, 254, 255, 7, 70, 2, 2, 255, 256, 7, 34, 2, 2, 256, 257, 7,
	81, 2, 2, 257, 258, 7, 72, 2, 2, 258, 259, 7, 34, 2, 2, 259, 260, 7, 70,
	2, 2, 260, 261, 7, 67, 2, 2, 261, 262, 7, 86, 2, 2, 262, 263, 7, 67, 2,
	2, 263, 64, 3, 2, 2, 2, 264, 266, 9, 2, 2, 2, 265, 264, 3, 2, 2, 2, 266,
	267, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 66, 3,
	2, 2, 2, 269, 271, 9, 3, 2, 2, 270, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2,
	2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 68, 3, 2, 2, 2, 274,
	276, 9, 4, 2, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275,
	3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 8, 35,
	2, 2, 280, 70, 3, 2, 2, 2, 7, 2, 244, 267, 272, 277, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'.'", "'('", "')'", "'HSP'", "'OUTPUT'", "'INPUT'", "'JUMP'",
	"'TO'", "'READ-ITEM'", "'WRITE-ITEM'", "'TRANSFER'", "'IF'", "'OTHERWISE'",
	"'REWIND'", "'STOP'", "'COMPARE'", "'WITH'", "'MOVE'", "'SET'", "'TEST'",
	"'AGAINST'", "'CLOSE-OUT'", "'FILES'", "'OPERATION'", "'EQUAL'", "'GREATER'",
	"", "", "'GO TO'", "'END OF DATA'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "END", "ZERO", "GOTO", "EOD", "ID",
	"NUM", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "END", "ZERO", "GOTO", "EOD", "ID", "NUM", "WS",
}

type flowmaticLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewflowmaticLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *flowmaticLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewflowmaticLexer(input antlr.CharStream) *flowmaticLexer {
	l := new(flowmaticLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "flowmatic.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// flowmaticLexer tokens.
const (
	flowmaticLexerT__0  = 1
	flowmaticLexerT__1  = 2
	flowmaticLexerT__2  = 3
	flowmaticLexerT__3  = 4
	flowmaticLexerT__4  = 5
	flowmaticLexerT__5  = 6
	flowmaticLexerT__6  = 7
	flowmaticLexerT__7  = 8
	flowmaticLexerT__8  = 9
	flowmaticLexerT__9  = 10
	flowmaticLexerT__10 = 11
	flowmaticLexerT__11 = 12
	flowmaticLexerT__12 = 13
	flowmaticLexerT__13 = 14
	flowmaticLexerT__14 = 15
	flowmaticLexerT__15 = 16
	flowmaticLexerT__16 = 17
	flowmaticLexerT__17 = 18
	flowmaticLexerT__18 = 19
	flowmaticLexerT__19 = 20
	flowmaticLexerT__20 = 21
	flowmaticLexerT__21 = 22
	flowmaticLexerT__22 = 23
	flowmaticLexerT__23 = 24
	flowmaticLexerT__24 = 25
	flowmaticLexerT__25 = 26
	flowmaticLexerT__26 = 27
	flowmaticLexerEND   = 28
	flowmaticLexerZERO  = 29
	flowmaticLexerGOTO  = 30
	flowmaticLexerEOD   = 31
	flowmaticLexerID    = 32
	flowmaticLexerNUM   = 33
	flowmaticLexerWS    = 34
)
