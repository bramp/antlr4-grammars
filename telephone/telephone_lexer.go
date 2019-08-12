// Code generated from telephone.g4 by ANTLR 4.7.2. DO NOT EDIT.

package telephone

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 37, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 6, 7, 32, 10, 7, 13, 7, 14, 7, 33, 3, 7, 3,
	7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 4, 3, 2, 50, 59,
	5, 2, 12, 12, 15, 15, 34, 34, 2, 37, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 3, 15, 3, 2, 2, 2, 5, 18, 3, 2, 2, 2, 7, 20, 3, 2, 2, 2, 9, 24, 3, 2,
	2, 2, 11, 28, 3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 16, 7, 45, 2, 2, 16,
	17, 7, 51, 2, 2, 17, 4, 3, 2, 2, 2, 18, 19, 7, 45, 2, 2, 19, 6, 3, 2, 2,
	2, 20, 21, 7, 50, 2, 2, 21, 22, 7, 51, 2, 2, 22, 23, 7, 51, 2, 2, 23, 8,
	3, 2, 2, 2, 24, 25, 7, 50, 2, 2, 25, 26, 7, 51, 2, 2, 26, 27, 7, 50, 2,
	2, 27, 10, 3, 2, 2, 2, 28, 29, 9, 2, 2, 2, 29, 12, 3, 2, 2, 2, 30, 32,
	9, 3, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	33, 34, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 8, 7, 2, 2, 36, 14, 3,
	2, 2, 2, 4, 2, 33, 3, 8, 2, 2,
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
	"", "'+1'", "'+'", "'011'", "'010'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "DIGIT", "WS",
}

type telephoneLexer struct {
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

func NewtelephoneLexer(input antlr.CharStream) *telephoneLexer {

	l := new(telephoneLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "telephone.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// telephoneLexer tokens.
const (
	telephoneLexerT__0  = 1
	telephoneLexerT__1  = 2
	telephoneLexerT__2  = 3
	telephoneLexerT__3  = 4
	telephoneLexerDIGIT = 5
	telephoneLexerWS    = 6
)
