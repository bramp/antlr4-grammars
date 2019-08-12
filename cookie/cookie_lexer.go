// Code generated from cookie.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cookie

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 39, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 22, 10, 4, 12, 4, 14, 4, 25,
	11, 4, 3, 4, 3, 4, 3, 5, 6, 5, 30, 10, 5, 13, 5, 14, 5, 31, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	3, 2, 5, 4, 2, 12, 12, 36, 36, 8, 2, 34, 34, 46, 47, 49, 60, 67, 92, 97,
	97, 99, 124, 4, 2, 11, 12, 15, 15, 2, 40, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 3, 15, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 19, 3, 2, 2, 2, 9, 29,
	3, 2, 2, 2, 11, 33, 3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 16, 7, 61, 2, 2,
	16, 4, 3, 2, 2, 2, 17, 18, 7, 63, 2, 2, 18, 6, 3, 2, 2, 2, 19, 23, 7, 36,
	2, 2, 20, 22, 10, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23,
	21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 23, 3, 2, 2,
	2, 26, 27, 7, 36, 2, 2, 27, 8, 3, 2, 2, 2, 28, 30, 9, 3, 2, 2, 29, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 10, 3, 2, 2, 2, 33, 34, 4, 50, 59, 2, 34, 12, 3, 2, 2, 2, 35, 36, 9,
	4, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 8, 7, 2, 2, 38, 14, 3, 2, 2, 2, 5,
	2, 23, 31, 3, 8, 2, 2,
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
	"", "';'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "STRING", "TOKEN", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "STRING", "TOKEN", "DIGIT", "WS",
}

type cookieLexer struct {
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

func NewcookieLexer(input antlr.CharStream) *cookieLexer {

	l := new(cookieLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "cookie.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cookieLexer tokens.
const (
	cookieLexerT__0   = 1
	cookieLexerT__1   = 2
	cookieLexerSTRING = 3
	cookieLexerTOKEN  = 4
	cookieLexerDIGIT  = 5
	cookieLexerWS     = 6
)
