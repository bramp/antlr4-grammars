// Code generated from refal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package refal

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 186,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 6, 18, 109, 10, 18, 13, 18, 14, 18, 110,
	3, 19, 3, 19, 7, 19, 115, 10, 19, 12, 19, 14, 19, 118, 11, 19, 3, 20, 3,
	20, 7, 20, 122, 10, 20, 12, 20, 14, 20, 125, 11, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 7, 21, 131, 10, 21, 12, 21, 14, 21, 134, 11, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 155, 10, 22, 3, 23, 3,
	23, 7, 23, 159, 10, 23, 12, 23, 14, 23, 162, 11, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 7, 24, 170, 10, 24, 12, 24, 14, 24, 173, 11, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 6, 25, 181, 10, 25, 13, 25, 14,
	25, 182, 3, 25, 3, 25, 3, 171, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	3, 2, 9, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 7, 2, 47, 47, 50, 59, 67,
	92, 97, 97, 99, 124, 3, 2, 36, 36, 3, 2, 41, 41, 4, 2, 12, 12, 15, 15,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 198, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 3, 51, 3, 2, 2,
	2, 5, 53, 3, 2, 2, 2, 7, 60, 3, 2, 2, 2, 9, 62, 3, 2, 2, 2, 11, 64, 3,
	2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 82, 3, 2, 2, 2, 17, 89, 3, 2, 2, 2, 19,
	91, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 95, 3, 2, 2, 2, 25, 97, 3, 2, 2,
	2, 27, 99, 3, 2, 2, 2, 29, 101, 3, 2, 2, 2, 31, 103, 3, 2, 2, 2, 33, 105,
	3, 2, 2, 2, 35, 108, 3, 2, 2, 2, 37, 112, 3, 2, 2, 2, 39, 119, 3, 2, 2,
	2, 41, 128, 3, 2, 2, 2, 43, 154, 3, 2, 2, 2, 45, 156, 3, 2, 2, 2, 47, 165,
	3, 2, 2, 2, 49, 180, 3, 2, 2, 2, 51, 52, 7, 61, 2, 2, 52, 4, 3, 2, 2, 2,
	53, 54, 7, 38, 2, 2, 54, 55, 7, 71, 2, 2, 55, 56, 7, 80, 2, 2, 56, 57,
	7, 86, 2, 2, 57, 58, 7, 84, 2, 2, 58, 59, 7, 91, 2, 2, 59, 6, 3, 2, 2,
	2, 60, 61, 7, 125, 2, 2, 61, 8, 3, 2, 2, 2, 62, 63, 7, 127, 2, 2, 63, 10,
	3, 2, 2, 2, 64, 65, 7, 38, 2, 2, 65, 66, 7, 71, 2, 2, 66, 67, 7, 90, 2,
	2, 67, 68, 7, 86, 2, 2, 68, 69, 7, 71, 2, 2, 69, 70, 7, 84, 2, 2, 70, 71,
	7, 80, 2, 2, 71, 72, 7, 67, 2, 2, 72, 73, 7, 78, 2, 2, 73, 12, 3, 2, 2,
	2, 74, 75, 7, 38, 2, 2, 75, 76, 7, 71, 2, 2, 76, 77, 7, 90, 2, 2, 77, 78,
	7, 86, 2, 2, 78, 79, 7, 71, 2, 2, 79, 80, 7, 84, 2, 2, 80, 81, 7, 80, 2,
	2, 81, 14, 3, 2, 2, 2, 82, 83, 7, 38, 2, 2, 83, 84, 7, 71, 2, 2, 84, 85,
	7, 90, 2, 2, 85, 86, 7, 86, 2, 2, 86, 87, 7, 84, 2, 2, 87, 88, 7, 80, 2,
	2, 88, 16, 3, 2, 2, 2, 89, 90, 7, 46, 2, 2, 90, 18, 3, 2, 2, 2, 91, 92,
	7, 63, 2, 2, 92, 20, 3, 2, 2, 2, 93, 94, 7, 60, 2, 2, 94, 22, 3, 2, 2,
	2, 95, 96, 7, 62, 2, 2, 96, 24, 3, 2, 2, 2, 97, 98, 7, 64, 2, 2, 98, 26,
	3, 2, 2, 2, 99, 100, 7, 117, 2, 2, 100, 28, 3, 2, 2, 2, 101, 102, 7, 48,
	2, 2, 102, 30, 3, 2, 2, 2, 103, 104, 7, 118, 2, 2, 104, 32, 3, 2, 2, 2,
	105, 106, 7, 103, 2, 2, 106, 34, 3, 2, 2, 2, 107, 109, 9, 2, 2, 2, 108,
	107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111,
	3, 2, 2, 2, 111, 36, 3, 2, 2, 2, 112, 116, 9, 3, 2, 2, 113, 115, 9, 4,
	2, 2, 114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 38, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 123,
	7, 36, 2, 2, 120, 122, 10, 5, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3,
	2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2,
	2, 125, 123, 3, 2, 2, 2, 126, 127, 7, 36, 2, 2, 127, 40, 3, 2, 2, 2, 128,
	132, 7, 41, 2, 2, 129, 131, 10, 6, 2, 2, 130, 129, 3, 2, 2, 2, 131, 134,
	3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2,
	2, 2, 134, 132, 3, 2, 2, 2, 135, 136, 7, 41, 2, 2, 136, 42, 3, 2, 2, 2,
	137, 138, 7, 94, 2, 2, 138, 155, 7, 41, 2, 2, 139, 140, 7, 94, 2, 2, 140,
	155, 7, 36, 2, 2, 141, 142, 7, 94, 2, 2, 142, 155, 7, 94, 2, 2, 143, 144,
	7, 94, 2, 2, 144, 155, 7, 112, 2, 2, 145, 146, 7, 94, 2, 2, 146, 155, 7,
	116, 2, 2, 147, 148, 7, 94, 2, 2, 148, 155, 7, 118, 2, 2, 149, 150, 7,
	94, 2, 2, 150, 151, 7, 122, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 9, 2,
	2, 2, 153, 155, 9, 2, 2, 2, 154, 137, 3, 2, 2, 2, 154, 139, 3, 2, 2, 2,
	154, 141, 3, 2, 2, 2, 154, 143, 3, 2, 2, 2, 154, 145, 3, 2, 2, 2, 154,
	147, 3, 2, 2, 2, 154, 149, 3, 2, 2, 2, 155, 44, 3, 2, 2, 2, 156, 160, 7,
	44, 2, 2, 157, 159, 10, 7, 2, 2, 158, 157, 3, 2, 2, 2, 159, 162, 3, 2,
	2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 163, 3, 2, 2, 2,
	162, 160, 3, 2, 2, 2, 163, 164, 8, 23, 2, 2, 164, 46, 3, 2, 2, 2, 165,
	166, 7, 49, 2, 2, 166, 167, 7, 44, 2, 2, 167, 171, 3, 2, 2, 2, 168, 170,
	11, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 172, 3, 2,
	2, 2, 171, 169, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2,
	174, 175, 7, 44, 2, 2, 175, 176, 7, 49, 2, 2, 176, 177, 3, 2, 2, 2, 177,
	178, 8, 24, 2, 2, 178, 48, 3, 2, 2, 2, 179, 181, 9, 8, 2, 2, 180, 179,
	3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2,
	2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 8, 25, 2, 2, 185, 50, 3, 2, 2, 2,
	11, 2, 110, 116, 123, 132, 154, 160, 171, 182, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'$ENTRY'", "'{'", "'}'", "'$EXTERNAL'", "'$EXTERN'", "'$EXTRN'",
	"','", "'='", "':'", "'<'", "'>'", "'s'", "'.'", "'t'", "'e'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "DIGITS",
	"IDENTIFER", "STRING", "STRING2", "CHAR", "LINE_COMMENT", "BLOCK_COMMENT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "DIGITS",
	"IDENTIFER", "STRING", "STRING2", "CHAR", "LINE_COMMENT", "BLOCK_COMMENT",
	"WS",
}

type refalLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewrefalLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *refalLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewrefalLexer(input antlr.CharStream) *refalLexer {
	l := new(refalLexer)
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
	l.GrammarFileName = "refal.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// refalLexer tokens.
const (
	refalLexerT__0          = 1
	refalLexerT__1          = 2
	refalLexerT__2          = 3
	refalLexerT__3          = 4
	refalLexerT__4          = 5
	refalLexerT__5          = 6
	refalLexerT__6          = 7
	refalLexerT__7          = 8
	refalLexerT__8          = 9
	refalLexerT__9          = 10
	refalLexerT__10         = 11
	refalLexerT__11         = 12
	refalLexerT__12         = 13
	refalLexerT__13         = 14
	refalLexerT__14         = 15
	refalLexerT__15         = 16
	refalLexerDIGITS        = 17
	refalLexerIDENTIFER     = 18
	refalLexerSTRING        = 19
	refalLexerSTRING2       = 20
	refalLexerCHAR          = 21
	refalLexerLINE_COMMENT  = 22
	refalLexerBLOCK_COMMENT = 23
	refalLexerWS            = 24
)
