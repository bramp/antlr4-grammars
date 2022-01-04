// Code generated from tl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 47, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 4, 27, 10, 4, 13, 4, 14, 4, 28, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 42, 10, 10, 13, 10, 14, 10, 43, 3,
	10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 3, 2, 4, 3, 2, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 48, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 26, 3, 2, 2,
	2, 9, 30, 3, 2, 2, 2, 11, 32, 3, 2, 2, 2, 13, 34, 3, 2, 2, 2, 15, 36, 3,
	2, 2, 2, 17, 38, 3, 2, 2, 2, 19, 41, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22,
	4, 3, 2, 2, 2, 23, 24, 7, 43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 27, 9, 2, 2,
	2, 26, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29,
	3, 2, 2, 2, 29, 8, 3, 2, 2, 2, 30, 31, 7, 73, 2, 2, 31, 10, 3, 2, 2, 2,
	32, 33, 7, 74, 2, 2, 33, 12, 3, 2, 2, 2, 34, 35, 7, 8746, 2, 2, 35, 14,
	3, 2, 2, 2, 36, 37, 7, 8871, 2, 2, 37, 16, 3, 2, 2, 2, 38, 39, 7, 8978,
	2, 2, 39, 18, 3, 2, 2, 2, 40, 42, 9, 3, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2,
	45, 46, 8, 10, 2, 2, 46, 20, 3, 2, 2, 2, 5, 2, 28, 43, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "", "'G'", "'H'", "'\u2228'", "'\u22A5'", "'\u2310'",
}

var lexerSymbolicNames = []string{
	"", "", "", "ATOMIC", "TL_ALWAYS", "TL_WAS", "TL_OR", "TL_UPTACK", "TL_NOT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "ATOMIC", "TL_ALWAYS", "TL_WAS", "TL_OR", "TL_UPTACK",
	"TL_NOT", "WS",
}

type tlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewtlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *tlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtlLexer(input antlr.CharStream) *tlLexer {
	l := new(tlLexer)
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
	l.GrammarFileName = "tl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tlLexer tokens.
const (
	tlLexerT__0      = 1
	tlLexerT__1      = 2
	tlLexerATOMIC    = 3
	tlLexerTL_ALWAYS = 4
	tlLexerTL_WAS    = 5
	tlLexerTL_OR     = 6
	tlLexerTL_UPTACK = 7
	tlLexerTL_NOT    = 8
	tlLexerWS        = 9
)
