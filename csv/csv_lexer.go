// Code generated from CSV.g4 by ANTLR 4.9.3. DO NOT EDIT.

package csv

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 35, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 21, 10, 5, 13, 5, 14, 5, 22, 3,
	6, 3, 6, 3, 6, 3, 6, 7, 6, 29, 10, 6, 12, 6, 14, 6, 32, 11, 6, 3, 6, 3,
	6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 4, 6, 2, 12, 12, 15, 15,
	36, 36, 46, 46, 3, 2, 36, 36, 2, 37, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2,
	2, 5, 15, 3, 2, 2, 2, 7, 17, 3, 2, 2, 2, 9, 20, 3, 2, 2, 2, 11, 24, 3,
	2, 2, 2, 13, 14, 7, 46, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 15, 2, 2, 16,
	6, 3, 2, 2, 2, 17, 18, 7, 12, 2, 2, 18, 8, 3, 2, 2, 2, 19, 21, 10, 2, 2,
	2, 20, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23,
	3, 2, 2, 2, 23, 10, 3, 2, 2, 2, 24, 30, 7, 36, 2, 2, 25, 26, 7, 36, 2,
	2, 26, 29, 7, 36, 2, 2, 27, 29, 10, 3, 2, 2, 28, 25, 3, 2, 2, 2, 28, 27,
	3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2,
	31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34, 7, 36, 2, 2, 34, 12, 3,
	2, 2, 2, 6, 2, 22, 28, 30, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'\r'", "'\n'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "TEXT", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "TEXT", "STRING",
}

type CSVLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCSVLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CSVLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCSVLexer(input antlr.CharStream) *CSVLexer {
	l := new(CSVLexer)
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
	l.GrammarFileName = "CSV.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CSVLexer tokens.
const (
	CSVLexerT__0   = 1
	CSVLexerT__1   = 2
	CSVLexerT__2   = 3
	CSVLexerTEXT   = 4
	CSVLexerSTRING = 5
)
