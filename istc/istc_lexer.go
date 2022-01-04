// Code generated from istc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package istc

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 31, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 26,
	10, 6, 13, 6, 14, 6, 27, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 3, 2, 5, 4, 2, 34, 34, 47, 47, 4, 2, 50, 59, 67, 72, 4, 2, 11, 12, 15,
	15, 2, 31, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 18, 3, 2, 2, 2, 7,
	20, 3, 2, 2, 2, 9, 22, 3, 2, 2, 2, 11, 25, 3, 2, 2, 2, 13, 14, 7, 75, 2,
	2, 14, 15, 7, 85, 2, 2, 15, 16, 7, 86, 2, 2, 16, 17, 7, 69, 2, 2, 17, 4,
	3, 2, 2, 2, 18, 19, 7, 34, 2, 2, 19, 6, 3, 2, 2, 2, 20, 21, 9, 2, 2, 2,
	21, 8, 3, 2, 2, 2, 22, 23, 9, 3, 2, 2, 23, 10, 3, 2, 2, 2, 24, 26, 9, 4,
	2, 2, 25, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28,
	3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 30, 8, 6, 2, 2, 30, 12, 3, 2, 2, 2,
	4, 2, 27, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'ISTC'", "' '",
}

var lexerSymbolicNames = []string{
	"", "", "", "SEP", "CHAR", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "SEP", "CHAR", "WS",
}

type istcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewistcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *istcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewistcLexer(input antlr.CharStream) *istcLexer {
	l := new(istcLexer)
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
	l.GrammarFileName = "istc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// istcLexer tokens.
const (
	istcLexerT__0 = 1
	istcLexerT__1 = 2
	istcLexerSEP  = 3
	istcLexerCHAR = 4
	istcLexerWS   = 5
)
