// Code generated from scotty.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scotty

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 55, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3,
	2, 5, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 54, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9,
	35, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 41, 3, 2, 2,
	2, 17, 43, 3, 2, 2, 2, 19, 45, 3, 2, 2, 2, 21, 47, 3, 2, 2, 2, 23, 49,
	3, 2, 2, 2, 25, 51, 3, 2, 2, 2, 27, 28, 7, 63, 2, 2, 28, 4, 3, 2, 2, 2,
	29, 30, 7, 104, 2, 2, 30, 31, 7, 119, 2, 2, 31, 32, 7, 112, 2, 2, 32, 6,
	3, 2, 2, 2, 33, 34, 7, 45, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 47, 2, 2,
	36, 10, 3, 2, 2, 2, 37, 38, 7, 44, 2, 2, 38, 12, 3, 2, 2, 2, 39, 40, 7,
	49, 2, 2, 40, 14, 3, 2, 2, 2, 41, 42, 7, 42, 2, 2, 42, 16, 3, 2, 2, 2,
	43, 44, 7, 43, 2, 2, 44, 18, 3, 2, 2, 2, 45, 46, 7, 48, 2, 2, 46, 20, 3,
	2, 2, 2, 47, 48, 9, 2, 2, 2, 48, 22, 3, 2, 2, 2, 49, 50, 9, 3, 2, 2, 50,
	24, 3, 2, 2, 2, 51, 52, 9, 4, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 8, 13,
	2, 2, 54, 26, 3, 2, 2, 2, 3, 2, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'fun'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "DIGIT", "LETTER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"DIGIT", "LETTER", "WS",
}

type scottyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewscottyLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *scottyLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewscottyLexer(input antlr.CharStream) *scottyLexer {
	l := new(scottyLexer)
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
	l.GrammarFileName = "scotty.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// scottyLexer tokens.
const (
	scottyLexerT__0   = 1
	scottyLexerT__1   = 2
	scottyLexerT__2   = 3
	scottyLexerT__3   = 4
	scottyLexerT__4   = 5
	scottyLexerT__5   = 6
	scottyLexerT__6   = 7
	scottyLexerT__7   = 8
	scottyLexerT__8   = 9
	scottyLexerDIGIT  = 10
	scottyLexerLETTER = 11
	scottyLexerWS     = 12
)
