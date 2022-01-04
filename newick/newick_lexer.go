// Code generated from newick.g4 by ANTLR 4.9.3. DO NOT EDIT.

package newick

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 56, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 6, 7, 31, 10, 7, 13, 7, 14, 7, 32, 3, 7, 3, 7, 6,
	7, 37, 10, 7, 13, 7, 14, 7, 38, 5, 7, 41, 10, 7, 3, 8, 3, 8, 7, 8, 45,
	10, 8, 12, 8, 14, 8, 48, 11, 8, 3, 9, 6, 9, 51, 10, 9, 13, 9, 14, 9, 52,
	3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 3, 2, 6, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92,
	99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 60, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5,
	21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 27, 3, 2, 2,
	2, 13, 30, 3, 2, 2, 2, 15, 42, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 20,
	7, 61, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 6, 3, 2, 2, 2,
	23, 24, 7, 43, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 46, 2, 2, 26, 10, 3,
	2, 2, 2, 27, 28, 7, 60, 2, 2, 28, 12, 3, 2, 2, 2, 29, 31, 9, 2, 2, 2, 30,
	29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2,
	2, 33, 40, 3, 2, 2, 2, 34, 36, 7, 48, 2, 2, 35, 37, 9, 2, 2, 2, 36, 35,
	3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2,
	39, 41, 3, 2, 2, 2, 40, 34, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 14, 3,
	2, 2, 2, 42, 46, 9, 3, 2, 2, 43, 45, 9, 4, 2, 2, 44, 43, 3, 2, 2, 2, 45,
	48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 16, 3, 2, 2,
	2, 48, 46, 3, 2, 2, 2, 49, 51, 9, 5, 2, 2, 50, 49, 3, 2, 2, 2, 51, 52,
	3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2,
	54, 55, 8, 9, 2, 2, 55, 18, 3, 2, 2, 2, 8, 2, 32, 38, 40, 46, 52, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'('", "')'", "','", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "NUMBER", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "NUMBER", "STRING", "WS",
}

type newickLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewnewickLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *newickLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnewickLexer(input antlr.CharStream) *newickLexer {
	l := new(newickLexer)
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
	l.GrammarFileName = "newick.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// newickLexer tokens.
const (
	newickLexerT__0   = 1
	newickLexerT__1   = 2
	newickLexerT__2   = 3
	newickLexerT__3   = 4
	newickLexerT__4   = 5
	newickLexerNUMBER = 6
	newickLexerSTRING = 7
	newickLexerWS     = 8
)
