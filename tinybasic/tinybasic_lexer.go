// Code generated from tinybasic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tinybasic

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 152,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 133, 10,
	22, 12, 22, 14, 22, 136, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 6, 25, 145, 10, 25, 13, 25, 14, 25, 146, 3, 26, 3, 26, 3, 26,
	3, 26, 2, 2, 27, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 3, 2, 5, 5,
	2, 12, 12, 15, 15, 36, 36, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34,
	2, 153, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 53, 3, 2, 2, 2, 5, 59, 3,
	2, 2, 2, 7, 62, 3, 2, 2, 2, 9, 67, 3, 2, 2, 2, 11, 72, 3, 2, 2, 2, 13,
	78, 3, 2, 2, 2, 15, 82, 3, 2, 2, 2, 17, 84, 3, 2, 2, 2, 19, 90, 3, 2, 2,
	2, 21, 97, 3, 2, 2, 2, 23, 103, 3, 2, 2, 2, 25, 108, 3, 2, 2, 2, 27, 112,
	3, 2, 2, 2, 29, 116, 3, 2, 2, 2, 31, 118, 3, 2, 2, 2, 33, 120, 3, 2, 2,
	2, 35, 122, 3, 2, 2, 2, 37, 124, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 128,
	3, 2, 2, 2, 43, 130, 3, 2, 2, 2, 45, 139, 3, 2, 2, 2, 47, 141, 3, 2, 2,
	2, 49, 144, 3, 2, 2, 2, 51, 148, 3, 2, 2, 2, 53, 54, 7, 82, 2, 2, 54, 55,
	7, 84, 2, 2, 55, 56, 7, 75, 2, 2, 56, 57, 7, 80, 2, 2, 57, 58, 7, 86, 2,
	2, 58, 4, 3, 2, 2, 2, 59, 60, 7, 75, 2, 2, 60, 61, 7, 72, 2, 2, 61, 6,
	3, 2, 2, 2, 62, 63, 7, 86, 2, 2, 63, 64, 7, 74, 2, 2, 64, 65, 7, 71, 2,
	2, 65, 66, 7, 80, 2, 2, 66, 8, 3, 2, 2, 2, 67, 68, 7, 73, 2, 2, 68, 69,
	7, 81, 2, 2, 69, 70, 7, 86, 2, 2, 70, 71, 7, 81, 2, 2, 71, 10, 3, 2, 2,
	2, 72, 73, 7, 75, 2, 2, 73, 74, 7, 80, 2, 2, 74, 75, 7, 82, 2, 2, 75, 76,
	7, 87, 2, 2, 76, 77, 7, 86, 2, 2, 77, 12, 3, 2, 2, 2, 78, 79, 7, 78, 2,
	2, 79, 80, 7, 71, 2, 2, 80, 81, 7, 86, 2, 2, 81, 14, 3, 2, 2, 2, 82, 83,
	7, 63, 2, 2, 83, 16, 3, 2, 2, 2, 84, 85, 7, 73, 2, 2, 85, 86, 7, 81, 2,
	2, 86, 87, 7, 85, 2, 2, 87, 88, 7, 87, 2, 2, 88, 89, 7, 68, 2, 2, 89, 18,
	3, 2, 2, 2, 90, 91, 7, 84, 2, 2, 91, 92, 7, 71, 2, 2, 92, 93, 7, 86, 2,
	2, 93, 94, 7, 87, 2, 2, 94, 95, 7, 84, 2, 2, 95, 96, 7, 80, 2, 2, 96, 20,
	3, 2, 2, 2, 97, 98, 7, 69, 2, 2, 98, 99, 7, 78, 2, 2, 99, 100, 7, 71, 2,
	2, 100, 101, 7, 67, 2, 2, 101, 102, 7, 84, 2, 2, 102, 22, 3, 2, 2, 2, 103,
	104, 7, 78, 2, 2, 104, 105, 7, 75, 2, 2, 105, 106, 7, 85, 2, 2, 106, 107,
	7, 86, 2, 2, 107, 24, 3, 2, 2, 2, 108, 109, 7, 84, 2, 2, 109, 110, 7, 87,
	2, 2, 110, 111, 7, 80, 2, 2, 111, 26, 3, 2, 2, 2, 112, 113, 7, 71, 2, 2,
	113, 114, 7, 80, 2, 2, 114, 115, 7, 70, 2, 2, 115, 28, 3, 2, 2, 2, 116,
	117, 7, 46, 2, 2, 117, 30, 3, 2, 2, 2, 118, 119, 7, 45, 2, 2, 119, 32,
	3, 2, 2, 2, 120, 121, 7, 47, 2, 2, 121, 34, 3, 2, 2, 2, 122, 123, 7, 44,
	2, 2, 123, 36, 3, 2, 2, 2, 124, 125, 7, 49, 2, 2, 125, 38, 3, 2, 2, 2,
	126, 127, 7, 62, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 7, 64, 2, 2, 129,
	42, 3, 2, 2, 2, 130, 134, 7, 36, 2, 2, 131, 133, 10, 2, 2, 2, 132, 131,
	3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2,
	2, 2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 138, 7, 36, 2, 2,
	138, 44, 3, 2, 2, 2, 139, 140, 4, 50, 59, 2, 140, 46, 3, 2, 2, 2, 141,
	142, 4, 67, 92, 2, 142, 48, 3, 2, 2, 2, 143, 145, 9, 3, 2, 2, 144, 143,
	3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2,
	2, 2, 147, 50, 3, 2, 2, 2, 148, 149, 9, 4, 2, 2, 149, 150, 3, 2, 2, 2,
	150, 151, 8, 26, 2, 2, 151, 52, 3, 2, 2, 2, 5, 2, 134, 146, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'PRINT'", "'IF'", "'THEN'", "'GOTO'", "'INPUT'", "'LET'", "'='", "'GOSUB'",
	"'RETURN'", "'CLEAR'", "'LIST'", "'RUN'", "'END'", "','", "'+'", "'-'",
	"'*'", "'/'", "'<'", "'>'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "STRING", "DIGIT", "VAR", "CR", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "STRING", "DIGIT", "VAR", "CR", "WS",
}

type tinybasicLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewtinybasicLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *tinybasicLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtinybasicLexer(input antlr.CharStream) *tinybasicLexer {
	l := new(tinybasicLexer)
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
	l.GrammarFileName = "tinybasic.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tinybasicLexer tokens.
const (
	tinybasicLexerT__0   = 1
	tinybasicLexerT__1   = 2
	tinybasicLexerT__2   = 3
	tinybasicLexerT__3   = 4
	tinybasicLexerT__4   = 5
	tinybasicLexerT__5   = 6
	tinybasicLexerT__6   = 7
	tinybasicLexerT__7   = 8
	tinybasicLexerT__8   = 9
	tinybasicLexerT__9   = 10
	tinybasicLexerT__10  = 11
	tinybasicLexerT__11  = 12
	tinybasicLexerT__12  = 13
	tinybasicLexerT__13  = 14
	tinybasicLexerT__14  = 15
	tinybasicLexerT__15  = 16
	tinybasicLexerT__16  = 17
	tinybasicLexerT__17  = 18
	tinybasicLexerT__18  = 19
	tinybasicLexerT__19  = 20
	tinybasicLexerSTRING = 21
	tinybasicLexerDIGIT  = 22
	tinybasicLexerVAR    = 23
	tinybasicLexerCR     = 24
	tinybasicLexerWS     = 25
)
