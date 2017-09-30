// Generated from tiny.g4 by ANTLR 4.7.

package tiny

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 69, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 57, 10, 10, 13, 10, 14,
	10, 58, 3, 11, 6, 11, 62, 10, 11, 13, 11, 14, 11, 63, 3, 12, 3, 12, 3,
	12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 3, 2, 4, 4, 2, 67, 92, 99, 124, 5, 2, 12, 12, 15,
	15, 34, 34, 2, 70, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 3, 25, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 35, 3, 2, 2, 2, 9, 38, 3,
	2, 2, 2, 11, 43, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17,
	53, 3, 2, 2, 2, 19, 56, 3, 2, 2, 2, 21, 61, 3, 2, 2, 2, 23, 65, 3, 2, 2,
	2, 25, 26, 7, 68, 2, 2, 26, 27, 7, 71, 2, 2, 27, 28, 7, 73, 2, 2, 28, 29,
	7, 75, 2, 2, 29, 30, 7, 80, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 71, 2,
	2, 32, 33, 7, 80, 2, 2, 33, 34, 7, 70, 2, 2, 34, 6, 3, 2, 2, 2, 35, 36,
	7, 60, 2, 2, 36, 37, 7, 63, 2, 2, 37, 8, 3, 2, 2, 2, 38, 39, 7, 84, 2,
	2, 39, 40, 7, 71, 2, 2, 40, 41, 7, 67, 2, 2, 41, 42, 7, 70, 2, 2, 42, 10,
	3, 2, 2, 2, 43, 44, 7, 89, 2, 2, 44, 45, 7, 84, 2, 2, 45, 46, 7, 75, 2,
	2, 46, 47, 7, 86, 2, 2, 47, 48, 7, 71, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50,
	7, 46, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52, 7, 47, 2, 2, 52, 16, 3, 2, 2,
	2, 53, 54, 7, 45, 2, 2, 54, 18, 3, 2, 2, 2, 55, 57, 9, 2, 2, 2, 56, 55,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 20, 3, 2, 2, 2, 60, 62, 4, 50, 59, 2, 61, 60, 3, 2, 2, 2, 62, 63, 3,
	2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 22, 3, 2, 2, 2, 65,
	66, 9, 3, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 8, 12, 2, 2, 68, 24, 3, 2,
	2, 2, 5, 2, 58, 63, 3, 8, 2, 2,
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
	"", "'BEGIN'", "'END'", "':='", "'READ'", "'WRITE'", "','", "'-'", "'+'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ID", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "ID", "NUMBER",
	"WS",
}

type tinyLexer struct {
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

func NewtinyLexer(input antlr.CharStream) *tinyLexer {

	l := new(tinyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "tiny.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tinyLexer tokens.
const (
	tinyLexerT__0   = 1
	tinyLexerT__1   = 2
	tinyLexerT__2   = 3
	tinyLexerT__3   = 4
	tinyLexerT__4   = 5
	tinyLexerT__5   = 6
	tinyLexerT__6   = 7
	tinyLexerT__7   = 8
	tinyLexerID     = 9
	tinyLexerNUMBER = 10
	tinyLexerWS     = 11
)
