// Code generated from fen.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 115,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 26, 6, 26, 105, 10, 26, 13, 26, 14, 26, 106, 3,
	27, 6, 27, 110, 10, 27, 13, 27, 14, 27, 111, 3, 27, 3, 27, 2, 2, 28, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 2, 51, 26, 53, 27, 3, 2, 3, 4, 2, 11, 12, 15,
	15, 2, 115, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 3, 55, 3, 2, 2, 2, 5, 57,
	3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 61, 3, 2, 2, 2, 11, 63, 3, 2, 2, 2, 13,
	65, 3, 2, 2, 2, 15, 67, 3, 2, 2, 2, 17, 69, 3, 2, 2, 2, 19, 71, 3, 2, 2,
	2, 21, 73, 3, 2, 2, 2, 23, 75, 3, 2, 2, 2, 25, 77, 3, 2, 2, 2, 27, 79,
	3, 2, 2, 2, 29, 81, 3, 2, 2, 2, 31, 83, 3, 2, 2, 2, 33, 85, 3, 2, 2, 2,
	35, 87, 3, 2, 2, 2, 37, 89, 3, 2, 2, 2, 39, 91, 3, 2, 2, 2, 41, 93, 3,
	2, 2, 2, 43, 95, 3, 2, 2, 2, 45, 97, 3, 2, 2, 2, 47, 99, 3, 2, 2, 2, 49,
	101, 3, 2, 2, 2, 51, 104, 3, 2, 2, 2, 53, 109, 3, 2, 2, 2, 55, 56, 7, 34,
	2, 2, 56, 4, 3, 2, 2, 2, 57, 58, 7, 121, 2, 2, 58, 6, 3, 2, 2, 2, 59, 60,
	7, 100, 2, 2, 60, 8, 3, 2, 2, 2, 61, 62, 7, 47, 2, 2, 62, 10, 3, 2, 2,
	2, 63, 64, 7, 83, 2, 2, 64, 12, 3, 2, 2, 2, 65, 66, 7, 77, 2, 2, 66, 14,
	3, 2, 2, 2, 67, 68, 7, 109, 2, 2, 68, 16, 3, 2, 2, 2, 69, 70, 7, 115, 2,
	2, 70, 18, 3, 2, 2, 2, 71, 72, 7, 99, 2, 2, 72, 20, 3, 2, 2, 2, 73, 74,
	7, 101, 2, 2, 74, 22, 3, 2, 2, 2, 75, 76, 7, 102, 2, 2, 76, 24, 3, 2, 2,
	2, 77, 78, 7, 103, 2, 2, 78, 26, 3, 2, 2, 2, 79, 80, 7, 104, 2, 2, 80,
	28, 3, 2, 2, 2, 81, 82, 7, 105, 2, 2, 82, 30, 3, 2, 2, 2, 83, 84, 7, 106,
	2, 2, 84, 32, 3, 2, 2, 2, 85, 86, 7, 49, 2, 2, 86, 34, 3, 2, 2, 2, 87,
	88, 7, 114, 2, 2, 88, 36, 3, 2, 2, 2, 89, 90, 7, 116, 2, 2, 90, 38, 3,
	2, 2, 2, 91, 92, 7, 112, 2, 2, 92, 40, 3, 2, 2, 2, 93, 94, 7, 82, 2, 2,
	94, 42, 3, 2, 2, 2, 95, 96, 7, 84, 2, 2, 96, 44, 3, 2, 2, 2, 97, 98, 7,
	80, 2, 2, 98, 46, 3, 2, 2, 2, 99, 100, 7, 68, 2, 2, 100, 48, 3, 2, 2, 2,
	101, 102, 4, 50, 59, 2, 102, 50, 3, 2, 2, 2, 103, 105, 5, 49, 25, 2, 104,
	103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 52, 3, 2, 2, 2, 108, 110, 9, 2, 2, 2, 109, 108, 3, 2,
	2, 2, 110, 111, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2,
	112, 113, 3, 2, 2, 2, 113, 114, 8, 27, 2, 2, 114, 54, 3, 2, 2, 2, 5, 2,
	106, 111, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "' '", "'w'", "'b'", "'-'", "'Q'", "'K'", "'k'", "'q'", "'a'", "'c'",
	"'d'", "'e'", "'f'", "'g'", "'h'", "'/'", "'p'", "'r'", "'n'", "'P'", "'R'",
	"'N'", "'B'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "DIGIT", "NUMBER",
	"WS",
}

type fenLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewfenLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *fenLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfenLexer(input antlr.CharStream) *fenLexer {
	l := new(fenLexer)
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
	l.GrammarFileName = "fen.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fenLexer tokens.
const (
	fenLexerT__0   = 1
	fenLexerT__1   = 2
	fenLexerT__2   = 3
	fenLexerT__3   = 4
	fenLexerT__4   = 5
	fenLexerT__5   = 6
	fenLexerT__6   = 7
	fenLexerT__7   = 8
	fenLexerT__8   = 9
	fenLexerT__9   = 10
	fenLexerT__10  = 11
	fenLexerT__11  = 12
	fenLexerT__12  = 13
	fenLexerT__13  = 14
	fenLexerT__14  = 15
	fenLexerT__15  = 16
	fenLexerT__16  = 17
	fenLexerT__17  = 18
	fenLexerT__18  = 19
	fenLexerT__19  = 20
	fenLexerT__20  = 21
	fenLexerT__21  = 22
	fenLexerT__22  = 23
	fenLexerNUMBER = 24
	fenLexerWS     = 25
)
