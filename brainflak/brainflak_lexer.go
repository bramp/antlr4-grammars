// Code generated from brainflak.g4 by ANTLR 4.9.3. DO NOT EDIT.

package brainflak

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 41, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 3, 2, 2, 2, 40, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2,
	2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 29,
	3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 35, 3, 2, 2, 2,
	19, 37, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7,
	43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 64, 2, 2, 26, 8, 3, 2, 2, 2, 27,
	28, 7, 62, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 125, 2, 2, 30, 12, 3, 2,
	2, 2, 31, 32, 7, 127, 2, 2, 32, 14, 3, 2, 2, 2, 33, 34, 7, 93, 2, 2, 34,
	16, 3, 2, 2, 2, 35, 36, 7, 95, 2, 2, 36, 18, 3, 2, 2, 2, 37, 38, 11, 2,
	2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 8, 10, 2, 2, 40, 20, 3, 2, 2, 2, 3, 2,
	3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'>'", "'<'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "GT", "LT", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"WS",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "GT", "LT", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"WS",
}

type brainflakLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewbrainflakLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *brainflakLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbrainflakLexer(input antlr.CharStream) *brainflakLexer {
	l := new(brainflakLexer)
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
	l.GrammarFileName = "brainflak.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// brainflakLexer tokens.
const (
	brainflakLexerLPAREN = 1
	brainflakLexerRPAREN = 2
	brainflakLexerGT     = 3
	brainflakLexerLT     = 4
	brainflakLexerLBRACE = 5
	brainflakLexerRBRACE = 6
	brainflakLexerLBRACK = 7
	brainflakLexerRBRACK = 8
	brainflakLexerWS     = 9
)
