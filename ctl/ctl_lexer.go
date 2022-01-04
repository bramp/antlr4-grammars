// Code generated from ctl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ctl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 91, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 53, 10, 6, 13, 6, 14, 6, 54, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 86, 10, 21, 13, 21,
	14, 21, 87, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 4, 3, 2, 99, 124, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 92, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45, 3, 2,
	2, 2, 7, 47, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 52, 3, 2, 2, 2, 13, 56,
	3, 2, 2, 2, 15, 58, 3, 2, 2, 2, 17, 60, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2,
	21, 64, 3, 2, 2, 2, 23, 66, 3, 2, 2, 2, 25, 68, 3, 2, 2, 2, 27, 70, 3,
	2, 2, 2, 29, 72, 3, 2, 2, 2, 31, 74, 3, 2, 2, 2, 33, 76, 3, 2, 2, 2, 35,
	78, 3, 2, 2, 2, 37, 80, 3, 2, 2, 2, 39, 82, 3, 2, 2, 2, 41, 85, 3, 2, 2,
	2, 43, 44, 7, 93, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 95, 2, 2, 46, 6,
	3, 2, 2, 2, 47, 48, 7, 42, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7, 43, 2, 2,
	50, 10, 3, 2, 2, 2, 51, 53, 9, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 12, 3, 2, 2, 2, 56,
	57, 7, 87, 2, 2, 57, 14, 3, 2, 2, 2, 58, 59, 7, 73, 2, 2, 59, 16, 3, 2,
	2, 2, 60, 61, 7, 72, 2, 2, 61, 18, 3, 2, 2, 2, 62, 63, 7, 90, 2, 2, 63,
	20, 3, 2, 2, 2, 64, 65, 7, 67, 2, 2, 65, 22, 3, 2, 2, 2, 66, 67, 7, 71,
	2, 2, 67, 24, 3, 2, 2, 2, 68, 69, 7, 8711, 2, 2, 69, 26, 3, 2, 2, 2, 70,
	71, 7, 8662, 2, 2, 71, 28, 3, 2, 2, 2, 72, 73, 7, 8660, 2, 2, 73, 30, 3,
	2, 2, 2, 74, 75, 7, 8745, 2, 2, 75, 32, 3, 2, 2, 2, 76, 77, 7, 8746, 2,
	2, 77, 34, 3, 2, 2, 2, 78, 79, 7, 8870, 2, 2, 79, 36, 3, 2, 2, 2, 80, 81,
	7, 8871, 2, 2, 81, 38, 3, 2, 2, 2, 82, 83, 7, 8978, 2, 2, 83, 40, 3, 2,
	2, 2, 84, 86, 9, 3, 2, 2, 85, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85,
	3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 8, 21, 2, 2,
	90, 42, 3, 2, 2, 2, 5, 2, 54, 87, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'('", "')'", "", "'U'", "'G'", "'F'", "'X'", "'A'",
	"'E'", "'\u2205'", "'\u21D4'", "'\u21D2'", "'\u2227'", "'\u2228'", "'\u22A4'",
	"'\u22A5'", "'\u2310'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "ATOMIC", "CTL_UNTIL", "CTL_GLOBALLY", "CTL_FINALLY",
	"CTL_NEXT", "CTL_INEVITABLE", "CTL_EXISTS", "CTL_PROPOSITION", "CTL_LEFT_RIGHT_DOUBLE_ARROW",
	"CTL_RIGHTWARDS_DOUBLE_ARROW", "CTL_AND", "CTL_OR", "CTL_DOWNTACK", "CTL_UPTACK",
	"CTL_NOT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "ATOMIC", "CTL_UNTIL", "CTL_GLOBALLY",
	"CTL_FINALLY", "CTL_NEXT", "CTL_INEVITABLE", "CTL_EXISTS", "CTL_PROPOSITION",
	"CTL_LEFT_RIGHT_DOUBLE_ARROW", "CTL_RIGHTWARDS_DOUBLE_ARROW", "CTL_AND",
	"CTL_OR", "CTL_DOWNTACK", "CTL_UPTACK", "CTL_NOT", "WS",
}

type ctlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewctlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ctlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewctlLexer(input antlr.CharStream) *ctlLexer {
	l := new(ctlLexer)
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
	l.GrammarFileName = "ctl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ctlLexer tokens.
const (
	ctlLexerT__0                        = 1
	ctlLexerT__1                        = 2
	ctlLexerT__2                        = 3
	ctlLexerT__3                        = 4
	ctlLexerATOMIC                      = 5
	ctlLexerCTL_UNTIL                   = 6
	ctlLexerCTL_GLOBALLY                = 7
	ctlLexerCTL_FINALLY                 = 8
	ctlLexerCTL_NEXT                    = 9
	ctlLexerCTL_INEVITABLE              = 10
	ctlLexerCTL_EXISTS                  = 11
	ctlLexerCTL_PROPOSITION             = 12
	ctlLexerCTL_LEFT_RIGHT_DOUBLE_ARROW = 13
	ctlLexerCTL_RIGHTWARDS_DOUBLE_ARROW = 14
	ctlLexerCTL_AND                     = 15
	ctlLexerCTL_OR                      = 16
	ctlLexerCTL_DOWNTACK                = 17
	ctlLexerCTL_UPTACK                  = 18
	ctlLexerCTL_NOT                     = 19
	ctlLexerWS                          = 20
)
