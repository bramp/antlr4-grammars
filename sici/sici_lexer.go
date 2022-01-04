// Code generated from sici.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sici

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 56, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 6, 13, 51, 10, 13, 13, 13, 14, 13, 52, 3, 13, 3, 13, 2, 2,
	14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 3, 2, 5, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 56, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2,
	7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 37, 3, 2,
	2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 43, 3, 2, 2, 2, 21, 45,
	3, 2, 2, 2, 23, 47, 3, 2, 2, 2, 25, 50, 3, 2, 2, 2, 27, 28, 7, 47, 2, 2,
	28, 4, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 49,
	2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 43, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36,
	7, 60, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 7, 62, 2, 2, 38, 14, 3, 2, 2,
	2, 39, 40, 7, 64, 2, 2, 40, 16, 3, 2, 2, 2, 41, 42, 7, 48, 2, 2, 42, 18,
	3, 2, 2, 2, 43, 44, 7, 61, 2, 2, 44, 20, 3, 2, 2, 2, 45, 46, 9, 2, 2, 2,
	46, 22, 3, 2, 2, 2, 47, 48, 9, 3, 2, 2, 48, 24, 3, 2, 2, 2, 49, 51, 9,
	4, 2, 2, 50, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 8, 13, 2, 2, 55, 26, 3, 2,
	2, 2, 4, 2, 52, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'-'", "'('", "'/'", "')'", "':'", "'<'", "'>'", "'.'", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "LETTER", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"LETTER", "DIGIT", "WS",
}

type siciLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewsiciLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *siciLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsiciLexer(input antlr.CharStream) *siciLexer {
	l := new(siciLexer)
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
	l.GrammarFileName = "sici.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// siciLexer tokens.
const (
	siciLexerT__0   = 1
	siciLexerT__1   = 2
	siciLexerT__2   = 3
	siciLexerT__3   = 4
	siciLexerT__4   = 5
	siciLexerT__5   = 6
	siciLexerT__6   = 7
	siciLexerT__7   = 8
	siciLexerT__8   = 9
	siciLexerLETTER = 10
	siciLexerDIGIT  = 11
	siciLexerWS     = 12
)