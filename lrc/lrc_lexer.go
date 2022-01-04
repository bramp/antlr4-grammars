// Code generated from lrcLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lrc

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 55, 8,
	1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4,
	7, 9, 7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 32, 10, 5, 13, 5, 14, 5, 33, 3, 6, 6,
	6, 37, 10, 6, 13, 6, 14, 6, 38, 3, 6, 3, 6, 3, 7, 6, 7, 44, 10, 7, 13,
	7, 14, 7, 45, 3, 8, 6, 8, 49, 10, 8, 13, 8, 14, 8, 50, 3, 8, 3, 8, 3, 8,
	2, 2, 9, 4, 3, 6, 4, 8, 5, 10, 6, 12, 7, 14, 8, 16, 9, 4, 2, 3, 5, 3, 2,
	50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 57, 2, 4,
	3, 2, 2, 2, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 12,
	3, 2, 2, 2, 3, 14, 3, 2, 2, 2, 3, 16, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6,
	20, 3, 2, 2, 2, 8, 24, 3, 2, 2, 2, 10, 31, 3, 2, 2, 2, 12, 36, 3, 2, 2,
	2, 14, 43, 3, 2, 2, 2, 16, 48, 3, 2, 2, 2, 18, 19, 7, 93, 2, 2, 19, 5,
	3, 2, 2, 2, 20, 21, 7, 95, 2, 2, 21, 22, 3, 2, 2, 2, 22, 23, 8, 3, 2, 2,
	23, 7, 3, 2, 2, 2, 24, 25, 5, 10, 5, 2, 25, 26, 7, 60, 2, 2, 26, 27, 5,
	10, 5, 2, 27, 28, 7, 48, 2, 2, 28, 29, 5, 10, 5, 2, 29, 9, 3, 2, 2, 2,
	30, 32, 9, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3,
	2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 11, 3, 2, 2, 2, 35, 37, 9, 3, 2, 2, 36,
	35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2,
	2, 39, 40, 3, 2, 2, 2, 40, 41, 8, 6, 3, 2, 41, 13, 3, 2, 2, 2, 42, 44,
	10, 4, 2, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2,
	45, 46, 3, 2, 2, 2, 46, 15, 3, 2, 2, 2, 47, 49, 9, 4, 2, 2, 48, 47, 3,
	2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 53, 8, 8, 3, 2, 53, 54, 8, 8, 4, 2, 54, 17, 3, 2, 2,
	2, 8, 2, 3, 33, 38, 45, 50, 5, 7, 3, 2, 8, 2, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "LINETEXT",
}

var lexerLiteralNames = []string{
	"", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "LB", "RB", "TIMESTAMP", "NUM", "WS", "TEXT", "EOL",
}

var lexerRuleNames = []string{
	"LB", "RB", "TIMESTAMP", "NUM", "WS", "TEXT", "EOL",
}

type lrcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewlrcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *lrcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlrcLexer(input antlr.CharStream) *lrcLexer {
	l := new(lrcLexer)
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
	l.GrammarFileName = "lrcLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lrcLexer tokens.
const (
	lrcLexerLB        = 1
	lrcLexerRB        = 2
	lrcLexerTIMESTAMP = 3
	lrcLexerNUM       = 4
	lrcLexerWS        = 5
	lrcLexerTEXT      = 6
	lrcLexerEOL       = 7
)

// lrcLexerLINETEXT is the lrcLexer mode.
const lrcLexerLINETEXT = 1
