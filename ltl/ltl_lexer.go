// Code generated from ltl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ltl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 82, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 6, 6, 52, 10, 6, 13, 6, 14, 6, 53, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 6, 17, 77, 10, 17, 13,
	17, 14, 17, 78, 3, 17, 3, 17, 2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 3, 2, 4, 3, 2, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 83,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 40, 3, 2, 2, 2, 7, 46, 3, 2, 2, 2, 9, 48,
	3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13, 55, 3, 2, 2, 2, 15, 57, 3, 2, 2, 2,
	17, 59, 3, 2, 2, 2, 19, 61, 3, 2, 2, 2, 21, 63, 3, 2, 2, 2, 23, 65, 3,
	2, 2, 2, 25, 67, 3, 2, 2, 2, 27, 69, 3, 2, 2, 2, 29, 71, 3, 2, 2, 2, 31,
	73, 3, 2, 2, 2, 33, 76, 3, 2, 2, 2, 35, 36, 7, 118, 2, 2, 36, 37, 7, 116,
	2, 2, 37, 38, 7, 119, 2, 2, 38, 39, 7, 103, 2, 2, 39, 4, 3, 2, 2, 2, 40,
	41, 7, 104, 2, 2, 41, 42, 7, 99, 2, 2, 42, 43, 7, 110, 2, 2, 43, 44, 7,
	117, 2, 2, 44, 45, 7, 103, 2, 2, 45, 6, 3, 2, 2, 2, 46, 47, 7, 42, 2, 2,
	47, 8, 3, 2, 2, 2, 48, 49, 7, 43, 2, 2, 49, 10, 3, 2, 2, 2, 50, 52, 9,
	2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53,
	54, 3, 2, 2, 2, 54, 12, 3, 2, 2, 2, 55, 56, 7, 87, 2, 2, 56, 14, 3, 2,
	2, 2, 57, 58, 7, 73, 2, 2, 58, 16, 3, 2, 2, 2, 59, 60, 7, 72, 2, 2, 60,
	18, 3, 2, 2, 2, 61, 62, 7, 90, 2, 2, 62, 20, 3, 2, 2, 2, 63, 64, 7, 89,
	2, 2, 64, 22, 3, 2, 2, 2, 65, 66, 7, 84, 2, 2, 66, 24, 3, 2, 2, 2, 67,
	68, 7, 8596, 2, 2, 68, 26, 3, 2, 2, 2, 69, 70, 7, 8745, 2, 2, 70, 28, 3,
	2, 2, 2, 71, 72, 7, 8746, 2, 2, 72, 30, 3, 2, 2, 2, 73, 74, 7, 8978, 2,
	2, 74, 32, 3, 2, 2, 2, 75, 77, 9, 3, 2, 2, 76, 75, 3, 2, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 81, 8, 17, 2, 2, 81, 34, 3, 2, 2, 2, 5, 2, 53, 78, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'true'", "'false'", "'('", "')'", "", "'U'", "'G'", "'F'", "'X'",
	"'W'", "'R'", "'\u2192'", "'\u2227'", "'\u2228'", "'\u2310'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "ATOMIC", "LTL_UNTIL", "LTL_GLOBALLY", "LTL_FINALLY",
	"LTL_NEXT", "LTL_WEAK", "LTL_RELEASE", "LTL_RIGHTWARDS_SINGLE_ARROW", "LTL_AND",
	"LTL_OR", "LTL_NOT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "ATOMIC", "LTL_UNTIL", "LTL_GLOBALLY",
	"LTL_FINALLY", "LTL_NEXT", "LTL_WEAK", "LTL_RELEASE", "LTL_RIGHTWARDS_SINGLE_ARROW",
	"LTL_AND", "LTL_OR", "LTL_NOT", "WS",
}

type ltlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewltlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ltlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewltlLexer(input antlr.CharStream) *ltlLexer {
	l := new(ltlLexer)
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
	l.GrammarFileName = "ltl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ltlLexer tokens.
const (
	ltlLexerT__0                        = 1
	ltlLexerT__1                        = 2
	ltlLexerT__2                        = 3
	ltlLexerT__3                        = 4
	ltlLexerATOMIC                      = 5
	ltlLexerLTL_UNTIL                   = 6
	ltlLexerLTL_GLOBALLY                = 7
	ltlLexerLTL_FINALLY                 = 8
	ltlLexerLTL_NEXT                    = 9
	ltlLexerLTL_WEAK                    = 10
	ltlLexerLTL_RELEASE                 = 11
	ltlLexerLTL_RIGHTWARDS_SINGLE_ARROW = 12
	ltlLexerLTL_AND                     = 13
	ltlLexerLTL_OR                      = 14
	ltlLexerLTL_NOT                     = 15
	ltlLexerWS                          = 16
)
