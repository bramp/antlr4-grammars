// Code generated from bcl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 26, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 21, 10, 5, 13, 5, 14, 5, 22, 3,
	5, 3, 5, 2, 2, 6, 3, 3, 5, 4, 7, 5, 9, 6, 3, 2, 3, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 26, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 14, 3, 2, 2, 2, 7, 17, 3, 2, 2, 2,
	9, 20, 3, 2, 2, 2, 11, 12, 7, 50, 2, 2, 12, 13, 7, 50, 2, 2, 13, 4, 3,
	2, 2, 2, 14, 15, 7, 50, 2, 2, 15, 16, 7, 51, 2, 2, 16, 6, 3, 2, 2, 2, 17,
	18, 7, 51, 2, 2, 18, 8, 3, 2, 2, 2, 19, 21, 9, 2, 2, 2, 20, 19, 3, 2, 2,
	2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 24,
	3, 2, 2, 2, 24, 25, 8, 5, 2, 2, 25, 10, 3, 2, 2, 2, 4, 2, 22, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'00'", "'01'", "'1'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "WS",
}

type bclLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewbclLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *bclLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbclLexer(input antlr.CharStream) *bclLexer {
	l := new(bclLexer)
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
	l.GrammarFileName = "bcl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// bclLexer tokens.
const (
	bclLexerT__0 = 1
	bclLexerT__1 = 2
	bclLexerT__2 = 3
	bclLexerWS   = 4
)
