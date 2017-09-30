// Generated from fol.g4 by ANTLR 4.7.

package fol

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 90, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 58, 10, 10, 3,
	10, 7, 10, 61, 10, 10, 12, 10, 14, 10, 64, 11, 10, 3, 11, 5, 11, 67, 10,
	11, 3, 11, 7, 11, 70, 10, 11, 12, 11, 14, 11, 73, 11, 11, 3, 12, 3, 12,
	7, 12, 77, 10, 12, 12, 12, 14, 12, 80, 11, 12, 3, 13, 3, 13, 3, 14, 6,
	14, 85, 10, 14, 13, 14, 14, 14, 86, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 2,
	27, 14, 3, 2, 5, 4, 2, 50, 59, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 92, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5,
	31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 37, 3, 2, 2,
	2, 13, 39, 3, 2, 2, 2, 15, 41, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2, 19, 55,
	3, 2, 2, 2, 21, 66, 3, 2, 2, 2, 23, 74, 3, 2, 2, 2, 25, 81, 3, 2, 2, 2,
	27, 84, 3, 2, 2, 2, 29, 30, 7, 46, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7,
	42, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 43, 2, 2, 34, 8, 3, 2, 2, 2, 35,
	36, 7, 40, 2, 2, 36, 10, 3, 2, 2, 2, 37, 38, 7, 126, 2, 2, 38, 12, 3, 2,
	2, 2, 39, 40, 7, 35, 2, 2, 40, 14, 3, 2, 2, 2, 41, 42, 7, 72, 2, 2, 42,
	43, 7, 113, 2, 2, 43, 44, 7, 116, 2, 2, 44, 45, 7, 99, 2, 2, 45, 46, 7,
	110, 2, 2, 46, 47, 7, 110, 2, 2, 47, 16, 3, 2, 2, 2, 48, 49, 7, 71, 2,
	2, 49, 50, 7, 122, 2, 2, 50, 51, 7, 107, 2, 2, 51, 52, 7, 117, 2, 2, 52,
	53, 7, 118, 2, 2, 53, 54, 7, 117, 2, 2, 54, 18, 3, 2, 2, 2, 55, 57, 7,
	65, 2, 2, 56, 58, 9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 62, 3, 2, 2, 2, 59,
	61, 5, 25, 13, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2,
	2, 2, 62, 63, 3, 2, 2, 2, 63, 20, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 67,
	9, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 71, 3, 2, 2, 2, 68, 70, 5, 25, 13,
	2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 22, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 78, 4, 67, 92,
	2, 75, 77, 5, 25, 13, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76,
	3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 24, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2,
	81, 82, 9, 3, 2, 2, 82, 26, 3, 2, 2, 2, 83, 85, 9, 4, 2, 2, 84, 83, 3,
	2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87,
	88, 3, 2, 2, 2, 88, 89, 8, 14, 2, 2, 89, 28, 3, 2, 2, 2, 9, 2, 57, 62,
	66, 71, 78, 86, 3, 8, 2, 2,
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
	"", "','", "'('", "')'", "'&'", "'|'", "'!'", "'Forall'", "'Exists'",
}

var lexerSymbolicNames = []string{
	"", "", "LPAREN", "RPAREN", "AND", "OR", "NOT", "FORALL", "EXISTS", "VARIABLE",
	"CONSTANT", "PREPOSITION", "WS",
}

var lexerRuleNames = []string{
	"T__0", "LPAREN", "RPAREN", "AND", "OR", "NOT", "FORALL", "EXISTS", "VARIABLE",
	"CONSTANT", "PREPOSITION", "CHARACTER", "WS",
}

type folLexer struct {
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

func NewfolLexer(input antlr.CharStream) *folLexer {

	l := new(folLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "fol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// folLexer tokens.
const (
	folLexerT__0        = 1
	folLexerLPAREN      = 2
	folLexerRPAREN      = 3
	folLexerAND         = 4
	folLexerOR          = 5
	folLexerNOT         = 6
	folLexerFORALL      = 7
	folLexerEXISTS      = 8
	folLexerVARIABLE    = 9
	folLexerCONSTANT    = 10
	folLexerPREPOSITION = 11
	folLexerWS          = 12
)
