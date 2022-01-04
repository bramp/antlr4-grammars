// Code generated from PMMN.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pmmn

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 13, 6, 13, 63, 10, 13, 13, 13, 14, 13, 64, 3, 13, 3, 13,
	2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 68, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 35, 3, 2,
	2, 2, 11, 37, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 41, 3, 2, 2, 2, 17, 44,
	3, 2, 2, 2, 19, 49, 3, 2, 2, 2, 21, 55, 3, 2, 2, 2, 23, 59, 3, 2, 2, 2,
	25, 62, 3, 2, 2, 2, 27, 28, 7, 125, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30, 7,
	127, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 107, 2, 2, 32, 33, 7, 112, 2,
	2, 33, 34, 7, 101, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 42, 2, 2, 36, 10,
	3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 12, 3, 2, 2, 2, 39, 40, 7, 61, 2,
	2, 40, 14, 3, 2, 2, 2, 41, 42, 7, 107, 2, 2, 42, 43, 7, 104, 2, 2, 43,
	16, 3, 2, 2, 2, 44, 45, 7, 103, 2, 2, 45, 46, 7, 110, 2, 2, 46, 47, 7,
	117, 2, 2, 47, 48, 7, 103, 2, 2, 48, 18, 3, 2, 2, 2, 49, 50, 7, 121, 2,
	2, 50, 51, 7, 106, 2, 2, 51, 52, 7, 107, 2, 2, 52, 53, 7, 110, 2, 2, 53,
	54, 7, 103, 2, 2, 54, 20, 3, 2, 2, 2, 55, 56, 7, 102, 2, 2, 56, 57, 7,
	103, 2, 2, 57, 58, 7, 101, 2, 2, 58, 22, 3, 2, 2, 2, 59, 60, 4, 50, 59,
	2, 60, 24, 3, 2, 2, 2, 61, 63, 9, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2,
	66, 67, 8, 13, 2, 2, 67, 26, 3, 2, 2, 2, 4, 2, 64, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'", "'inc'", "'('", "')'", "';'", "'if'", "'else'", "'while'",
	"'dec'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "DIGIT", "WS",
}

type PMMNLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPMMNLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PMMNLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPMMNLexer(input antlr.CharStream) *PMMNLexer {
	l := new(PMMNLexer)
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
	l.GrammarFileName = "PMMN.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PMMNLexer tokens.
const (
	PMMNLexerT__0  = 1
	PMMNLexerT__1  = 2
	PMMNLexerT__2  = 3
	PMMNLexerT__3  = 4
	PMMNLexerT__4  = 5
	PMMNLexerT__5  = 6
	PMMNLexerT__6  = 7
	PMMNLexerT__7  = 8
	PMMNLexerT__8  = 9
	PMMNLexerT__9  = 10
	PMMNLexerDIGIT = 11
	PMMNLexerWS    = 12
)
