// Code generated from tnt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tnt

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 77, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 2, 2, 20, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 3, 2, 3, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 76, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2,
	2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53,
	3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 59, 3, 2, 2, 2,
	25, 61, 3, 2, 2, 2, 27, 63, 3, 2, 2, 2, 29, 65, 3, 2, 2, 2, 31, 67, 3,
	2, 2, 2, 33, 69, 3, 2, 2, 2, 35, 71, 3, 2, 2, 2, 37, 73, 3, 2, 2, 2, 39,
	40, 7, 63, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 45, 2, 2, 42, 6, 3, 2, 2,
	2, 43, 44, 7, 44, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 42, 2, 2, 46, 10,
	3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 128, 2,
	2, 50, 14, 3, 2, 2, 2, 51, 52, 7, 60, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54,
	7, 50, 2, 2, 54, 18, 3, 2, 2, 2, 55, 56, 7, 85, 2, 2, 56, 20, 3, 2, 2,
	2, 57, 58, 7, 99, 2, 2, 58, 22, 3, 2, 2, 2, 59, 60, 7, 100, 2, 2, 60, 24,
	3, 2, 2, 2, 61, 62, 7, 101, 2, 2, 62, 26, 3, 2, 2, 2, 63, 64, 7, 102, 2,
	2, 64, 28, 3, 2, 2, 2, 65, 66, 7, 103, 2, 2, 66, 30, 3, 2, 2, 2, 67, 68,
	7, 41, 2, 2, 68, 32, 3, 2, 2, 2, 69, 70, 7, 67, 2, 2, 70, 34, 3, 2, 2,
	2, 71, 72, 7, 71, 2, 2, 72, 36, 3, 2, 2, 2, 73, 74, 9, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 76, 8, 19, 2, 2, 76, 38, 3, 2, 2, 2, 3, 2, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'+'", "'*'", "'('", "')'", "'~'", "':'", "'0'", "'S'", "'a'",
	"'b'", "'c'", "'d'", "'e'", "'''", "'A'", "'E'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ZERO", "SUCCESSOR", "A", "B", "C", "D",
	"E", "PRIME", "FOREVERY", "EXISTS", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ZERO", "SUCCESSOR",
	"A", "B", "C", "D", "E", "PRIME", "FOREVERY", "EXISTS", "WS",
}

type tntLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewtntLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *tntLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtntLexer(input antlr.CharStream) *tntLexer {
	l := new(tntLexer)
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
	l.GrammarFileName = "tnt.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tntLexer tokens.
const (
	tntLexerT__0      = 1
	tntLexerT__1      = 2
	tntLexerT__2      = 3
	tntLexerT__3      = 4
	tntLexerT__4      = 5
	tntLexerT__5      = 6
	tntLexerT__6      = 7
	tntLexerZERO      = 8
	tntLexerSUCCESSOR = 9
	tntLexerA         = 10
	tntLexerB         = 11
	tntLexerC         = 12
	tntLexerD         = 13
	tntLexerE         = 14
	tntLexerPRIME     = 15
	tntLexerFOREVERY  = 16
	tntLexerEXISTS    = 17
	tntLexerWS        = 18
)
