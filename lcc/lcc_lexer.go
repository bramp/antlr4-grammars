// Code generated from lcc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package lcc

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 28, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 23, 10, 6, 13, 6, 14,
	6, 24, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 3, 4,
	2, 11, 12, 15, 15, 2, 28, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15,
	3, 2, 2, 2, 7, 17, 3, 2, 2, 2, 9, 19, 3, 2, 2, 2, 11, 22, 3, 2, 2, 2, 13,
	14, 7, 34, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 48, 2, 2, 16, 6, 3, 2, 2,
	2, 17, 18, 4, 50, 59, 2, 18, 8, 3, 2, 2, 2, 19, 20, 4, 67, 92, 2, 20, 10,
	3, 2, 2, 2, 21, 23, 9, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2,
	24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27, 8,
	6, 2, 2, 27, 12, 3, 2, 2, 2, 4, 2, 24, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "' '", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "DIGIT", "LETTER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "DIGIT", "LETTER", "WS",
}

type lccLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewlccLexer(input antlr.CharStream) *lccLexer {

	l := new(lccLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "lcc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lccLexer tokens.
const (
	lccLexerT__0   = 1
	lccLexerT__1   = 2
	lccLexerDIGIT  = 3
	lccLexerLETTER = 4
	lccLexerWS     = 5
)
