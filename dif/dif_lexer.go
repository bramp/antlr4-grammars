// Code generated from dif.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dif

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 87, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 5, 10, 65, 10, 10, 3, 10, 6, 10, 68, 10, 10, 13, 10, 14,
	10, 69, 3, 11, 3, 11, 7, 11, 74, 10, 11, 12, 11, 14, 11, 77, 11, 11, 3,
	11, 3, 11, 3, 12, 6, 12, 82, 10, 12, 13, 12, 14, 12, 83, 3, 12, 3, 12,
	3, 75, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 90, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25,
	3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 39, 3, 2, 2, 2, 9, 46, 3, 2, 2, 2, 11,
	51, 3, 2, 2, 2, 13, 55, 3, 2, 2, 2, 15, 59, 3, 2, 2, 2, 17, 61, 3, 2, 2,
	2, 19, 64, 3, 2, 2, 2, 21, 71, 3, 2, 2, 2, 23, 81, 3, 2, 2, 2, 25, 26,
	7, 86, 2, 2, 26, 27, 7, 67, 2, 2, 27, 28, 7, 68, 2, 2, 28, 29, 7, 78, 2,
	2, 29, 30, 7, 71, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 88, 2, 2, 32, 33,
	7, 71, 2, 2, 33, 34, 7, 69, 2, 2, 34, 35, 7, 86, 2, 2, 35, 36, 7, 81, 2,
	2, 36, 37, 7, 84, 2, 2, 37, 38, 7, 85, 2, 2, 38, 6, 3, 2, 2, 2, 39, 40,
	7, 86, 2, 2, 40, 41, 7, 87, 2, 2, 41, 42, 7, 82, 2, 2, 42, 43, 7, 78, 2,
	2, 43, 44, 7, 71, 2, 2, 44, 45, 7, 85, 2, 2, 45, 8, 3, 2, 2, 2, 46, 47,
	7, 70, 2, 2, 47, 48, 7, 67, 2, 2, 48, 49, 7, 86, 2, 2, 49, 50, 7, 67, 2,
	2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 68, 2, 2, 52, 53, 7, 81, 2, 2, 53, 54,
	7, 86, 2, 2, 54, 12, 3, 2, 2, 2, 55, 56, 7, 71, 2, 2, 56, 57, 7, 81, 2,
	2, 57, 58, 7, 70, 2, 2, 58, 14, 3, 2, 2, 2, 59, 60, 7, 88, 2, 2, 60, 16,
	3, 2, 2, 2, 61, 62, 7, 46, 2, 2, 62, 18, 3, 2, 2, 2, 63, 65, 7, 47, 2,
	2, 64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 68,
	9, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 20, 3, 2, 2, 2, 71, 75, 7, 36, 2, 2, 72, 74, 11,
	2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 75,
	73, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 36,
	2, 2, 79, 22, 3, 2, 2, 2, 80, 82, 9, 3, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83,
	3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2,
	85, 86, 8, 12, 2, 2, 86, 24, 3, 2, 2, 2, 7, 2, 64, 69, 75, 83, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'TABLE'", "'VECTORS'", "'TUPLES'", "'DATA'", "'BOT'", "'EOD'", "'V'",
	"','",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "NUM", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "NUM",
	"STRING", "WS",
}

type difLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewdifLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *difLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdifLexer(input antlr.CharStream) *difLexer {
	l := new(difLexer)
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
	l.GrammarFileName = "dif.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// difLexer tokens.
const (
	difLexerT__0   = 1
	difLexerT__1   = 2
	difLexerT__2   = 3
	difLexerT__3   = 4
	difLexerT__4   = 5
	difLexerT__5   = 6
	difLexerT__6   = 7
	difLexerT__7   = 8
	difLexerNUM    = 9
	difLexerSTRING = 10
	difLexerWS     = 11
)
