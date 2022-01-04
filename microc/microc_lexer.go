// Code generated from microc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package microc

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 79, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 67, 10, 14, 13, 14,
	14, 14, 68, 3, 15, 6, 15, 72, 10, 15, 13, 15, 14, 15, 73, 3, 16, 3, 16,
	3, 16, 3, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 5, 3,
	2, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 80, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2,
	5, 36, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 47, 3, 2, 2, 2, 11, 49, 3, 2,
	2, 2, 13, 51, 3, 2, 2, 2, 15, 53, 3, 2, 2, 2, 17, 55, 3, 2, 2, 2, 19, 57,
	3, 2, 2, 2, 21, 59, 3, 2, 2, 2, 23, 61, 3, 2, 2, 2, 25, 63, 3, 2, 2, 2,
	27, 66, 3, 2, 2, 2, 29, 71, 3, 2, 2, 2, 31, 75, 3, 2, 2, 2, 33, 34, 7,
	107, 2, 2, 34, 35, 7, 104, 2, 2, 35, 4, 3, 2, 2, 2, 36, 37, 7, 103, 2,
	2, 37, 38, 7, 110, 2, 2, 38, 39, 7, 117, 2, 2, 39, 40, 7, 103, 2, 2, 40,
	6, 3, 2, 2, 2, 41, 42, 7, 121, 2, 2, 42, 43, 7, 106, 2, 2, 43, 44, 7, 107,
	2, 2, 44, 45, 7, 110, 2, 2, 45, 46, 7, 103, 2, 2, 46, 8, 3, 2, 2, 2, 47,
	48, 7, 125, 2, 2, 48, 10, 3, 2, 2, 2, 49, 50, 7, 127, 2, 2, 50, 12, 3,
	2, 2, 2, 51, 52, 7, 61, 2, 2, 52, 14, 3, 2, 2, 2, 53, 54, 7, 42, 2, 2,
	54, 16, 3, 2, 2, 2, 55, 56, 7, 43, 2, 2, 56, 18, 3, 2, 2, 2, 57, 58, 7,
	63, 2, 2, 58, 20, 3, 2, 2, 2, 59, 60, 7, 62, 2, 2, 60, 22, 3, 2, 2, 2,
	61, 62, 7, 45, 2, 2, 62, 24, 3, 2, 2, 2, 63, 64, 7, 47, 2, 2, 64, 26, 3,
	2, 2, 2, 65, 67, 9, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68,
	66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 28, 3, 2, 2, 2, 70, 72, 9, 3, 2,
	2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 30, 3, 2, 2, 2, 75, 76, 9, 4, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 78, 8, 16, 2, 2, 78, 32, 3, 2, 2, 2, 5, 2, 68, 73, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'if'", "'else'", "'while'", "'{'", "'}'", "';'", "'('", "')'", "'='",
	"'<'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "STRING", "INT", "WS",
}

type microcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmicrocLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *microcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmicrocLexer(input antlr.CharStream) *microcLexer {
	l := new(microcLexer)
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
	l.GrammarFileName = "microc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// microcLexer tokens.
const (
	microcLexerT__0   = 1
	microcLexerT__1   = 2
	microcLexerT__2   = 3
	microcLexerT__3   = 4
	microcLexerT__4   = 5
	microcLexerT__5   = 6
	microcLexerT__6   = 7
	microcLexerT__7   = 8
	microcLexerT__8   = 9
	microcLexerT__9   = 10
	microcLexerT__10  = 11
	microcLexerT__11  = 12
	microcLexerSTRING = 13
	microcLexerINT    = 14
	microcLexerWS     = 15
)
