// Code generated from morsecode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package morsecode

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 21, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 3, 3, 5, 4, 7, 5, 9, 6,
	3, 2, 3, 4, 2, 11, 12, 15, 15, 2, 20, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 13, 3, 2,
	2, 2, 7, 15, 3, 2, 2, 2, 9, 17, 3, 2, 2, 2, 11, 12, 7, 48, 2, 2, 12, 4,
	3, 2, 2, 2, 13, 14, 7, 47, 2, 2, 14, 6, 3, 2, 2, 2, 15, 16, 7, 34, 2, 2,
	16, 8, 3, 2, 2, 2, 17, 18, 9, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 20, 8, 5,
	2, 2, 20, 10, 3, 2, 2, 2, 3, 2, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'-'", "' '",
}

var lexerSymbolicNames = []string{
	"", "DOT", "DASH", "SPACE", "WS",
}

var lexerRuleNames = []string{
	"DOT", "DASH", "SPACE", "WS",
}

type morsecodeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmorsecodeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *morsecodeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmorsecodeLexer(input antlr.CharStream) *morsecodeLexer {
	l := new(morsecodeLexer)
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
	l.GrammarFileName = "morsecode.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// morsecodeLexer tokens.
const (
	morsecodeLexerDOT   = 1
	morsecodeLexerDASH  = 2
	morsecodeLexerSPACE = 3
	morsecodeLexerWS    = 4
)
