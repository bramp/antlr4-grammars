// Code generated from p.g4 by ANTLR 4.7.2. DO NOT EDIT.

package p

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 25, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 2, 2, 7,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	24, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15, 3, 2, 2, 2, 7, 17, 3, 2,
	2, 2, 9, 19, 3, 2, 2, 2, 11, 21, 3, 2, 2, 2, 13, 14, 7, 42, 2, 2, 14, 4,
	3, 2, 2, 2, 15, 16, 7, 43, 2, 2, 16, 6, 3, 2, 2, 2, 17, 18, 7, 84, 2, 2,
	18, 8, 3, 2, 2, 2, 19, 20, 7, 957, 2, 2, 20, 10, 3, 2, 2, 2, 21, 22, 9,
	2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 24, 8, 6, 2, 2, 24, 12, 3, 2, 2, 2, 3,
	2, 3, 8, 2, 2,
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
	"", "'('", "')'", "'R'", "'\u03BB'",
}

var lexerSymbolicNames = []string{
	"", "", "", "R", "L", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "R", "L", "WS",
}

type pLexer struct {
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

func NewpLexer(input antlr.CharStream) *pLexer {

	l := new(pLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "p.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// pLexer tokens.
const (
	pLexerT__0 = 1
	pLexerT__1 = 2
	pLexerR    = 3
	pLexerL    = 4
	pLexerWS   = 5
)
