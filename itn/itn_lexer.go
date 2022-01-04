// Code generated from itn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package itn

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 30, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 6,
	3, 15, 10, 3, 13, 3, 14, 3, 16, 3, 4, 6, 4, 20, 10, 4, 13, 4, 14, 4, 21,
	3, 5, 6, 5, 25, 10, 5, 13, 5, 14, 5, 26, 3, 5, 3, 5, 2, 2, 6, 3, 3, 5,
	4, 7, 5, 9, 6, 3, 2, 5, 3, 2, 50, 59, 6, 2, 34, 34, 50, 59, 67, 92, 99,
	124, 4, 2, 11, 12, 15, 15, 2, 32, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 14, 3, 2, 2, 2,
	7, 19, 3, 2, 2, 2, 9, 24, 3, 2, 2, 2, 11, 12, 7, 126, 2, 2, 12, 4, 3, 2,
	2, 2, 13, 15, 9, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 14,
	3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 6, 3, 2, 2, 2, 18, 20, 9, 3, 2, 2,
	19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3,
	2, 2, 2, 22, 8, 3, 2, 2, 2, 23, 25, 9, 4, 2, 2, 24, 23, 3, 2, 2, 2, 25,
	26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2,
	2, 28, 29, 8, 5, 2, 2, 29, 10, 3, 2, 2, 2, 6, 2, 16, 21, 26, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'|'",
}

var lexerSymbolicNames = []string{
	"", "", "NUM", "TEXT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "NUM", "TEXT", "WS",
}

type itnLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewitnLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *itnLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewitnLexer(input antlr.CharStream) *itnLexer {
	l := new(itnLexer)
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
	l.GrammarFileName = "itn.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// itnLexer tokens.
const (
	itnLexerT__0 = 1
	itnLexerNUM  = 2
	itnLexerTEXT = 3
	itnLexerWS   = 4
)
