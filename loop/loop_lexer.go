// Code generated from loop.g4 by ANTLR 4.9.3. DO NOT EDIT.

package loop

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 79, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 49,
	10, 9, 12, 9, 14, 9, 52, 11, 9, 3, 10, 6, 10, 55, 10, 10, 13, 10, 14, 10,
	56, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 63, 10, 11, 12, 11, 14, 11, 66,
	11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 6, 12, 74, 10, 12, 13,
	12, 14, 12, 75, 3, 12, 3, 12, 3, 64, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 6, 4, 2, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15,
	15, 34, 34, 2, 82, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 30, 3, 2, 2, 2, 9, 32, 3,
	2, 2, 2, 11, 34, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 42, 3, 2, 2, 2, 17,
	46, 3, 2, 2, 2, 19, 54, 3, 2, 2, 2, 21, 58, 3, 2, 2, 2, 23, 73, 3, 2, 2,
	2, 25, 26, 7, 61, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 60, 2, 2, 28, 29,
	7, 63, 2, 2, 29, 6, 3, 2, 2, 2, 30, 31, 7, 45, 2, 2, 31, 8, 3, 2, 2, 2,
	32, 33, 7, 47, 2, 2, 33, 10, 3, 2, 2, 2, 34, 35, 7, 78, 2, 2, 35, 36, 7,
	81, 2, 2, 36, 37, 7, 81, 2, 2, 37, 38, 7, 82, 2, 2, 38, 12, 3, 2, 2, 2,
	39, 40, 7, 70, 2, 2, 40, 41, 7, 81, 2, 2, 41, 14, 3, 2, 2, 2, 42, 43, 7,
	71, 2, 2, 43, 44, 7, 80, 2, 2, 44, 45, 7, 70, 2, 2, 45, 16, 3, 2, 2, 2,
	46, 50, 9, 2, 2, 2, 47, 49, 9, 3, 2, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 18, 3, 2, 2, 2, 52,
	50, 3, 2, 2, 2, 53, 55, 9, 4, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2,
	2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 20, 3, 2, 2, 2, 58, 59,
	7, 49, 2, 2, 59, 60, 7, 44, 2, 2, 60, 64, 3, 2, 2, 2, 61, 63, 11, 2, 2,
	2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 68, 7, 44, 2, 2,
	68, 69, 7, 49, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 8, 11, 2, 2, 71, 22, 3,
	2, 2, 2, 72, 74, 9, 5, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 8, 12,
	2, 2, 78, 24, 3, 2, 2, 2, 7, 2, 50, 56, 64, 75, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "':='", "'+'", "'-'", "'LOOP'", "'DO'", "'END'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ID", "NUMBER", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "NUMBER",
	"COMMENT", "WS",
}

type loopLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewloopLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *loopLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewloopLexer(input antlr.CharStream) *loopLexer {
	l := new(loopLexer)
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
	l.GrammarFileName = "loop.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// loopLexer tokens.
const (
	loopLexerT__0    = 1
	loopLexerT__1    = 2
	loopLexerT__2    = 3
	loopLexerT__3    = 4
	loopLexerT__4    = 5
	loopLexerT__5    = 6
	loopLexerT__6    = 7
	loopLexerID      = 8
	loopLexerNUMBER  = 9
	loopLexerCOMMENT = 10
	loopLexerWS      = 11
)
